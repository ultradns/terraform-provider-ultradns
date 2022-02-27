package tcpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func resourceTCPoolSchema() map[string]*schema.Schema {
	tcPoolSchema := rrset.ResourceRRSetSchema()

	delete(tcPoolSchema, "record_data")

	tcPoolSchema["rdata_info"] = &schema.Schema{
		Type:     schema.TypeSet,
		Required: true,
		Elem:     rdataInfoResource(),
	}
	tcPoolSchema["backup_record"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     backupRecordResource(),
	}
	tcPoolSchema["pool_description"] = &schema.Schema{
		Type:             schema.TypeString,
		Optional:         true,
		DiffSuppressFunc: helper.ComputedDescriptionDiffSuppress,
	}
	tcPoolSchema["run_probes"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
	}
	tcPoolSchema["act_on_probes"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
	}
	tcPoolSchema["failure_threshold"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}
	tcPoolSchema["max_to_lb"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}
	tcPoolSchema["status"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return tcPoolSchema
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
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"threshold": {
				Type:     schema.TypeInt,
				Required: true,
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
