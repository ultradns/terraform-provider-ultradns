package sbpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sbpool"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func flattenSBPool(resList *sdkrrset.ResponseList, rd *schema.ResourceData) error {
	if err := rrset.FlattenRRSet(resList, rd); err != nil {
		return err
	}

	profile := resList.RRSets[0].Profile.(*sbpool.Profile)

	if err := rd.Set("rdata_info", getRDataInfoSet(resList.RRSets[0])); err != nil {
		return err
	}

	if err := rd.Set("backup_record", getBackupDataInfoSet(profile.BackupRecords)); err != nil {
		return err
	}

	if err := rd.Set("run_probes", profile.RunProbes); err != nil {
		return err
	}

	if err := rd.Set("act_on_probes", profile.ActOnProbes); err != nil {
		return err
	}

	if err := rd.Set("order", profile.Order); err != nil {
		return err
	}

	if err := rd.Set("failure_threshold", profile.FailureThreshold); err != nil {
		return err
	}

	if err := rd.Set("max_active", profile.MaxActive); err != nil {
		return err
	}

	if err := rd.Set("max_served", profile.MaxServed); err != nil {
		return err
	}

	if err := rd.Set("pool_description", profile.Description); err != nil {
		return err
	}

	if err := rd.Set("status", profile.Status); err != nil {
		return err
	}

	return nil
}

func getRDataInfoSet(rrSetData *sdkrrset.RRSet) *schema.Set {
	set := &schema.Set{F: schema.HashResource(rdataInfoResource())}

	rdataInfoListData := rrSetData.Profile.(*sbpool.Profile).RDataInfo
	for i, rdataInfoData := range rdataInfoListData {
		rdataInfo := make(map[string]interface{})
		rdataInfo["state"] = rdataInfoData.State
		rdataInfo["available_to_serve"] = rdataInfoData.AvailableToServe
		rdataInfo["run_probes"] = rdataInfoData.RunProbes
		rdataInfo["priority"] = rdataInfoData.Priority
		rdataInfo["threshold"] = rdataInfoData.Threshold
		rdataInfo["failover_delay"] = rdataInfoData.FailoverDelay
		rdataInfo["rdata"] = rrSetData.RData[i]
		set.Add(rdataInfo)
	}

	return set
}

func getBackupDataInfoSet(backupRecordDataList []*pool.BackupRecord) *schema.Set {
	set := &schema.Set{F: schema.HashResource(backupRecordResource())}

	for _, backupRecordData := range backupRecordDataList {
		backupRecord := make(map[string]interface{})
		backupRecord["rdata"] = backupRecordData.RData
		backupRecord["failover_delay"] = backupRecordData.FailOverDelay
		backupRecord["available_to_serve"] = backupRecordData.AvailableToServe
		set.Add(backupRecord)
	}

	return set
}
