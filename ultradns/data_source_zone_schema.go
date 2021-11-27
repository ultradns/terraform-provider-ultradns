package ultradns

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceZoneSchema() map[string]*schema.Schema {
	zoneSchema := resultInfoSchema()

	zoneSchema["zones"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     zoneResource(),
	}

	return zoneSchema
}

func zoneResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
