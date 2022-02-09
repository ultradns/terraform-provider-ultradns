package rdpool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
)

func resourceRDPoolSchema() map[string]*schema.Schema {
	rdPoolSchema := rrset.ResourceRRSetSchema()

	rdPoolSchema["order"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	rdPoolSchema["description"] = &schema.Schema{
		Type:             schema.TypeString,
		Optional:         true,
		DiffSuppressFunc: helper.ComputedDescriptionDiffSuppress,
	}

	return rdPoolSchema
}
