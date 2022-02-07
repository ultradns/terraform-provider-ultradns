package slbpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/slbpool"
)

func flattenSLBPool(resList *sdkrrset.ResponseList, rd *schema.ResourceData) error {
	if err := rrset.FlattenRRSet(resList, rd); err != nil {
		return err
	}

	profile, ok := resList.RRSets[0].Profile.(*slbpool.Profile)
	profileSchema := resList.RRSets[0].Profile.GetContext()

	if !ok || slbpool.Schema != profileSchema {
		return errors.ResourceTypeMismatched(slbpool.Schema, profileSchema)
	}

	if err := rd.Set("monitor", pool.GetMonitorList(profile.Monitor)); err != nil {
		return err
	}

	if err := rd.Set("all_fail_record", getAllFailRecordList(profile.AllFailRecord)); err != nil {
		return err
	}

	if err := rd.Set("rdata_info", getRDataInfoSet(resList.RRSets[0])); err != nil {
		return err
	}

	if err := rd.Set("region_failure_sensitivity", profile.RegionFailureSensitivity); err != nil {
		return err
	}

	if err := rd.Set("response_method", profile.ResponseMethod); err != nil {
		return err
	}

	if err := rd.Set("serving_preference", profile.ServingPreference); err != nil {
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

func getAllFailRecordList(allFailRecordData *slbpool.AllFailRecord) []interface{} {
	var list []interface{}

	if allFailRecordData != nil {
		list = make([]interface{}, 1)
		allFailRecord := make(map[string]interface{})
		allFailRecord["rdata"] = allFailRecordData.RData
		allFailRecord["serving"] = allFailRecordData.Serving
		allFailRecord["description"] = allFailRecordData.Description
		list[0] = allFailRecord
	}

	return list
}

func getRDataInfoSet(rrSetData *sdkrrset.RRSet) *schema.Set {
	set := &schema.Set{F: schema.HashResource(rdataInfoResource())}

	rdataInfoListData := rrSetData.Profile.(*slbpool.Profile).RDataInfo
	for i, rdataInfoData := range rdataInfoListData {
		rdataInfo := make(map[string]interface{})
		rdataInfo["forced_state"] = rdataInfoData.ForcedState
		rdataInfo["available_to_serve"] = rdataInfoData.AvailableToServe
		rdataInfo["probing_enabled"] = rdataInfoData.ProbingEnabled
		rdataInfo["description"] = rdataInfoData.Description
		rdataInfo["rdata"] = rrSetData.RData[i]
		set.Add(rdataInfo)
	}

	return set
}
