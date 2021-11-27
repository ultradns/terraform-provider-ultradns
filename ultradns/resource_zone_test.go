package ultradns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func TestAccZone(t *testing.T) {
	zoneName := fmt.Sprintf("test-acc-%s.com.", acctest.RandString(5))
	tc := resource.TestCase{
		PreCheck:     func() { TestAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccZonePrimary_create_new(zoneName, testUsername),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ultradns_zone.primary", "name", zoneName),
					resource.TestCheckResourceAttr("ultradns_zone.primary", "account_name", testUsername),
					resource.TestCheckResourceAttr("ultradns_zone.primary", "type", "PRIMARY"),
					testAccCheckZoneExists("ultradns_zone.primary"),
				),
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

		client := testAccProvider.Meta().(*ultradns.Client)
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

		client := testAccProvider.Meta().(*ultradns.Client)
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
