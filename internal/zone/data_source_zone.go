package zone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
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

	switch zoneResponse.Properties.Type {
	case zone.Primary:
		if err := flattenPrimaryZoneInfo(zoneResponse, rd); err != nil {
			return diag.FromErr(err)
		}
	case zone.Secondary:
		if err := flattenSecondaryZoneInfo(zoneResponse, rd); err != nil {
			return diag.FromErr(err)
		}
	case zone.Alias:
		if err := rd.Set("original_zone_name", zoneResponse.OriginalZoneName); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}
