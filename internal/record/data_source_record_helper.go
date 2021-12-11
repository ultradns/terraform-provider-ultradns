package record

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func flattenRRSets(rs []*rrset.RRSet) []map[string]interface{} {
	var rrSets []map[string]interface{}

	for _, rrset := range rs {
		data := make(map[string]interface{})
		data["owner_name"] = rrset.OwnerName
		data["record_type"] = rrset.RRType
		data["ttl"] = rrset.TTL
		set := &schema.Set{F: schema.HashString}
		for _, val := range rrset.RData {
			set.Add(val)
		}
		data["record_data"] = set
		rrSets = append(rrSets, data)
	}

	return rrSets
}
