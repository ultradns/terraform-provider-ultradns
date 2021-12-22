package record

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenRecordData(recordData []string) *schema.Set {
	set := &schema.Set{F: schema.HashString}

	for _, data := range recordData {
		set.Add(data)
	}

	return set
}
