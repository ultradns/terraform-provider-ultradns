package sfpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func dataSourceSFPoolSchema() map[string]*schema.Schema {
	sfPoolSchema := rrset.DataSourceRRSetSchema()

	sfPoolSchema["monitor"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     pool.MonitorResource(),
	}
	sfPoolSchema["backup_record"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     backupRecordResource(),
	}
	sfPoolSchema["status"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	sfPoolSchema["region_failure_sensitivity"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	sfPoolSchema["pool_description"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	sfPoolSchema["live_record_description"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return sfPoolSchema
}
