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
	rd.Set("name", geoGroup.Name)
	rd.Set("account_name", helper.GetAccountNameFromURI(uri))
	rd.Set("description", geoGroup.Description)
	rd.Set("codes", geoGroup.Codes)
	return diags

}
