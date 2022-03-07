package probehttp

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
)

func resourceProbeHTTPSchema() map[string]*schema.Schema {
	probeHTTPSchema := probe.ResourceProbeSchema()

	probeHTTPSchema["transaction"] = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		Elem:     httpTransactionResource(),
	}

	probeHTTPSchema["total_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     probe.LimitResource(),
	}

	return probeHTTPSchema
}

func httpTransactionResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"method": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"transmitted_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"follow_redirects": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"expected_response": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_string": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     probe.SearchStringResource(),
			},
			"connect_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     probe.LimitResource(),
			},
			"avg_connect_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     probe.LimitResource(),
			},
			"run_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     probe.LimitResource(),
			},
			"avg_run_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     probe.LimitResource(),
			},
		},
	}
}
