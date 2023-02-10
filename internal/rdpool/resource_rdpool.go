package rdpool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/rdpool"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func ResourceRDPool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRDPoolCreate,
		ReadContext:   resourceRDPoolRead,
		UpdateContext: resourceRDPoolUpdate,
		DeleteContext: resourceRDPoolDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceRDPoolSchema(),
	}
}

func resourceRDPoolCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := getNewRDPool(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	_, err := services.RecordService.Create(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.RecordID())

	return resourceRDPoolRead(ctx, rd, meta)
}

func resourceRDPoolRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := rrset.GetRRSetKeyFromID(rd.Id())
	rrSetKey.PType = pool.RD

	_, resList, err := services.RecordService.Read(rrSetKey)
	if err != nil {
		rd.SetId("")

		return nil
	}

	if len(resList.RRSets) > 0 {
		if err = flattenRDPool(resList, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceRDPoolUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := getNewRDPool(rd)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, err := services.RecordService.Update(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceRDPoolRead(ctx, rd, meta)
}

func resourceRDPoolDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, err := services.RecordService.Delete(rrSetKeyData)
	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	rd.SetId("")

	return diags
}

func getNewRDPool(rd *schema.ResourceData) *sdkrrset.RRSet {
	rrSetData := rrset.NewRRSetWithRecordData(rd)
	profile := &rdpool.Profile{}
	rrSetData.Profile = profile

	if val, ok := rd.GetOk("order"); ok {
		profile.Order = val.(string)
	}

	if val, ok := rd.GetOk("description"); ok {
		profile.Description = val.(string)
	}

	return rrSetData
}
