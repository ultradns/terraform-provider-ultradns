package probedns

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
)

func resourceProbeDNSSchema() map[string]*schema.Schema {
	probeDNSSchema := probe.ResourceProbeSchema()

	probeDNSSchema["port"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Default:  53,
	}

	probeDNSSchema["tcp_only"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	}

	probeDNSSchema["type"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Default:  "NULL",
	}

	probeDNSSchema["query_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	probeDNSSchema["response"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     probe.SearchStringResource(),
	}

	probeDNSSchema["run_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     probe.LimitResource(),
	}

	probeDNSSchema["avg_run_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     probe.LimitResource(),
	}

	return probeDNSSchema
}
