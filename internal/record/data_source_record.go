package record

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
)

func DataSourceRecord() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceRecordRead,

		Schema: dataSourceRecordSchema(),
	}
}

func dataSourceRecordRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)

	rrSetKeyData := newRRSetKey(rd)
	_, resList, err := services.RecordService.ReadRecord(rrSetKeyData)

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.URI())

	if resList.QueryInfo != nil {
		if err := rd.Set("query", resList.QueryInfo.Query); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("sort", resList.QueryInfo.Sort); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("reverse", resList.QueryInfo.Reverse); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("limit", resList.QueryInfo.Limit); err != nil {
			return diag.FromErr(err)
		}
	}

	if resList.ResultInfo != nil {
		if err := rd.Set("total_count", resList.ResultInfo.TotalCount); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("returned_count", resList.ResultInfo.ReturnedCount); err != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("offset", resList.ResultInfo.Offset); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := rd.Set("record_sets", flattenRRSets(resList.RRSets)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
