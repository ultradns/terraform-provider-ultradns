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

	query := getQueryInfo(rd)

	_, zlr, err := services.ZoneService.ListZone(query)

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(query.URI())

	if zlr.QueryInfo != nil {
		if err := rd.Set("query", zlr.QueryInfo.Query); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("sort", zlr.QueryInfo.Sort); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("cursor", zlr.QueryInfo.Cursor); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("reverse", zlr.QueryInfo.Reverse); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("limit", zlr.QueryInfo.Limit); err != nil {
			return diag.FromErr(err)
		}
	}

	// if zlr.ResultInfo != nil {
	// 	if err := rd.Set("total_count", zlr.ResultInfo.TotalCount); err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// 	if err := rd.Set("returned_count", zlr.ResultInfo.ReturnedCount); err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// 	if err := rd.Set("offset", zlr.ResultInfo.Offset); err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// }

	if zlr.CursorInfo != nil {
		if err := rd.Set("next", zlr.CursorInfo.Next); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("previous", zlr.CursorInfo.Previous); err != nil {
			return diag.FromErr(err)
		}
		if err := rd.Set("first", zlr.CursorInfo.First); err != nil {
			return diag.FromErr(err)
		}
		if err := rd.Set("last", zlr.CursorInfo.Last); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := rd.Set("zones", flattenZones(zlr.Zones)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
