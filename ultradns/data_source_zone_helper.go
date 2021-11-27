package ultradns

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func getZoneListUrlParameters(rd *schema.ResourceData) string {
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

	if val, ok := rd.GetOk("offset"); ok {
		param = param + "&offset=" + strconv.Itoa(val.(int))
	}

	return param
}

func flattenZones(zlr []*ultradns.ZoneResponse) []map[string]interface{} {
	var zones []map[string]interface{}

	for _, zone := range zlr {
		data := map[string]interface{}{
			"name":         zone.Properties.Name,
			"account_name": zone.Properties.AccountName,
			"type":         zone.Properties.Type,
		}
		zones = append(zones, data)
	}
	return zones
}
