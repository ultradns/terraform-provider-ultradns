package dirpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func dataSourceDIRPoolSchema() map[string]*schema.Schema {
	dirPoolSchema := rrset.DataSourceRRSetSchema()

	delete(dirPoolSchema, "record_data")

	dirPoolSchema["rdata_info"] = &schema.Schema{
		Type:     schema.TypeSet,
		Computed: true,
		Elem:     rdataInfoResource(),
	}
	dirPoolSchema["no_response"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     noResponseResource(),
	}
	dirPoolSchema["pool_description"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	dirPoolSchema["conflict_resolve"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	dirPoolSchema["ignore_ecs"] = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}

	return dirPoolSchema
}
