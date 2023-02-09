package rdpool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

func DataSourceRDPool() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRDPoolRead,

		Schema: dataSourceRDPoolSchema(),
	}
}

func dataSourceRDPoolRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)

	rrSetKeyData := rrset.NewRRSetKey(rd)
	rrSetKeyData.PType = pool.RD
	_, resList, err := services.RecordService.Read(rrSetKeyData)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.RecordID())

	if len(resList.RRSets) > 0 {
		if err = flattenRDPool(resList, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}
