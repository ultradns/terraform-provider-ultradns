package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rdpool"
	"github.com/ultradns/terraform-provider-ultradns/internal/record"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/terraform-provider-ultradns/internal/sfpool"
	"github.com/ultradns/terraform-provider-ultradns/internal/slbpool"
	"github.com/ultradns/terraform-provider-ultradns/internal/version"
	"github.com/ultradns/terraform-provider-ultradns/internal/zone"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{

		ConfigureContextFunc: providerConfigureContext,

		Schema: providerSchema(),

		ResourcesMap: map[string]*schema.Resource{
			"ultradns_zone":    zone.ResourceZone(),
			"ultradns_record":  record.ResourceRecord(),
			"ultradns_rdpool":  rdpool.ResourceRDPool(),
			"ultradns_sfpool":  sfpool.ResourceSFPool(),
			"ultradns_slbpool": slbpool.ResourceSLBPool(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ultradns_zone":    zone.DataSourceZone(),
			"ultradns_record":  record.DataSourceRecord(),
			"ultradns_rdpool":  rdpool.DataSourceRDPool(),
			"ultradns_sfpool":  sfpool.DataSourceSFPool(),
			"ultradns_slbpool": slbpool.DataSourceSLBPool(),
		},
	}
}

func providerConfigureContext(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	cnf := client.Config{
		Username:  rd.Get("username").(string),
		Password:  rd.Get("password").(string),
		HostURL:   rd.Get("hosturl").(string),
		UserAgent: version.GetProviderVersion(),
	}

	client, err := client.NewClient(cnf)

	if err != nil {
		return nil, diag.FromErr(err)
	}

	service, err := service.NewService(client)

	if err != nil {
		return nil, diag.FromErr(err)
	}

	return service, diags
}
