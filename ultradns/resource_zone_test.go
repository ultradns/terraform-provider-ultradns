package ultradns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func TestAccZone(t *testing.T) {
	var primaryZone ultradns.ZoneResponse
	resourceName := "ultradns_zone.primary"
	tc := resource.TestCase{
		PreCheck:  func() { TestAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccZonePrimary_create_new(testZoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", testZoneName),
					resource.TestCheckResourceAttr(resourceName, "account_name", "teamrest"),
					resource.TestCheckResourceAttr(resourceName, "type", "PRIMARY"),
					testAccCheckZonePrimary(resourceName, &primaryZone),
					testAccCheckZonePrimary_tsig(resourceName, &primaryZone),
					testAccCheckZonePrimary_notify_address(resourceName, &primaryZone),
					estAccCheckZonePrimary_restrict_ip(resourceName, &primaryZone),
				),
			},
		},
	}
	resource.Test(t, tc)
}

func testAccCheckZonePrimary(n string, primaryZone *ultradns.ZoneResponse) resource.TestCheckFunc {
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

		primaryZone = zoneResponse

		return nil
	}
}

func testAccCheckZonePrimary_tsig(n string, primaryZone *ultradns.ZoneResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		return nil
	}
}

func testAccCheckZonePrimary_notify_address(n string, primaryZone *ultradns.ZoneResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return nil
	}
}

func estAccCheckZonePrimary_restrict_ip(n string, primaryZone *ultradns.ZoneResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return nil
	}
}

func testAccZonePrimary_create_new(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "primary" {
		name        = "%s"
		account_name = "teamrest"
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
	`, zoneName)
}
