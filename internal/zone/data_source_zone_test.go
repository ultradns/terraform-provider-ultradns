package zone_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
)

func TestAccZoneDataSource(t *testing.T) {
	zoneName := fmt.Sprintf("test-acc-%s.com.", tfacctest.RandString(5))
	tc := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceZone(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists("ultradns_zone.primarydata"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "limit", "1"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "returned_count", "1"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.#", "1"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.0.name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.0.account_name", acctest.TestUsername),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.0.type", primaryZoneType),
				),
			},
		},
	}
	resource.Test(t, tc)
}

func testAccDataSourceZone(zoneName string) string {
	return fmt.Sprintf(`
	data "ultradns_zone" "testzone" {
		name = %s"
	}
	`, zoneName)
}
