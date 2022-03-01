package tcpool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkpool "github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/tcpool"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func ResourceTCPool() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceTCPoolCreate,
		ReadContext:   resourceTCPoolRead,
		UpdateContext: resourceTCPoolUpdate,
		DeleteContext: resourceTCPoolDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceTCPoolSchema(),
	}
}

func resourceTCPoolCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := getNewTCPoolRRSet(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	_, err := services.RecordService.Create(rrSetKeyData, rrSetData)

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.RecordID())

	return resourceTCPoolRead(ctx, rd, meta)
}

func resourceTCPoolRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := rrset.GetRRSetKeyFromID(rd.Id())
	rrSetKey.PType = sdkpool.TC
	_, resList, err := services.RecordService.Read(rrSetKey)

	if err != nil {
		rd.SetId("")

		return nil
	}

	if len(resList.RRSets) > 0 {
		if err = flattenTCPool(resList, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceTCPoolUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := getNewTCPoolRRSet(rd)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, err := services.RecordService.Update(rrSetKeyData, rrSetData)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceTCPoolRead(ctx, rd, meta)
}

func resourceTCPoolDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func getNewTCPoolRRSet(rd *schema.ResourceData) *sdkrrset.RRSet {
	rrSetData := rrset.NewRRSetWithRecordDataInfo(rd)
	profile := &tcpool.Profile{}
	rrSetData.Profile = profile

	if val, ok := rd.GetOk("pool_description"); ok {
		profile.Description = val.(string)
	}

	if val, ok := rd.GetOk("run_probes"); ok {
		profile.RunProbes = val.(bool)
	}

	if val, ok := rd.GetOk("act_on_probes"); ok {
		profile.ActOnProbes = val.(bool)
	}

	if val, ok := rd.GetOk("failure_threshold"); ok {
		profile.FailureThreshold = val.(int)
	}

	if val, ok := rd.GetOk("max_to_lb"); ok {
		profile.MaxToLB = val.(int)
	}

	if val, ok := rd.GetOk("backup_record"); ok && len(val.([]interface{})) > 0 {
		backupRecordData := val.([]interface{})[0].(map[string]interface{})
		profile.BackupRecord = pool.GetBackupRecord(backupRecordData)
	}

	if val, ok := rd.GetOk("rdata_info"); ok {
		rdataInfoDataList := val.(*schema.Set).List()
		profile.RDataInfo = getRDataInfoList(rdataInfoDataList)
	}

	return rrSetData
}

func getRDataInfoList(rdataInfoDataList []interface{}) []*sdkpool.RDataInfo {
	rdataInfoList := make([]*sdkpool.RDataInfo, len(rdataInfoDataList))

	for i, d := range rdataInfoDataList {
		rdataInfoData := d.(map[string]interface{})
		rdataInfoList[i] = getRDataInfo(rdataInfoData)
	}

	return rdataInfoList
}

func getRDataInfo(rdataInfoData map[string]interface{}) *sdkpool.RDataInfo {
	rdataInfo := &sdkpool.RDataInfo{}

	if val, ok := rdataInfoData["state"]; ok {
		rdataInfo.State = val.(string)
	}

	if val, ok := rdataInfoData["run_probes"]; ok {
		rdataInfo.RunProbes = val.(bool)
	}

	if val, ok := rdataInfoData["priority"]; ok {
		rdataInfo.Priority = val.(int)
	}

	if val, ok := rdataInfoData["weight"]; ok {
		rdataInfo.Weight = val.(int)
	}

	if val, ok := rdataInfoData["threshold"]; ok {
		rdataInfo.Threshold = val.(int)
	}

	if val, ok := rdataInfoData["failover_delay"]; ok {
		rdataInfo.FailoverDelay = val.(int)
	}

	return rdataInfo
}
