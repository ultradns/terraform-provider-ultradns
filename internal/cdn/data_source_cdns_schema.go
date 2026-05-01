package cdn

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceCDNsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"page": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
			ValidateFunc: validation.IntAtLeast(1),
		},
		"size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
			ValidateFunc: validation.IntAtLeast(1),
		},
		"total_pages": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"total_elements": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"cdns": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"fqdn": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"resource_id": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"name": {
					Type:     schema.TypeString,
					Computed: true,
				},
			}},
		},
	}
}
