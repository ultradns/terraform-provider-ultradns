package probehttp

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
	sdkprobe "github.com/ultradns/ultradns-go-sdk/pkg/probe"
	sdkprobehelper "github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/http"
)

func flattenProbeHTTP(probeData *sdkprobe.Probe, rd *schema.ResourceData) error {
	if err := probe.FlattenProbe(probeData, rd); err != nil {
		return err
	}

	details := probeData.Details.(*http.Details)

	if err := rd.Set("total_limit", probe.GetLimitList(details.TotalLimits)); err != nil {
		return err
	}

	if err := rd.Set("transaction", getTransactionsList(details.Transactions)); err != nil {
		return err
	}

	return nil
}

func getTransactionsList(transactionListData []*http.Transaction) []interface{} {
	var list []interface{}

	if len(transactionListData) > 0 {
		list = make([]interface{}, len(transactionListData))

		for i, transactionData := range transactionListData {
			transaction := make(map[string]interface{})
			transaction["method"] = transactionData.Method
			transaction["protocol_version"] = transactionData.ProtocolVersion
			transaction["url"] = transactionData.URL
			transaction["transmitted_data"] = transactionData.TransmittedData
			transaction["follow_redirects"] = transactionData.FollowRedirects
			transaction["expected_response"] = transactionData.ExpectedResponse
			setTransactionLmits(transactionData.Limits, transaction)
			list[i] = transaction
		}
	}

	return list
}

func setTransactionLmits(limitsInfoData *sdkprobehelper.LimitsInfo, transaction map[string]interface{}) {
	if limitsInfoData != nil {
		transaction["search_string"] = probe.GetSearchStringList(limitsInfoData.SearchString)
		transaction["connect_limit"] = probe.GetLimitList(limitsInfoData.Connect)
		transaction["avg_connect_limit"] = probe.GetLimitList(limitsInfoData.AvgConnect)
		transaction["run_limit"] = probe.GetLimitList(limitsInfoData.Run)
		transaction["avg_run_limit"] = probe.GetLimitList(limitsInfoData.AvgRun)
	}
}
