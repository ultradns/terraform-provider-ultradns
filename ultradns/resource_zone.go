package ultradns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ultradns "terraform-provider-ultradns/udnssdk"
)

func resourceZone() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceZoneCreate,
		ReadContext:   resourceZoneRead,
		UpdateContext: resourceZoneUpdate,
		DeleteContext: resourceZoneDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"change_comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_create_info": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"create_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"force_import": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"original_zone_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"inherit": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"nameserver": {
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
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
							},
						},
						"tsig": {
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
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
							},
						},
						"restrict_ip": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
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
							},
						},
						"notify_address": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
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
							},
						},
					},
				},
			},
			"secondary_create_info": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification_email_address": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"primary_name_server": {
							Type:     schema.TypeSet,
							Required: true,
							MaxItems: 3,
							Elem: &schema.Resource{
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
							},
						},
					},
				},
			},
			"alias_create_info": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"original_zone_name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func newZone(rd *schema.ResourceData) (ultradns.Zone, error) {

	var zoneType string
	zone := ultradns.Zone{}
	properties := ultradns.ZoneProperties{}

	if val, ok := rd.GetOk("name"); ok {
		properties.Name = val.(string)
	}

	if val, ok := rd.GetOk("account_name"); ok {
		properties.AccountName = val.(string)
	}

	if val, ok := rd.GetOk("type"); ok {
		properties.Type = val.(string)
		zoneType = val.(string)
	}

	if val, ok := rd.GetOk("change_comment"); ok {
		zone.ChangeComment = val.(string)
	}

	switch zoneType {
	case "PRIMARY":
		zone.PrimaryCreateInfo = getPrimaryCreateInfo(rd)
	case "SECONDARY":
		zone.SecondaryCreateInfo = getSecondaryCreateInfo(rd)
	case "ALIAS":
		zone.AliasCreateInfo = getAliasCreateInfo(rd)
	}

	zone.Properties = &properties
	return zone, nil
}

func getPrimaryCreateInfo(rd *schema.ResourceData) *ultradns.PrimaryZone {
	primaryCreateInfo := &ultradns.PrimaryZone{}
	if val, ok := rd.GetOk("primary_create_info"); ok {
		data := val.(*schema.Set).List()[0].(map[string]interface{})

		if val, ok = data["create_type"]; ok {
			primaryCreateInfo.CreateType = val.(string)
		}

		if val, ok = data["force_import"]; ok {
			primaryCreateInfo.ForceImport = val.(bool)
		}
	}
	return primaryCreateInfo
}

func getSecondaryCreateInfo(rd *schema.ResourceData) *ultradns.SecondaryZone {
	secondaryCreateInfo := &ultradns.SecondaryZone{}
	return secondaryCreateInfo
}

func getAliasCreateInfo(rd *schema.ResourceData) *ultradns.AliasZone {
	aliasCreateInfo := &ultradns.AliasZone{}
	if val, ok := rd.GetOk("alias_create_info"); ok {
		data := val.(*schema.Set).List()[0].(map[string]interface{})
		aliasCreateInfo.OriginalZoneName = data["original_zone_name"].(string)
	}
	return aliasCreateInfo
}

func mapPrimaryZoneSchema(zr *ultradns.ZoneResponse, rd *schema.ResourceData) {

}

func resourceZoneCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*ultradns.Client)
	zone, er := newZone(rd)
	if er != nil {
		return diag.FromErr(er)
	}
	res, err := client.CreateZone(zone)

	if err != nil {
		return diag.FromErr(err)
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return diag.Errorf("Not a successfull response while creating the zone from rest api : returned response code - %v", res.StatusCode)
	}

	rd.SetId(zone.Properties.Name)
	resourceZoneRead(ctx, rd, meta)

	return diags
}

func resourceZoneRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*ultradns.Client)
	zoneId := rd.Id()

	res, zoneType, zoneResponse, er := client.ReadZone(zoneId)

	if er != nil {
		return diag.FromErr(er)
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return diag.Errorf("Not a successfull response while reading the zone from rest api : returned response code - %v", res.StatusCode)
	}

	switch zoneType {
	case "PRIMARY":
		mapPrimaryZoneSchema(zoneResponse, rd)
	case "SECONDARY":
		mapPrimaryZoneSchema(zoneResponse, rd)
	case "ALIAS":
		mapPrimaryZoneSchema(zoneResponse, rd)
	}

	return diags
}

func resourceZoneUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*ultradns.Client)
	zoneId := rd.Id()

	zone, er := newZone(rd)
	if er != nil {
		return diag.FromErr(er)
	}

	res, err := client.UpdateZone(zoneId, zone)

	if err != nil {
		return diag.FromErr(err)
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return diag.Errorf("Not a successfull response while reading the zone from rest api : returned response code - %v", res.StatusCode)
	}

	return diags
}

func resourceZoneDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*ultradns.Client)
	zoneId := rd.Id()

	_, err := client.DeleteZone(zoneId)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId("")
	return diags
}
