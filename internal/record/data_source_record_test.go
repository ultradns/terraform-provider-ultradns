package record_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
)

func TestAccDataSourceRecord(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()

	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRecordA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_a"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "record_type", "1"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "record_data.0", "192.168.1.1"),
				),
			},

			{
				Config: testAccDataSourceRecordCNAME(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_cname"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "record_type", "CNAME"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "record_data.0", "example.com."),
				),
			},
			{
				Config: testAccDataSourceRecordPTR(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_ptr"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "record_type", "12"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "record_data.0", "example.com."),
				),
			},
			{
				Config: testAccDataSourceRecordHINFO(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_hinfo"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_hinfo", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_hinfo", "record_type", "HINFO"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_hinfo", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_hinfo", "record_data.0", "\"PC\" \"Linux\""),
				),
			},
			{
				Config: testAccDataSourceRecordMX(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_mx"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_mx", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("data.ultradns_record.data_mx", "record_type", "15"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_mx", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_mx", "record_data.0", "2 example.com."),
				),
			},
			{
				Config: testAccDataSourceRecordTXT(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_txt"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "record_type", "TXT"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "record_data.0", "example.com."),
				),
			},
			{
				Config: testAccDataSourceRecordRP(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_rp"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_rp", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_rp", "record_type", "17"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_rp", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_rp", "record_data.0", "test.example.com. example.128/134.123.178.178.in-addr.arpa."),
				),
			},
			{
				Config: testAccDataSourceRecordAAAA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_aaaa"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_aaaa", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_aaaa", "record_data.0", "2001:db8:85a3:0:0:8a2e:370:7334"),
				),
			},
			{
				Config: testAccDataSourceRecordSRV(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_srv"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_srv", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("data.ultradns_record.data_srv", "record_type", "33"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_srv", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_srv", "record_data.0", "5 6 7 example.com."),
				),
			},

			{
				Config: testAccDataSourceRecordNAPTR(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_naptr"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_naptr", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("data.ultradns_record.data_naptr", "record_type", "NAPTR"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_naptr", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_naptr", "record_data.0", "1 2 \"3\" \"test\" \"\" test.com."),
				),
			},
			{
				Config: testAccDataSourceRecordSSHFP(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_sshfp"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_sshfp", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_sshfp", "record_type", "44"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_sshfp", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_sshfp", "record_data.0", "1 2 54B5E539EAF593AEA410F80737530B71CCDE8B6C3D241184A1372E98BC7EDB37"),
				),
			},
			{
				Config: testAccDataSourceRecordTLSA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_tlsa"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_tlsa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_tlsa", "record_type", "TLSA"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_tlsa", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_tlsa", "record_data.0", "0 0 0 aaaaaaaa"),
				),
			},
			{
				Config: testAccDataSourceRecordSPF(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_spf"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_spf", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("data.ultradns_record.data_spf", "record_type", "99"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_spf", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_spf", "record_data.0", "v=spf1 ip4:1.2.3.4 ~all"),
				),
			},
			{
				Config: testAccDataSourceRecordCAA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_caa"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "record_type", "CAA"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "record_data.0", "1 issue \"test\""),
				),
			},
			{
				Config: testAccDataSourceRecordAPEXALIAS(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("data.ultradns_record.data_apex"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "record_type", "65282"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "record_data.0", "example.com."),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func testAccDataSourceRecordA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_a" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.a.owner_name}"
		record_type = "1"
	}
	`, testAccResourceRecordA(zoneName), strings.TrimSuffix(zoneName, "."))
}

func testAccDataSourceRecordCNAME(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_cname" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.cname.owner_name}"
		record_type = "CNAME"
	}
	`, testAccResourceRecordCNAME(zoneName), strings.TrimSuffix(zoneName, "."))
}

func testAccDataSourceRecordPTR(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_ptr" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.ptr.owner_name}"
		record_type = "12"
	}
	`, testAccResourceRecordPTR(zoneName), zoneName)
}

func testAccDataSourceRecordHINFO(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_hinfo" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.hinfo.owner_name}"
		record_type = "HINFO"
	}
	`, testAccResourceRecordHINFO(zoneName), zoneName)
}

func testAccDataSourceRecordMX(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_mx" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.mx.owner_name}"
		record_type = "15"
	}
	`, testAccResourceRecordMX(zoneName), strings.TrimSuffix(zoneName, "."))
}

func testAccDataSourceRecordTXT(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_txt" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.txt.owner_name}"
		record_type = "TXT"
	}
	`, testAccResourceRecordTXT(zoneName), strings.TrimSuffix(zoneName, "."))
}

func testAccDataSourceRecordRP(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_rp" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.rp.owner_name}"
		record_type = "17"
	}
	`, testAccResourceRecordRP(zoneName), zoneName)
}

func testAccDataSourceRecordAAAA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_aaaa" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.aaaa.owner_name}"
		record_type = "AAAA"
	}
	`, testAccResourceRecordAAAA(zoneName), zoneName)
}

func testAccDataSourceRecordSRV(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_srv" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.srv.owner_name}"
		record_type = "33"
	}
	`, testAccResourceRecordSRV(zoneName), strings.TrimSuffix(zoneName, "."))
}

func testAccDataSourceRecordNAPTR(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_naptr" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.naptr.owner_name}"
		record_type = "NAPTR"
	}
	`, testAccResourceRecordNAPTR(zoneName), strings.TrimSuffix(zoneName, "."))
}

func testAccDataSourceRecordSSHFP(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_sshfp" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.sshfp.owner_name}"
		record_type = "44"
	}
	`, testAccResourceRecordSSHFP(zoneName), zoneName)
}

func testAccDataSourceRecordTLSA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_tlsa" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.tlsa.owner_name}"
		record_type = "TLSA"
	}
	`, testAccResourceRecordTLSA(zoneName), zoneName)
}

func testAccDataSourceRecordSPF(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_spf" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.spf.owner_name}"
		record_type = "99"
	}
	`, testAccResourceRecordSPF(zoneName), strings.TrimSuffix(zoneName, "."))
}

func testAccDataSourceRecordCAA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_caa" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.caa.owner_name}"
		record_type = "CAA"
	}
	`, testAccResourceRecordCAA(zoneName), strings.TrimSuffix(zoneName, "."))
}

func testAccDataSourceRecordAPEXALIAS(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_record" "data_apex" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_record.apex.owner_name}"
		record_type = "65282"
	}
	`, testAccResourceRecordAPEXALIAS(zoneName), zoneName)
}
