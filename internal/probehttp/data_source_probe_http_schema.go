package probehttp

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
)

func dataSourceprobeHTTPSchema() map[string]*schema.Schema {
	probeHTTPSchema := probe.DataSourceProbeSchema()

	probeHTTPSchema["transaction"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     httpTransactionResource(),
	}

	probeHTTPSchema["total_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     probe.LimitResource(),
	}

	return probeHTTPSchema
}
