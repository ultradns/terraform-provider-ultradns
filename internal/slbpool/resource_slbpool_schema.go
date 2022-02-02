package slbpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func resourceSLBPoolSchema() map[string]*schema.Schema {
	slbPoolSchema := rrset.ResourceRRSetSchema()

	delete(slbPoolSchema, "record_data")

	slbPoolSchema["region_failure_sensitivity"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	slbPoolSchema["response_method"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	slbPoolSchema["serving_preference"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	slbPoolSchema["pool_description"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	slbPoolSchema["status"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	slbPoolSchema["monitor"] = &schema.Schema{
		Type:     schema.TypeSet,
		MaxItems: 1,
		Required: true,
		Elem:     pool.MonitorResource(),
	}
	slbPoolSchema["all_fail_record"] = &schema.Schema{
		Type:     schema.TypeSet,
		Required: true,
		MaxItems: 1,
		Elem:     allFailRecordResource(),
	}
	slbPoolSchema["rdata_info"] = &schema.Schema{
		Type:     schema.TypeSet,
		Required: true,
		MaxItems: 5,
		Set:      helper.HashResourceByStringField("rdata"),
		Elem:     rdataInfoResource(),
	}

	return slbPoolSchema
}

func allFailRecordResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rdata": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"serving": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func rdataInfoResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rdata": {
				Type:     schema.TypeString,
				Required: true,
			},
			"forced_state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NOT_FORCED",
			},
			"probing_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"availabe_to_serve": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}
