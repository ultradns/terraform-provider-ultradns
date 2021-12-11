package zone_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/test"
)

func TestAccZoneDataSource(t *testing.T) {
	zoneName := fmt.Sprintf("test-acc-%s.com.", test.GetRandomString())
	tc := resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceZone(zoneName, acctest.TestUsername),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists("ultradns_zone.primarydata"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "limit", "1"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "returned_count", "1"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.#", "1"),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.0.name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.0.account_name", acctest.TestUsername),
					resource.TestCheckResourceAttr("data.ultradns_zone.testzone", "zones.0.type", "PRIMARY"),
				),
			},
		},
	}
	resource.Test(t, tc)
}

func testAccDataSourceZone(zoneName, accountName string) string {
	return fmt.Sprintf(`
	data "ultradns_zone" "testzone" {
		limit = 1
		reverse = true
		query = "name:%s"
	}
	`, zoneName)
}
