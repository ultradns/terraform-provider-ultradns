package slbpool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkpool "github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/slbpool"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func ResourceSLBPool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSLBPoolCreate,
		ReadContext:   resourceSLBPoolRead,
		UpdateContext: resourceSLBPoolUpdate,
		DeleteContext: resourceSLBPoolDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceSLBPoolSchema(),
	}
}

func resourceSLBPoolCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Simple Load Balancing pool resource create context invoked")
	services := meta.(*service.Service)
	rrSetData := getNewSLBPoolRRSet(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	_, err := services.RecordService.Create(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.RecordID())

	return resourceSLBPoolRead(ctx, rd, meta)
}

func resourceSLBPoolRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Simple Load Balancing pool resource read context invoked")
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := rrset.GetRRSetKeyFromID(rd.Id())
	rrSetKey.PType = sdkpool.SLB
	res, resList, err := services.RecordService.Read(rrSetKey)
	if err != nil && res != nil && res.Status == helper.RESOURCE_NOT_FOUND {
		tflog.Warn(ctx, errors.ResourceNotFoundError(rd.Id()).Error())
		rd.SetId("")
		return nil
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if len(resList.RRSets) > 0 {
		if err = flattenSLBPool(resList, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceSLBPoolUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Simple Load Balancing pool resource update context invoked")
	services := meta.(*service.Service)
	rrSetData := getNewSLBPoolRRSet(rd)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, err := services.RecordService.Update(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSLBPoolRead(ctx, rd, meta)
}

func resourceSLBPoolDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Simple Load Balancing pool resource delete context invoked")
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

func getNewSLBPoolRRSet(rd *schema.ResourceData) *sdkrrset.RRSet {
	rrSetData := rrset.NewRRSetWithRecordDataInfo(rd)
	profile := &slbpool.Profile{}
	rrSetData.Profile = profile

	if val, ok := rd.GetOk("region_failure_sensitivity"); ok {
		profile.RegionFailureSensitivity = val.(string)
	}

	if val, ok := rd.GetOk("response_method"); ok {
		profile.ResponseMethod = val.(string)
	}

	if val, ok := rd.GetOk("serving_preference"); ok {
		profile.ServingPreference = val.(string)
	}

	if val, ok := rd.GetOk("pool_description"); ok {
		profile.Description = val.(string)
	}

	if val, ok := rd.GetOk("monitor"); ok && len(val.([]interface{})) > 0 {
		monitorData := val.([]interface{})[0].(map[string]interface{})
		profile.Monitor = pool.GetMonitor(monitorData)
	}

	if val, ok := rd.GetOk("all_fail_record"); ok && len(val.([]interface{})) > 0 {
		allFailRecordData := val.([]interface{})[0].(map[string]interface{})
		profile.AllFailRecord = getAllFailRecord(allFailRecordData)
	}

	if val, ok := rd.GetOk("rdata_info"); ok {
		rdataInfoDataList := val.(*schema.Set).List()
		profile.RDataInfo = getRDataInfoList(rdataInfoDataList)
	}

	return rrSetData
}

func getAllFailRecord(allFailRecordData map[string]interface{}) *slbpool.AllFailRecord {
	allFailRecord := &slbpool.AllFailRecord{}

	if val, ok := allFailRecordData["rdata"]; ok {
		allFailRecord.RData = val.(string)
	}

	if val, ok := allFailRecordData["serving"]; ok {
		allFailRecord.Serving = val.(bool)
	}

	if val, ok := allFailRecordData["description"]; ok {
		allFailRecord.Description = val.(string)
	}

	return allFailRecord
}

func getRDataInfoList(rdataInfoDataList []interface{}) []*slbpool.RDataInfo {
	rdataInfoList := make([]*slbpool.RDataInfo, len(rdataInfoDataList))

	for i, d := range rdataInfoDataList {
		rdataInfoData := d.(map[string]interface{})
		rdataInfoList[i] = getRDataInfo(rdataInfoData)
	}

	return rdataInfoList
}

func getRDataInfo(rdataInfoData map[string]interface{}) *slbpool.RDataInfo {
	rdataInfo := &slbpool.RDataInfo{}

	if val, ok := rdataInfoData["forced_state"]; ok {
		rdataInfo.ForcedState = val.(string)
	}

	if val, ok := rdataInfoData["available_to_serve"]; ok {
		rdataInfo.AvailableToServe = val.(bool)
	}

	if val, ok := rdataInfoData["probing_enabled"]; ok {
		rdataInfo.ProbingEnabled = val.(bool)
	}

	if val, ok := rdataInfoData["description"]; ok {
		rdataInfo.Description = val.(string)
	}

	return rdataInfo
}
