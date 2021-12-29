package record

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRecordSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"owner_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"record_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ttl": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"record_data": {
			Type:     schema.TypeSet,
			Computed: true,
			Set:      schema.HashString,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}
