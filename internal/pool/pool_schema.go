package pool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
)

func MonitorResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: helper.URIDiffSuppress,
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
