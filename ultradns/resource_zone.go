package ultradns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func ResourceZone() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceZoneCreate,
		ReadContext:   resourceZoneRead,
		UpdateContext: resourceZoneUpdate,
		DeleteContext: resourceZoneDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: zoneSchema(),
	}
}

func resourceZoneCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ultradns.Client)
	zone := newZone(rd)

	_, err := client.CreateZone(zone)

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(zone.Properties.Name)

	return resourceZoneRead(ctx, rd, meta)
}

func resourceZoneRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*ultradns.Client)
	zoneId := rd.Id()

	_, zoneResponse, err := client.ReadZone(zoneId)

	if err != nil {
		rd.SetId("")
		return nil
	}
	var zoneType string
	if zoneResponse.Properties != nil {
		zoneType = zoneResponse.Properties.Type
	}
	switch zoneType {
	case "PRIMARY":
		if er := mapPrimaryZoneSchema(zoneResponse, rd); er != nil {
			return diag.FromErr(er)
		}
	case "SECONDARY":
		if er := mapSecondaryZoneSchema(zoneResponse, rd); er != nil {
			return diag.FromErr(er)
		}
	case "ALIAS":
		if er := mapAliasZoneSchema(zoneResponse, rd); er != nil {
			return diag.FromErr(er)
		}
	}

	return diags
}

func resourceZoneUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ultradns.Client)
	zoneId := rd.Id()

	zone := newZone(rd)

	_, err := client.UpdateZone(zoneId, zone)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceZoneRead(ctx, rd, meta)
}

func resourceZoneDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*ultradns.Client)
	zoneId := rd.Id()

	_, err := client.DeleteZone(zoneId)
	if err != nil {
		rd.SetId("")
		return diag.FromErr(err)
	}

	rd.SetId("")
	return diags
}

func newZone(rd *schema.ResourceData) ultradns.Zone {

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
	return zone
}

func getPrimaryCreateInfo(rd *schema.ResourceData) *ultradns.PrimaryZone {
	primaryCreateInfo := &ultradns.PrimaryZone{}
	if val, ok := rd.GetOk("primary_create_info"); ok && val.(*schema.Set).Len() > 0 {
		createInfoData := val.(*schema.Set).List()[0].(map[string]interface{})

		if val, ok := createInfoData["create_type"]; ok {
			primaryCreateInfo.CreateType = val.(string)
		}

		if val, ok := createInfoData["force_import"]; ok {
			primaryCreateInfo.ForceImport = val.(bool)
		}

		if val, ok := createInfoData["original_zone_name"]; ok {
			primaryCreateInfo.OriginalZoneName = val.(string)
		}

		if val, ok := createInfoData["inherit"]; ok {
			primaryCreateInfo.Inherit = val.(string)
		}

		if val, ok := createInfoData["name_server"]; ok && val.(*schema.Set).Len() > 0 {
			nameServerData := val.(*schema.Set).List()[0].(map[string]interface{})
			nameServer := &ultradns.NameServerIp{}
			primaryCreateInfo.NameServer = nameServer

			if val, ok := nameServerData["ip"]; ok {
				nameServer.Ip = val.(string)
			}

			if val, ok := nameServerData["tsig_key"]; ok {
				nameServer.TsigKey = val.(string)
			}

			if val, ok := nameServerData["tsig_key_value"]; ok {
				nameServer.TsigKeyValue = val.(string)
			}

			if val, ok := nameServerData["tsig_algorithm"]; ok {
				nameServer.TsigAlgorithm = val.(string)
			}
		}

		if val, ok := createInfoData["tsig"]; ok && val.(*schema.Set).Len() > 0 {
			tsigData := val.(*schema.Set).List()[0].(map[string]interface{})
			tsig := &ultradns.Tsig{}
			primaryCreateInfo.Tsig = tsig

			if val, ok := tsigData["tsig_key_name"]; ok {
				tsig.TsigKeyName = val.(string)
			}

			if val, ok := tsigData["tsig_key_value"]; ok {
				tsig.TsigKeyValue = val.(string)
			}

			if val, ok := tsigData["tsig_algorithm"]; ok {
				tsig.TsigAlgorithm = val.(string)
			}

			if val, ok := tsigData["description"]; ok {
				tsig.Description = val.(string)
			}
		}

		if val, ok := createInfoData["restrict_ip"]; ok && val.(*schema.Set).Len() > 0 {
			restrictIpDataList := val.(*schema.Set).List()
			restrictIpList := make([]ultradns.RestrictIp, len(restrictIpDataList))
			primaryCreateInfo.RestrictIPList = &restrictIpList
			for i, d := range restrictIpDataList {
				restrictIpData := d.(map[string]interface{})
				restrictIp := ultradns.RestrictIp{}

				if val, ok := restrictIpData["start_ip"]; ok {
					restrictIp.StartIp = val.(string)
				}

				if val, ok := restrictIpData["end_ip"]; ok {
					restrictIp.EndIp = val.(string)
				}

				if val, ok := restrictIpData["cidr"]; ok {
					restrictIp.Cidr = val.(string)
				}

				if val, ok := restrictIpData["single_ip"]; ok {
					restrictIp.SingleIp = val.(string)
				}

				if val, ok := restrictIpData["comment"]; ok {
					restrictIp.Comment = val.(string)
				}
				restrictIpList[i] = restrictIp
			}
		}

		if val, ok := createInfoData["notify_addresses"]; ok && val.(*schema.Set).Len() > 0 {
			notifyAddressDataList := val.(*schema.Set).List()
			notifyAddressList := make([]ultradns.NotifyAddress, len(notifyAddressDataList))
			primaryCreateInfo.NotifyAddresses = &notifyAddressList

			for i, d := range notifyAddressDataList {
				notifyAddressData := d.(map[string]interface{})
				notifyAddress := ultradns.NotifyAddress{}

				if val, ok := notifyAddressData["notify_address"]; ok {
					notifyAddress.NotifyAddress = val.(string)
				}

				if val, ok := notifyAddressData["description"]; ok {
					notifyAddress.Description = val.(string)
				}

				notifyAddressList[i] = notifyAddress
			}
		}

	}
	return primaryCreateInfo
}

func getSecondaryCreateInfo(rd *schema.ResourceData) *ultradns.SecondaryZone {
	secondaryCreateInfo := &ultradns.SecondaryZone{}
	if val, ok := rd.GetOk("secondary_create_info"); ok && val.(*schema.Set).Len() > 0 {
		createInfoData := val.(*schema.Set).List()[0].(map[string]interface{})

		if val, ok := createInfoData["notification_email_address"]; ok {
			secondaryCreateInfo.NotificationEmailAddress = val.(string)
		}

		if val, ok := createInfoData["primary_name_server"]; ok && val.(*schema.Set).Len() > 0 {
			primaryNameServerList := val.(*schema.Set).List()
			length := len(primaryNameServerList)

			primaryNameServers := &ultradns.PrimaryNameServers{}
			nameServerIpList := &ultradns.NameServerIpList{}

			secondaryCreateInfo.PrimaryNameServers = primaryNameServers
			primaryNameServers.NameServerIpList = nameServerIpList

			for i := 0; i < length; i++ {
				primaryNameserver := primaryNameServerList[i].(map[string]interface{})
				nameServerIp := &ultradns.NameServerIp{}

				if val, ok := primaryNameserver["ip"]; ok {
					nameServerIp.Ip = val.(string)
				}

				if val, ok := primaryNameserver["tsig_key"]; ok {
					nameServerIp.TsigKey = val.(string)
				}

				if val, ok := primaryNameserver["tsig_key_value"]; ok {
					nameServerIp.TsigKeyValue = val.(string)
				}

				if val, ok := primaryNameserver["tsig_algorithm"]; ok {
					nameServerIp.TsigAlgorithm = val.(string)
				}

				switch i {
				case 0:
					nameServerIpList.NameServerIp1 = nameServerIp
				case 1:
					nameServerIpList.NameServerIp1 = nameServerIp
				case 2:
					nameServerIpList.NameServerIp1 = nameServerIp
				}
			}
		}
	}
	return secondaryCreateInfo
}

func getAliasCreateInfo(rd *schema.ResourceData) *ultradns.AliasZone {
	aliasCreateInfo := &ultradns.AliasZone{}
	if val, ok := rd.GetOk("alias_create_info"); ok && val.(*schema.Set).Len() > 0 {
		data := val.(*schema.Set).List()[0].(map[string]interface{})
		aliasCreateInfo.OriginalZoneName = data["original_zone_name"].(string)
	}
	return aliasCreateInfo
}
