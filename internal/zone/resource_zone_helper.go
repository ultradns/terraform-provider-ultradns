package zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func flattenZoneProperties(zoneResponse *zone.Response, rd *schema.ResourceData) error {
	currentSchemaZoneName := rd.Get("name").(string)

	if helper.GetZoneFQDN(currentSchemaZoneName) != zoneResponse.Properties.Name {
		if err := rd.Set("name", zoneResponse.Properties.Name); err != nil {
			return err
		}
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
	set := &schema.Set{F: schema.HashResource(primaryZoneCreateInfoResource())}
	primaryCreateInfo := make(map[string]interface{})

	if val, ok := rd.GetOk("primary_create_info"); ok && val.(*schema.Set).Len() > 0 {
		primaryCreateInfo = val.(*schema.Set).List()[0].(map[string]interface{})
	}

	if zoneResponse.Tsig != nil {
		primaryCreateInfo["tsig"] = getTsigSet(zoneResponse.Tsig)
	}

	if len(zoneResponse.RestrictIPList) > 0 {
		primaryCreateInfo["restrict_ip"] = getRestrictIPListSet(zoneResponse.RestrictIPList)
	}

	if len(zoneResponse.NotifyAddresses) > 0 {
		primaryCreateInfo["notify_addresses"] = getNotifyAddressesSet(zoneResponse.NotifyAddresses)
	}

	set.Add(primaryCreateInfo)

	if err := rd.Set("primary_create_info", set); err != nil {
		return err
	}

	return nil
}

func flattenSecondaryZone(zoneResponse *zone.Response, rd *schema.ResourceData) error {
	set := &schema.Set{F: schema.HashResource(secondaryZoneCreateInfoResource())}
	secondaryCreateInfo := make(map[string]interface{})

	if val, ok := rd.GetOk("secondary_create_info"); ok && val.(*schema.Set).Len() > 0 {
		secondaryCreateInfo = val.(*schema.Set).List()[0].(map[string]interface{})
	}

	if zoneResponse.NotificationEmailAddress != "" {
		secondaryCreateInfo["notification_email_address"] = zoneResponse.NotificationEmailAddress
	}

	if zoneResponse.PrimaryNameServers != nil && zoneResponse.PrimaryNameServers.NameServerIPList != nil {
		if zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP1 != nil {
			secondaryCreateInfo["primary_name_server_1"] = getNameServerSet(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP1)
		}

		if zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP2 != nil {
			secondaryCreateInfo["primary_name_server_2"] = getNameServerSet(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP2)
		}

		if zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP3 != nil {
			secondaryCreateInfo["primary_name_server_3"] = getNameServerSet(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP3)
		}
	}

	set.Add(secondaryCreateInfo)

	if err := rd.Set("secondary_create_info", set); err != nil {
		return err
	}

	return nil
}

func flattenAliasZone(zoneResponse *zone.Response, rd *schema.ResourceData) error {
	set := &schema.Set{F: schema.HashResource(aliasZoneCreateInfoResource())}
	aliasCreateInfo := make(map[string]interface{})

	if val, ok := rd.GetOk("alias_create_info"); ok && val.(*schema.Set).Len() > 0 {
		aliasCreateInfo = val.(*schema.Set).List()[0].(map[string]interface{})
	}

	currentSchemaOriginalZoneName, ok := aliasCreateInfo["original_zone_name"].(string)

	if !ok || helper.GetZoneFQDN(currentSchemaOriginalZoneName) != zoneResponse.OriginalZoneName {
		aliasCreateInfo["original_zone_name"] = zoneResponse.OriginalZoneName
	}

	set.Add(aliasCreateInfo)

	if err := rd.Set("alias_create_info", set); err != nil {
		return err
	}

	return nil
}

func getNameServerSet(nameServerData *zone.NameServer) *schema.Set {
	set := &schema.Set{F: schema.HashResource(nameServerResource())}
	nameserver := make(map[string]interface{})
	nameserver["ip"] = nameServerData.IP
	nameserver["tsig_key"] = nameServerData.TsigKey
	nameserver["tsig_key_value"] = nameServerData.TsigKeyValue
	nameserver["tsig_algorithm"] = nameServerData.TsigAlgorithm
	set.Add(nameserver)

	return set
}

func getTsigSet(tsigData *zone.Tsig) *schema.Set {
	set := &schema.Set{F: schema.HashResource(tsigResource())}
	tsig := make(map[string]interface{})
	tsig["tsig_key_name"] = tsigData.TsigKeyName
	tsig["tsig_key_value"] = tsigData.TsigKeyValue
	tsig["tsig_algorithm"] = tsigData.TsigAlgorithm
	tsig["description"] = tsigData.Description
	set.Add(tsig)

	return set
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
