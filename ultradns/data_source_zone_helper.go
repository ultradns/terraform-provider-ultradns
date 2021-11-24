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

	return param
}

func mapZoneDataSourceSchema(zlr *ultradns.ZoneListResponse, rd *schema.ResourceData) error {

	if err := rd.Set("sort", zlr.QueryInfo.Sort); err != nil {
		return err
	}

	if err := rd.Set("reverse", zlr.QueryInfo.Reverse); err != nil {
		return err
	}

	if err := rd.Set("limit", zlr.QueryInfo.Limit); err != nil {
		return err
	}

	if err := rd.Set("total_count", zlr.ResultInfo.TotalCount); err != nil {
		return err
	}

	if err := rd.Set("returned_count", zlr.ResultInfo.ReturnedCount); err != nil {
		return err
	}

	if err := rd.Set("offset", zlr.ResultInfo.Offset); err != nil {
		return err
	}

	zones := make([]interface{}, zlr.ResultInfo.ReturnedCount)

	for i, zone := range *zlr.Zones {
		prop := make(map[string]interface{})

		prop["name"] = zone.Properties.Name
		prop["account_name"] = zone.Properties.AccountName
		prop["type"] = zone.Properties.Type

		zones[i] = prop

	}

	if err := rd.Set("zones", zones); err != nil {
		return err
	}

	return nil
}
