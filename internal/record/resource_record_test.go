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
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_record", ""),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRecordA(zoneName, ownerNameTypeA),
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
				Config: testAccResourceUpdateRecordA(zoneName, ownerNameTypeA),
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
				Config: testAccResourceRecordCNAME(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.cname", ""),
					resource.TestCheckResourceAttr("ultradns_record.cname", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.cname", "record_type", "CNAME"),
					resource.TestCheckResourceAttr("ultradns_record.cname", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.cname", "record_data.0", "example.com."),
				),
			},
			{
				Config: testAccResourceRecordPTR(zoneName),
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
				Config: testAccResourceRecordTXT(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.txt", ""),
					resource.TestCheckResourceAttr("ultradns_record.txt", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.txt", "record_type", "TXT"),
					resource.TestCheckResourceAttr("ultradns_record.txt", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.txt", "record_data.0", "example.com."),
				),
			},
			{
				Config: testAccResourceRecordAAAA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.aaaa", ""),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "record_data.0", "2001:db8:85a3:0:0:8a2e:370:7334"),
				),
			},
			{
				Config: testAccResourceRecordSRV(zoneName),
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
				Config: testAccResourceRecordAPEXALIAS(zoneName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.apex", ""),
					resource.TestCheckResourceAttr("ultradns_record.apex", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.apex", "record_type", "APEXALIAS"),
					resource.TestCheckResourceAttr("ultradns_record.apex", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_record.apex", "record_data.0", "example.com."),
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

func testAccResourceUpdateRecordA(zoneName, ownerName string) string {
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
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "APEXALIAS"
		ttl = 800
		record_data = ["example.com."]
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), strings.TrimSuffix(zoneName, "."), tfacctest.RandString(3))
}
