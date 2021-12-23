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

	rd.SetId(rrSetKeyData.ID())

	if len(resList.RRSets) > 0 {
		if err = flattenRecord(resList, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}
