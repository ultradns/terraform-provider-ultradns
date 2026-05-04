package cdn

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	cdnresource "github.com/ultradns/ultradns-go-sdk/pkg/cdn/resource"
)

// clientCdnIDPattern mirrors the Java validator: ^[A-Za-z0-9\-_]{1,64}$
const clientCdnIDPattern = `^[A-Za-z0-9\-_]{1,64}$`

// nameMaxLen mirrors the Java validator constant NAME_MAX_LENGTH = 64
const nameMaxLen = 64

func mustCompileRegex(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func resourceCDNSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"fqdn": {
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			DiffSuppressFunc: helper.ZoneFQDNDiffSuppress,
			StateFunc:        helper.CaseInSensitiveState,
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{cdnresource.TypeBYOD, cdnresource.TypeSynthetic}, false),
		},
		"name": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringLenBetween(1, nameMaxLen),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ttl": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"content_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		// cdn_providers maps to configs.cdns in the API payload.
		// Required: validator enforces configs.cdns to be non-empty.
		"cdn_providers": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"client_cdn_id": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringMatch(mustCompileRegex(clientCdnIDPattern), "must match "+clientCdnIDPattern),
					},
					"cdn_name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"description": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"fqdn": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: helper.ZoneFQDNDiffSuppress,
						StateFunc:        helper.CaseInSensitiveState,
					},
				},
			},
		},
		// config_properties maps to configs.additionalProperties (inline via @JsonAnyGetter).
		// Values are JSON-encoded strings. Required: validator enforces non-empty.
		"config_properties": {
			Type:     schema.TypeMap,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		// preference_properties maps to preferences.additionalProperties (inline via @JsonAnyGetter).
		// Values are JSON-encoded strings. Required: validator enforces non-empty.
		"preference_properties": {
			Type:     schema.TypeMap,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		// Computed fields returned by the API (read-only).
		"resource_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"version": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_updated": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"owner_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
