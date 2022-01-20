package rrset

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func NewRRSet(rd *schema.ResourceData) *rrset.RRSet {
	rrSetData := &rrset.RRSet{}

	if val, ok := rd.GetOk("owner_name"); ok {
		rrSetData.OwnerName = val.(string)
	}

	if val, ok := rd.GetOk("record_type"); ok {
		rrSetData.RRType = val.(string)
	}

	if val, ok := rd.GetOk("ttl"); ok {
		rrSetData.TTL = val.(int)
	}

	if val, ok := rd.GetOk("record_data"); ok {
		recordData := val.(*schema.Set).List()
		rrSetData.RData = make([]string, len(recordData))

		for i, record := range recordData {
			rrSetData.RData[i] = record.(string)
		}
	}

	return rrSetData
}

func NewRRSetKey(rd *schema.ResourceData) *rrset.RRSetKey {
	rrSetKeyData := &rrset.RRSetKey{}

	if val, ok := rd.GetOk("zone_name"); ok {
		rrSetKeyData.Zone = val.(string)
	}

	if val, ok := rd.GetOk("owner_name"); ok {
		rrSetKeyData.Name = val.(string)
	}

	if val, ok := rd.GetOk("record_type"); ok {
		rrSetKeyData.Type = val.(string)
	}

	return rrSetKeyData
}

func GetRRSetKeyFromID(id string) *rrset.RRSetKey {
	rrSetKeyData := &rrset.RRSetKey{}
	splitStringData := strings.Split(id, ":")

	if len(splitStringData) == 3 {
		rrSetKeyData.Name = splitStringData[0]
		rrSetKeyData.Zone = splitStringData[1]
		rrSetKeyData.Type = helper.GetRecordTypeString(splitStringData[2])
	}

	return rrSetKeyData
}

func FlattenRRSet(resList *rrset.ResponseList, rd *schema.ResourceData) error {
	currentSchemaZoneName := rd.Get("zone_name").(string)

	if helper.GetZoneFQDN(currentSchemaZoneName) != helper.GetZoneFQDN(resList.ZoneName) {
		if err := rd.Set("zone_name", resList.ZoneName); err != nil {
			return err
		}
	}

	currentSchemaOwnerName := rd.Get("owner_name").(string)

	if helper.GetOwnerFQDN(currentSchemaOwnerName, resList.ZoneName) != resList.RRSets[0].OwnerName {
		if err := rd.Set("owner_name", resList.RRSets[0].OwnerName); err != nil {
			return err
		}
	}

	currentSchemaRecordType := rd.Get("record_type").(string)

	if helper.GetRecordTypeFullString(currentSchemaRecordType) != resList.RRSets[0].RRType {
		if err := rd.Set("record_type", helper.GetRecordTypeString(resList.RRSets[0].RRType)); err != nil {
			return err
		}
	}

	if err := rd.Set("ttl", resList.RRSets[0].TTL); err != nil {
		return err
	}

	if err := rd.Set("record_data", flattenRRSetData(resList.RRSets[0].RData)); err != nil {
		return err
	}

	return nil
}

func flattenRRSetData(recordData []string) *schema.Set {
	set := &schema.Set{F: schema.HashString}

	for _, data := range recordData {
		set.Add(data)
	}

	return set
}
