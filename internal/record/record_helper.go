package record

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
)

func formatRecordData(rec, recType string) string {
	switch recType {
	case record.CAA:
		return helper.FormatCAARecord(rec)
	case record.SVCB, record.HTTPS:
		return helper.FormatSVCRecord(rec)
	}
	return rec
}

func getMatchedRecordData(state []interface{}, server []string, recType string) []string {
	res := []string{}
	resMap := make(map[string]bool)

	for _, val := range state {
		resMap[formatRecordData(val.(string), recType)] = true
	}

	for _, val := range server {
		if resMap[val] {
			res = append(res, val)
		}
	}

	return res
}

func getUnMatchedRecordData(state []interface{}, server []string, recType string) []string {
	res := []string{}
	resMap := make(map[string]bool)

	for _, val := range state {
		resMap[formatRecordData(val.(string), recType)] = true
	}

	for _, val := range server {
		if !resMap[val] {
			res = append(res, val)
		}
	}

	return res
}

func getDiffRecordData(first, second []interface{}, recType string) []string {
	res := []string{}
	resMap := make(map[string]bool)

	for _, val := range first {
		resMap[formatRecordData(val.(string), recType)] = true
	}

	for _, val := range second {
		if !resMap[formatRecordData(val.(string), recType)] {
			res = append(res, val.(string))
		}
	}

	return res
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
	target = append(target, data...)

	return target
}

func escapeSOAEmail(email string) string {
	index1 := strings.Index(email, "@")
	index2 := strings.LastIndex(email[:index1], ".")

	if index2 == -1 {
		return email[:index1] + "." + email[index1+1:]
	}

	return strings.Replace(email[:index2]+"\\"+email[index2:], "@", ".", 1)
}

func formatSOAEmail(email string) string {
	index := strings.Index(email, "\\.")

	if index == -1 {
		return strings.Replace(email, ".", "@", 1)
	}

	return email[:index] + "." + strings.Replace(email[index+2:], ".", "@", 1)
}

func isRecordTypeShareCommonOwnerName(recordType string) bool {
	rrtypeWithCommonOwner := map[string]bool{
		"NS (2)":       true,
		"HINFO (13)":   true,
		"MX (15)":      true,
		"TXT (16)":     true,
		"RP (17)":      true,
		"SRV (33)":     true,
		"DS (43)":      true,
		"CDS (59)":     true,
		"CDNSKEY (60)": true,
		"SPF (99)":     true,
		"CAA (257)":    true,
	}

	if _, ok := rrtypeWithCommonOwner[recordType]; ok {
		return true
	}
	return false
}

func formatCAARecord(ctx context.Context, recData []string) []string {
	tflog.Debug(ctx, fmt.Sprintf("CAA Record before formatting - %s\n", recData))
	for i, v := range recData {
		recData[i] = helper.FormatCAARecord(v)
	}
	tflog.Debug(ctx, fmt.Sprintf("CAA Record after formatting - %s\n", recData))
	return recData
}

func formatSVCRecord(ctx context.Context, recData []string) []string {
	tflog.Debug(ctx, fmt.Sprintf("SVC Record before formatting - %s\n", recData))
	recData[0] = helper.FormatSVCRecord(recData[0])
	tflog.Debug(ctx, fmt.Sprintf("SVC Record after formatting - %s\n", recData))
	return recData
}
