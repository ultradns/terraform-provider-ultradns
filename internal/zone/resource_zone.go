package zone

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	provhelper "github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
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
	services := meta.(*service.Service)
	zoneData := newZone(rd)

	_, err := services.ZoneService.CreateZone(zoneData)
	if err != nil {
		return diag.FromErr(err)
	}

	id := helper.GetZoneFQDN(zoneData.Properties.Name)
	rd.SetId(id)

	return resourceZoneRead(ctx, rd, meta)
}

func resourceZoneRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	zoneID := rd.Id()

	res, zoneResponse, err := services.ZoneService.ReadZone(zoneID)
	if err != nil && res != nil && res.Status == provhelper.RESOURCE_NOT_FOUND {
		rd.SetId("")
		tflog.Debug(ctx, err.Error())
		return nil
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if zoneResponse.Properties != nil {
		if err := flattenZoneProperties(zoneResponse, rd); err != nil {
			return diag.FromErr(err)
		}

		switch zoneResponse.Properties.Type {
		case zone.Primary:
			if err := flattenPrimaryZone(zoneResponse, rd); err != nil {
				return diag.FromErr(err)
			}
		case zone.Secondary:
			if err := flattenSecondaryZone(zoneResponse, rd); err != nil {
				return diag.FromErr(err)
			}
		case zone.Alias:
			if err := flattenAliasZone(zoneResponse, rd); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return diags
}

func resourceZoneUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	zoneName := rd.Id()

	zoneData := newZone(rd)

	_, err := services.ZoneService.UpdateZone(zoneName, zoneData)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceZoneRead(ctx, rd, meta)
}

func resourceZoneDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	zoneName := rd.Id()

	_, err := services.ZoneService.DeleteZone(zoneName)
	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	rd.SetId("")

	return diags
}

func newZone(rd *schema.ResourceData) *zone.Zone {
	zoneData := &zone.Zone{}
	properties := getZoneProperties(rd)

	if val, ok := rd.GetOk("change_comment"); ok {
		zoneData.ChangeComment = val.(string)
	}

	switch properties.Type {
	case zone.Primary:
		zoneData.PrimaryCreateInfo = getPrimaryCreateInfo(rd)
	case zone.Secondary:
		zoneData.SecondaryCreateInfo = getSecondaryCreateInfo(rd)
	case zone.Alias:
		zoneData.AliasCreateInfo = getAliasCreateInfo(rd)
	}

	zoneData.Properties = properties

	return zoneData
}

func getZoneProperties(rd *schema.ResourceData) *zone.Properties {
	properties := &zone.Properties{}

	if val, ok := rd.GetOk("name"); ok {
		properties.Name = strings.ToLower(val.(string))
	}

	if val, ok := rd.GetOk("account_name"); ok {
		properties.AccountName = val.(string)
	}

	if val, ok := rd.GetOk("type"); ok {
		properties.Type = val.(string)
	}

	return properties
}

func getPrimaryCreateInfo(rd *schema.ResourceData) *zone.PrimaryZone {
	primaryCreateInfo := &zone.PrimaryZone{}

	if val, ok := rd.GetOk("primary_create_info"); ok && len(val.([]interface{})) > 0 {
		createInfoData := val.([]interface{})[0].(map[string]interface{})

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

		if val, ok := createInfoData["name_server"]; ok && len(val.([]interface{})) > 0 {
			nameServerData := val.([]interface{})[0].(map[string]interface{})
			primaryCreateInfo.NameServer = getNameServer(nameServerData)
		}

		if val, ok := createInfoData["tsig"]; ok && len(val.([]interface{})) > 0 {
			tsigData := val.([]interface{})[0].(map[string]interface{})
			primaryCreateInfo.Tsig = getTsig(tsigData)
		}

		if val, ok := createInfoData["restrict_ip"]; ok {
			restrictIPDataList := val.(*schema.Set).List()
			primaryCreateInfo.RestrictIPList = getRestrictIPList(restrictIPDataList)
		}

		if val, ok := createInfoData["notify_addresses"]; ok {
			notifyAddressDataList := val.(*schema.Set).List()
			primaryCreateInfo.NotifyAddresses = getNotifyAddresses(notifyAddressDataList)
		}
	}

	return primaryCreateInfo
}

func getSecondaryCreateInfo(rd *schema.ResourceData) *zone.SecondaryZone {
	nameServerIPList := &zone.NameServerIPList{}
	primaryNameServers := &zone.PrimaryNameServers{NameServerIPList: nameServerIPList}
	secondaryCreateInfo := &zone.SecondaryZone{PrimaryNameServers: primaryNameServers}

	if val, ok := rd.GetOk("secondary_create_info"); ok && len(val.([]interface{})) > 0 {
		createInfoData := val.([]interface{})[0].(map[string]interface{})

		if val, ok := createInfoData["notification_email_address"]; ok {
			secondaryCreateInfo.NotificationEmailAddress = val.(string)
		}

		if val, ok := createInfoData["primary_name_server_1"]; ok && len(val.([]interface{})) > 0 {
			nameServerData := val.([]interface{})[0].(map[string]interface{})
			secondaryCreateInfo.PrimaryNameServers.NameServerIPList.NameServerIP1 = getNameServer(nameServerData)
		}

		if val, ok := createInfoData["primary_name_server_2"]; ok && len(val.([]interface{})) > 0 {
			nameServerData := val.([]interface{})[0].(map[string]interface{})
			secondaryCreateInfo.PrimaryNameServers.NameServerIPList.NameServerIP2 = getNameServer(nameServerData)
		}

		if val, ok := createInfoData["primary_name_server_3"]; ok && len(val.([]interface{})) > 0 {
			nameServerData := val.([]interface{})[0].(map[string]interface{})
			secondaryCreateInfo.PrimaryNameServers.NameServerIPList.NameServerIP3 = getNameServer(nameServerData)
		}
	}

	return secondaryCreateInfo
}

func getAliasCreateInfo(rd *schema.ResourceData) *zone.AliasZone {
	aliasCreateInfo := &zone.AliasZone{}

	if val, ok := rd.GetOk("alias_create_info"); ok && len(val.([]interface{})) > 0 {
		createInfoData := val.([]interface{})[0].(map[string]interface{})
		aliasCreateInfo.OriginalZoneName = createInfoData["original_zone_name"].(string)
	}

	return aliasCreateInfo
}

func getNameServer(nameServerData map[string]interface{}) *zone.NameServer {
	nameServer := &zone.NameServer{}

	if val, ok := nameServerData["ip"]; ok {
		nameServer.IP = val.(string)
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

	return nameServer
}

func getTsig(tsigData map[string]interface{}) *zone.Tsig {
	tsig := &zone.Tsig{}

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

	return tsig
}

func getRestrictIPList(restrictIPDataList []interface{}) []*zone.RestrictIP {
	restrictIPList := make([]*zone.RestrictIP, len(restrictIPDataList))

	for i, d := range restrictIPDataList {
		restrictIPData := d.(map[string]interface{})
		restrictIPList[i] = getRestrictIP(restrictIPData)
	}

	return restrictIPList
}

func getRestrictIP(restrictIPData map[string]interface{}) *zone.RestrictIP {
	restrictIP := &zone.RestrictIP{}

	if val, ok := restrictIPData["start_ip"]; ok {
		restrictIP.StartIP = val.(string)
	}

	if val, ok := restrictIPData["end_ip"]; ok {
		restrictIP.EndIP = val.(string)
	}

	if val, ok := restrictIPData["cidr"]; ok {
		restrictIP.Cidr = val.(string)
	}

	if val, ok := restrictIPData["single_ip"]; ok {
		restrictIP.SingleIP = val.(string)
	}

	if val, ok := restrictIPData["comment"]; ok {
		restrictIP.Comment = val.(string)
	}

	return restrictIP
}

func getNotifyAddresses(notifyAddressDataList []interface{}) []*zone.NotifyAddress {
	notifyAddressList := make([]*zone.NotifyAddress, len(notifyAddressDataList))

	for i, d := range notifyAddressDataList {
		notifyAddressData := d.(map[string]interface{})
		notifyAddressList[i] = getNotifyAddress(notifyAddressData)
	}

	return notifyAddressList
}

func getNotifyAddress(notifyAddressData map[string]interface{}) *zone.NotifyAddress {
	notifyAddress := &zone.NotifyAddress{}

	if val, ok := notifyAddressData["notify_address"]; ok {
		notifyAddress.NotifyAddress = val.(string)
	}

	if val, ok := notifyAddressData["description"]; ok {
		notifyAddress.Description = val.(string)
	}

	return notifyAddress
}
