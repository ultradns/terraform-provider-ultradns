package zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func flattenZoneProperties(zoneResponse *zone.Response, rd *schema.ResourceData) error {
	if err := rd.Set("name", zoneResponse.Properties.Name); err != nil {
		return err
	}

	if err := rd.Set("account_name", zoneResponse.Properties.AccountName); err != nil {
		return err
	}

	if err := rd.Set("type", zoneResponse.Properties.Type); err != nil {
		return err
	}

	if err := rd.Set("dnssec_status", zoneResponse.Properties.DNSSecStatus); err != nil {
		return err
	}

	if err := rd.Set("resource_record_count", zoneResponse.Properties.ResourceRecordCount); err != nil {
		return err
	}

	if err := rd.Set("last_modified_time", zoneResponse.Properties.LastModifiedDateTime); err != nil {
		return err
	}

	if err := rd.Set("status", zoneResponse.Properties.Status); err != nil {
		return err
	}

	if err := rd.Set("owner", zoneResponse.Properties.Owner); err != nil {
		return err
	}

	return nil
}

func flattenPrimaryZone(zoneResponse *zone.Response, rd *schema.ResourceData) error {
	list := make([]interface{}, 1)
	primaryCreateInfo := make(map[string]interface{})

	if val, ok := rd.GetOk("primary_create_info"); ok {
		primaryCreateInfo = val.([]interface{})[0].(map[string]interface{})
	}

	primaryCreateInfo["inherit"] = zoneResponse.Inherit
	primaryCreateInfo["tsig"] = getTsigList(zoneResponse.Tsig)
	primaryCreateInfo["restrict_ip"] = getRestrictIPListSet(zoneResponse.RestrictIPList)
	primaryCreateInfo["notify_addresses"] = getNotifyAddressesSet(zoneResponse.NotifyAddresses)

	list[0] = primaryCreateInfo

	if err := rd.Set("primary_create_info", list); err != nil {
		return err
	}

	if err := rd.Set("registrar_info", getRegistrarInfoList(zoneResponse.RegistrarInfo)); err != nil {
		return err
	}

	return nil
}

func flattenSecondaryZone(zoneResponse *zone.Response, rd *schema.ResourceData) error {
	list := make([]interface{}, 1)
	secondaryCreateInfo := make(map[string]interface{})

	if zoneResponse.NotificationEmailAddress != "" {
		secondaryCreateInfo["notification_email_address"] = zoneResponse.NotificationEmailAddress
	}

	if zoneResponse.PrimaryNameServers != nil && zoneResponse.PrimaryNameServers.NameServerIPList != nil {
		secondaryCreateInfo["primary_name_server_1"] = getNameServerList(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP1)
		secondaryCreateInfo["primary_name_server_2"] = getNameServerList(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP2)
		secondaryCreateInfo["primary_name_server_3"] = getNameServerList(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP3)
	}

	list[0] = secondaryCreateInfo

	if err := rd.Set("secondary_create_info", list); err != nil {
		return err
	}

	if err := rd.Set("transfer_status_details", getTransferStatusDetailsList(zoneResponse.TransferStatusDetails)); err != nil {
		return err
	}

	return nil
}

func flattenAliasZone(zoneResponse *zone.Response, rd *schema.ResourceData) error {
	list := make([]interface{}, 1)
	aliasCreateInfo := make(map[string]interface{})
	aliasCreateInfo["original_zone_name"] = zoneResponse.OriginalZoneName
	list[0] = aliasCreateInfo

	if err := rd.Set("alias_create_info", list); err != nil {
		return err
	}

	return nil
}

func getNameServerList(nameServerData *zone.NameServer) []interface{} {
	var list []interface{}

	if nameServerData != nil {
		list = make([]interface{}, 1)
		nameServer := make(map[string]interface{})
		nameServer["ip"] = nameServerData.IP
		nameServer["tsig_key"] = nameServerData.TsigKey
		nameServer["tsig_key_value"] = nameServerData.TsigKeyValue
		nameServer["tsig_algorithm"] = nameServerData.TsigAlgorithm
		list[0] = nameServer
	}

	return list
}

func getTsigList(tsigData *zone.Tsig) []interface{} {
	var list []interface{}

	if tsigData != nil {
		list = make([]interface{}, 1)
		tsig := make(map[string]interface{})
		tsig["tsig_key_name"] = tsigData.TsigKeyName
		tsig["tsig_key_value"] = tsigData.TsigKeyValue
		tsig["tsig_algorithm"] = tsigData.TsigAlgorithm
		tsig["description"] = tsigData.Description
		list[0] = tsig
	}

	return list
}

func getRestrictIPListSet(restrictIPDataList []*zone.RestrictIP) *schema.Set {
	set := &schema.Set{F: schema.HashResource(restrictIPResource())}

	for _, restrictIPData := range restrictIPDataList {
		restrictIP := make(map[string]interface{})
		restrictIP["start_ip"] = restrictIPData.StartIP
		restrictIP["end_ip"] = restrictIPData.EndIP
		restrictIP["cidr"] = restrictIPData.Cidr
		restrictIP["single_ip"] = restrictIPData.SingleIP
		restrictIP["comment"] = restrictIPData.Comment
		set.Add(restrictIP)
	}

	return set
}

func getNotifyAddressesSet(notifyAddressDataList []*zone.NotifyAddress) *schema.Set {
	set := &schema.Set{F: schema.HashResource(notifyAddressResource())}

	for _, notifyAddressData := range notifyAddressDataList {
		notifyAddress := make(map[string]interface{})
		notifyAddress["notify_address"] = notifyAddressData.NotifyAddress
		notifyAddress["description"] = notifyAddressData.Description
		set.Add(notifyAddress)
	}

	return set
}

func getRegistrarInfoList(registrarInfoData *zone.RegistrarInfo) []interface{} {
	var list []interface{}

	if registrarInfoData != nil {
		list = make([]interface{}, 1)
		registrarInfo := make(map[string]interface{})
		registrarInfo["registrar"] = registrarInfoData.Registrar
		registrarInfo["who_is_expiration"] = registrarInfoData.WhoIsExpiration
		registrarInfo["name_servers"] = getRegistrarInfoNameServersList(registrarInfoData.NameServers)
		list[0] = registrarInfo
	}

	return list
}

func getRegistrarInfoNameServersList(nameServersList *zone.NameServersList) []interface{} {
	var list []interface{}

	if nameServersList != nil {
		list = make([]interface{}, 1)
		registrarInfoNameServersList := make(map[string]interface{})
		registrarInfoNameServersList["ok"] = nameServersList.Ok
		registrarInfoNameServersList["unknown"] = nameServersList.Unknown
		registrarInfoNameServersList["missing"] = nameServersList.Missing
		registrarInfoNameServersList["incorrect"] = nameServersList.Incorrect
		list[0] = registrarInfoNameServersList
	}

	return list
}

func getTransferStatusDetailsList(transferDetailsData *zone.TransferStatusDetails) []interface{} {
	var list []interface{}

	if transferDetailsData != nil {
		list = make([]interface{}, 1)
		transferDetails := make(map[string]interface{})
		transferDetails["last_refresh"] = transferDetailsData.LastRefresh
		transferDetails["next_refresh"] = transferDetailsData.NextRefresh
		transferDetails["last_refresh_status"] = transferDetailsData.LastRefreshStatus
		transferDetails["last_refresh_status_message"] = transferDetailsData.LastRefreshStatusMessage
		list[0] = transferDetails
	}

	return list
}
