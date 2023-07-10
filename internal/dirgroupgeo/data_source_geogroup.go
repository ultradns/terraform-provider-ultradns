package dirgroupgeo

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

func DataSourceGeoGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGeoGroupRead,

		Schema: dataSourceGeoSchema(),
	}
}

func dataSourceGeoGroupRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	services := meta.(*service.Service)
	geoGroupData := newGeoGroup(rd)

	_, geoGroup, uri, err := services.DirGroupGeoService.Read(geoGroupData.DirGroupGeoID())
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(geoGroupData.DirGroupGeoID())
	if err := rd.Set("name", geoGroup.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := rd.Set("account_name", helper.GetAccountNameFromURI(uri)); err != nil {
		return diag.FromErr(err)
	}
	if err := rd.Set("description", geoGroup.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := rd.Set("codes", geoGroup.Codes); err != nil {
		return diag.FromErr(err)
	}
	return diags

}
