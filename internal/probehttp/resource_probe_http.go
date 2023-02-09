package probehttp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkprobe "github.com/ultradns/ultradns-go-sdk/pkg/probe"
	sdkprobehelper "github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/http"
)

func ResourceProbeHTTP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProbeHTTPCreate,
		ReadContext:   resourceProbeHTTPRead,
		UpdateContext: resourceProbeHTTPUpdate,
		DeleteContext: resourceProbeHTTPDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceProbeHTTPSchema(),
	}
}

func resourceProbeHTTPCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	probeData := getNewProbeHTTP(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	rrSetKeyData.RecordType = probe.RecordTypeA

	res, err := services.ProbeService.Create(rrSetKeyData, probeData)
	if err != nil {
		return diag.FromErr(err)
	}

	uri := res.Header.Get("Location")

	rrSetKeyData.ID = helper.GetProbeIDFromURI(uri)

	rd.SetId(rrSetKeyData.PID())

	return resourceProbeHTTPRead(ctx, rd, meta)
}

func resourceProbeHTTPRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := probe.GetRRSetKeyFromID(rd.Id())
	rrSetKey.PType = sdkprobe.HTTP
	_, probeData, err := services.ProbeService.Read(rrSetKey)
	if err != nil {
		rd.SetId("")

		return nil
	}

	if err = probe.FlattenRRSetKey(rrSetKey, rd); err != nil {
		return diag.FromErr(err)
	}

	if err = flattenProbeHTTP(probeData, rd); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceProbeHTTPUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	probeData := getNewProbeHTTP(rd)
	rrSetKeyData := probe.GetRRSetKeyFromID(rd.Id())

	_, err := services.ProbeService.Update(rrSetKeyData, probeData)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceProbeHTTPRead(ctx, rd, meta)
}

func resourceProbeHTTPDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKeyData := probe.GetRRSetKeyFromID(rd.Id())

	_, err := services.ProbeService.Delete(rrSetKeyData)
	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	rd.SetId("")

	return diags
}

func getNewProbeHTTP(rd *schema.ResourceData) *sdkprobe.Probe {
	probeData := probe.NewProbe(rd, sdkprobe.HTTP)
	details := &http.Details{}
	probeData.Details = details

	if val, ok := rd.GetOk("transaction"); ok {
		transactionsDataList := val.([]interface{})
		details.Transactions = getTransactions(transactionsDataList)
	}

	if val, ok := rd.GetOk("total_limit"); ok && len(val.([]interface{})) > 0 {
		totalLimitData := val.([]interface{})[0].(map[string]interface{})
		details.TotalLimits = probe.GetLimit(totalLimitData)
	}

	return probeData
}

func getTransactions(transactionsDataList []interface{}) []*http.Transaction {
	transactionList := make([]*http.Transaction, len(transactionsDataList))

	for i, d := range transactionsDataList {
		transactionData := d.(map[string]interface{})
		transactionList[i] = getTransaction(transactionData)
	}

	return transactionList
}

func getTransaction(transactionData map[string]interface{}) *http.Transaction {
	transaction := &http.Transaction{}

	if val, ok := transactionData["method"]; ok {
		transaction.Method = val.(string)
	}

	if val, ok := transactionData["protocol_version"]; ok {
		transaction.ProtocolVersion = val.(string)
	}

	if val, ok := transactionData["url"]; ok {
		transaction.URL = val.(string)
	}

	if val, ok := transactionData["transmitted_data"]; ok {
		transaction.TransmittedData = val.(string)
	}

	if val, ok := transactionData["expected_response"]; ok {
		transaction.ExpectedResponse = val.(string)
	}

	if val, ok := transactionData["follow_redirects"]; ok {
		transaction.FollowRedirects = val.(bool)
	}

	transaction.Limits = getTransactionLimitInfo(transactionData)

	return transaction
}

func getTransactionLimitInfo(transactionData map[string]interface{}) *sdkprobehelper.LimitsInfo {
	limitsInfo := &sdkprobehelper.LimitsInfo{}

	if val, ok := transactionData["search_string"]; ok && len(val.([]interface{})) > 0 {
		searchStringData := val.([]interface{})[0].(map[string]interface{})
		limitsInfo.SearchString = probe.GetSearchString(searchStringData)
	}

	if val, ok := transactionData["connect_limit"]; ok && len(val.([]interface{})) > 0 {
		connectLimitData := val.([]interface{})[0].(map[string]interface{})
		limitsInfo.Connect = probe.GetLimit(connectLimitData)
	}

	if val, ok := transactionData["avg_connect_limit"]; ok && len(val.([]interface{})) > 0 {
		avgConnectLimitData := val.([]interface{})[0].(map[string]interface{})
		limitsInfo.AvgConnect = probe.GetLimit(avgConnectLimitData)
	}

	if val, ok := transactionData["run_limit"]; ok && len(val.([]interface{})) > 0 {
		runLimitData := val.([]interface{})[0].(map[string]interface{})
		limitsInfo.Run = probe.GetLimit(runLimitData)
	}

	if val, ok := transactionData["avg_run_limit"]; ok && len(val.([]interface{})) > 0 {
		avgRunLimitData := val.([]interface{})[0].(map[string]interface{})
		limitsInfo.AvgRun = probe.GetLimit(avgRunLimitData)
	}

	return limitsInfo
}
