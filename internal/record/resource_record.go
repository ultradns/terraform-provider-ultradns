package record

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkhelper "github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
)

const (
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
	tflog.Trace(ctx, "Record resource create context invoked")
	services := meta.(*service.Service)
	rrSetData := rrset.NewRRSetWithRecordData(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	switch sdkhelper.GetRecordTypeFullString(rrSetKeyData.RecordType) {
	case record.SOA:
		rd.SetId(rrSetKeyData.RecordID())
		return resourceSOARecordUpdate(ctx, rd, meta)
	case record.CAA:
		rrSetData.RData = formatCAARecord(ctx, rrSetData.RData)
	case record.SVCB, record.HTTPS:
		rrSetData.RData = formatSVCRecord(ctx, rrSetData.RData)
	}

	_, err := services.RecordService.Create(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(rrSetKeyData.RecordID())

	return resourceRecordRead(ctx, rd, meta)
}

func resourceRecordRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Record resource read context invoked")
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := rrset.GetRRSetKeyFromID(rd.Id())

	res, resList, err := services.RecordService.Read(rrSetKey)

	if err != nil && res != nil && res.Status == helper.RESOURCE_NOT_FOUND {
		tflog.Warn(ctx, errors.ResourceNotFoundError(rd.Id()).Error())
		rd.SetId("")
		return nil
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if len(resList.RRSets) > 0 {
		if err = rrset.FlattenRRSet(resList, rd); err != nil {
			return diag.FromErr(err)
		}

		recData := resList.RRSets[0].RData
		recType := resList.RRSets[0].RRType

		if recType == record.SOA {
			recDataArr := strings.Split(recData[0], " ")
			recDataArr[1] = formatSOAEmail(recDataArr[1])
			recDataArr = append(recDataArr[:2], recDataArr[3:]...)
			recData = []string{strings.Join(recDataArr, " ")}
		}

		if isRecordTypeShareCommonOwnerName(recType) {
			var oldRecData []interface{}

			if val, ok := rd.GetOk("record_data"); ok {
				oldRecData = val.(*schema.Set).List()
			}

			recData = getMatchedRecordData(oldRecData, recData, recType)

			if len(oldRecData) == 0 {
				recData = resList.RRSets[0].RData
			}
		}

		if err := rd.Set("record_data", helper.GetSchemaSetFromList(recData)); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceRecordUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Record resource update context invoked")
	services := meta.(*service.Service)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	if sdkhelper.GetRecordTypeFullString(rrSetKeyData.RecordType) == record.SOA {
		return resourceSOARecordUpdate(ctx, rd, meta)
	}

	if isRecordTypeShareCommonOwnerName(sdkhelper.GetRecordTypeFullString(rrSetKeyData.RecordType)) {
		return resourceCommonOwnerRecordUpdate(ctx, rd, meta)
	}

	rrSetData := rrset.NewRRSetWithRecordData(rd)

	if sdkhelper.GetRecordTypeFullString(rrSetKeyData.RecordType) == record.SVCB ||
		sdkhelper.GetRecordTypeFullString(rrSetKeyData.RecordType) == record.HTTPS {
		rrSetData.RData = formatSVCRecord(ctx, rrSetData.RData)
	}

	_, err := services.RecordService.Update(rrSetKeyData, rrSetData)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceRecordRead(ctx, rd, meta)
}

func resourceCommonOwnerRecordUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Record resource update context invoked")
	services := meta.(*service.Service)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, resList, err := services.RecordService.Read(rrSetKeyData)
	if err != nil {
		return diag.FromErr(err)
	}

	rrSetData := rrset.NewRRSetWithRecordData(rd)

	old, new := rd.GetChange("record_data")

	rmData := getDiffRecordData(new.(*schema.Set).List(), old.(*schema.Set).List(), resList.RRSets[0].RRType)
	addData := getDiffRecordData(old.(*schema.Set).List(), new.(*schema.Set).List(), resList.RRSets[0].RRType)

	rrSetData.RData = resList.RRSets[0].RData

	rrSetData.RData = rmRecordData(rmData, rrSetData.RData)
	rrSetData.RData = addRecordData(addData, rrSetData.RData)

	if sdkhelper.GetRecordTypeFullString(rrSetKeyData.RecordType) == record.CAA {
		rrSetData.RData = formatCAARecord(ctx, rrSetData.RData)
	}

	_, er := services.RecordService.Update(rrSetKeyData, rrSetData)

	if er != nil {
		return diag.FromErr(er)
	}

	return resourceRecordRead(ctx, rd, meta)
}

func resourceSOARecordUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "SOA Record resource update context invoked")
	services := meta.(*service.Service)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, resList, err := services.RecordService.Read(rrSetKeyData)
	if err != nil {
		return diag.FromErr(err)
	}

	rrSetData := rrset.NewRRSetWithRecordData(rd)

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

	return resourceRecordRead(ctx, rd, meta)
}

func resourceRecordDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Record resource delete context invoked")
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	if sdkhelper.GetRecordTypeFullString(rrSetKeyData.RecordType) == record.SOA {
		rd.SetId("")
		return diags
	}

	if isRecordTypeShareCommonOwnerName(sdkhelper.GetRecordTypeFullString(rrSetKeyData.RecordType)) {
		return resourceCommonOwnerRecordDelete(rd, services)
	}

	_, err := services.RecordService.Delete(rrSetKeyData)
	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	rd.SetId("")

	return diags
}

func resourceCommonOwnerRecordDelete(rd *schema.ResourceData, services *service.Service) diag.Diagnostics {
	var diags diag.Diagnostics

	rrSetKeyData := rrset.GetRRSetKeyFromID(rd.Id())

	_, resList, err := services.RecordService.Read(rrSetKeyData)
	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	var oldRecData []interface{}

	if val, ok := rd.GetOk("record_data"); ok {
		oldRecData = val.(*schema.Set).List()
	}

	rrSetData := resList.RRSets[0]
	rrSetData.RData = getUnMatchedRecordData(oldRecData, resList.RRSets[0].RData, resList.RRSets[0].RRType)

	if len(rrSetData.RData) == 0 {
		_, err := services.RecordService.Delete(rrSetKeyData)
		if err != nil {
			rd.SetId("")
			return diag.FromErr(err)
		}
		return diags
	}

	_, er := services.RecordService.Update(rrSetKeyData, rrSetData)

	if er != nil {
		rd.SetId("")
		return diag.FromErr(er)
	}

	rd.SetId("")
	return diags
}
