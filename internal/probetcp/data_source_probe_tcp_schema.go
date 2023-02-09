package probetcp

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
)

func dataSourceprobeTCPSchema() map[string]*schema.Schema {
	probeTCPSchema := probe.DataSourceProbeSchema()

	probeTCPSchema["port"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}

	probeTCPSchema["control_ip"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	probeTCPSchema["connect_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     probe.LimitResource(),
	}

	probeTCPSchema["avg_connect_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     probe.LimitResource(),
	}

	return probeTCPSchema
}
