package sfpool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/sfpool"
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

	_, err := services.SFPoolService.CreateSFPool(rrSetKeyData, rrSetData)

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.ID())

	return resourceSFPoolRead(ctx, rd, meta)
}

func resourceSFPoolRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := rrset.GetRRSetKeyFromID(rd.Id())
	_, resList, err := services.SFPoolService.ReadSFPool(rrSetKey)

	if err != nil {
		rd.SetId("")

		return nil
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

	_, err := services.SFPoolService.UpdateSFPool(rrSetKeyData, rrSetData)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSFPoolRead(ctx, rd, meta)
}

func resourceSFPoolDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, err := services.SFPoolService.DeleteSFPool(rrSetKeyData)

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

	if val, ok := rd.GetOk("monitor"); ok {
		monitorData := val.([]interface{})[0].(map[string]interface{})
		profile.Monitor = pool.GetMonitor(monitorData)
	}

	if val, ok := rd.GetOk("backup_record"); ok {
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
