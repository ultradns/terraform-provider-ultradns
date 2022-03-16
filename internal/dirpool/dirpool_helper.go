package dirpool

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/dirpool"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func flattenDIRPool(resList *sdkrrset.ResponseList, rd *schema.ResourceData) error {
	if err := rrset.FlattenRRSet(resList, rd); err != nil {
		return err
	}

	profile := resList.RRSets[0].Profile.(*dirpool.Profile)

	if err := rd.Set("rdata_info", getRDataInfoSet(resList.RRSets[0])); err != nil {
		return err
	}

	if err := rd.Set("no_response", getNoResponseList(profile.NoResponse)); err != nil {
		return err
	}

	if err := rd.Set("pool_description", profile.Description); err != nil {
		return err
	}

	if profile.ConflictResolve == "" {
		profile.ConflictResolve = "GEO"
	}

	if err := rd.Set("conflict_resolve", profile.ConflictResolve); err != nil {
		return err
	}

	if err := rd.Set("ignore_ecs", profile.IgnoreECS); err != nil {
		return err
	}

	return nil
}

func getRDataInfoSet(rrSetData *sdkrrset.RRSet) *schema.Set {
	set := &schema.Set{F: schema.HashResource(rdataInfoResource())}

	rdataInfoListData := rrSetData.Profile.(*dirpool.Profile).RDataInfo
	for i, rdataInfoData := range rdataInfoListData {
		rdataInfo := make(map[string]interface{})
		rdataInfo["type"] = rdataInfoData.Type
		rdataInfo["ttl"] = rdataInfoData.TTL
		rdataInfo["all_non_configured"] = rdataInfoData.AllNonConfigured
		rdataInfo["rdata"] = rrSetData.RData[i]

		if rdataInfoData.GeoInfo != nil {
			rdataInfo["geo_group_name"] = rdataInfoData.GeoInfo.Name
			rdataInfo["geo_codes"] = getGEOCodesSet(rdataInfoData.GeoInfo.Codes)
		}

		if rdataInfoData.IPInfo != nil {
			rdataInfo["ip_group_name"] = rdataInfoData.IPInfo.Name
			rdataInfo["ip"] = getSourceIPInfoSet(rdataInfoData.IPInfo.IPs)
		}

		set.Add(rdataInfo)
	}

	return set
}

func getGEOCodesSet(geoCodes []string) *schema.Set {
	set := &schema.Set{F: schema.HashString}

	for _, data := range geoCodes {
		set.Add(data)
	}

	return set
}

func getSourceIPInfoSet(sourceIPDataList []*dirpool.IPAddress) *schema.Set {
	set := &schema.Set{F: schema.HashResource(sourceIPResource())}

	for _, sourceIPData := range sourceIPDataList {
		sourceIP := make(map[string]interface{})
		sourceIP["start"] = strings.ToLower(sourceIPData.Start)
		sourceIP["end"] = strings.ToLower(sourceIPData.End)
		sourceIP["cidr"] = strings.ToLower(sourceIPData.Cidr)
		sourceIP["address"] = strings.ToLower(sourceIPData.Address)
		set.Add(sourceIP)
	}

	return set
}

func getNoResponseList(noResponseData *dirpool.RDataInfo) []interface{} {
	var list []interface{}

	if noResponseData != nil {
		list = make([]interface{}, 1)
		noRespone := make(map[string]interface{})
		noRespone["all_non_configured"] = noResponseData.AllNonConfigured

		if noResponseData.GeoInfo != nil {
			noRespone["geo_group_name"] = noResponseData.GeoInfo.Name
			noRespone["geo_codes"] = getGEOCodesSet(noResponseData.GeoInfo.Codes)
		}

		if noResponseData.IPInfo != nil {
			noRespone["ip_group_name"] = noResponseData.IPInfo.Name
			noRespone["ip"] = getSourceIPInfoSet(noResponseData.IPInfo.IPs)
		}

		list[0] = noRespone
	}

	return list
}
