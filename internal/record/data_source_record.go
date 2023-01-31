package record

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
)

func DataSourceRecord() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceRecordRead,

		Schema: rrset.DataSourceRRSetSchema(),
	}
}

func dataSourceRecordRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)

	rrSetKeyData := rrset.NewRRSetKey(rd)
	_, resList, err := services.RecordService.Read(rrSetKeyData)

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.RecordID())

	if len(resList.RRSets) > 0 {
		if err = rrset.FlattenRRSetWithRecordData(resList, rd); err != nil {
			return diag.FromErr(err)
		}

		if resList.RRSets[0].RRType == "SOA (6)" {
			recordDataArr := strings.Split(resList.RRSets[0].RData[0], " ")
			recordDataArr[1] = formatSOAEmail(recordDataArr[1])
			recordDataArr = append(recordDataArr[:2], recordDataArr[3:]...)
			recordData := []string{strings.Join(recordDataArr, " ")}
			
			if err := rd.Set("record_data", helper.GetSchemaSetFromList(recordData)); err != nil {
				return diag.FromErr(err)
			}
		}
		
	}

	return diags
}
