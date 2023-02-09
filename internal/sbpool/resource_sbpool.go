package sbpool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkpool "github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sbpool"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func ResourceSBPool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSBPoolCreate,
		ReadContext:   resourceSBPoolRead,
		UpdateContext: resourceSBPoolUpdate,
		DeleteContext: resourceSBPoolDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceSBPoolSchema(),
	}
}

func resourceSBPoolCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := getNewSBPoolRRSet(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	_, err := services.RecordService.Create(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.RecordID())

	return resourceSBPoolRead(ctx, rd, meta)
}

func resourceSBPoolRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := rrset.GetRRSetKeyFromID(rd.Id())
	rrSetKey.PType = sdkpool.SB
	_, resList, err := services.RecordService.Read(rrSetKey)
	if err != nil {
		rd.SetId("")

		return nil
	}

	if len(resList.RRSets) > 0 {
		if err = flattenSBPool(resList, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceSBPoolUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := getNewSBPoolRRSet(rd)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, err := services.RecordService.Update(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSBPoolRead(ctx, rd, meta)
}

func resourceSBPoolDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func getNewSBPoolRRSet(rd *schema.ResourceData) *sdkrrset.RRSet {
	rrSetData := rrset.NewRRSetWithRecordDataInfo(rd)
	profile := &sbpool.Profile{}
	rrSetData.Profile = profile

	if val, ok := rd.GetOk("pool_description"); ok {
		profile.Description = val.(string)
	}

	if val, ok := rd.GetOk("order"); ok {
		profile.Order = val.(string)
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

	if val, ok := rd.GetOk("max_active"); ok {
		profile.MaxActive = val.(int)
	}

	if val, ok := rd.GetOk("max_served"); ok {
		profile.MaxServed = val.(int)
	}

	if val, ok := rd.GetOk("backup_record"); ok {
		backupRecordDataList := val.(*schema.Set).List()
		profile.BackupRecords = getBackupRecordList(backupRecordDataList)
	}

	if val, ok := rd.GetOk("rdata_info"); ok {
		rdataInfoDataList := val.(*schema.Set).List()
		profile.RDataInfo = getRDataInfoList(rdataInfoDataList)
	}

	return rrSetData
}

func getBackupRecordList(backupRecordDataList []interface{}) []*sdkpool.BackupRecord {
	backupRecordList := make([]*sdkpool.BackupRecord, len(backupRecordDataList))

	for i, d := range backupRecordDataList {
		backupRecordData := d.(map[string]interface{})
		backupRecordList[i] = pool.GetBackupRecord(backupRecordData)
	}

	return backupRecordList
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

	if val, ok := rdataInfoData["threshold"]; ok {
		rdataInfo.Threshold = val.(int)
	}

	if val, ok := rdataInfoData["failover_delay"]; ok {
		rdataInfo.FailoverDelay = val.(int)
	}

	return rdataInfo
}
