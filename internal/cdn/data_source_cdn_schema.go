package cdn

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceCDNSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"fqdn": {
			Type:     schema.TypeString,
			Required: true,
		},
		// All remaining fields are read-only; values come from the API.
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ttl": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"content_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"resource_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"version": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_updated": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"owner_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		// cdn_providers exposes the configs.cdns list from the API response.
		"cdn_providers": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"client_cdn_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"cdn_name": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"description": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"fqdn": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
		// config_properties exposes configs.additionalProperties as JSON strings.
		"config_properties": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		// preference_properties exposes preferences.additionalProperties as JSON strings.
		"preference_properties": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}
