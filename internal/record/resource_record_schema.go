package record

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceRecordSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
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
			Optional: true,
			Default:  86400,
		},
		"record_data": {
			Type:     schema.TypeSet,
			Required: true,
			Set:      schema.HashString,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}
