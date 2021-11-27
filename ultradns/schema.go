package ultradns

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resultInfoSchema() map[string]*schema.Schema {
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
		"offset": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"total_count": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"returned_count": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}
