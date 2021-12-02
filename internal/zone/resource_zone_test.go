package zone_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func TestAccZoneResource(t *testing.T) {
	zoneName := fmt.Sprintf("test-acc-%s.com.", tfacctest.RandString(5))
	tc := resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccZonePrimary_create_new(zoneName, acctest.TestUsername),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists("ultradns_zone.primary"),
					resource.TestCheckResourceAttr("ultradns_zone.primary", "name", zoneName),
					resource.TestCheckResourceAttr("ultradns_zone.primary", "account_name", acctest.TestUsername),
					resource.TestCheckResourceAttr("ultradns_zone.primary", "type", "PRIMARY"),
				),
			},
			{
				Config: testAccZoneSecondary_create_new("d100-permission.com.", acctest.TestUsername),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists("ultradns_zone.secondary"),
					resource.TestCheckResourceAttr("ultradns_zone.secondary", "name", "d100-permission.com."),
					resource.TestCheckResourceAttr("ultradns_zone.secondary", "account_name", acctest.TestUsername),
					resource.TestCheckResourceAttr("ultradns_zone.secondary", "type", "SECONDARY"),
				),
			},
			{
				Config: testAccZoneAlias_create_new(zoneName, acctest.TestUsername),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists("ultradns_zone.alias"),
					resource.TestCheckResourceAttr("ultradns_zone.alias", "name", zoneName),
					resource.TestCheckResourceAttr("ultradns_zone.alias", "account_name", acctest.TestUsername),
					resource.TestCheckResourceAttr("ultradns_zone.alias", "type", "ALIAS"),
				),
			},
			{
				Config: testAccZone_import(zoneName, acctest.TestUsername),
			},
			{
				ResourceName:     "ultradns_zone.importdata",
				ImportState:      true,
				ImportStateCheck: testAccZoneImportStateCheck(zoneName, acctest.TestUsername),
			},
		},
	}
	resource.Test(t, tc)
}

func testAccCheckZoneExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		client := acctest.TestAccProvider.Meta().(*ultradns.Client)
		_, zoneResponse, err := client.ReadZone(rs.Primary.ID)

		if err != nil {
			return err
		}

		if zoneResponse.Properties == nil {
			return fmt.Errorf("zone properties are nil")
		}

		if zoneResponse.Properties.Name != rs.Primary.ID {
			return fmt.Errorf("zone name mismactched expected : %v - returned : %v", rs.Primary.ID, zoneResponse.Properties.Name)
		}

		if zoneTypeExpected, ok := rs.Primary.Attributes["type"]; !ok || zoneTypeExpected != zoneResponse.Properties.Type {
			return fmt.Errorf("zone type mismactched expected : %v - returned : %v", zoneTypeExpected, zoneResponse.Properties.Type)
		}

		if zoneAccountExpected, ok := rs.Primary.Attributes["account_name"]; !ok || zoneAccountExpected != zoneResponse.Properties.AccountName {
			return fmt.Errorf("zone account mismactched expected : %v - returned : %v", zoneAccountExpected, zoneResponse.Properties.AccountName)
		}

		return nil
	}
}

func testAccCheckZoneDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ultradns_zone" {
			continue
		}

		client := acctest.TestAccProvider.Meta().(*ultradns.Client)
		res, zoneResponse, err := client.ReadZone(rs.Primary.ID)
		if err == nil {
			if zoneResponse.Properties != nil && zoneResponse.Properties.Name == rs.Primary.ID {
				return fmt.Errorf("zone %v not destroyed!", rs.Primary.ID)
			}
			return nil
		}
		if res.StatusCode != 404 {
			return err
		}
	}

	return nil
}

func testAccZoneImportStateCheck(zoneName, accountName string) func(is []*terraform.InstanceState) error {
	return func(is []*terraform.InstanceState) error {
		if len(is) > 0 {
			state := is[0]
			if zoneName != state.ID {
				return fmt.Errorf("zone name mismactched expected : %v - returned : %v", zoneName, state.ID)
			}
			if accountName != state.Attributes["account_name"] {
				return fmt.Errorf("zone account mismactched expected : %v - returned : %v", accountName, state.Attributes["account_name"])
			}
			if "PRIMARY" != state.Attributes["type"] {
				return fmt.Errorf("zone type mismactched expected : %v - returned : %v", "PRIMARY", state.Attributes["type"])
			}
			return nil
		}
		return fmt.Errorf("length of instance state is %v while checking import state", len(is))
	}
}

func testAccZonePrimary_create_new(zoneName, accountName string) string {
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
			tsig {
				tsig_key_name = "0-0-0-0-0antony.com.0.3276735349282751.key."
				tsig_key_value = "ZWFlY2U1MTBlNTZhYTRmM2Y0NGQ5MTlmYTdmZTE0Njc="
				tsig_algorithm  = "hmac-md5"
			}
			restrict_ip {
				single_ip = "192.168.1.3"
			}
			restrict_ip {
				single_ip = "192.168.1.4"
			}
		}
	}
	`, zoneName, accountName)
}

func testAccZoneSecondary_create_new(zoneName, accountName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "secondary" {
		name        = "%s"
		account_name = "%s"
		type        = "SECONDARY"
		secondary_create_info {
			primary_name_server_1 {
				ip = "e2e-bind-useast1a01-01.dev.ultradns.net"
			}
		}
	}
	`, zoneName, accountName)
}

func testAccZoneAlias_create_new(zoneName, accountName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "alias" {
		name        = "%s"
		account_name = "%s"
		type        = "ALIAS"
	  
		alias_create_info {
			  original_zone_name = "0-0-0-0-0antony.com."
		}
	  }
	`, zoneName, accountName)
}

func testAccZone_import(zoneName, accountName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "importdata" {
		name = "%s"
		account_name = "%s"
		type = "PRIMARY"
		primary_create_info {
			create_type = "NEW"
			notify_addresses {
				notify_address = "192.168.1.1"
			}
		}
	}
	`, zoneName, accountName)
}
