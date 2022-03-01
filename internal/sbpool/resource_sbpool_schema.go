package sbpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func resourceSBPoolSchema() map[string]*schema.Schema {
	sbPoolSchema := rrset.ResourceRRSetSchema()

	delete(sbPoolSchema, "record_data")

	sbPoolSchema["rdata_info"] = &schema.Schema{
		Type:     schema.TypeSet,
		Required: true,
		Elem:     rdataInfoResource(),
	}
	sbPoolSchema["backup_record"] = &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem:     backupRecordResource(),
	}
	sbPoolSchema["pool_description"] = &schema.Schema{
		Type:             schema.TypeString,
		Optional:         true,
		DiffSuppressFunc: helper.ComputedDescriptionDiffSuppress,
	}
	sbPoolSchema["run_probes"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
	}
	sbPoolSchema["act_on_probes"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
	}
	sbPoolSchema["order"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Default:  "ROUND_ROBIN",
	}
	sbPoolSchema["failure_threshold"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}
	sbPoolSchema["max_active"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Default:  1,
	}
	sbPoolSchema["max_served"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	}
	sbPoolSchema["status"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return sbPoolSchema
}

func rdataInfoResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rdata": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"failover_delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NORMAL",
			},
			"run_probes": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"available_to_serve": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func backupRecordResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rdata": {
				Type:     schema.TypeString,
				Required: true,
			},
			"failover_delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"available_to_serve": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}
