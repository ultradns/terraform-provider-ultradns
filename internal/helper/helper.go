package helper

import (
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

func ZoneFQDNDiffSuppress(k, old, new string, rd *schema.ResourceData) bool {
	return helper.GetZoneFQDN(old) == helper.GetZoneFQDN(new)
}

func OwnerFQDNDiffSuppress(k, old, new string, rd *schema.ResourceData) bool {
	zoneName := ""

	if val, ok := rd.GetOk("zone_name"); ok {
		zoneName = val.(string)
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

func URIDiffSuppress(k, old, new string, rd *schema.ResourceData) bool {
	return old == new || old == strings.TrimSuffix(new, "/") || strings.TrimSuffix(old, "/") == new
}

func ComputedDescriptionDiffSuppress(k, old, new string, rd *schema.ResourceData) bool {
	zoneName := ""
	ownerName := ""

	if val, ok := rd.GetOk("zone_name"); ok {
		zoneName = val.(string)
	}

	if val, ok := rd.GetOk("owner_name"); ok {
		ownerName = val.(string)
	}

	return old == helper.GetOwnerFQDN(ownerName, zoneName) && new == ""
}

func RecordTypeValidation(i interface{}, p cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	var supportedRRType = map[string]bool{
		"A": true, "1": true,
		"NS": true, "2": true,
		"CNAME": true, "5": true,
		"PTR": true, "12": true,
		"MX": true, "15": true,
		"TXT": true, "16": true,
		"AAAA": true, "28": true,
		"SRV": true, "33": true,
		"SSHFP": true, "44": true,
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
