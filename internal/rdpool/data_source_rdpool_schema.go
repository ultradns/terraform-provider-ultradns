package rdpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func dataSourceRDPoolSchema() map[string]*schema.Schema {
	rdPoolSchema := rrset.DataSourceRRSetSchema()

	rdPoolSchema["order"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	rdPoolSchema["description"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return rdPoolSchema
}
