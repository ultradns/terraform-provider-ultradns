package zone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
)

func DataSourceZone() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceZoneRead,

		Schema: dataSourceZoneSchema(),
	}
}

func dataSourceZoneRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)

	zoneName := ""

	if val, ok := rd.GetOk("name"); ok {
		zoneName = val.(string)
	}

	_, zoneResponse, err := services.ZoneService.ReadZone(zoneName)

	if err != nil {
		return diag.FromErr(err)
	}

	if zoneResponse.Properties != nil {
		rd.SetId(zoneResponse.Properties.Name)

		if err := rd.Set("name", zoneResponse.Properties.Name); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("account_name", zoneResponse.Properties.AccountName); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("type", zoneResponse.Properties.Type); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("dnssec_status", zoneResponse.Properties.DNSSecStatus); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("resource_record_count", zoneResponse.Properties.ResourceRecordCount); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("last_modified_time", zoneResponse.Properties.LastModifiedDateTime); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("status", zoneResponse.Properties.Status); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("owner", zoneResponse.Properties.Owner); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := rd.Set("inherit", zoneResponse.Inherit); err != nil {
		return diag.FromErr(err)
	}

	if err := rd.Set("notification_email_address", zoneResponse.NotificationEmailAddress); err != nil {
		return diag.FromErr(err)
	}

	if err := rd.Set("original_zone_name", zoneResponse.OriginalZoneName); err != nil {
		return diag.FromErr(err)
	}

	if zoneResponse.Tsig != nil {
		if err := rd.Set("tsig", flattenTsig(zoneResponse.Tsig)); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.RestrictIPList != nil {
		if err := rd.Set("restrict_ip", flattenRestrictIP(zoneResponse.RestrictIPList)); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.NotifyAddresses != nil {
		if err := rd.Set("notify_addresses", flattenNotifyAddresses(zoneResponse.NotifyAddresses)); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.RegistrarInfo != nil {
		if err := rd.Set("registrar_info", flattenRegistrarInfo(zoneResponse.RegistrarInfo)); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.TransferStatusDetails != nil {
		if err := rd.Set("transfer_status_details", flattenTransferStatusDetails(zoneResponse.TransferStatusDetails)); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.PrimaryNameServers != nil && zoneResponse.PrimaryNameServers.NameServerIPList != nil && zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP1 != nil {
		if err := rd.Set("primary_name_server_1", flattenNameServer(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP1)); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.PrimaryNameServers != nil && zoneResponse.PrimaryNameServers.NameServerIPList != nil && zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP2 != nil {
		if err := rd.Set("primary_name_server_2", flattenNameServer(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP2)); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.PrimaryNameServers != nil && zoneResponse.PrimaryNameServers.NameServerIPList != nil && zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP3 != nil {
		if err := rd.Set("primary_name_server_3", flattenNameServer(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP3)); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}
