package rrset

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
)

func ResourceRRSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_name": {
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			DiffSuppressFunc: helper.ZoneFQDNDiffSuppress,
		},
		"owner_name": {
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			DiffSuppressFunc: helper.OwnerFQDNDiffSuppress,
		},
		"record_type": {
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			DiffSuppressFunc: helper.RecordTypeDiffSuppress,
			ValidateDiagFunc: helper.RecordTypeValidation,
		},
		"ttl": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  86400,
		},
		"record_data": {
			Type:     schema.TypeSet,
			Required: true,
			Set:      schema.HashString,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}
