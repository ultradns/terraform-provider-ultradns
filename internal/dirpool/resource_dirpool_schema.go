package dirpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func resourceDIRPoolSchema() map[string]*schema.Schema {
	dirPoolSchema := rrset.ResourceRRSetSchema()

	delete(dirPoolSchema, "record_data")

	dirPoolSchema["ttl"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}

	dirPoolSchema["rdata_info"] = &schema.Schema{
		Type:     schema.TypeSet,
		Required: true,
		Elem:     rdataInfoResource(),
	}
	dirPoolSchema["no_response"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     noResponseResource(),
	}
	dirPoolSchema["pool_description"] = &schema.Schema{
		Type:             schema.TypeString,
		Optional:         true,
		DiffSuppressFunc: helper.ComputedDescriptionDiffSuppress,
	}
	dirPoolSchema["conflict_resolve"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Default:  "GEO",
	}
	dirPoolSchema["ignore_ecs"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	}

	return dirPoolSchema
}

func rdataInfoResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rdata": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"all_non_configured": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"geo_group_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"geo_codes": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ip_group_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     sourceIPResource(),
			},
		},
	}
}

func noResponseResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_non_configured": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"geo_group_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"geo_codes": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ip_group_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     sourceIPResource(),
			},
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
