package zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func flattenTsig(tsigData *zone.Tsig, rd *schema.ResourceData) error {
	set := &schema.Set{F: zeroIndexHash}
	tsig := make(map[string]interface{})
	tsig["tsig_key_name"] = tsigData.TsigKeyName
	tsig["tsig_key_value"] = tsigData.TsigKeyValue
	tsig["tsig_algorithm"] = tsigData.TsigAlgorithm
	tsig["description"] = tsigData.Description
	set.Add(tsig)

	if err := rd.Set("tsig", set); err != nil {
		return err
	}

	return nil
}

func flattenRestrictIPList(restrictIPListData []*zone.RestrictIP, rd *schema.ResourceData) error {
	set := &schema.Set{F: schema.HashResource(restrictIPResource())}

	for _, restrictIPData := range restrictIPListData {
		restrictIP := make(map[string]interface{})
		restrictIP["start_ip"] = restrictIPData.StartIP
		restrictIP["end_ip"] = restrictIPData.EndIP
		restrictIP["cidr"] = restrictIPData.Cidr
		restrictIP["single_ip"] = restrictIPData.SingleIP
		restrictIP["comment"] = restrictIPData.Comment
		set.Add(restrictIP)
	}

	if err := rd.Set("restrict_ip", set); err != nil {
		return err
	}

	return nil
}

func flattenNotifyAddresses(notifyAddressesData []*zone.NotifyAddress, rd *schema.ResourceData) error {
	set := &schema.Set{F: schema.HashResource(notifyAddressResource())}

	for _, notifyAddressData := range notifyAddressesData {
		notifyAddress := make(map[string]interface{})
		notifyAddress["notify_address"] = notifyAddressData.NotifyAddress
		notifyAddress["description"] = notifyAddressData.Description
		set.Add(notifyAddress)
	}

	if err := rd.Set("notify_addresses", set); err != nil {
		return err
	}

	return nil
}

func flattenRegistrarInfo(registrarInfoData *zone.RegistrarInfo, rd *schema.ResourceData) error {
	set := &schema.Set{F: zeroIndexHash}
	registrarInfo := make(map[string]interface{})
	registrarInfo["registrar"] = registrarInfoData.Registrar
	registrarInfo["who_is_expiration"] = registrarInfoData.WhoIsExpiration
	registrarInfo["name_servers"] = flattenRegistrarInfoNameServers(registrarInfoData.NameServers)
	set.Add(registrarInfo)

	if err := rd.Set("registrar_info", set); err != nil {
		return err
	}

	return nil
}

func flattenRegistrarInfoNameServers(nameServersList *zone.NameServersList) *schema.Set {
	set := &schema.Set{F: zeroIndexHash}
	registrarInfoNameServersList := make(map[string]interface{})
	registrarInfoNameServersList["ok"] = nameServersList.Ok
	registrarInfoNameServersList["unknown"] = nameServersList.Unknown
	registrarInfoNameServersList["missing"] = nameServersList.Missing
	registrarInfoNameServersList["incorrect"] = nameServersList.Incorrect
	set.Add(registrarInfoNameServersList)

	return set
}

func flattenTransferStatusDetails(transferDetailsData *zone.TransferStatusDetails, rd *schema.ResourceData) error {
	set := &schema.Set{F: zeroIndexHash}
	transferDetails := make(map[string]interface{})
	transferDetails["last_refresh"] = transferDetailsData.LastRefresh
	transferDetails["next_refresh"] = transferDetailsData.NextRefresh
	transferDetails["last_refresh_status"] = transferDetailsData.LastRefreshStatus
	transferDetails["last_refresh_status_message"] = transferDetailsData.LastRefreshStatusMessage
	set.Add(transferDetails)

	if err := rd.Set("transfer_status_details", set); err != nil {
		return err
	}

	return nil
}

func flattenPrimaryNameServers(nameServerIPListData *zone.NameServerIPList, rd *schema.ResourceData) error {
	if nameServerIPListData.NameServerIP1 != nil {
		if err := rd.Set("primary_name_server_1", flattenNameServer(nameServerIPListData.NameServerIP1)); err != nil {
			return err
		}
	}

	if nameServerIPListData.NameServerIP2 != nil {
		if err := rd.Set("primary_name_server_2", flattenNameServer(nameServerIPListData.NameServerIP2)); err != nil {
			return err
		}
	}

	if nameServerIPListData.NameServerIP3 != nil {
		if err := rd.Set("primary_name_server_3", flattenNameServer(nameServerIPListData.NameServerIP3)); err != nil {
			return err
		}
	}

	return nil
}

func flattenNameServer(nameServerData *zone.NameServer) *schema.Set {
	set := &schema.Set{F: zeroIndexHash}
	nameServer := make(map[string]interface{})
	nameServer["ip"] = nameServerData.IP
	nameServer["tsig_key"] = nameServerData.TsigKey
	nameServer["tsig_key_value"] = nameServerData.TsigKeyValue
	nameServer["tsig_algorithm"] = nameServerData.TsigAlgorithm
	set.Add(nameServer)

	return set
}

func zeroIndexHash(i interface{}) int {
	return 0
}
