package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		// "apiversion": {
		// 	Type:        schema.TypeString,
		// 	Description: "Api version for UltraDNS rest api.",
		// 	Optional:    true,
		// 	DefaultFunc: schema.EnvDefaultFunc("ULTRADNS_API_VERSION", nil),
		// },
		"useragent": {
			Type:        schema.TypeString,
			Description: "User agent for UltraDNS rest api.",
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("ULTRADNS_USER_AGENT", "terraform-provider-ultrdns"),
		},
	}
}
