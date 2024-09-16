package probe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceProbeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"owner_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"pool_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "A",
		},
		"guid": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"agents": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Set:      schema.HashString,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"pool_record": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}
