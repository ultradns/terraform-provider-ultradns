package record

import (
	"context"
	"strings"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const (
	recordTypeStringNS = "NS"
	recordTypeStringSOA = "SOA"
	recordTypeNumberSOA = "6"
	errSOAInvalidFormat = "SOA record format is Invalid. Expected: '<Nameserver> <E-Mail> <REFRESH> <RETRY> <EXPIRE> <MINIMUM>' Found:"
)

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

	if rrSetKeyData.RecordType == recordTypeStringSOA || rrSetKeyData.RecordType == recordTypeNumberSOA {
		return resourceSOARecordUpdate(rd, services, rrSetData, rrSetKeyData)
	}

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
		return resourceNSRecordRead(rd, services)
	}

	if rrSetKey.RecordType == recordTypeStringSOA {
		return resourceSOARecordRead(rd, services)
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
		return resourceNSRecordUpdate(rd, services, rrSetData)
	}

	if rrSetKeyData.RecordType == recordTypeStringSOA {
		return resourceSOARecordUpdate(rd, services, rrSetData, rrSetKeyData)
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
		return resourceNSRecordDelete(rd, services)
	}

	if rrSetKeyData.RecordType == recordTypeStringSOA {
		rd.SetId("")

		return diags
	}

	_, err := services.RecordService.Delete(rrSetKeyData)

	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	rd.SetId("")

	return diags
}

func resourceNSRecordRead(rd *schema.ResourceData, services *service.Service) diag.Diagnostics {
	var diags diag.Diagnostics

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

func resourceNSRecordUpdate(rd *schema.ResourceData, services *service.Service, rrSetData *sdkrrset.RRSet) diag.Diagnostics {
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

	return resourceNSRecordRead(rd, services)
}

func resourceNSRecordDelete(rd *schema.ResourceData, services *service.Service) diag.Diagnostics {
	var diags diag.Diagnostics

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

func resourceSOARecordRead(rd *schema.ResourceData, services *service.Service) diag.Diagnostics {
	var diags diag.Diagnostics

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

		recordDataArr := strings.Split(resList.RRSets[0].RData[0], " ")
		recordDataArr[1] = formatSOAEmail(recordDataArr[1])
		recordDataArr = append(recordDataArr[:2], recordDataArr[3:]...)
		recordData := []string{strings.Join(recordDataArr, " ")}

		if err := rd.Set("record_data", helper.GetSchemaSetFromList(recordData)); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceSOARecordUpdate(rd *schema.ResourceData, services *service.Service, rrSetData *sdkrrset.RRSet,rrSetKeyData *sdkrrset.RRSetKey) diag.Diagnostics {
	_, resList, err := services.RecordService.Read(rrSetKeyData)

	if err != nil {
		return diag.FromErr(err)
	}

	_, new := rd.GetChange("record_data")

	serverRData := strings.Split(resList.RRSets[0].RData[0], " ")
	newRData := strings.Split(new.(*schema.Set).List()[0].(string), " ")

	if len(newRData) != 6 {
		return diag.FromErr(fmt.Errorf("%s %s", errSOAInvalidFormat, strings.Join(newRData, " ")))
	}

	serverRData[0] = newRData[0]
	serverRData[1] = escapeSOAEmail(newRData[1])
	serverRData[3] = newRData[2]
	serverRData[4] = newRData[3]
	serverRData[5] = newRData[4]
	serverRData[6] = newRData[5]

	rrSetData.RData = []string{strings.Join(serverRData, " ")}

	_, er := services.RecordService.Update(rrSetKeyData, rrSetData)

	if er != nil {
		return diag.FromErr(er)
	}

	rd.SetId(rrSetKeyData.RecordID())

	return resourceSOARecordRead(rd, services)
}