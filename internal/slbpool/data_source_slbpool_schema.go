package slbpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/pool"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func dataSourceSLBPoolSchema() map[string]*schema.Schema {
	slbPoolSchema := rrset.DataSourceRRSetSchema()

	delete(slbPoolSchema, "record_data")

	slbPoolSchema["region_failure_sensitivity"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	slbPoolSchema["response_method"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	slbPoolSchema["serving_preference"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	slbPoolSchema["pool_description"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	slbPoolSchema["status"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	slbPoolSchema["monitor"] = &schema.Schema{
		Type:     schema.TypeSet,
		Computed: true,
		Set:      helper.HashSingleSetResource,
		Elem:     pool.MonitorResource(),
	}
	slbPoolSchema["all_fail_record"] = &schema.Schema{
		Type:     schema.TypeSet,
		Computed: true,
		Set:      helper.HashSingleSetResource,
		Elem:     allFailRecordResource(),
	}
	slbPoolSchema["rdata_info"] = &schema.Schema{
		Type:     schema.TypeSet,
		Computed: true,
		Set:      helper.HashSingleSetResource,
		Elem:     rdataInfoResource(),
	}

	return slbPoolSchema
}
