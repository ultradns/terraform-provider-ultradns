package zone_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

const (
	zoneResourceName    = "primary"
	defaultCount        = "2"
	defaultZoneStatus   = "ACTIVE"
	defaultDNSSECStatus = "UNSIGNED"
)

var testNameServer = os.Getenv("ULTRADNS_UNIT_TEST_NAME_SERVER")

func TestAccResourceZonePrimary(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	resourceName := "ultradns_zone.primary"

	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", zoneName),
					resource.TestCheckResourceAttr(resourceName, "account_name", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "type", zone.Primary),
					resource.TestCheckResourceAttr(resourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(resourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(resourceName, "owner", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(resourceName, "primary_create_info.0.notify_addresses.#", defaultCount),
					resource.TestCheckResourceAttr(resourceName, "primary_create_info.0.restrict_ip.#", defaultCount),
				),
			},
			{
				Config: testAccResourceUpdateZonePrimary(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", zoneName),
					resource.TestCheckResourceAttr(resourceName, "account_name", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "type", zone.Primary),
					resource.TestCheckResourceAttr(resourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(resourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(resourceName, "owner", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(resourceName, "primary_create_info.0.notify_addresses.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "primary_create_info.0.restrict_ip.#", "0"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"primary_create_info.0.create_type"},
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func TestAccResourceZoneSecondary(t *testing.T) {
	zoneName := acctest.GetRandomSecondaryZoneName()
	resourceName := "ultradns_zone.secondary"

	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceZoneSecondary(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", zoneName),
					resource.TestCheckResourceAttr(resourceName, "account_name", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "type", zone.Secondary),
					resource.TestCheckResourceAttr(resourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(resourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(resourceName, "owner", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(resourceName, "secondary_create_info.0.primary_name_server_1.0.ip", testNameServer),
					resource.TestCheckResourceAttr(resourceName, "secondary_create_info.0.notification_email_address", "test@ultradns.com"),
					resource.TestCheckResourceAttr(resourceName, "transfer_status_details.0.last_refresh_status", "SUCCESSFUL"),
				),
			},
			{
				Config: testAccResourceUpdateZoneSecondary(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", zoneName),
					resource.TestCheckResourceAttr(resourceName, "account_name", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "type", zone.Secondary),
					resource.TestCheckResourceAttr(resourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(resourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(resourceName, "owner", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(resourceName, "secondary_create_info.0.primary_name_server_1.0.ip", testNameServer),
					resource.TestCheckResourceAttr(resourceName, "secondary_create_info.0.notification_email_address", "testing@ultradns.com"),
					resource.TestCheckResourceAttr(resourceName, "transfer_status_details.0.last_refresh_status", "SUCCESSFUL"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func TestAccResourceZoneAlias(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	primaryZoneName := acctest.GetRandomZoneNameWithSpecialChar()
	resourceName := "ultradns_zone.alias"

	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceZoneAlias(zoneName, primaryZoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", zoneName),
					resource.TestCheckResourceAttr(resourceName, "account_name", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "type", zone.Alias),
					resource.TestCheckResourceAttr(resourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(resourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(resourceName, "owner", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(resourceName, "alias_create_info.0.original_zone_name", primaryZoneName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func testAccCheckZoneExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return errors.ResourceNotFoundError(resourceName)
		}

		services := acctest.TestAccProvider.Meta().(*service.Service)
		_, _, err := services.ZoneService.ReadZone(rs.Primary.ID)

		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckZoneDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ultradns_zone" {
			continue
		}

		services := acctest.TestAccProvider.Meta().(*service.Service)
		_, zoneResponse, err := services.ZoneService.ReadZone(rs.Primary.ID)

		if err == nil {
			if zoneResponse.Properties != nil && zoneResponse.Properties.Name == rs.Primary.ID {
				return errors.ResourceNotDestroyedError(rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccResourceUpdateZonePrimary(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "primary" {
		name        = "%s"
		account_name = "%s"
		type        = "PRIMARY"
		primary_create_info {
			create_type = "NEW"
		
		}
	}
	`, zoneName, acctest.TestAccount)
}

func testAccResourceZoneSecondary(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "secondary" {
		name        = "%s"
		account_name = "%s"
		type        = "SECONDARY"
		secondary_create_info {
			notification_email_address = "test@ultradns.com"
			primary_name_server_1 {
				ip = "%s"
			} 
		}
	}
	`, strings.TrimSuffix(zoneName, "."), acctest.TestAccount, testNameServer)
}

func testAccResourceUpdateZoneSecondary(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "secondary" {
		name        = "%s"
		account_name = "%s"
		type        = "SECONDARY"
		secondary_create_info {
			notification_email_address = "testing@ultradns.com"
			primary_name_server_1 {
				ip = "%s"
			} 
		}
	}
	`, zoneName, acctest.TestAccount, testNameServer)
}

func testAccResourceZoneAlias(zoneName, primaryZoneName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_zone" "alias" {
		name        = "%s"
		account_name = "%s"
		type        = "ALIAS"
		alias_create_info {
			  original_zone_name = "${resource.ultradns_zone.primary.id}"
		}
	  }
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, primaryZoneName), zoneName, acctest.TestAccount)
}
