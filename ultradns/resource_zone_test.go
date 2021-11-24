package ultradns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func TestAccZone(t *testing.T) {

	resourceName := "ultradns_zone.primary"
	tc := resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccZonePrimary_create(testZoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", testZoneName),
					resource.TestCheckResourceAttr(resourceName, "account_name", "teamrest"),
					resource.TestCheckResourceAttr(resourceName, "type", "PRIMARY"),
					testAccCheckZonePrimary_create(resourceName),
				),
			},
		},
	}
	resource.Test(t, tc)
}

func testAccCheckZonePrimary_create(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		client := testAccProvider.Meta().(*ultradns.Client)
		_, zoneType, zoneResponse, err := client.ReadZone(rs.Primary.ID)

		if err != nil {
			return err
		}

		if zoneResponse.Properties.Name != rs.Primary.ID {
			return fmt.Errorf("zone name mismactched expected : %v - returned : %v", rs.Primary.ID, zoneResponse.Properties.Name)
		}

		if zoneTypeExpected, ok := rs.Primary.Attributes["type"]; !ok || zoneTypeExpected != zoneType {
			return fmt.Errorf("zone type mismactched expected : %v - returned : %v", zoneTypeExpected, zoneType)
		}

		return nil
	}
}

func testAccZonePrimary_create(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "primary" {
		name        = "%s"
		account_name = "teamrest"
		type        = "PRIMARY"
		primary_create_info {
			create_type = "NEW"
		}
	}
	`, zoneName)
}
