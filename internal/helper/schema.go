package helper

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func ResultInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"query": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"sort": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"reverse": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},
		"limit": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},
		"offset": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},
		"cursor": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
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
