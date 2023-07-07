package probetcp

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
)

func resourceProbeTCPSchema() map[string]*schema.Schema {
	probeTCPSchema := probe.ResourceProbeSchema()

	probeTCPSchema["port"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Default:  443,
	}

	probeTCPSchema["control_ip"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	probeTCPSchema["connect_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem:     probe.LimitResource(),
	}

	probeTCPSchema["avg_connect_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     probe.LimitResource(),
	}

	return probeTCPSchema
}
