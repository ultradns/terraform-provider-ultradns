package ultradns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccZoneDataSource(t *testing.T) {
	zoneName := fmt.Sprintf("test-acc-%s.com.", acctest.RandString(5))
	tc := resource.TestCase{
		PreCheck:     func() { TestAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceZone(zoneName, testUsername),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists("ultradns_zone.primarydata"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "limit", "1"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "returned_count", "1"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.#", "1"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.0.name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.0.account_name", testUsername),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.0.type", "PRIMARY"),
				),
			},
		},
	}
	resource.Test(t, tc)
}

func testAccDataSourceZone(zoneName, accountName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "primarydata" {
		name        = "%s"
		account_name = "%s"
		type        = "PRIMARY"
		primary_create_info {
			create_type = "NEW"
		}
	}
	data "ultradns_zone" "testzone" {
		limit = 1
		reverse = true
		query = "name:${ultradns_zone.primarydata.id}"
	}
	`, zoneName, accountName)
}
