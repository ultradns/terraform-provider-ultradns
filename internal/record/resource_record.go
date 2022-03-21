package record

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
)

const recordTypeStringNS = "NS"

func ResourceRecord() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceRecordCreate,
		ReadContext:   resourceRecordRead,
		UpdateContext: resourceRecordUpdate,
		DeleteContext: resourceRecordDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: rrset.ResourceRRSetSchema(),
	}
}

func resourceRecordCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := rrset.NewRRSetWithRecordData(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	_, err := services.RecordService.Create(rrSetKeyData, rrSetData)

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.RecordID())

	return resourceRecordRead(ctx, rd, meta)
}

func resourceRecordRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := rrset.GetRRSetKeyFromID(rd.Id())

	if rrSetKey.RecordType == recordTypeStringNS {
		return resourceNSRecordRead(ctx, rd, meta)
	}

	_, resList, err := services.RecordService.Read(rrSetKey)

	if err != nil {
		rd.SetId("")

		return nil
	}

	if len(resList.RRSets) > 0 {
		if err = rrset.FlattenRRSetWithRecordData(resList, rd); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceRecordUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := rrset.NewRRSetWithRecordData(rd)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	if rrSetKeyData.RecordType == recordTypeStringNS {
		return resourceNSRecordUpdate(ctx, rd, meta)
	}

	_, err := services.RecordService.Update(rrSetKeyData, rrSetData)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceRecordRead(ctx, rd, meta)
}

func resourceRecordDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	if rrSetKeyData.RecordType == recordTypeStringNS {
		return resourceNSRecordDelete(ctx, rd, meta)
	}

	_, err := services.RecordService.Delete(rrSetKeyData)

	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	rd.SetId("")

	return diags
}

func resourceNSRecordRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := rrset.GetRRSetKeyFromID(rd.Id())
	_, resList, err := services.RecordService.Read(rrSetKey)

	if err != nil {
		rd.SetId("")

		return nil
	}

	if len(resList.RRSets) > 0 {
		if err = rrset.FlattenRRSet(resList, rd); err != nil {
			return diag.FromErr(err)
		}

		var oldRecordData []interface{}

		if val, ok := rd.GetOk("record_data"); ok {
			oldRecordData = val.(*schema.Set).List()
		}

		recordData := getMatchedRecordData(oldRecordData, resList.RRSets[0].RData)

		if len(oldRecordData) == 0 {
			recordData = resList.RRSets[0].RData
		}

		if err := rd.Set("record_data", helper.GetSchemaSetFromList(recordData)); err != nil {
			return diag.FromErr(err)
		}

	}

	return diags
}

func resourceNSRecordUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	rrSetData := rrset.NewRRSetWithRecordData(rd)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, resList, err := services.RecordService.Read(rrSetKeyData)

	if err != nil {
		return diag.FromErr(err)
	}

	old, new := rd.GetChange("record_data")

	rmData := getDiffRecordData(new.(*schema.Set).List(), old.(*schema.Set).List())
	addData := getDiffRecordData(old.(*schema.Set).List(), new.(*schema.Set).List())

	rrSetData.RData = resList.RRSets[0].RData

	rrSetData.RData = rmRecordData(rmData, rrSetData.RData)
	rrSetData.RData = addRecordData(addData, rrSetData.RData)

	_, er := services.RecordService.Update(rrSetKeyData, rrSetData)

	if er != nil {
		return diag.FromErr(er)
	}

	return resourceNSRecordRead(ctx, rd, meta)
}

func resourceNSRecordDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, resList, err := services.RecordService.Read(rrSetKeyData)

	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	var oldRecordData []interface{}

	if val, ok := rd.GetOk("record_data"); ok {
		oldRecordData = val.(*schema.Set).List()
	}

	if len(oldRecordData) == len(resList.RRSets[0].RData) {
		_, err := services.RecordService.Delete(rrSetKeyData)

		if err != nil {
			rd.SetId("")

			return diag.FromErr(err)
		}

		return diags
	}

	rrSetData := resList.RRSets[0]
	rrSetData.RData = getUnMatchedRecordData(oldRecordData, resList.RRSets[0].RData)
	_, er := services.RecordService.Update(rrSetKeyData, rrSetData)

	if er != nil {
		rd.SetId("")

		return diag.FromErr(er)
	}

	rd.SetId("")

	return diags
}

func getMatchedRecordData(state []interface{}, server []string) []string {
	data := []string{}
	dataMap := make(map[string]bool)

	for _, val := range state {
		dataMap[val.(string)] = true
	}

	for _, val := range server {
		if dataMap[val] {
			data = append(data, val)
		}
	}

	return data
}

func getUnMatchedRecordData(state []interface{}, server []string) []string {
	data := []string{}
	dataMap := make(map[string]bool)

	for _, val := range state {
		dataMap[val.(string)] = true
	}

	for _, val := range server {
		if !dataMap[val] {
			data = append(data, val)
		}
	}

	return data
}

func getDiffRecordData(first []interface{}, second []interface{}) []string {
	data := []string{}
	dataMap := make(map[string]bool)

	for _, val := range first {
		dataMap[val.(string)] = true
	}

	for _, val := range second {
		if !dataMap[val.(string)] {
			data = append(data, val.(string))
		}
	}

	return data
}

func rmRecordData(data, target []string) []string {
	dataMap := make(map[string]bool)

	for _, val := range data {
		dataMap[val] = true
	}

	for i, val := range target {
		if dataMap[val] {
			target[i] = target[len(target)-1]
			target[len(target)-1] = ""
			target = target[:len(target)-1]
		}
	}

	return target
}

func addRecordData(data, target []string) []string {
	for _, val := range data {
		target = append(target, val)
	}
	return target
}
