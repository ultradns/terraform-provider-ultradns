package zone

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func getZoneListURLParameters(rd *schema.ResourceData) string {
	param := "?"

	if val, ok := rd.GetOk("query"); ok {
		param = param + "&q=" + val.(string)
	}

	if val, ok := rd.GetOk("sort"); ok {
		param = param + "&sort=" + val.(string)
	}

	if val, ok := rd.GetOk("reverse"); ok {
		param = param + "&reverse=" + strconv.FormatBool(val.(bool))
	}

	if val, ok := rd.GetOk("limit"); ok {
		param = param + "&limit=" + strconv.Itoa(val.(int))
	}

	if val, ok := rd.GetOk("offset"); ok {
		param = param + "&offset=" + strconv.Itoa(val.(int))
	}

	if val, ok := rd.GetOk("cursor"); ok {
		param = param + "&cursor=" + val.(string)
	}

	return param
}

func flattenZones(zlr []*ultradns.ZoneResponse) []map[string]interface{} {
	var zones []map[string]interface{}

	for _, zone := range zlr {
		data := make(map[string]interface{})

		if zone.Properties != nil {
			data["name"] = zone.Properties.Name
			data["account_name"] = zone.Properties.AccountName
			data["type"] = zone.Properties.Type
			data["dnssec_status"] = zone.Properties.DnsSecStatus
			data["status"] = zone.Properties.Status
			data["owner"] = zone.Properties.Owner
			data["resource_record_count"] = zone.Properties.ResourceRecordCount
			data["last_modified_time"] = zone.Properties.LastModifiedDateTime
		}

		//data["inherit"] =
		data["notification_email_address"] = zone.NotificationEmailAddress
		data["original_zone_name"] = zone.OriginalZoneName

		if zone.Tsig != nil {
			data["tsig"] = flattenTsig(zone.Tsig)
		}

		if zone.RestrictIPList != nil {
			data["restrict_ip"] = flattenRestrictIP(zone.RestrictIPList)
		}

		if zone.NotifyAddresses != nil {
			data["notify_addresses"] = flattenNotifyAddresses(zone.NotifyAddresses)
		}

		if zone.RegistrarInfo != nil {
			data["registrar_info"] = flattenRegistrarInfo(zone.RegistrarInfo)
		}

		if zone.PrimaryNameServers != nil && zone.PrimaryNameServers.NameServerIpList != nil {
			if zone.PrimaryNameServers.NameServerIpList.NameServerIp1 != nil {
				data["primary_name_server_1"] = flattenNameServer(zone.PrimaryNameServers.NameServerIpList.NameServerIp1)
			}

			if zone.PrimaryNameServers.NameServerIpList.NameServerIp2 != nil {
				data["primary_name_server_2"] = flattenNameServer(zone.PrimaryNameServers.NameServerIpList.NameServerIp2)
			}

			if zone.PrimaryNameServers.NameServerIpList.NameServerIp3 != nil {
				data["primary_name_server_3"] = flattenNameServer(zone.PrimaryNameServers.NameServerIpList.NameServerIp3)
			}
		}

		if zone.TransferStatusDetails != nil {
			data["transfer_status_details"] = flattenTransferStatusDetails(zone.TransferStatusDetails)
		}

		zones = append(zones, data)
	}
	return zones
}

func flattenTsig(t *ultradns.Tsig) *schema.Set {
	set := &schema.Set{F: zeroIndexHash}
	tsig := make(map[string]interface{})

	tsig["tsig_key_name"] = t.TsigKeyName
	tsig["tsig_key_value"] = t.TsigKeyValue
	tsig["tsig_algorithm"] = t.TsigAlgorithm
	tsig["description"] = t.Description

	set.Add(tsig)
	return set
}

func flattenRestrictIP(ri []*ultradns.RestrictIp) *schema.Set {
	set := &schema.Set{F: zeroIndexHash}

	for _, restrictIPData := range ri {
		restrictIP := make(map[string]interface{})

		restrictIP["start_ip"] = restrictIPData.StartIp
		restrictIP["end_ip"] = restrictIPData.EndIp
		restrictIP["cidr"] = restrictIPData.Cidr
		restrictIP["single_ip"] = restrictIPData.SingleIp
		restrictIP["comment"] = restrictIPData.Comment

		set.Add(restrictIP)
	}
	return set
}

func flattenNotifyAddresses(na []*ultradns.NotifyAddress) *schema.Set {
	set := &schema.Set{F: zeroIndexHash}

	for _, notifyAddressData := range na {
		notifyAddress := make(map[string]interface{})

		notifyAddress["notify_address"] = notifyAddressData.NotifyAddress
		notifyAddress["description"] = notifyAddressData.Description

		set.Add(notifyAddress)
	}
	return set
}

func flattenRegistrarInfo(ri *ultradns.RegistrarInfo) *schema.Set {
	set := &schema.Set{F: zeroIndexHash}

	registrarInfo := make(map[string]interface{})

	registrarInfo["registrar"] = ri.Registrar
	registrarInfo["who_is_expiration"] = ri.WhoIsExpiration
	registrarInfo["name_servers"] = flattenRegistrarInfoNameServer(ri.NameServers)

	set.Add(registrarInfo)
	return set
}

func flattenRegistrarInfoNameServer(nsl *ultradns.NameServersList) *schema.Set {
	set := &schema.Set{F: zeroIndexHash}

	RegistrarInfoNameServersList := make(map[string]interface{})

	RegistrarInfoNameServersList["ok"] = nsl.Ok
	RegistrarInfoNameServersList["unknown"] = nsl.Unknown
	RegistrarInfoNameServersList["missing"] = nsl.Missing
	RegistrarInfoNameServersList["incorrect"] = nsl.Incorrect

	set.Add(RegistrarInfoNameServersList)
	return set
}

func flattenNameServer(ns *ultradns.NameServerIp) *schema.Set {
	set := &schema.Set{F: zeroIndexHash}

	nameServer := make(map[string]interface{})

	nameServer["ip"] = ns.Ip
	nameServer["tsig_key"] = ns.TsigKey
	nameServer["tsig_key_value"] = ns.TsigKeyValue
	nameServer["tsig_algorithm"] = ns.TsigAlgorithm

	set.Add(nameServer)
	return set
}

func flattenTransferStatusDetails(tsd *ultradns.TransferStatusDetails) *schema.Set {
	set := &schema.Set{F: zeroIndexHash}
	transferDetails := make(map[string]interface{})

	transferDetails["last_refresh"] = tsd.LastRefresh
	transferDetails["next_refresh"] = tsd.NextRefresh
	transferDetails["last_refresh_status"] = tsd.LastRefreshStatus
	transferDetails["last_refresh_status_message"] = tsd.LastRefreshStatusMessage

	set.Add(transferDetails)
	return set
}

func zeroIndexHash(i interface{}) int {
	return 0
}
