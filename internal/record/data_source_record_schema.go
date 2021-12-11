package record

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
)

func dataSourceRecordSchema() map[string]*schema.Schema {
	recordSchema := make(map[string]*schema.Schema)
	queryInfo := helper.QueryInfoSchema()
	for k, v := range queryInfo {
		recordSchema[k] = v
	}
	resultInfo := helper.ResultInfoSchema()
	for k, v := range resultInfo {
		recordSchema[k] = v
	}

	recordSchema["zone_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}

	recordSchema["owner_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	recordSchema["record_type"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	recordSchema["record_sets"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem:     RecordResource(),
	}

	return recordSchema
}

func RecordResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"owner_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"record_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"record_data": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
