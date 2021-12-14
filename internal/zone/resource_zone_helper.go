package zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func flattenPrimaryZone(zr *zone.Response, rd *schema.ResourceData) *schema.Set {
	set := &schema.Set{
		F: schema.HashResource(primaryZoneCreateInfoResource()),
	}

	primaryCreateInfo := make(map[string]interface{})

	if val, ok := rd.GetOk("primary_create_info"); ok && val.(*schema.Set).Len() > 0 {
		primaryCreateInfo = val.(*schema.Set).List()[0].(map[string]interface{})
	}

	if len(zr.NotifyAddresses) > 0 {
		s := &schema.Set{
			F: schema.HashResource(notifyAddressResource()),
		}

		for _, notifyAddressData := range zr.NotifyAddresses {
			notifyAddress := make(map[string]interface{})

			notifyAddress["notify_address"] = notifyAddressData.NotifyAddress
			notifyAddress["description"] = notifyAddressData.Description

			s.Add(notifyAddress)
		}

		primaryCreateInfo["notify_addresses"] = s
	}

	if len(zr.RestrictIPList) > 0 {
		s := &schema.Set{
			F: schema.HashResource(restrictIPResource()),
		}

		for _, restrictIPData := range zr.RestrictIPList {
			restrictIP := make(map[string]interface{})

			restrictIP["start_ip"] = restrictIPData.StartIP
			restrictIP["end_ip"] = restrictIPData.EndIP
			restrictIP["cidr"] = restrictIPData.Cidr
			restrictIP["single_ip"] = restrictIPData.SingleIP
			restrictIP["comment"] = restrictIPData.Comment

			s.Add(restrictIP)
		}

		primaryCreateInfo["restrict_ip"] = s
	}

	if zr.Tsig != nil {
		tsig := make(map[string]interface{})

		tsig["tsig_key_name"] = zr.Tsig.TsigKeyName
		tsig["tsig_key_value"] = zr.Tsig.TsigKeyValue
		tsig["tsig_algorithm"] = zr.Tsig.TsigAlgorithm
		tsig["description"] = zr.Tsig.Description

		s := &schema.Set{
			F: schema.HashResource(tsigResource()),
		}

		s.Add(tsig)

		primaryCreateInfo["tsig"] = s
	}

	set.Add(primaryCreateInfo)

	return set
}

func flattenSecondaryZone(zr *zone.Response, rd *schema.ResourceData) *schema.Set {
	set := &schema.Set{
		F: schema.HashResource(secondaryZoneCreateInfoResource()),
	}

	secondaryCreateInfo := make(map[string]interface{})

	if val, ok := rd.GetOk("secondary_create_info"); ok && val.(*schema.Set).Len() > 0 {
		secondaryCreateInfo = val.(*schema.Set).List()[0].(map[string]interface{})
	}

	if zr.NotificationEmailAddress != "" {
		secondaryCreateInfo["notification_email_address"] = zr.NotificationEmailAddress
	}

	if zr.PrimaryNameServers != nil && zr.PrimaryNameServers.NameServerIPList != nil {
		if zr.PrimaryNameServers.NameServerIPList.NameServerIP1 != nil {
			s := &schema.Set{
				F: schema.HashResource(nameServerResource()),
			}
			s.Add(getNameServer(zr.PrimaryNameServers.NameServerIPList.NameServerIP1))
			secondaryCreateInfo["primary_name_server_1"] = s
		}

		if zr.PrimaryNameServers.NameServerIPList.NameServerIP2 != nil {
			s := &schema.Set{
				F: schema.HashResource(nameServerResource()),
			}
			s.Add(getNameServer(zr.PrimaryNameServers.NameServerIPList.NameServerIP2))
			secondaryCreateInfo["primary_name_server_2"] = s
		}

		if zr.PrimaryNameServers.NameServerIPList.NameServerIP3 != nil {
			s := &schema.Set{
				F: schema.HashResource(nameServerResource()),
			}
			s.Add(getNameServer(zr.PrimaryNameServers.NameServerIPList.NameServerIP3))
			secondaryCreateInfo["primary_name_server_3"] = s
		}
	}

	set.Add(secondaryCreateInfo)

	return set
}

func flattenAliasZone(zr *zone.Response) *schema.Set {
	set := &schema.Set{
		F: schema.HashResource(aliasZoneCreateInfoResource()),
	}

	aliasCreateInfo := make(map[string]interface{})
	aliasCreateInfo["original_zone_name"] = zr.OriginalZoneName

	set.Add(aliasCreateInfo)

	return set
}

func getNameServer(ns *zone.NameServer) map[string]interface{} {
	nameserver := make(map[string]interface{})
	nameserver["ip"] = ns.IP
	nameserver["tsig_key"] = ns.TsigKey
	nameserver["tsig_key_value"] = ns.TsigKeyValue
	nameserver["tsig_algorithm"] = ns.TsigAlgorithm

	return nameserver
}

func validateZoneName(i interface{}, s string) (warns []string, errs []error) {
	zoneName := i.(string)

	if len(zoneName) > 0 {
		if lastChar := zoneName[len(zoneName)-1]; lastChar != '.' {
			errs = append(errs, helper.ZoneNameValidationError())
		}
	}

	return
}
