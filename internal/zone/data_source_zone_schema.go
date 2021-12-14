package zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceZoneSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validateZoneName,
		},
		"account_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"dnssec_status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"owner": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"resource_record_count": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"last_modified_time": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"inherit": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"tsig": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem:     tsigResource(),
		},
		"restrict_ip": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem:     restrictIPResource(),
		},
		"notify_addresses": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem:     notifyAddressResource(),
		},
		"registrar_info": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem:     registrarInfoResource(),
		},
		"notification_email_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"primary_name_server_1": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem:     nameServerResource(),
		},
		"primary_name_server_2": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem:     nameServerResource(),
		},
		"primary_name_server_3": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem:     nameServerResource(),
		},
		"transfer_status_details": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem:     transferStatusResource(),
		},
		"original_zone_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func registrarInfoResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"registrar": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"who_is_expiration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name_servers": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     registrarInfoNameServerResource(),
			},
		},
	}
}

func registrarInfoNameServerResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ok": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"unknown": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"missing": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"incorrect": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func transferStatusResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_refresh": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"next_refresh": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_refresh_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_refresh_status_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
