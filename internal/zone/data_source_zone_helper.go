package zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func flattenPrimaryZoneInfo(zoneResponse *zone.Response, rd *schema.ResourceData) error {
	if err := rd.Set("inherit", zoneResponse.Inherit); err != nil {
		return err
	}

	if zoneResponse.Tsig != nil {
		if err := rd.Set("tsig", getTsigSet(zoneResponse.Tsig)); err != nil {
			return err
		}
	}

	if zoneResponse.RestrictIPList != nil {
		if err := rd.Set("restrict_ip", getRestrictIPListSet(zoneResponse.RestrictIPList)); err != nil {
			return err
		}
	}

	if zoneResponse.NotifyAddresses != nil {
		if err := rd.Set("notify_addresses", getNotifyAddressesSet(zoneResponse.NotifyAddresses)); err != nil {
			return err
		}
	}

	if zoneResponse.RegistrarInfo != nil {
		if err := flattenRegistrarInfo(zoneResponse.RegistrarInfo, rd); err != nil {
			return err
		}
	}

	return nil
}

func flattenSecondaryZoneInfo(zoneResponse *zone.Response, rd *schema.ResourceData) error {
	if err := rd.Set("notification_email_address", zoneResponse.NotificationEmailAddress); err != nil {
		return err
	}

	if zoneResponse.TransferStatusDetails != nil {
		if err := flattenTransferStatusDetails(zoneResponse.TransferStatusDetails, rd); err != nil {
			return err
		}
	}

	if zoneResponse.PrimaryNameServers != nil && zoneResponse.PrimaryNameServers.NameServerIPList != nil {
		if err := flattenPrimaryNameServers(zoneResponse.PrimaryNameServers.NameServerIPList, rd); err != nil {
			return err
		}
	}

	return nil
}

func flattenRegistrarInfo(registrarInfoData *zone.RegistrarInfo, rd *schema.ResourceData) error {
	set := &schema.Set{F: helper.HashSingleSetResource}
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
	set := &schema.Set{F: helper.HashSingleSetResource}
	registrarInfoNameServersList := make(map[string]interface{})
	registrarInfoNameServersList["ok"] = nameServersList.Ok
	registrarInfoNameServersList["unknown"] = nameServersList.Unknown
	registrarInfoNameServersList["missing"] = nameServersList.Missing
	registrarInfoNameServersList["incorrect"] = nameServersList.Incorrect
	set.Add(registrarInfoNameServersList)

	return set
}

func flattenTransferStatusDetails(transferDetailsData *zone.TransferStatusDetails, rd *schema.ResourceData) error {
	set := &schema.Set{F: helper.HashSingleSetResource}
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
		if err := rd.Set("primary_name_server_1", getNameServerSet(nameServerIPListData.NameServerIP1)); err != nil {
			return err
		}
	}

	if nameServerIPListData.NameServerIP2 != nil {
		if err := rd.Set("primary_name_server_2", getNameServerSet(nameServerIPListData.NameServerIP2)); err != nil {
			return err
		}
	}

	if nameServerIPListData.NameServerIP3 != nil {
		if err := rd.Set("primary_name_server_3", getNameServerSet(nameServerIPListData.NameServerIP3)); err != nil {
			return err
		}
	}

	return nil
}
