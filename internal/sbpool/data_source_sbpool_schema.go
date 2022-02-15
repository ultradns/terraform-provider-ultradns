package sbpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func dataSourceSBPoolSchema() map[string]*schema.Schema {
	sbPoolSchema := rrset.DataSourceRRSetSchema()

	delete(sbPoolSchema, "record_data")

	sbPoolSchema["rdata_info"] = &schema.Schema{
		Type:     schema.TypeSet,
		Computed: true,
		Elem:     rdataInfoResource(),
	}
	sbPoolSchema["backup_record"] = &schema.Schema{
		Type:     schema.TypeSet,
		Computed: true,
		Elem:     backupRecordResource(),
	}
	sbPoolSchema["pool_description"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	sbPoolSchema["run_probes"] = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
	sbPoolSchema["act_on_probes"] = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
	sbPoolSchema["order"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	sbPoolSchema["failure_threshold"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
	sbPoolSchema["max_active"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
	sbPoolSchema["max_served"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
	sbPoolSchema["status"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return sbPoolSchema
}
