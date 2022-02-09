package sfpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func resourceSFPoolSchema() map[string]*schema.Schema {
	sfPoolSchema := rrset.ResourceRRSetSchema()

	sfPoolSchema["monitor"] = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem:     pool.MonitorResource(),
	}
	sfPoolSchema["backup_record"] = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem:     backupRecordResource(),
	}
	sfPoolSchema["region_failure_sensitivity"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	sfPoolSchema["live_record_state"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	sfPoolSchema["pool_description"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	sfPoolSchema["live_record_description"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	sfPoolSchema["status"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return sfPoolSchema
}

func backupRecordResource() *schema.Resource {
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
		},
	}
}
