package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/dirpool"
	"github.com/ultradns/terraform-provider-ultradns/internal/probedns"
	"github.com/ultradns/terraform-provider-ultradns/internal/probehttp"
	"github.com/ultradns/terraform-provider-ultradns/internal/probeping"
	"github.com/ultradns/terraform-provider-ultradns/internal/probetcp"
	"github.com/ultradns/terraform-provider-ultradns/internal/rdpool"
	"github.com/ultradns/terraform-provider-ultradns/internal/record"
	"github.com/ultradns/terraform-provider-ultradns/internal/sbpool"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/terraform-provider-ultradns/internal/sfpool"
	"github.com/ultradns/terraform-provider-ultradns/internal/slbpool"
	"github.com/ultradns/terraform-provider-ultradns/internal/tcpool"
	"github.com/ultradns/terraform-provider-ultradns/internal/version"
	"github.com/ultradns/terraform-provider-ultradns/internal/zone"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: providerConfigureContext,

		Schema: providerSchema(),

		ResourcesMap: map[string]*schema.Resource{
			"ultradns_zone":       zone.ResourceZone(),
			"ultradns_record":     record.ResourceRecord(),
			"ultradns_rdpool":     rdpool.ResourceRDPool(),
			"ultradns_sfpool":     sfpool.ResourceSFPool(),
			"ultradns_slbpool":    slbpool.ResourceSLBPool(),
			"ultradns_sbpool":     sbpool.ResourceSBPool(),
			"ultradns_tcpool":     tcpool.ResourceTCPool(),
			"ultradns_dirpool":    dirpool.ResourceDIRPool(),
			"ultradns_probe_http": probehttp.ResourceProbeHTTP(),
			"ultradns_probe_ping": probeping.ResourceProbePING(),
			"ultradns_probe_dns":  probedns.ResourceProbeDNS(),
			"ultradns_probe_tcp":  probetcp.ResourceProbeTCP(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ultradns_zone":       zone.DataSourceZone(),
			"ultradns_record":     record.DataSourceRecord(),
			"ultradns_rdpool":     rdpool.DataSourceRDPool(),
			"ultradns_sfpool":     sfpool.DataSourceSFPool(),
			"ultradns_slbpool":    slbpool.DataSourceSLBPool(),
			"ultradns_sbpool":     sbpool.DataSourceSBPool(),
			"ultradns_tcpool":     tcpool.DataSourceTCPool(),
			"ultradns_dirpool":    dirpool.DataSourceDIRPool(),
			"ultradns_probe_http": probehttp.DataSourceprobeHTTP(),
			"ultradns_probe_ping": probeping.DataSourceprobePING(),
			"ultradns_probe_dns":  probedns.DataSourceprobeDNS(),
			"ultradns_probe_tcp":  probetcp.DataSourceprobeTCP(),
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
