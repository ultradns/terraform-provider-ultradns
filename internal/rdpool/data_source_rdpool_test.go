package rdpool_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
)

func TestAccDataSourceRDPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()

	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckRDPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRDPoolA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRDPoolExists("data.ultradns_rdpool.data_rdpool_a"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_rdpool_a", "zone_name", strings.TrimSuffix(zoneName, ".")),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_rdpool_a", "record_type", "A"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_rdpool_a", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_rdpool_a", "record_data.0", "192.168.1.1"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_rdpool_a", "order", "FIXED"),
				),
			},
			{
				Config: testAccDataSourceRDPoolAAAA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRDPoolExists("data.ultradns_rdpool.data_rdpool_aaaa"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_rdpool_aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_rdpool_aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_rdpool_aaaa", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_rdpool_aaaa", "record_data.0", "2001:db8:85a3:0:0:8a2e:370:7334"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_rdpool_aaaa", "order", "ROUND_ROBIN"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func testAccDataSourceRDPoolA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_rdpool" "data_rdpool_a" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_rdpool.a.owner_name}"
		record_type = "A"
	}
	`, testAccResourceRDPoolA(zoneName), strings.TrimSuffix(zoneName, "."))
}

func testAccDataSourceRDPoolAAAA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_rdpool" "data_rdpool_aaaa" {
		zone_name = "%s"
		owner_name = "${resource.ultradns_rdpool.aaaa.owner_name}"
		record_type = "AAAA"
	}
	`, testAccResourceRDPoolAAAA(zoneName), zoneName)
}
