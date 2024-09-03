package record_test

import (
	"fmt"
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
)

const zoneResourceName = "primary_record"

func TestAccResourceRecord(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerNameTypeA := tfacctest.RandString(3)
	ownerNameTypeNS := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_record", ""),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRecordA(zoneName, strings.ToUpper(ownerNameTypeA)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.a", ""),
					resource.TestCheckResourceAttr("ultradns_record.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_record.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_record.a", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.a", "record_data.0", "192.168.1.1"),
				),
			},
			{
				Config: testAccUpdateResourceRecordA(strings.ToUpper(zoneName), ownerNameTypeA),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.a", ""),
					resource.TestCheckResourceAttr("ultradns_record.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_record.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_record.a", "ttl", "850"),
					resource.TestCheckResourceAttr("ultradns_record.a", "record_data.0", "192.168.1.2"),
				),
			},
			{
				ResourceName:      "ultradns_record.a",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordNS(strings.ToUpper(zoneName)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.ns", ""),
					resource.TestCheckResourceAttr("ultradns_record.ns", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ns", "owner_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ns", "record_type", "NS"),
					resource.TestCheckResourceAttr("ultradns_record.ns", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.ns", "record_data.0", "ns11.example.com."),
				),
			},
			{
				Config: testAccUpdateResourceRecordNS(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.ns", ""),
					resource.TestCheckResourceAttr("ultradns_record.ns", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ns", "owner_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ns", "record_type", "NS"),
					resource.TestCheckResourceAttr("ultradns_record.ns", "ttl", "900"),
					resource.TestCheckResourceAttr("ultradns_record.ns", "record_data.#", "2"),
				),
			},
			{
				ResourceName:            "ultradns_record.ns",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"record_data"},
			},
			{
				Config: testAccResourceRecordNSwithOwner(strings.ToUpper(zoneName), ownerNameTypeNS),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.ns_owner", ""),
					resource.TestCheckResourceAttr("ultradns_record.ns_owner", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ns_owner", "owner_name", ownerNameTypeNS+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ns_owner", "record_type", "NS"),
					resource.TestCheckResourceAttr("ultradns_record.ns_owner", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.ns_owner", "record_data.#", "2"),
				),
			},
			{
				ResourceName:      "ultradns_record.ns_owner",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordCNAME(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.cname", ""),
					resource.TestCheckResourceAttr("ultradns_record.cname", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.cname", "record_type", "CNAME"),
					resource.TestCheckResourceAttr("ultradns_record.cname", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.cname", "record_data.0", "example.com."),
				),
			},
			// {
			// 	Config: testAccResourceRecordSOA(strings.ToUpper(zoneName)),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		acctest.TestAccCheckRecordResourceExists("ultradns_record.soa", ""),
			// 		resource.TestCheckResourceAttr("ultradns_record.soa", "zone_name", zoneName),
			// 		resource.TestCheckResourceAttr("ultradns_record.soa", "owner_name", zoneName),
			// 		resource.TestCheckResourceAttr("ultradns_record.soa", "record_type", "SOA"),
			// 		resource.TestCheckResourceAttr("ultradns_record.soa", "ttl", "800"),
			// 		resource.TestCheckResourceAttr("ultradns_record.soa", "record_data.0", "udns1.ultradns.net. sample@example.com. 10800 3600 2592000 10800"),
			// 	),
			// },
			// {
			// 	ResourceName:      "ultradns_record.soa",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
			// {
			// 	Config: testAccUpdateResourceRecordSOA(zoneName),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		acctest.TestAccCheckRecordResourceExists("ultradns_record.soa", ""),
			// 		resource.TestCheckResourceAttr("ultradns_record.soa", "zone_name", zoneName),
			// 		resource.TestCheckResourceAttr("ultradns_record.soa", "owner_name", zoneName),
			// 		resource.TestCheckResourceAttr("ultradns_record.soa", "record_type", "SOA"),
			// 		resource.TestCheckResourceAttr("ultradns_record.soa", "ttl", "800"),
			// 		resource.TestCheckResourceAttr("ultradns_record.soa", "record_data.0", "udns1.ultradns.net. test.sample@example.com. 10800 3600 2592000 10800"),
			// 	),
			// },
			{
				Config: testAccResourceRecordPTR(strings.ToUpper(zoneName)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.ptr", ""),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "record_type", "PTR"),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "record_data.0", "example.com."),
				),
			},
			{
				ResourceName:      "ultradns_record.ptr",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordMX(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.mx", ""),
					resource.TestCheckResourceAttr("ultradns_record.mx", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.mx", "record_type", "MX"),
					resource.TestCheckResourceAttr("ultradns_record.mx", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.mx", "record_data.0", "2 example.com."),
				),
			},
			{
				ResourceName:      "ultradns_record.mx",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordTXT(strings.ToUpper(zoneName)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.txt", ""),
					resource.TestCheckResourceAttr("ultradns_record.txt", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.txt", "record_type", "TXT"),
					resource.TestCheckResourceAttr("ultradns_record.txt", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.txt", "record_data.0", "example.com."),
				),
			},
			{
				Config: testAccResourceRecordAAAA(strings.ToUpper(zoneName)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.aaaa", ""),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "record_data.0", "2001:db8:85a3:0:0:8a2e:370:7334"),
				),
			},
			{
				Config: testAccResourceRecordSRV(strings.ToUpper(zoneName)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.srv", ""),
					resource.TestCheckResourceAttr("ultradns_record.srv", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.srv", "record_type", "SRV"),
					resource.TestCheckResourceAttr("ultradns_record.srv", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.srv", "record_data.0", "5 6 7 example.com."),
				),
			},
			{
				ResourceName:      "ultradns_record.srv",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordSSHFP(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.sshfp", ""),
					resource.TestCheckResourceAttr("ultradns_record.sshfp", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.sshfp", "record_type", "SSHFP"),
					resource.TestCheckResourceAttr("ultradns_record.sshfp", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.sshfp", "record_data.0", "1 2 54B5E539EAF593AEA410F80737530B71CCDE8B6C3D241184A1372E98BC7EDB37"),
				),
			},
			{
				ResourceName:      "ultradns_record.sshfp",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordAPEXALIAS(strings.ToUpper(zoneName)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.apex", ""),
					resource.TestCheckResourceAttr("ultradns_record.apex", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.apex", "record_type", "APEXALIAS"),
					resource.TestCheckResourceAttr("ultradns_record.apex", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.apex", "record_data.0", "example.com."),
				),
			},
			{
				Config: testAccResourceRecordDS(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.ds", ""),
					resource.TestCheckResourceAttr("ultradns_record.ds", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ds", "owner_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ds", "record_type", "DS"),
					resource.TestCheckResourceAttr("ultradns_record.ds", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.ds", "record_data.0", "25286 1 1 340437DC66C3DFAD0B3E849740D2CF1A4151671D"),
				),
			},
			{
				Config: testAccResourceRecordCAA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.caa", ""),
					resource.TestCheckResourceAttr("ultradns_record.caa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.caa", "owner_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.caa", "record_type", "CAA"),
					resource.TestCheckResourceAttr("ultradns_record.caa", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.caa", "record_data.0", "0 issue ultradns"),
				),
			},
			{
				ResourceName:      "ultradns_record.caa",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordHTTPS(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.https", ""),
					resource.TestCheckResourceAttr("ultradns_record.https", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.https", "record_type", "HTTPS"),
					resource.TestCheckResourceAttr("ultradns_record.https", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.https", "record_data.0", "1 www.ultradns.com. ech=dGVzdA== mandatory=alpn,key65444 no-default-alpn port=8080 ipv4hint=1.2.3.4,9.8.7.6 key65444=privateKeyTesting ipv6hint=2001:db8:3333:4444:5555:6666:7777:8888,2001:db8:3333:4444:cccc:dddd:eeee:ffff alpn=h3,h3-29,h2"),
				),
			},
			{
				ResourceName:      "ultradns_record.https",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordSVCB(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.svcb", ""),
					resource.TestCheckResourceAttr("ultradns_record.svcb", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.svcb", "record_type", "SVCB"),
					resource.TestCheckResourceAttr("ultradns_record.svcb", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.svcb", "record_data.0", "0 www.ultradns.com."),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func testAccResourceRecordA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "a" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "A"
		ttl = 800
		record_data = ["192.168.1.1"]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccUpdateResourceRecordA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "a" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s"
		record_type = "1"
		ttl = 850
		record_data = ["192.168.1.2"]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccResourceRecordNS(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "ns" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "${resource.ultradns_zone.primary_record.id}"
		record_type = "NS"
		ttl = 800
		record_data = ["ns11.example.com."]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName))
}

func testAccUpdateResourceRecordNS(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "ns" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "${resource.ultradns_zone.primary_record.id}"
		record_type = "NS"
		ttl = 900
		record_data = ["ns12.example.com.","ns13.example.com."]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName))
}

func testAccResourceRecordNSwithOwner(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "ns_owner" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "2"
		ttl = 800
		record_data = ["ns12.example.com.","ns13.example.com."]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccResourceRecordCNAME(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "cname" {
		zone_name = "%s"
		owner_name = "%s"
		record_type = "5"
		ttl = 800
		record_data = ["example.com."]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), strings.TrimSuffix(zoneName, "."), tfacctest.RandString(3))
}

func testAccResourceRecordSOA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "soa" {
		zone_name = "%s"
		owner_name = "%s"
		record_type = "6"
		ttl = 800
		record_data = ["udns1.ultradns.net. sample@example.com. 10800 3600 2592000 10800"]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), strings.TrimSuffix(zoneName, "."), zoneName)
}

func testAccUpdateResourceRecordSOA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "soa" {
		zone_name = "%s"
		owner_name = "%s"
		record_type = "SOA"
		ttl = 800
		record_data = ["udns1.ultradns.net. test.sample@example.com. 10800 3600 2592000 10800"]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), zoneName, strings.TrimSuffix(zoneName, "."))
}

func testAccResourceRecordPTR(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "ptr" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "PTR"
		ttl = 800
		record_data = ["example.com."]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordMX(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "mx" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "MX"
		ttl = 800
		record_data = ["2 example.com."]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordTXT(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "txt" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s"
		record_type = "16"
		ttl = 800
		record_data = ["example.com."]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordAAAA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "aaaa" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s"
		record_type = "28"
		ttl = 800
		record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordSRV(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "srv" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "SRV"
		ttl = 800
		record_data = ["5 6 7 example.com."]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordSSHFP(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "sshfp" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "SSHFP"
		ttl = 800
		record_data = ["1 2 54B5E539EAF593AEA410F80737530B71CCDE8B6C3D241184A1372E98BC7EDB37"]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordAPEXALIAS(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "apex" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_zone.primary_record.id}"
		record_type = "APEXALIAS"
		ttl = 800
		record_data = ["example.com."]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), strings.TrimSuffix(zoneName, "."))
}

func testAccResourceRecordDS(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "ds" {
		zone_name = "%s"
		owner_name = "%s"
		record_type = "DS"
		ttl = 800
		record_data = ["25286 1 1 340437DC66C3DFAD0B3E849740D2CF1A4151671D"]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), zoneName, zoneName)
}

func testAccResourceRecordCAA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "caa" {
		zone_name = "%s"
		owner_name = "%s"
		record_type = "CAA"
		ttl = 800
		record_data = ["0 issue ultradns"]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), zoneName, zoneName)
}

func testAccResourceRecordHTTPS(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "https" {
		zone_name = "%s"
		owner_name = "%s"
		record_type = "HTTPS"
		ttl = 800
		record_data = ["1 www.ultradns.com. ech=dGVzdA== mandatory=alpn,key65444 no-default-alpn port=8080 ipv4hint=1.2.3.4,9.8.7.6 key65444=privateKeyTesting ipv6hint=2001:db8:3333:4444:5555:6666:7777:8888,2001:db8:3333:4444:cccc:dddd:eeee:ffff alpn=h3,h3-29,h2"]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), zoneName, tfacctest.RandString(3))
}

func testAccResourceRecordSVCB(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "svcb" {
		zone_name = "%s"
		owner_name = "%s"
		record_type = "SVCB"
		ttl = 800
		record_data = ["0 www.ultradns.com."]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), zoneName, tfacctest.RandString(3))
}
