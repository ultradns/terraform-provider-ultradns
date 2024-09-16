package slbpool_test

import (
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

func TestAccDataSourceSLBPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerNameTypeA := tfacctest.RandString(3)
	ownerNameTypeAAAA := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_slbpool", pool.SLB),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_slbpool",
					"a",
					zoneName,
					strings.ToUpper(ownerNameTypeA+"."+zoneName),
					"1",
					testAccResourceSLBPoolA(strings.ToUpper(zoneName), ownerNameTypeA),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_slbpool.data_a", pool.SLB),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "record_type", "A"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "ttl", "120"),
					// resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "status", "OK"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "serving_preference", "AUTO_SELECT"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "response_method", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "region_failure_sensitivity", "HIGH"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "pool_description", "SLB Pool Resource of Type A"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "monitor.0.method", "POST"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "monitor.0.search_string", "test"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "monitor.0.transmitted_data", "foo=bar"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "all_fail_record.0.rdata", "192.168.1.6"),
					// resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "all_fail_record.0.serving", "false"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "all_fail_record.0.description", "All Fail Record"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_a", "rdata_info.#", "5"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_slbpool",
					"aaaa",
					strings.ToUpper(strings.TrimSuffix(zoneName, ".")),
					strings.ToUpper(ownerNameTypeAAAA),
					"AAAA",
					testAccResourceSLBPoolAAAA(zoneName, ownerNameTypeAAAA),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_slbpool.data_aaaa", pool.SLB),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "owner_name", ownerNameTypeAAAA+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "ttl", "120"),
					// resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "status", "OK"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "serving_preference", "SERVE_ALL_FAIL"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "response_method", "PRIORITY_HUNT"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "region_failure_sensitivity", "LOW"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "pool_description", "SLB Pool Resource of Type AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "monitor.0.method", "GET"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "monitor.0.search_string", ""),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "monitor.0.transmitted_data", ""),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "all_fail_record.0.rdata", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"),
					// resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "all_fail_record.0.serving", "false"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "all_fail_record.0.description", "All fail record"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "rdata_info.#", "1"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "rdata_info.0.rdata", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:2222"),
					// resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "rdata_info.0.available_to_serve", "true"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "rdata_info.0.forced_state", "FORCED_ACTIVE"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "rdata_info.0.probing_enabled", "false"),
					resource.TestCheckResourceAttr("data.ultradns_slbpool.data_aaaa", "rdata_info.0.description", "RData of type AAAA"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}
