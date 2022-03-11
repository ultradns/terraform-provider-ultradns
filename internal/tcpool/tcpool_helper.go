package tcpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/tcpool"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func flattenTCPool(resList *sdkrrset.ResponseList, rd *schema.ResourceData) error {
	if err := rrset.FlattenRRSet(resList, rd); err != nil {
		return err
	}

	profile := resList.RRSets[0].Profile.(*tcpool.Profile)

	if err := rd.Set("rdata_info", getRDataInfoSet(resList.RRSets[0])); err != nil {
		return err
	}

	if err := rd.Set("backup_record", getBackupRecordDataList(profile.BackupRecord)); err != nil {
		return err
	}

	if err := rd.Set("run_probes", profile.RunProbes); err != nil {
		return err
	}

	if err := rd.Set("act_on_probes", profile.ActOnProbes); err != nil {
		return err
	}

	if err := rd.Set("failure_threshold", profile.FailureThreshold); err != nil {
		return err
	}

	if err := rd.Set("max_to_lb", profile.MaxToLB); err != nil {
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

	rdataInfoListData := rrSetData.Profile.(*tcpool.Profile).RDataInfo
	for i, rdataInfoData := range rdataInfoListData {
		rdataInfo := make(map[string]interface{})
		rdataInfo["state"] = rdataInfoData.State
		rdataInfo["available_to_serve"] = rdataInfoData.AvailableToServe
		rdataInfo["run_probes"] = rdataInfoData.RunProbes
		rdataInfo["priority"] = rdataInfoData.Priority
		rdataInfo["weight"] = rdataInfoData.Weight
		rdataInfo["threshold"] = rdataInfoData.Threshold
		rdataInfo["failover_delay"] = rdataInfoData.FailoverDelay
		rdataInfo["rdata"] = rrSetData.RData[i]
		set.Add(rdataInfo)
	}

	return set
}

func getBackupRecordDataList(backupRecordData *pool.BackupRecord) []interface{} {
	var list []interface{}

	if backupRecordData != nil {
		list = make([]interface{}, 1)
		backupRecord := make(map[string]interface{})
		backupRecord["rdata"] = backupRecordData.RData
		backupRecord["failover_delay"] = backupRecordData.FailOverDelay
		backupRecord["available_to_serve"] = backupRecordData.AvailableToServe
		list[0] = backupRecord
	}

	return list
}
