package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/zone"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func Provider() *schema.Provider {
	return &schema.Provider{

		ConfigureContextFunc: providerConfigureContext,

		Schema: providerSchema(),

		ResourcesMap: map[string]*schema.Resource{
			"ultradns_zone": zone.ResourceZone(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ultradns_zone": zone.DataSourceZone(),
		},
	}
}

func providerConfigureContext(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	client, err := ultradns.NewClient(
		rd.Get("username").(string),
		rd.Get("password").(string),
		rd.Get("hosturl").(string),
		rd.Get("apiversion").(string),
		rd.Get("useragent").(string),
	)

	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, diags
}
