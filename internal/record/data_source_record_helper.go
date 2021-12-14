package record

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenRecordData(recordData []string) *schema.Set {
	set := &schema.Set{F: schema.HashString}

	for _, data := range recordData {
		set.Add(data)
	}

	return set
}

func getRecordTypeString(key string) string {
	var rrTypes = map[string]string{
		"A (1)":             "A",
		"NS (2)":            "NS",
		"CNAME (5)":         "CNAME",
		"SOA (6)":           "SOA",
		"PTR (12)":          "PTR",
		"HINFO (13)":        "HINFO",
		"MX (15)":           "MX",
		"TXT (16)":          "TXT",
		"RP (17)":           "RP",
		"AAAA (28)":         "AAAA",
		"SRV (33)":          "SRV",
		"NAPTR (35)":        "NAPTR",
		"DS (43)":           "DS",
		"SSHFP (44)":        "SSHFP",
		"TLSA (52)":         "TLSA",
		"SPF (99)":          "SPF",
		"CAA (257)":         "CAA",
		"APEXALIAS (65282)": "APEXALIAS",
	}

	return rrTypes[key]
}
