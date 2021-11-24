package ultradns

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceZoneSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"query": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"sort": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"reverse": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"limit": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"total_count": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"returned_count": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"offset": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"zones": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"account_name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"type": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}
}
