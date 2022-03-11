package probeping

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
)

func resourceProbePINGSchema() map[string]*schema.Schema {
	probePINGSchema := probe.ResourceProbeSchema()

	probePINGSchema["packets"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Default:  3,
	}

	probePINGSchema["packet_size"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Default:  56,
	}

	probePINGSchema["loss_percent_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     probe.LimitResource(),
	}

	probePINGSchema["total_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     probe.LimitResource(),
	}

	probePINGSchema["average_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     probe.LimitResource(),
	}

	probePINGSchema["run_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     probe.LimitResource(),
	}

	probePINGSchema["avg_run_limit"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     probe.LimitResource(),
	}

	return probePINGSchema
}
