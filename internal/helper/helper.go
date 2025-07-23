package helper

import (
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
)

const RESOURCE_NOT_FOUND = "404 Not Found"

func CaseInSensitiveState(val any) string {
	return strings.ToLower(val.(string))
}

func ZoneFQDNDiffSuppress(k, old, new string, rd *schema.ResourceData) bool {
	if len(old) == 0 || len(new) == 0 {
		return false
	}

	return helper.GetZoneFQDN(old) == helper.GetZoneFQDN(new)
}

func OwnerFQDNDiffSuppress(k, old, new string, rd *schema.ResourceData) bool {
	if len(old) == 0 || len(new) == 0 {
		return false
	}

	zoneName := ""

	if val, ok := rd.GetOk("zone_name"); ok {
		zoneName = strings.ToLower(val.(string))
	}

	return helper.GetOwnerFQDN(old, zoneName) == helper.GetOwnerFQDN(new, zoneName)
}

func RecordTypeDiffSuppress(k, old, new string, rd *schema.ResourceData) bool {
	oldRecordType := helper.GetRecordTypeFullString(old)
	newRecordType := helper.GetRecordTypeFullString(new)

	if len(oldRecordType) == 0 && len(newRecordType) == 0 {
		return false
	}

	return oldRecordType == newRecordType || oldRecordType == new
}

func RecordDataDiffSuppress(k, old, new string, rd *schema.ResourceData) bool {
	recType := ""

	if val, ok := rd.GetOk("record_type"); ok {
		recType = helper.GetRecordTypeFullString(val.(string))
	}

	if recType == record.CAA || recType == record.SVCB || recType == record.HTTPS {
		return fmtDiffSuppress(recType, rd)
	}

	return false
}

func fmtDiffSuppress(recordType string, rd *schema.ResourceData) bool {
	o, n := rd.GetChange("record_data")
	oldData := o.(*schema.Set)
	newData := n.(*schema.Set)

	if oldData.Len() == 0 || newData.Len() == 0 {
		return false
	}

	switch recordType {
	case record.CAA:
		return formatCAARecordSet(oldData).Equal(formatCAARecordSet(newData))
	case record.SVCB, record.HTTPS:
		return formatSVCRecordSet(oldData).Equal(formatSVCRecordSet(newData))
	}

	return false
}

func URIDiffSuppress(k, old, new string, rd *schema.ResourceData) bool {
	return old == new || old == strings.TrimSuffix(new, "/") || strings.TrimSuffix(old, "/") == new
}

func ComputedDescriptionDiffSuppress(k, old, new string, rd *schema.ResourceData) bool {
	zoneName := ""
	ownerName := ""

	if val, ok := rd.GetOk("zone_name"); ok {
		zoneName = strings.ToLower(val.(string))
	}

	if val, ok := rd.GetOk("owner_name"); ok {
		ownerName = strings.ToLower(val.(string))
	}

	return old == helper.GetOwnerFQDN(ownerName, zoneName) && new == ""
}

func RecordTypeValidation(i interface{}, p cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	supportedRRType := map[string]bool{
		"A": true, "1": true,
		"NS": true, "2": true,
		"CNAME": true, "5": true,
		"SOA": true, "6": true,
		"PTR": true, "12": true,
		"MX": true, "15": true,
		"TXT": true, "16": true,
		"AAAA": true, "28": true,
		"SRV": true, "33": true,
		"DS": true, "43": true,
		"SSHFP": true, "44": true,
		"SVCB": true, "64": true,
		"HTTPS": true, "65": true,
		"CAA": true, "257": true,
		"APEXALIAS": true, "65282": true,
	}

	recordType := i.(string)
	_, ok := supportedRRType[recordType]

	if !ok {
		return diag.Errorf("invalid or unsupported record type")
	}

	return diags
}

func GetSchemaSetFromList(dataList []string) *schema.Set {
	set := &schema.Set{F: schema.HashString}

	for _, data := range dataList {
		set.Add(data)
	}

	return set
}

func GetProbeIDFromURI(uri string) string {
	return splitURI(uri, "probes/")
}

func GetGeoIdFromURI(uri string) string {
	return splitURI(uri, "geo/")
}

func GetIPIdFromURI(uri string) string {
	return splitURI(uri, "ip/")
}

func splitURI(uri, split string) string {
	splitStringData := strings.Split(uri, split)

	if len(splitStringData) == 2 {
		return splitStringData[1]
	}

	return ""
}

func formatCAARecordSet(data *schema.Set) *schema.Set {
	result := &schema.Set{F: schema.HashString}

	for _, d := range data.List() {
		result.Add(FormatCAARecord(d.(string)))
	}

	return result
}

func FormatCAARecord(rec string) string {
	splitStringData := strings.SplitN(rec, " ", 3)
	if len(splitStringData) == 3 {
		splitStringData[2] = strings.Trim(splitStringData[2], "\"")
		splitStringData[2] = "\"" + splitStringData[2] + "\""
	}
	return strings.Join(splitStringData, " ")
}

func formatSVCRecordSet(data *schema.Set) *schema.Set {
	result := &schema.Set{F: schema.HashString}

	for _, d := range data.List() {
		result.Add(FormatSVCRecord(d.(string)))
	}

	return result
}

func FormatSVCRecord(rec string) string {
	svcDataSplt := strings.SplitN(rec, " ", 3)

	if len(svcDataSplt) == 3 {
		svcDataSplt[2] = formatSVCParams(svcDataSplt[2])
	}

	return strings.Join(svcDataSplt, " ")
}

func formatSVCParams(svcParams string) string {
	svcParams = strings.TrimPrefix(svcParams, "(")
	svcParams = strings.TrimSuffix(svcParams, ")")
	svcParams = strings.TrimSpace(svcParams)
	svcParamsSplit := strings.Split(svcParams, " ")
	svcParamsMap := make(map[int]string)

	for _, v := range svcParamsSplit {
		key := -1
		value := ""
		paramSplit := strings.SplitN(v, "=", 2)
		if len(paramSplit) > 0 {
			key = getSvcKeyNumber(paramSplit[0])
		}
		if len(paramSplit) > 1 {
			value = strings.Trim(paramSplit[1], "\"")
		}
		svcParamsMap[key] = value
	}

	keys := make([]int, 0)
	for k := range svcParamsMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	result := ""
	for _, k := range keys {
		result += getSvcParamText(k, svcParamsMap[k])
	}

	return strings.TrimSpace(result)
}

func getSvcKeyNumber(s string) int {
	svcKey := strings.ToUpper(s)
	switch svcKey {
	case "MANDATORY":
		return 0
	case "ALPN":
		return 1
	case "NO-DEFAULT-ALPN":
		return 2
	case "PORT":
		return 3
	case "IPV4HINT":
		return 4
	case "ECH":
		return 5
	case "IPV6HINT":
		return 6
	case "DOHPATH":
		return 7
	case "OHTTP":
		return 8
	}

	if strings.HasPrefix(svcKey, "KEY") {
		splt := strings.Split(svcKey, "KEY")
		i, err := strconv.Atoi(splt[1])
		if err != nil {
			return -1
		}
		return i
	}
	return -1
}

func getSvcParamText(key int, value string) string {
	switch key {
	case 0:
		return "mandatory=" + value + " "
	case 1:
		return "alpn=\"" + value + "\" "
	case 2:
		return "no-default-alpn "
	case 3:
		return "port=" + value + " "
	case 4:
		return "ipv4hint=" + value + " "
	case 5:
		return "ech=" + value + " "
	case 6:
		return "ipv6hint=" + value + " "
	case 7:
		return "dohpath=\"" + value + "\" "
	case 8:
		return "ohttp "
	}

	if key > 8 {
		keyStr := strconv.Itoa(key)
		if len(value) == 0 {
			return "key" + keyStr + " "
		} else {
			return "key" + keyStr + "=\"" + value + "\" "
		}
	}
	return ""
}
