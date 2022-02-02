package slbpool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
)

func DataSourceSLBPool() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceSLBPoolRead,

		Schema: dataSourceSLBPoolSchema(),
	}
}

func dataSourceSLBPoolRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)

	rrSetKeyData := rrset.NewRRSetKey(rd)
	_, resList, err := services.SLBPoolService.ReadSLBPool(rrSetKeyData)

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.ID())

	if len(resList.RRSets) > 0 {
		if err = flattenSLBPool(resList, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}
