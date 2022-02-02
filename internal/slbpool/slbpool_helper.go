package slbpool

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/slbpool"
)

const profileType = "*slbpool.Profile"

func flattenSLBPool(resList *sdkrrset.ResponseList, rd *schema.ResourceData) error {
	if err := rrset.FlattenRRSet(resList, rd); err != nil {
		return err
	}

	profile, ok := resList.RRSets[0].Profile.(*slbpool.Profile)

	if !ok {
		return errors.ResourceTypeMismatched(profileType, fmt.Sprintf("%T", profile))
	}

	if profile.Monitor != nil {
		if err := rd.Set("monitor", pool.GetMonitorSet(profile.Monitor)); err != nil {
			return err
		}
	}

	if profile.AllFailRecord != nil {
		if err := flattenAllFailRecord(profile.AllFailRecord, rd); err != nil {
			return err
		}
	}

	if len(profile.RDataInfo) > 0 {
		if err := flattenRDataInfo(resList.RRSets[0], rd); err != nil {
			return err
		}
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

func flattenRDataInfo(rrSetData *sdkrrset.RRSet, rd *schema.ResourceData) error {
	set := &schema.Set{F: helper.HashResourceByStringField("rdata")}

	rdataInfoListData := rrSetData.Profile.(*slbpool.Profile).RDataInfo
	for i, rdataInfoData := range rdataInfoListData {
		rdataInfo := make(map[string]interface{})
		rdataInfo["forced_state"] = rdataInfoData.ForcedState
		rdataInfo["availabe_to_serve"] = rdataInfoData.AvailableToServe
		rdataInfo["probing_enabled"] = rdataInfoData.ProbingEnabled
		rdataInfo["description"] = rdataInfoData.Description
		rdataInfo["rdata"] = rrSetData.RData[i]
		set.Add(rdataInfo)
	}

	if err := rd.Set("rdata_info", set); err != nil {
		return err
	}

	return nil
}

func flattenAllFailRecord(allFailRecordData *slbpool.AllFailRecord, rd *schema.ResourceData) error {
	set := &schema.Set{F: schema.HashResource(allFailRecordResource())}
	allFailRecord := make(map[string]interface{})
	allFailRecord["rdata"] = allFailRecordData.RData
	allFailRecord["serving"] = allFailRecordData.Serving
	allFailRecord["description"] = allFailRecordData.Description
	set.Add(allFailRecord)

	if err := rd.Set("all_fail_record", set); err != nil {
		return err
	}

	return nil
}
