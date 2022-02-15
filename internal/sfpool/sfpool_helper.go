package sfpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sfpool"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func flattenSFPool(resList *sdkrrset.ResponseList, rd *schema.ResourceData) error {
	if err := rrset.FlattenRRSetWithRecordData(resList, rd); err != nil {
		return err
	}

	profile, ok := resList.RRSets[0].Profile.(*sfpool.Profile)
	profileSchema := resList.RRSets[0].Profile.GetContext()

	if !ok || sfpool.Schema != profileSchema {
		return errors.ResourceTypeMismatched(sfpool.Schema, profileSchema)
	}

	if err := rd.Set("monitor", pool.GetMonitorList(profile.Monitor)); err != nil {
		return err
	}

	if err := rd.Set("backup_record", getBackupRecordList(profile.BackupRecord)); err != nil {
		return err
	}

	if err := rd.Set("region_failure_sensitivity", profile.RegionFailureSensitivity); err != nil {
		return err
	}

	if err := rd.Set("live_record_description", profile.LiveRecordDescription); err != nil {
		return err
	}

	if err := rd.Set("pool_description", profile.PoolDescription); err != nil {
		return err
	}

	if err := rd.Set("status", profile.Status); err != nil {
		return err
	}

	return nil
}

func getBackupRecordList(backupRecordData *sfpool.BackupRecord) []interface{} {
	var list []interface{}

	if backupRecordData != nil {
		list = make([]interface{}, 1)
		backupRecord := make(map[string]interface{})
		backupRecord["rdata"] = backupRecordData.RData
		backupRecord["description"] = backupRecordData.Description
		list[0] = backupRecord
	}

	return list
}
