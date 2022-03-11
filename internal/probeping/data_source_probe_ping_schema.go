package probeping

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
)

func dataSourceprobePINGSchema() map[string]*schema.Schema {
	probePINGSchema := probe.DataSourceProbeSchema()

	probePINGSchema["packets"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}

	probePINGSchema["packet_size"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}

	probePINGSchema["loss_percent_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     probe.LimitResource(),
	}

	probePINGSchema["total_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     probe.LimitResource(),
	}

	probePINGSchema["average_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     probe.LimitResource(),
	}

	probePINGSchema["run_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     probe.LimitResource(),
	}

	probePINGSchema["avg_run_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     probe.LimitResource(),
	}

	return probePINGSchema
}
