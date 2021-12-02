package zone

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func DataSourceZone() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceZoneRead,

		Schema: dataSourceZoneSchema(),
	}
}

func dataSourceZoneRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*ultradns.Client)

	param := getZoneListURLParameters(rd)

	_, zlr, err := client.ListZone(param)

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	if zlr.QueryInfo != nil {
		if err := rd.Set("query", zlr.QueryInfo.Query); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("sort", zlr.QueryInfo.Sort); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("reverse", zlr.QueryInfo.Reverse); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("limit", zlr.QueryInfo.Limit); err != nil {
			return diag.FromErr(err)
		}
	}

	if zlr.ResultInfo != nil {
		if err := rd.Set("total_count", zlr.ResultInfo.TotalCount); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("returned_count", zlr.ResultInfo.ReturnedCount); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("offset", zlr.ResultInfo.Offset); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := rd.Set("zones", flattenZones(zlr.Zones)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
