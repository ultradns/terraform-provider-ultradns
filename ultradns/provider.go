package ultradns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func Provider() *schema.Provider {
	return &schema.Provider{

		ConfigureContextFunc: providerConfigureContext,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Description: "User name for UltraDNS rest api.",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ULTRADNS_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password for UltraDNS rest api.",
				Sensitive:   true,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ULTRADNS_PASSWORD", nil),
			},
			"hosturl": {
				Type:        schema.TypeString,
				Description: "Host url for UltraDNS rest api.",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ULTRADNS_HOST_URL", nil),
			},
			"apiversion": {
				Type:        schema.TypeString,
				Description: "Api version for UltraDNS rest api.",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ULTRADNS_API_VERSION", nil),
			},
			"useragent": {
				Type:        schema.TypeString,
				Description: "User agent for UltraDNS rest api.",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ULTRADNS_USER_AGENT", "terraform-provider-ultrdns"),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"ultradns_zone": ResourceZone(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ultradns_zone": DataSourceZone(),
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
