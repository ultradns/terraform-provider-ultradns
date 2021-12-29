package record

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func flattenRecord(resList *rrset.ResponseList, rd *schema.ResourceData) error {
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

	if err := rd.Set("record_data", flattenRecordData(resList.RRSets[0].RData)); err != nil {
		return err
	}

	return nil
}

func flattenRecordData(recordData []string) *schema.Set {
	set := &schema.Set{F: schema.HashString}

	for _, data := range recordData {
		set.Add(data)
	}

	return set
}
