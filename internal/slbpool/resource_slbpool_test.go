package slbpool_test

import (
	"fmt"
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

const zoneResourceName = "primary_slbpool"

func TestAccResourceSLBPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerNameTypeA := tfacctest.RandString(3)
	ownerNameTypeAAAA := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_slbpool", pool.SLB),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSLBPoolA(strings.ToUpper(zoneName), strings.ToUpper(ownerNameTypeA)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_slbpool.a", pool.SLB),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "serving_preference", "AUTO_SELECT"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "response_method", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "region_failure_sensitivity", "HIGH"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "pool_description", "SLB Pool Resource of Type A"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "monitor.0.method", "POST"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "monitor.0.search_string", "test"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "monitor.0.transmitted_data", "foo=bar"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "all_fail_record.0.rdata", "192.168.1.6"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "all_fail_record.0.serving", "false"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "all_fail_record.0.description", "All Fail Record"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "rdata_info.#", "5"),
				),
			},
			{
				Config: testAccResourceUpdateSLBPoolA(strings.ToUpper(zoneName), ownerNameTypeA),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_slbpool.a", pool.SLB),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "ttl", "150"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "serving_preference", "SERVE_PRIMARY"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "response_method", "RANDOM"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "region_failure_sensitivity", "LOW"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "pool_description", ""),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "monitor.0.method", "GET"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "monitor.0.search_string", "testing"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "monitor.0.transmitted_data", ""),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "all_fail_record.0.rdata", "192.168.1.2"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "all_fail_record.0.serving", "false"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "all_fail_record.0.description", ""),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "rdata_info.#", "1"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "rdata_info.0.rdata", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "rdata_info.0.available_to_serve", "true"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "rdata_info.0.forced_state", "FORCED_INACTIVE"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "rdata_info.0.probing_enabled", "false"),
					resource.TestCheckResourceAttr("ultradns_slbpool.a", "rdata_info.0.description", "RData of type A"),
				),
			},
			{
				ResourceName:      "ultradns_slbpool.a",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceSLBPoolAAAA(zoneName, strings.ToUpper(ownerNameTypeAAAA)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_slbpool.aaaa", pool.SLB),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "owner_name", ownerNameTypeAAAA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "serving_preference", "SERVE_ALL_FAIL"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "response_method", "PRIORITY_HUNT"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "region_failure_sensitivity", "LOW"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "pool_description", "SLB Pool Resource of Type AAAA"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "monitor.0.method", "GET"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "monitor.0.search_string", ""),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "monitor.0.transmitted_data", ""),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "all_fail_record.0.rdata", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "all_fail_record.0.serving", "false"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "all_fail_record.0.description", "All fail record"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.#", "1"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.0.rdata", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:2222"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.0.available_to_serve", "true"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.0.forced_state", "FORCED_ACTIVE"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.0.probing_enabled", "false"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.0.description", "RData of type AAAA"),
				),
			},
			{
				Config: testAccResourceUpdateSLBPoolAAAA(strings.ToUpper(zoneName), ownerNameTypeAAAA),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_slbpool.aaaa", pool.SLB),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "owner_name", ownerNameTypeAAAA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "ttl", "150"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "serving_preference", "AUTO_SELECT"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "response_method", "RANDOM"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "region_failure_sensitivity", "HIGH"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "pool_description", "Update SLB Pool Resource of Type AAAA"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "monitor.0.method", "POST"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "monitor.0.search_string", "test"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "monitor.0.transmitted_data", "foo=bar"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "all_fail_record.0.rdata", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:2222"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "all_fail_record.0.serving", "false"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "all_fail_record.0.description", "Update All fail record"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.#", "1"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.0.rdata", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.0.available_to_serve", "true"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.0.forced_state", "NOT_FORCED"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.0.probing_enabled", "true"),
					resource.TestCheckResourceAttr("ultradns_slbpool.aaaa", "rdata_info.0.description", "Update RData of type AAAA"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func testAccResourceSLBPoolA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_slbpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_slbpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_slbpool.id}"
		record_type = "1"
		ttl = 120
		region_failure_sensitivity = "HIGH"
		serving_preference = "AUTO_SELECT"
    	response_method = "ROUND_ROBIN"
		pool_description = "SLB Pool Resource of Type A"
		monitor{
			url = "%s"
			method = "POST"
			search_string = "test"
			transmitted_data = "foo=bar"
		}
		all_fail_record{
			rdata = "192.168.1.6"
			description = "All Fail Record"
		}
		rdata_info{
			description = "one"
			rdata = "192.168.1.1"
			probing_enabled = true
		}
		rdata_info{
			description = "two"
			rdata = "192.168.1.2"
			probing_enabled = true
		}
		rdata_info{
			description = "three"
			rdata = "192.168.1.3"
			probing_enabled = true
		}
		rdata_info{
			description = "four"
			rdata = "192.168.1.4"
			probing_enabled = false
		}
		rdata_info{
			description = "five"
			rdata = "192.168.1.5"
			probing_enabled = false
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName, strings.TrimSuffix(acctest.TestHost, "/"))
}

func testAccResourceUpdateSLBPoolA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_slbpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_slbpool.id}"
		owner_name = "%s"
		record_type = "A"
		ttl = 150
		region_failure_sensitivity = "LOW"
		serving_preference = "SERVE_PRIMARY"
    	response_method = "RANDOM"
		monitor{
			url = "%s"
			method = "GET"
			search_string = "testing"
		}
		all_fail_record{
			rdata = "192.168.1.2"
		}
		rdata_info{
			rdata = "192.168.1.1"
			forced_state = "FORCED_INACTIVE"
			probing_enabled = false
			description = "RData of type A"
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName, acctest.TestHost)
}

func testAccResourceSLBPoolAAAA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_slbpool" "aaaa" {
		zone_name = "${resource.ultradns_zone.primary_slbpool.id}"
		owner_name = "%s"
		record_type = "AAAA"
		ttl = 120
		region_failure_sensitivity = "LOW"
		serving_preference = "SERVE_ALL_FAIL"
    	response_method = "PRIORITY_HUNT"
		pool_description = "SLB Pool Resource of Type AAAA"
		monitor{
			url = "%s"
			method = "GET"
		}
		all_fail_record{
			rdata = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"
			description = "All fail record"
		}
		rdata_info{
			rdata = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:2222"
			forced_state = "FORCED_ACTIVE"
			probing_enabled = false
			description = "RData of type AAAA"
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName, acctest.TestHost)
}

func testAccResourceUpdateSLBPoolAAAA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_slbpool" "aaaa" {
		zone_name = "${resource.ultradns_zone.primary_slbpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_slbpool.id}"
		record_type = "28"
		ttl = 150
		region_failure_sensitivity = "HIGH"
		serving_preference = "AUTO_SELECT"
    	response_method = "RANDOM"
		pool_description = "Update SLB Pool Resource of Type AAAA"
		monitor{
			url = "%s"
			method = "POST"
			search_string = "test"
			transmitted_data = "foo=bar"
		}
		all_fail_record{
			rdata = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:2222"
			description = "Update All fail record"
		}
		rdata_info{
			rdata = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"
			forced_state = "NOT_FORCED"
			probing_enabled = true
			description = "Update RData of type AAAA"
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName, strings.TrimSuffix(acctest.TestHost, "/"))
}
