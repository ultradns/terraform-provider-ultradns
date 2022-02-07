package zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func flattenPrimaryZoneInfo(zoneResponse *zone.Response, rd *schema.ResourceData) error {
	if err := rd.Set("inherit", zoneResponse.Inherit); err != nil {
		return err
	}

	if err := rd.Set("tsig", getTsigList(zoneResponse.Tsig)); err != nil {
		return err
	}

	if err := rd.Set("restrict_ip", getRestrictIPListSet(zoneResponse.RestrictIPList)); err != nil {
		return err
	}

	if err := rd.Set("notify_addresses", getNotifyAddressesSet(zoneResponse.NotifyAddresses)); err != nil {
		return err
	}

	if err := rd.Set("registrar_info", getRegistrarInfoList(zoneResponse.RegistrarInfo)); err != nil {
		return err
	}

	return nil
}

func flattenSecondaryZoneInfo(zoneResponse *zone.Response, rd *schema.ResourceData) error {
	if err := rd.Set("notification_email_address", zoneResponse.NotificationEmailAddress); err != nil {
		return err
	}

	if err := rd.Set("transfer_status_details", getTransferStatusDetailsList(zoneResponse.TransferStatusDetails)); err != nil {
		return err
	}

	if err := rd.Set("primary_name_server_1", getNameServerList(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP1)); err != nil {
		return err
	}

	if err := rd.Set("primary_name_server_2", getNameServerList(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP2)); err != nil {
		return err
	}

	if err := rd.Set("primary_name_server_3", getNameServerList(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP3)); err != nil {
		return err
	}

	return nil
}
