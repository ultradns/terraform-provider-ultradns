package record_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/record"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
)

func TestAccResourceRecord(t *testing.T) {
	zoneName := fmt.Sprintf("test-acc-%s.com.", tfacctest.RandString(5))

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
				Config: testAccResourceRecordAAAA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.aaaa"),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.aaaa", "record_data.0", "2001:db8:85a3:0:0:8a2e:370:7334"),
				),
			},
			{
				ResourceName:      "ultradns_record.aaaa",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordCNAME(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.cname"),
					resource.TestCheckResourceAttr("ultradns_record.cname", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.cname", "record_type", "CNAME"),
					resource.TestCheckResourceAttr("ultradns_record.cname", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.cname", "record_data.0", "google.com."),
				),
			},
			{
				ResourceName:      "ultradns_record.cname",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordMX(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.mx"),
					resource.TestCheckResourceAttr("ultradns_record.mx", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.mx", "record_type", "MX"),
					resource.TestCheckResourceAttr("ultradns_record.mx", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.mx", "record_data.0", "2 google.com."),
				),
			},
			{
				ResourceName:      "ultradns_record.mx",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordSRV(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.srv"),
					resource.TestCheckResourceAttr("ultradns_record.srv", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.srv", "record_type", "SRV"),
					resource.TestCheckResourceAttr("ultradns_record.srv", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.srv", "record_data.0", "5 6 7 google.com."),
				),
			},
			{
				ResourceName:      "ultradns_record.srv",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordTXT(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.txt"),
					resource.TestCheckResourceAttr("ultradns_record.txt", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.txt", "record_type", "TXT"),
					resource.TestCheckResourceAttr("ultradns_record.txt", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.txt", "record_data.0", "google.com."),
				),
			},
			{
				ResourceName:      "ultradns_record.txt",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRecordPTR(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists("ultradns_record.ptr"),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "record_type", "PTR"),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_record.ptr", "record_data.0", "google.com."),
				),
			},
			{
				ResourceName:      "ultradns_record.ptr",
				ImportState:       true,
				ImportStateVerify: true,
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
		rrSetKey := record.GetRRSetKey(rs.Primary.ID)
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
		rrSetKey := record.GetRRSetKey(rs.Primary.ID)
		_, recordResponse, err := services.RecordService.ReadRecord(rrSetKey)

		if err == nil {
			if len(recordResponse.RRSets) > 0 && recordResponse.RRSets[0].OwnerName == rrSetKey.Name {
				return fmt.Errorf("record - %v not destroyed.", rs.Primary.ID)
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

func testAccResourceRecordAAAA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "aaaa" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "AAAA"
		ttl = 120
		record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordCNAME(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "cname" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "CNAME"
		ttl = 120
		record_data = ["google.com."]
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
		record_data = ["2 google.com."]
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
		record_data = ["5 6 7 google.com."]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordTXT(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "txt" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "TXT"
		ttl = 120
		record_data = ["google.com."]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRecordPTR(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_record" "ptr" {
		zone_name = "${resource.ultradns_zone.primary_record.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_record.id}"
		record_type = "PTR"
		ttl = 120
		record_data = ["google.com."]
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}
