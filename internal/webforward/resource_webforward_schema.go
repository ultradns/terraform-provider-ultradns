package webforward

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	sdkwebforward "github.com/ultradns/ultradns-go-sdk/pkg/webforward"
)

func resourceWebForwardSchema() map[string]*schema.Schema {
	webForwardSchema := webForwardSchema()
	webForwardSchema["zone_name"] = &schema.Schema{
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: helper.ZoneFQDNDiffSuppress,
		StateFunc:        helper.CaseInSensitiveState,
	}
	webForwardSchema["request_to"].ForceNew = true

	return webForwardSchema
}

func dataSourceWebForwardSchema() map[string]*schema.Schema {
	webForwardSchema := webForwardSchema()
	webForwardSchema["zone_name"] = &schema.Schema{
		Type:             schema.TypeString,
		Required:         true,
		DiffSuppressFunc: helper.ZoneFQDNDiffSuppress,
		StateFunc:        helper.CaseInSensitiveState,
	}
	webForwardSchema["guid"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ExactlyOneOf: []string{"guid", "request_to"},
	}
	webForwardSchema["request_to"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ExactlyOneOf: []string{"guid", "request_to"},
	}

	for name, field := range webForwardSchema {
		if name != "zone_name" && name != "guid" && name != "request_to" {
			field.Required = false
			field.Optional = false
			field.Computed = true
			field.ForceNew = false
			field.StateFunc = nil
			field.DiffSuppressFunc = nil
			field.ValidateFunc = nil
			field.ValidateDiagFunc = nil
		}
	}

	return webForwardSchema
}

func webForwardSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"guid": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"request_to": {
			Type:     schema.TypeString,
			Required: true,
		},
		"default_redirect_to": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: helper.URIDiffSuppress,
		},
		"default_forward_type": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{sdkwebforward.HTTP301Redirect, sdkwebforward.HTTP302Redirect, sdkwebforward.HTTP303Redirect, sdkwebforward.HTTP307Redirect, sdkwebforward.Framed}, false),
		},
		"relative_forward_type": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{sdkwebforward.Path, sdkwebforward.Parameter, sdkwebforward.ParameterAndPath}, false),
		},
		"default_redirect_type": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"certificate_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"certificate_managed_type": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"certificate_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"expiration_days": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"state": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"error_description": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"record": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"redirect_to": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"forward_type": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{sdkwebforward.HTTP301Redirect, sdkwebforward.HTTP302Redirect, sdkwebforward.HTTP303Redirect, sdkwebforward.HTTP307Redirect, sdkwebforward.Framed}, false),
					},
					"priority": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"rule": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"header": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"match_criteria": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"value": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"case_insensitive": {
									Type:     schema.TypeBool,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
	}
}
