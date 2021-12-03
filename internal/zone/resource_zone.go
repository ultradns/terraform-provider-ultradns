package zone

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

		Schema: resourceZoneSchema(),
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
	zoneID := rd.Id()

	_, zr, err := client.ReadZone(zoneID)

	if err != nil {
		rd.SetId("")
		return nil
	}

	if zr.Properties != nil {

		if err := rd.Set("name", zr.Properties.Name); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("account_name", zr.Properties.AccountName); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("type", zr.Properties.Type); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("dnssec_status", zr.Properties.DnsSecStatus); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("resource_record_count", zr.Properties.ResourceRecordCount); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("last_modified_time", zr.Properties.LastModifiedDateTime); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("status", zr.Properties.Status); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("owner", zr.Properties.Owner); err != nil {
			return diag.FromErr(err)
		}

		if zr.RegistrarInfo != nil {
			if err := rd.Set("registrar_info", flattenRegistrarInfo(zr.RegistrarInfo)); err != nil {
				return diag.FromErr(err)
			}
		}

		if zr.TransferStatusDetails != nil {
			if err := rd.Set("transfer_status_details", flattenTransferStatusDetails(zr.TransferStatusDetails)); err != nil {
				return diag.FromErr(err)
			}
		}

		switch zr.Properties.Type {
		case "PRIMARY":
			if err := rd.Set("primary_create_info", flattenPrimaryZone(zr, rd)); err != nil {
				return diag.FromErr(err)
			}
		case "SECONDARY":
			if err := rd.Set("secondary_create_info", flattenSecondaryZone(zr, rd)); err != nil {
				return diag.FromErr(err)
			}
		case "ALIAS":
			if err := rd.Set("alias_create_info", flattenAliasZone(zr, rd)); err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return diags
}

func resourceZoneUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ultradns.Client)
	zoneID := rd.Id()

	zone := newZone(rd)

	_, err := client.UpdateZone(zoneID, zone)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceZoneRead(ctx, rd, meta)
}

func resourceZoneDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*ultradns.Client)
	zoneID := rd.Id()

	_, err := client.DeleteZone(zoneID)
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
	properties := &ultradns.ZoneProperties{}

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

	zone.Properties = properties
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
			restrictIPDataList := val.(*schema.Set).List()
			restrictIPList := make([]*ultradns.RestrictIp, len(restrictIPDataList))
			primaryCreateInfo.RestrictIPList = restrictIPList
			for i, d := range restrictIPDataList {
				restrictIPData := d.(map[string]interface{})
				restrictIP := ultradns.RestrictIp{}

				if val, ok := restrictIPData["start_ip"]; ok {
					restrictIP.StartIp = val.(string)
				}

				if val, ok := restrictIPData["end_ip"]; ok {
					restrictIP.EndIp = val.(string)
				}

				if val, ok := restrictIPData["cidr"]; ok {
					restrictIP.Cidr = val.(string)
				}

				if val, ok := restrictIPData["single_ip"]; ok {
					restrictIP.SingleIp = val.(string)
				}

				if val, ok := restrictIPData["comment"]; ok {
					restrictIP.Comment = val.(string)
				}
				restrictIPList[i] = &restrictIP
			}
		}

		if val, ok := createInfoData["notify_addresses"]; ok && val.(*schema.Set).Len() > 0 {
			notifyAddressDataList := val.(*schema.Set).List()
			notifyAddressList := make([]*ultradns.NotifyAddress, len(notifyAddressDataList))
			primaryCreateInfo.NotifyAddresses = notifyAddressList

			for i, d := range notifyAddressDataList {
				notifyAddressData := d.(map[string]interface{})
				notifyAddress := ultradns.NotifyAddress{}

				if val, ok := notifyAddressData["notify_address"]; ok {
					notifyAddress.NotifyAddress = val.(string)
				}

				if val, ok := notifyAddressData["description"]; ok {
					notifyAddress.Description = val.(string)
				}

				notifyAddressList[i] = &notifyAddress
			}
		}

	}
	return primaryCreateInfo
}

func getSecondaryCreateInfo(rd *schema.ResourceData) *ultradns.SecondaryZone {
	nameServerIPList := &ultradns.NameServerIpList{}
	primaryNameServers := &ultradns.PrimaryNameServers{NameServerIpList: nameServerIPList}
	secondaryCreateInfo := &ultradns.SecondaryZone{PrimaryNameServers: primaryNameServers}

	if val, ok := rd.GetOk("secondary_create_info"); ok && val.(*schema.Set).Len() > 0 {
		createInfoData := val.(*schema.Set).List()[0].(map[string]interface{})

		if val, ok := createInfoData["notification_email_address"]; ok {
			secondaryCreateInfo.NotificationEmailAddress = val.(string)
		}

		if val, ok := createInfoData["primary_name_server_1"]; ok && val.(*schema.Set).Len() > 0 {
			nameServerData := val.(*schema.Set).List()[0].(map[string]interface{})
			nameServer := &ultradns.NameServerIp{}
			secondaryCreateInfo.PrimaryNameServers.NameServerIpList.NameServerIp1 = nameServer

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

		if val, ok := createInfoData["primary_name_server_2"]; ok && val.(*schema.Set).Len() > 0 {
			nameServerData := val.(*schema.Set).List()[0].(map[string]interface{})
			nameServer := &ultradns.NameServerIp{}
			secondaryCreateInfo.PrimaryNameServers.NameServerIpList.NameServerIp2 = nameServer

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

		if val, ok := createInfoData["primary_name_server_3"]; ok && val.(*schema.Set).Len() > 0 {
			nameServerData := val.(*schema.Set).List()[0].(map[string]interface{})
			nameServer := &ultradns.NameServerIp{}
			secondaryCreateInfo.PrimaryNameServers.NameServerIpList.NameServerIp3 = nameServer

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
