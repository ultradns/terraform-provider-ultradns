package zone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
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
		id := helper.GetZoneFQDN(zoneResponse.Properties.Name)
		rd.SetId(id)

		if err := flattenZoneProperties(zoneResponse, rd); err != nil {
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
		if err := flattenTsig(zoneResponse.Tsig, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.RestrictIPList != nil {
		if err := flattenRestrictIPList(zoneResponse.RestrictIPList, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.NotifyAddresses != nil {
		if err := flattenNotifyAddresses(zoneResponse.NotifyAddresses, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.RegistrarInfo != nil {
		if err := flattenRegistrarInfo(zoneResponse.RegistrarInfo, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.TransferStatusDetails != nil {
		if err := flattenTransferStatusDetails(zoneResponse.TransferStatusDetails, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	if zoneResponse.PrimaryNameServers != nil && zoneResponse.PrimaryNameServers.NameServerIPList != nil {
		if zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP1 != nil {
			if err := rd.Set("primary_name_server_1", flattenNameServer(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP1)); err != nil {
				return diag.FromErr(err)
			}
		}

		if zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP2 != nil {
			if err := rd.Set("primary_name_server_2", flattenNameServer(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP2)); err != nil {
				return diag.FromErr(err)
			}
		}

		if zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP3 != nil {
			if err := rd.Set("primary_name_server_3", flattenNameServer(zoneResponse.PrimaryNameServers.NameServerIPList.NameServerIP3)); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return diags
}
