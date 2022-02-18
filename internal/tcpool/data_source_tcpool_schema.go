package tcpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func dataSourceTCPoolSchema() map[string]*schema.Schema {
	tcPoolSchema := rrset.DataSourceRRSetSchema()

	delete(tcPoolSchema, "record_data")

	tcPoolSchema["rdata_info"] = &schema.Schema{
		Type:     schema.TypeSet,
		Computed: true,
		Elem:     rdataInfoResource(),
	}
	tcPoolSchema["backup_record"] = &schema.Schema{
		Type:     schema.TypeSet,
		Computed: true,
		Elem:     backupRecordResource(),
	}
	tcPoolSchema["pool_description"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	tcPoolSchema["run_probes"] = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
	tcPoolSchema["act_on_probes"] = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
	tcPoolSchema["failure_threshold"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
	tcPoolSchema["max_to_lb"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
	tcPoolSchema["status"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return tcPoolSchema
}
