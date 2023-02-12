package dirgroupgeo

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/geo"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

func ResourceGeoGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGeoGroupCreate,
		ReadContext:   resourceGeoGroupRead,
		UpdateContext: resourceGeoGroupUpdate,
		DeleteContext: resourceGeoGroupDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceGeoGroupSchema(),
	}
}

func resourceGeoGroupCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	geoGroupData := newGeoGroup(rd)

	_, err := services.DirGroupGeoService.Create(geoGroupData)

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(geoGroupData.DirGroupGeoID())

	return resourceGeoGroupRead(ctx, rd, meta)
}

func resourceGeoGroupRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	geoID := rd.Id()

	_, geoGroup, _, err := services.DirGroupGeoService.Read(geoID)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.Set("name", geoGroup.Name)
	rd.Set("account_name", helper.GetAccountName(geoID))
	rd.Set("description", geoGroup.Description)
	rd.Set("codes", geoGroup.Codes)

	return diags
}

func resourceGeoGroupUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {

	services := meta.(*service.Service)
	geoGroupData := newGeoGroup(rd)

	_, err := services.DirGroupGeoService.Update(geoGroupData)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGeoGroupRead(ctx, rd, meta)
}

func resourceGeoGroupDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	geoID := rd.Id()
	_, err := services.DirGroupGeoService.Delete(geoID)

	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	rd.SetId("")

	return diags
}

func newGeoGroup(rd *schema.ResourceData) *geo.DirGroupGeo {
	geoData := &geo.DirGroupGeo{}

	if val, ok := rd.GetOk("name"); ok {
		geoData.Name = val.(string)
	}
	if val, ok := rd.GetOk("account_name"); ok {
		geoData.AccountName = val.(string)
	}
	if val, ok := rd.GetOk("description"); ok {
		geoData.Description = val.(string)
	}
	if val, ok := rd.GetOk("codes"); ok {
		geoCodesData := val.(*schema.Set).List()
		geoData.Codes = make([]string, len(geoCodesData))
		for i, geoCode := range geoCodesData {
			geoData.Codes[i] = geoCode.(string)
		}
	}

	return geoData
}
