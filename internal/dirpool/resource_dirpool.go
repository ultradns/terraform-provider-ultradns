package dirpool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/dirpool"
	sdkpool "github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func ResourceDIRPool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDIRPoolCreate,
		ReadContext:   resourceDIRPoolRead,
		UpdateContext: resourceDIRPoolUpdate,
		DeleteContext: resourceDIRPoolDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceDIRPoolSchema(),
	}
}

func resourceDIRPoolCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := getNewDIRPoolRRSet(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	_, err := services.RecordService.Create(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.RecordID())

	return resourceDIRPoolRead(ctx, rd, meta)
}

func resourceDIRPoolRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := rrset.GetRRSetKeyFromID(rd.Id())
	rrSetKey.PType = sdkpool.DIR
	_, resList, err := services.RecordService.Read(rrSetKey)
	if err != nil {
		rd.SetId("")
		tflog.Error(ctx, err.Error())
		return nil
	}

	if len(resList.RRSets) > 0 {
		if err = flattenDIRPool(resList, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceDIRPoolUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := getNewDIRPoolRRSet(rd)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, err := services.RecordService.Update(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceDIRPoolRead(ctx, rd, meta)
}

func resourceDIRPoolDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, err := services.RecordService.Delete(rrSetKeyData)
	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	rd.SetId("")

	return diags
}

func getNewDIRPoolRRSet(rd *schema.ResourceData) *sdkrrset.RRSet {
	rrSetData := rrset.NewRRSetWithRecordDataInfo(rd)
	profile := &dirpool.Profile{}
	rrSetData.Profile = profile

	if val, ok := rd.GetOk("pool_description"); ok {
		profile.Description = val.(string)
	}

	if val, ok := rd.GetOk("conflict_resolve"); ok {
		profile.ConflictResolve = val.(string)
	}

	if val, ok := rd.GetOk("ignore_ecs"); ok {
		profile.IgnoreECS = val.(bool)
	}

	if val, ok := rd.GetOk("no_response"); ok && len(val.([]interface{})) > 0 {
		noResponseData := val.([]interface{})[0].(map[string]interface{})
		profile.NoResponse = getRDataInfo(noResponseData)
	}

	if val, ok := rd.GetOk("rdata_info"); ok {
		rdataInfoDataList := val.(*schema.Set).List()
		profile.RDataInfo = getRDataInfoList(rdataInfoDataList)
	}

	return rrSetData
}

func getRDataInfoList(rdataInfoDataList []interface{}) []*dirpool.RDataInfo {
	rdataInfoList := make([]*dirpool.RDataInfo, len(rdataInfoDataList))

	for i, d := range rdataInfoDataList {
		rdataInfoData := d.(map[string]interface{})
		rdataInfoList[i] = getRDataInfo(rdataInfoData)
	}

	return rdataInfoList
}

func getRDataInfo(rdataInfoData map[string]interface{}) *dirpool.RDataInfo {
	rdataInfo := &dirpool.RDataInfo{}
	geoInfo := &dirpool.GEOInfo{}
	ipInfo := &dirpool.IPInfo{}

	rdataInfo.GeoInfo = geoInfo
	rdataInfo.IPInfo = ipInfo

	if val, ok := rdataInfoData["type"]; ok {
		rdataInfo.Type = val.(string)
	}

	if val, ok := rdataInfoData["ttl"]; ok {
		rdataInfo.TTL = val.(int)
	}

	if val, ok := rdataInfoData["all_non_configured"]; ok {
		rdataInfo.AllNonConfigured = val.(bool)
	}

	if val, ok := rdataInfoData["geo_group_name"]; ok {
		geoInfo.Name = val.(string)
	}

	if val, ok := rdataInfoData["geo_codes"]; ok {
		geoCodesData := val.(*schema.Set).List()
		geoInfo.Codes = make([]string, len(geoCodesData))

		for i, geoCode := range geoCodesData {
			geoInfo.Codes[i] = geoCode.(string)
		}
	}
	if val, ok := rdataInfoData["geo_account_level"]; ok {
		geoInfo.IsAccountLevel = val.(bool)
	}

	if val, ok := rdataInfoData["ip_group_name"]; ok {
		ipInfo.Name = val.(string)
	}

	if val, ok := rdataInfoData["ip"]; ok {
		sourceIPInfoDataList := val.(*schema.Set).List()
		ipInfo.IPs = getSourceIPInfoList(sourceIPInfoDataList)
	}

	if val, ok := rdataInfoData["ip_account_level"]; ok {
		ipInfo.IsAccountLevel = val.(bool)
	}

	return rdataInfo
}

func getSourceIPInfoList(sourceIPInfoDataList []interface{}) []*dirpool.IPAddress {
	sourceIPInfoList := make([]*dirpool.IPAddress, len(sourceIPInfoDataList))

	for i, d := range sourceIPInfoDataList {
		sourceIPInfoData := d.(map[string]interface{})
		sourceIPInfoList[i] = getSourceIPInfo(sourceIPInfoData)
	}

	return sourceIPInfoList
}

func getSourceIPInfo(sourceIPInfoData map[string]interface{}) *dirpool.IPAddress {
	sourceIPInfo := &dirpool.IPAddress{}

	if val, ok := sourceIPInfoData["start"]; ok {
		sourceIPInfo.Start = val.(string)
	}

	if val, ok := sourceIPInfoData["end"]; ok {
		sourceIPInfo.End = val.(string)
	}

	if val, ok := sourceIPInfoData["cidr"]; ok {
		sourceIPInfo.Cidr = val.(string)
	}

	if val, ok := sourceIPInfoData["address"]; ok {
		sourceIPInfo.Address = val.(string)
	}

	return sourceIPInfo
}
