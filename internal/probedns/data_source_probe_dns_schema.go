package probedns

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
)

func dataSourceprobeDNSSchema() map[string]*schema.Schema {
	probeDNSSchema := probe.DataSourceProbeSchema()

	probeDNSSchema["port"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}

	probeDNSSchema["tcp_only"] = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}

	probeDNSSchema["type"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	probeDNSSchema["query_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	probeDNSSchema["response"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     probe.SearchStringResource(),
	}

	probeDNSSchema["run_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     probe.LimitResource(),
	}

	probeDNSSchema["avg_run_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     probe.LimitResource(),
	}

	return probeDNSSchema
}
