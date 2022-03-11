package rdpool_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

const zoneResourceName = "primary_rdpool"

func TestAccResourceRDPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerNameTypeA := tfacctest.RandString(3)
	ownerNameTypeAAAA := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_rdpool", pool.RD),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRDPoolA(zoneName, ownerNameTypeA),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_rdpool.a", pool.RD),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "record_data.0", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "order", "FIXED"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "description", "RD Pool Resource of Type A"),
				),
			},
			{
				Config: testAccResourceUpdateRDPoolA(zoneName, ownerNameTypeA),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_rdpool.a", pool.RD),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "ttl", "850"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "record_data.0", "192.168.1.2"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "order", "RANDOM"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "description", ownerNameTypeA+"."+zoneName),
				),
			},
			{
				ResourceName:      "ultradns_rdpool.a",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRDPoolAAAA(zoneName, ownerNameTypeAAAA),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_rdpool.aaaa", pool.RD),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "owner_name", ownerNameTypeAAAA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "record_data.0", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:2222"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "order", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "description", ownerNameTypeAAAA+"."+zoneName),
				),
			},
			{
				Config: testAccResourceUpdateRDPoolAAAA(zoneName, ownerNameTypeAAAA),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_rdpool.aaaa", pool.RD),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "owner_name", ownerNameTypeAAAA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "ttl", "850"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "record_data.0", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "order", "FIXED"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "description", "RD Pool Resource of Type AAAA"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func testAccResourceRDPoolA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_rdpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_rdpool.id}"
		owner_name = "%s"
		record_type = "1"
		ttl = 800
		record_data = ["192.168.1.1"]
		order = "FIXED"
		description = "RD Pool Resource of Type A"
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccResourceUpdateRDPoolA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_rdpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_rdpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_rdpool.id}"
		record_type = "A"
		ttl = 850
		record_data = ["192.168.1.2"]
		order = "RANDOM"
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccResourceRDPoolAAAA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_rdpool" "aaaa" {
		zone_name = "${resource.ultradns_zone.primary_rdpool.id}"
		owner_name = "%s"
		record_type = "AAAA"
		ttl = 800
		record_data = ["aaaa:bbbb:cccc:dddd:eeee:ffff:1111:2222"]
		order = "ROUND_ROBIN"
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccResourceUpdateRDPoolAAAA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_rdpool" "aaaa" {
		zone_name = "${resource.ultradns_zone.primary_rdpool.id}"
		owner_name = "%s"
		record_type = "28"
		ttl = 850
		record_data = ["aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"]
		order = "FIXED"
		description = "RD Pool Resource of Type AAAA"
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}
