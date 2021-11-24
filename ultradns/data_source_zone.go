package ultradns

import (
	"context"
	"strconv"
	"time"

	"github.com/ultradns/ultradns-go-sdk/ultradns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

	param := getZoneListUrlParameters(rd)

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
