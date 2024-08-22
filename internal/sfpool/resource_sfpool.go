package sfpool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkpool "github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sfpool"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func ResourceSFPool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSFPoolCreate,
		ReadContext:   resourceSFPoolRead,
		UpdateContext: resourceSFPoolUpdate,
		DeleteContext: resourceSFPoolDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceSFPoolSchema(),
	}
}

func resourceSFPoolCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := getNewSFPoolRRSet(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	_, err := services.RecordService.Create(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.RecordID())

	return resourceSFPoolRead(ctx, rd, meta)
}

func resourceSFPoolRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := rrset.GetRRSetKeyFromID(rd.Id())
	rrSetKey.PType = sdkpool.SF
	res, resList, err := services.RecordService.Read(rrSetKey)
	if err != nil && res != nil && res.Status == helper.RESOURCE_NOT_FOUND {
		rd.SetId("")
		tflog.Debug(ctx, err.Error())
		return nil
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if len(resList.RRSets) > 0 {
		if err = flattenSFPool(resList, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceSFPoolUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := getNewSFPoolRRSet(rd)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, err := services.RecordService.Update(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSFPoolRead(ctx, rd, meta)
}

func resourceSFPoolDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func getNewSFPoolRRSet(rd *schema.ResourceData) *sdkrrset.RRSet {
	rrSetData := rrset.NewRRSetWithRecordData(rd)
	profile := &sfpool.Profile{}

	rrSetData.Profile = profile

	if val, ok := rd.GetOk("monitor"); ok && len(val.([]interface{})) > 0 {
		monitorData := val.([]interface{})[0].(map[string]interface{})
		profile.Monitor = pool.GetMonitor(monitorData)
	}

	if val, ok := rd.GetOk("backup_record"); ok && len(val.([]interface{})) > 0 {
		backupRecordData := val.([]interface{})[0].(map[string]interface{})
		profile.BackupRecord = getBackupRecord(backupRecordData)
	}

	if val, ok := rd.GetOk("region_failure_sensitivity"); ok {
		profile.RegionFailureSensitivity = val.(string)
	}

	if val, ok := rd.GetOk("live_record_state"); ok {
		profile.LiveRecordState = val.(string)
	}

	if val, ok := rd.GetOk("live_record_description"); ok {
		profile.LiveRecordDescription = val.(string)
	}

	if val, ok := rd.GetOk("pool_description"); ok {
		profile.PoolDescription = val.(string)
	}

	return rrSetData
}

func getBackupRecord(backupRecordData map[string]interface{}) *sfpool.BackupRecord {
	backupRecord := &sfpool.BackupRecord{}

	if val, ok := backupRecordData["rdata"]; ok {
		backupRecord.RData = val.(string)
	}

	if val, ok := backupRecordData["description"]; ok {
		backupRecord.Description = val.(string)
	}

	return backupRecord
}
