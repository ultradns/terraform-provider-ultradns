package rrset

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	sdkhelper "github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func newRRSet(rd *schema.ResourceData) *rrset.RRSet {
	rrSetData := &rrset.RRSet{}

	if val, ok := rd.GetOk("owner_name"); ok {
		rrSetData.OwnerName = strings.ToLower(val.(string))
	}

	if val, ok := rd.GetOk("record_type"); ok {
		rrSetData.RRType = val.(string)
	}

	if val, ok := rd.GetOk("ttl"); ok {
		rrSetData.TTL = val.(int)
	}

	return rrSetData
}

func NewRRSetWithRecordData(rd *schema.ResourceData) *rrset.RRSet {
	rrSetData := newRRSet(rd)

	if val, ok := rd.GetOk("record_data"); ok {
		recordData := val.(*schema.Set).List()
		rrSetData.RData = make([]string, len(recordData))

		for i, record := range recordData {
			rrSetData.RData[i] = record.(string)
		}
	}

	return rrSetData
}

func NewRRSetWithRecordDataInfo(rd *schema.ResourceData) *rrset.RRSet {
	rrSetData := newRRSet(rd)

	if val, ok := rd.GetOk("rdata_info"); ok {
		rDataInfoList := val.(*schema.Set).List()
		rrSetData.RData = make([]string, len(rDataInfoList))

		for i, rDataInfoData := range rDataInfoList {
			rDataInfo := rDataInfoData.(map[string]interface{})
			rrSetData.RData[i] = rDataInfo["rdata"].(string)
		}
	}

	return rrSetData
}

func NewRRSetKey(rd *schema.ResourceData) *rrset.RRSetKey {
	rrSetKeyData := &rrset.RRSetKey{}

	if val, ok := rd.GetOk("zone_name"); ok {
		rrSetKeyData.Zone = strings.ToLower(val.(string))
	}

	if val, ok := rd.GetOk("owner_name"); ok {
		rrSetKeyData.Owner = strings.ToLower(val.(string))
	}

	if val, ok := rd.GetOk("record_type"); ok {
		rrSetKeyData.RecordType = val.(string)
	}

	return rrSetKeyData
}

func GetRRSetKeyFromID(id string) *rrset.RRSetKey {
	rrSetKeyData := &rrset.RRSetKey{}
	splitStringData := strings.Split(id, ":")

	if len(splitStringData) == 3 {
		rrSetKeyData.Owner = splitStringData[0]
		rrSetKeyData.Zone = splitStringData[1]
		rrSetKeyData.RecordType = sdkhelper.GetRecordTypeString(splitStringData[2])
	}

	return rrSetKeyData
}

func FlattenRRSet(resList *rrset.ResponseList, rd *schema.ResourceData) error {
	if err := rd.Set("zone_name", sdkhelper.GetZoneFQDN(resList.ZoneName)); err != nil {
		return err
	}

	if err := rd.Set("owner_name", resList.RRSets[0].OwnerName); err != nil {
		return err
	}

	if err := rd.Set("record_type", sdkhelper.GetRecordTypeString(resList.RRSets[0].RRType)); err != nil {
		return err
	}

	if err := rd.Set("ttl", resList.RRSets[0].TTL); err != nil {
		return err
	}

	return nil
}

func FlattenRRSetWithRecordData(resList *rrset.ResponseList, rd *schema.ResourceData) error {
	if err := rd.Set("record_data", helper.GetSchemaSetFromList(resList.RRSets[0].RData)); err != nil {
		return err
	}

	return FlattenRRSet(resList, rd)
}
