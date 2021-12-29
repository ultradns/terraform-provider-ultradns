package zone_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
)

const (
	primaryZoneType     = "PRIMARY"
	secondaryZoneType   = "SECONDARY"
	aliasZoneType       = "ALIAS"
	defaultCount        = "2"
	defaultZoneStatus   = "ACTIVE"
	defaultDNSSECStatus = "UNSIGNED"
)

var testNameServer = os.Getenv("ULTRADNS_UNIT_TEST_NAME_SERVER")

func TestAccResourceZonePrimary(t *testing.T) {
	zoneName := fmt.Sprintf("test-acc-%s.com.", tfacctest.RandString(5))
	resourceName := "ultradns_zone.primary"

	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceZonePrimary(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", zoneName),
					resource.TestCheckResourceAttr(resourceName, "account_name", acctest.TestUsername),
					resource.TestCheckResourceAttr(resourceName, "type", primaryZoneType),
					resource.TestCheckResourceAttr(resourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(resourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(resourceName, "owner", acctest.TestUsername),
					resource.TestCheckResourceAttr(resourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(resourceName, "primary_create_info.0.notify_addresses.#", defaultCount),
					resource.TestCheckResourceAttr(resourceName, "primary_create_info.0.restrict_ip.#", defaultCount),
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
	zoneName := fmt.Sprintf("test-acc-%s.com.", tfacctest.RandString(5))
	resourceName := "ultradns_zone.secondary"

	testCase := resource.TestCase{
		PreCheck: func() {
			acctest.TestPreCheck(t)
			acctest.CreateOxfrZone(zoneName)
		},
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroyForSecondaryZone(zoneName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceZoneSecondary(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr(resourceName, "account_name", acctest.TestUsername),
					resource.TestCheckResourceAttr(resourceName, "type", secondaryZoneType),
					resource.TestCheckResourceAttr(resourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(resourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(resourceName, "owner", acctest.TestUsername),
					resource.TestCheckResourceAttr(resourceName, "resource_record_count", defaultCount),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func TestAccResourceZoneAlias(t *testing.T) {
	zoneName := fmt.Sprintf("test-acc-%s.com.", tfacctest.RandString(5))
	resourceName := "ultradns_zone.alias"

	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceZoneAlias(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", zoneName),
					resource.TestCheckResourceAttr(resourceName, "account_name", acctest.TestUsername),
					resource.TestCheckResourceAttr(resourceName, "type", aliasZoneType),
					resource.TestCheckResourceAttr(resourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(resourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(resourceName, "owner", acctest.TestUsername),
					resource.TestCheckResourceAttr(resourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(resourceName, "alias_create_info.0.original_zone_name", testAccGetPrimaryZoneNameForAlias(zoneName)),
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

func testAccCheckZoneDestroyForSecondaryZone(zoneName string) func(s *terraform.State) error {
	acctest.DeleteOxfrZone(zoneName)

	return testAccCheckZoneDestroy
}

func testAccResourceZonePrimary(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "primary" {
		name        = "%s"
		account_name = "%s"
		type        = "PRIMARY"
		primary_create_info {
			create_type = "NEW"
			notify_addresses {
				notify_address = "192.168.1.1"
			}
			notify_addresses {
				notify_address = "192.168.1.2"
			}
			restrict_ip {
				single_ip = "192.168.1.3"
			}
			restrict_ip {
				single_ip = "192.168.1.4"
			}
		}
	}
	`, zoneName, acctest.TestUsername)
}

func testAccResourceZoneSecondary(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "secondary" {
		name        = "%s"
		account_name = "%s"
		type        = "SECONDARY"
		secondary_create_info {
			primary_name_server_1 {
				ip = "%s"
			} 
		}
	}
	`, strings.TrimSuffix(zoneName, "."), acctest.TestUsername, testNameServer)
}

func testAccResourceZoneAlias(zoneName string) string {
	primaryZoneNameForAlias := testAccGetPrimaryZoneNameForAlias(zoneName)

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
	`, testAccResourceZonePrimary(primaryZoneNameForAlias), zoneName, acctest.TestUsername)
}

func testAccGetPrimaryZoneNameForAlias(zoneName string) string {
	return fmt.Sprintf("%[1]s/%[1]sin-addr.arpa.", zoneName)
}
