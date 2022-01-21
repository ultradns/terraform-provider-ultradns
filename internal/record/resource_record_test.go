package record_test

import (
	"fmt"
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
)

func TestAccResourceRecord(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()

	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRecordA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.a"),
					resource.TestCheckResourceAttr("ultradns_record.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_record.a", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.a", "record_data.0", "192.168.1.1"),
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
					testAccCheckRecordExists("ultradns_record.cname"),
					resource.TestCheckResourceAttr("ultradns_record.cname", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("ultradns_record.cname", "record_type", "5"),
					resource.TestCheckResourceAttr("ultradns_record.cname", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.cname", "record_data.0", "example.com."),
				),
			},
			{
				Config: testAccResourceRecordPTR(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.ptr"),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "record_type", "PTR"),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "record_data.0", "example.com."),
				),
			},
			{
				ResourceName:      "ultradns_record.ptr",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordHINFO(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.hinfo"),
					resource.TestCheckResourceAttr("ultradns_record.hinfo", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.hinfo", "record_type", "13"),
					resource.TestCheckResourceAttr("ultradns_record.hinfo", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.hinfo", "record_data.0", "\"PC\" \"Linux\""),
				),
			},
			{
				Config: testAccResourceRecordMX(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.mx"),
					resource.TestCheckResourceAttr("ultradns_record.mx", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.mx", "record_type", "MX"),
					resource.TestCheckResourceAttr("ultradns_record.mx", "ttl", "120"),
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
					testAccCheckRecordExists("ultradns_record.txt"),
					resource.TestCheckResourceAttr("ultradns_record.txt", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.txt", "record_type", "16"),
					resource.TestCheckResourceAttr("ultradns_record.txt", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.txt", "record_data.0", "example.com."),
				),
			},
			{
				Config: testAccResourceRecordRP(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.rp"),
					resource.TestCheckResourceAttr("ultradns_record.rp", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.rp", "record_type", "RP"),
					resource.TestCheckResourceAttr("ultradns_record.rp", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.rp", "record_data.0", "test.example.com. example.128/134.123.178.178.in-addr.arpa."),
				),
			},
			{
				ResourceName:      "ultradns_record.rp",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordAAAA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.aaaa"),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "record_type", "28"),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "record_data.0", "2001:db8:85a3:0:0:8a2e:370:7334"),
				),
			},
			{
				Config: testAccResourceRecordSRV(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.srv"),
					resource.TestCheckResourceAttr("ultradns_record.srv", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.srv", "record_type", "SRV"),
					resource.TestCheckResourceAttr("ultradns_record.srv", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.srv", "record_data.0", "5 6 7 example.com."),
				),
			},
			{
				ResourceName:      "ultradns_record.srv",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordNAPTR(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.naptr"),
					resource.TestCheckResourceAttr("ultradns_record.naptr", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.naptr", "record_type", "35"),
					resource.TestCheckResourceAttr("ultradns_record.naptr", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.naptr", "record_data.0", "1 2 \"3\" \"test\" \"\" test.com."),
				),
			},
			{
				Config: testAccResourceRecordSSHFP(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.sshfp"),
					resource.TestCheckResourceAttr("ultradns_record.sshfp", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.sshfp", "record_type", "SSHFP"),
					resource.TestCheckResourceAttr("ultradns_record.sshfp", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.sshfp", "record_data.0", "1 2 54B5E539EAF593AEA410F80737530B71CCDE8B6C3D241184A1372E98BC7EDB37"),
				),
			},
			{
				ResourceName:      "ultradns_record.sshfp",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordTLSA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.tlsa"),
					resource.TestCheckResourceAttr("ultradns_record.tlsa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.tlsa", "record_type", "52"),
					resource.TestCheckResourceAttr("ultradns_record.tlsa", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.tlsa", "record_data.0", "0 0 0 aaaaaaaa"),
				),
			},
			{
				Config: testAccResourceRecordSPF(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.spf"),
					resource.TestCheckResourceAttr("ultradns_record.spf", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.spf", "record_type", "SPF"),
					resource.TestCheckResourceAttr("ultradns_record.spf", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.spf", "record_data.0", "v=spf1 ip4:1.2.3.4 ~all"),
				),
			},
			{
				ResourceName:      "ultradns_record.spf",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordCAA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.caa"),
					resource.TestCheckResourceAttr("ultradns_record.caa", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("ultradns_record.caa", "record_type", "257"),
					resource.TestCheckResourceAttr("ultradns_record.caa", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.caa", "record_data.0", "1 issue \"test\""),
				),
			},
			{
				Config: testAccResourceRecordAPEXALIAS(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.apex"),
					resource.TestCheckResourceAttr("ultradns_record.apex", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("ultradns_record.apex", "record_type", "APEXALIAS"),
					resource.TestCheckResourceAttr("ultradns_record.apex", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.apex", "record_data.0", "example.com."),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func testAccCheckRecordExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return errors.ResourceNotFoundError(resourceName)
		}

		services := acctest.TestAccProvider.Meta().(*service.Service)
		rrSetKey := rrset.GetRRSetKeyFromID(rs.Primary.ID)
		_, _, err := services.RecordService.ReadRecord(rrSetKey)

		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckRecordDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ultradns_record" {
			continue
		}

		services := acctest.TestAccProvider.Meta().(*service.Service)
		rrSetKey := rrset.GetRRSetKeyFromID(rs.Primary.ID)
		_, recordResponse, err := services.RecordService.ReadRecord(rrSetKey)

		if err == nil {
			if len(recordResponse.RRSets) > 0 && recordResponse.RRSets[0].OwnerName == rrSetKey.Name {
				return errors.ResourceNotDestroyedError(rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccResourceZonePrimary(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "primary_record" {
		name        = "%s"
		account_name = "%s"
		type        = "PRIMARY"
		primary_create_info {
			create_type = "NEW"
		}
	}
	`, zoneName, acctest.TestUsername)
}

func testAccResourceRecordA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "a" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "A"
		ttl = 120
		record_data = ["192.168.1.1"]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordCNAME(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "cname" {
		zone_name = "%s"
		owner_name = "%s"
		record_type = "5"
		ttl = 120
		record_data = ["example.com."]
	}
	`, testAccResourceZonePrimary(zoneName), strings.TrimSuffix(zoneName, "."), tfacctest.RandString(3))
}

func testAccResourceRecordPTR(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "ptr" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "PTR"
		ttl = 120
		record_data = ["example.com."]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordHINFO(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "hinfo" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s"
		record_type = "13"
		ttl = 120
		record_data = ["\"PC\" \"Linux\""]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordMX(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "mx" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "MX"
		ttl = 120
		record_data = ["2 example.com."]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordTXT(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "txt" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s"
		record_type = "16"
		ttl = 120
		record_data = ["example.com."]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordRP(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "rp" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "RP"
		ttl = 120
		record_data = ["test.example.com. example.128/134.123.178.178.in-addr.arpa."]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordAAAA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "aaaa" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s"
		record_type = "28"
		ttl = 120
		record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordSRV(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "srv" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "SRV"
		ttl = 120
		record_data = ["5 6 7 example.com."]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordNAPTR(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "naptr" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s"
		record_type = "35"
		ttl = 120
		record_data = ["1 2 \"3\" \"test\" \"\" test.com."]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordSSHFP(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "sshfp" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "SSHFP"
		ttl = 120
		record_data = ["1 2 54B5E539EAF593AEA410F80737530B71CCDE8B6C3D241184A1372E98BC7EDB37"]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordTLSA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "tlsa" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "_23._tcp.%s"
		record_type = "52"
		ttl = 120
		record_data = ["0 0 0 aaaaaaaa"]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordSPF(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "spf" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "SPF"
		ttl = 120
		record_data = ["v=spf1 ip4:1.2.3.4 ~all"]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordCAA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "caa" {
		zone_name = "%s"
		owner_name = "%s"
		record_type = "257"
		ttl = 120
		record_data = ["1 issue \"test\""]
	}
	`, testAccResourceZonePrimary(zoneName), strings.TrimSuffix(zoneName, "."), tfacctest.RandString(3))
}

func testAccResourceRecordAPEXALIAS(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "apex" {
		zone_name = "%s"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "APEXALIAS"
		ttl = 120
		record_data = ["example.com."]
	}
	`, testAccResourceZonePrimary(zoneName), strings.TrimSuffix(zoneName, "."), tfacctest.RandString(3))
}
