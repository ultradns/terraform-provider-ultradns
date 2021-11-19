package zone

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func zoneSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"account_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"change_comment": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"primary_create_info": {
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Set:      zeroIndexHash,
			Elem:     &schema.Resource{Schema: primaryZoneCreateInfoSchema()},
		},
		"secondary_create_info": {
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Set:      zeroIndexHash,
			Elem:     &schema.Resource{Schema: secondaryZoneCreateInfo()},
		},
		"alias_create_info": {
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Set:      zeroIndexHash,
			Elem:     &schema.Resource{Schema: aliasZoneCreateInfo()},
		},
	}
}

func primaryZoneCreateInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"create_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"force_import": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"original_zone_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"inherit": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"name_server": {
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ip": {
						Type:     schema.TypeString,
						Required: true,
					},
					"tsig_key": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"tsig_key_value": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"tsig_algorithm": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"tsig": {
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"tsig_key_name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"tsig_key_value": {
						Type:     schema.TypeString,
						Required: true,
					},
					"tsig_algorithm": {
						Type:     schema.TypeString,
						Required: true,
					},
					"description": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"restrict_ip": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"start_ip": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"end_ip": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"cidr": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"single_ip": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"comment": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"notify_addresses": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"notify_address": {
						Type:     schema.TypeString,
						Required: true,
					},
					"description": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
	}
}

func secondaryZoneCreateInfo() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"notification_email_address": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"primary_name_server": {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			MaxItems: 3,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ip": {
						Type:     schema.TypeString,
						Required: true,
					},
					"tsig_key": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"tsig_key_value": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"tsig_algorithm": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
	}
}

func aliasZoneCreateInfo() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"original_zone_name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func zoneDsSchema() map[string]*schema.Schema {
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

func zeroIndexHash(v interface{}) int {
	return 0
}
