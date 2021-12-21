package zone_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
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
			return helper.ResourceNotFoundError(resourceName)
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
				return fmt.Errorf("zone %v not destroyed.\n", rs.Primary.ID)
			}
		}
	}

	return nil
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

func testAccResourceZoneAlias(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "alias_primary" {
		name        = "%s"
		account_name = "%s"
		type        = "PRIMARY"
		primary_create_info {
			create_type = "NEW"
		}
	}

	resource "ultradns_zone" "alias" {
		name        = "%s"
		account_name = "%s"
		type        = "ALIAS"
		alias_create_info {
			  original_zone_name = "${resource.ultradns_zone.alias_primary.id}"
		}
	  }
	`, testAccGetPrimaryZoneNameForAlias(zoneName), acctest.TestUsername, zoneName, acctest.TestUsername)
}

func testAccGetPrimaryZoneNameForAlias(zoneName string) string {
	return fmt.Sprintf("%s/%sin-addr.arpa.", zoneName, zoneName)
}
