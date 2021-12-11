package helper

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func QueryInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"query": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"sort": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"cursor": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"reverse": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
			Default:  false,
		},
		"limit": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
			Default:  100,
		},
		"offset": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},
	}
}

func ResultInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"total_count": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"returned_count": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func CursorInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"next": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"previous": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"first": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"last": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
	}
}
