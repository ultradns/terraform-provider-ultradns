package pool

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func MonitorResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"method": {
				Type:     schema.TypeString,
				Required: true,
			},
			"transmitted_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
