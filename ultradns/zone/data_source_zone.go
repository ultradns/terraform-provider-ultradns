package zone

import (
	"context"
	"strconv"
	ultradns "terraform-provider-ultradns/udnssdk"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceZone() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceZoneRead,

		Schema: zoneDsSchema(),
	}
}

func dataSourceZoneRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*ultradns.Client)

	param := getUrlParameters(rd)

	_, zoneListResponse, err := client.ListZone(param)

	if err != nil {
		return diag.FromErr(err)
	}

	er := mapZoneDataSourceSchema(zoneListResponse, rd)

	if er != nil {
		return diag.FromErr(er)
	}

	rd.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func getUrlParameters(rd *schema.ResourceData) string {
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

	return param
}
