package probe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
)

const RecordTypeA = "A"

func ResourceProbeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_name": {
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			DiffSuppressFunc: helper.ZoneFQDNDiffSuppress,
			StateFunc:        helper.CaseInSensitiveState,
		},
		"owner_name": {
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			DiffSuppressFunc: helper.OwnerFQDNDiffSuppress,
			StateFunc:        helper.CaseInSensitiveState,
		},
		"pool_type": {
			Type:             schema.TypeString,
			Optional:         true,
			ForceNew:         true,
			Default:          "A",
			DiffSuppressFunc: helper.RecordTypeDiffSuppress,
			ValidateDiagFunc: poolTypeValidation,
		},
		"interval": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "FIVE_MINUTES",
		},
		"agents": {
			Type:     schema.TypeSet,
			Required: true,
			Set:      schema.HashString,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"threshold": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"pool_record": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"guid": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func LimitResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"warning": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"critical": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fail": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func SearchStringResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"warning": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"critical": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fail": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
