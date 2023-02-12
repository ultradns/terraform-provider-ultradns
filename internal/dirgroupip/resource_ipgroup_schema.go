package dirgroupip

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPGroupSchema() map[string]*schema.Schema {
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
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ip": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     sourceIPResource(),
		},
	}
}

func sourceIPResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cidr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
