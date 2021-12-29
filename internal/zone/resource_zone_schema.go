package zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceZoneSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"account_name": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			DefaultFunc: schema.EnvDefaultFunc("ULTRADNS_ACCOUNT", nil),
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"change_comment": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"primary_create_info": {
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem:     primaryZoneCreateInfoResource(),
		},
		"secondary_create_info": {
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem:     secondaryZoneCreateInfoResource(),
		},
		"alias_create_info": {
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem:     aliasZoneCreateInfoResource(),
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
		"registrar_info": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem:     registrarInfoResource(),
		},
		"transfer_status_details": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem:     transferStatusResource(),
		},
	}
}

func primaryZoneCreateInfoResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"create_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"force_import": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"original_zone_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inherit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_server": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem:     nameServerResource(),
			},
			"tsig": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem:     tsigResource(),
			},
			"restrict_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     restrictIPResource(),
			},
			"notify_addresses": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     notifyAddressResource(),
			},
		},
	}
}

func secondaryZoneCreateInfoResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notification_email_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_name_server_1": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Elem:     nameServerResource(),
			},
			"primary_name_server_2": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Elem:     nameServerResource(),
			},
			"primary_name_server_3": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Elem:     nameServerResource(),
			},
		},
	}
}

func aliasZoneCreateInfoResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"original_zone_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func nameServerResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tsig_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tsig_key_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tsig_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func tsigResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tsig_key_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tsig_key_value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tsig_algorithm": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func restrictIPResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"start_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cidr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"single_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func notifyAddressResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notify_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
