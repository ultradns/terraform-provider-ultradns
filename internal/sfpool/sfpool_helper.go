package sfpool

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/sfpool"
)

const profileType = "*sfpool.Profile"

func flattenSFPool(resList *sdkrrset.ResponseList, rd *schema.ResourceData) error {
	if err := rrset.FlattenRRSetWithRecordData(resList, rd); err != nil {
		return err
	}

	profile, ok := resList.RRSets[0].Profile.(*sfpool.Profile)

	if !ok {
		return errors.ResourceTypeMismatched(profileType, fmt.Sprintf("%T", profile))
	}

	if profile.Monitor != nil {
		if err := rd.Set("monitor", pool.GetMonitorSet(profile.Monitor)); err != nil {
			return err
		}
	}

	if profile.BackupRecord != nil {
		if err := rd.Set("backup_record", getBackupRecordSet(profile.BackupRecord)); err != nil {
			return err
		}
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

func getBackupRecordSet(backupRecordData *sfpool.BackupRecord) *schema.Set {
	set := &schema.Set{F: schema.HashResource(backupRecordResource())}
	backupRecord := make(map[string]interface{})
	backupRecord["rdata"] = backupRecordData.RData
	backupRecord["description"] = backupRecordData.Description
	set.Add(backupRecord)

	return set
}
