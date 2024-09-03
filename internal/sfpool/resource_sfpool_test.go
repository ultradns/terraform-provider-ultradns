package sfpool_test

import (
	"fmt"
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

const zoneResourceName = "primary_sfpool"

func TestAccResourceSFPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerNameTypeA := tfacctest.RandString(3)
	ownerNameTypeAAAA := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_sfpool", pool.SF),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSFPoolA(strings.ToUpper(zoneName), strings.ToUpper(ownerNameTypeA)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_sfpool.a", pool.SF),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "record_data.0", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "backup_record.0.rdata", "192.168.1.2"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "backup_record.0.description", "Type A backup record"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "live_record_state", "FORCED_INACTIVE"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "live_record_description", "Maintenance Activity"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "pool_description", "SF Pool Resource of Type A"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "region_failure_sensitivity", "HIGH"),
					// resource.TestCheckResourceAttr("ultradns_sfpool.a", "status", "MANUAL"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "monitor.0.method", "POST"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "monitor.0.search_string", "test"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "monitor.0.transmitted_data", "foo=bar"),
				),
			},
			{
				Config: testAccResourceUpdateSFPoolA(strings.ToUpper(zoneName), ownerNameTypeA),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_sfpool.a", pool.SF),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "ttl", "150"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "record_data.0", "192.168.1.2"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "backup_record.0.rdata", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "backup_record.0.description", ""),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "monitor.0.method", "GET"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "monitor.0.search_string", ""),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "monitor.0.transmitted_data", ""),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "pool_description", ""),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "region_failure_sensitivity", "LOW"),
					// resource.TestCheckResourceAttr("ultradns_sfpool.a", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "live_record_state", "NOT_FORCED"),
					resource.TestCheckResourceAttr("ultradns_sfpool.a", "live_record_description", ""),
				),
			},
			{
				Config: testAccResourceSFPoolAAAA(zoneName, strings.ToUpper(ownerNameTypeAAAA)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_sfpool.aaaa", pool.SF),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "owner_name", ownerNameTypeAAAA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "record_data.0", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:2222"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "monitor.0.method", "GET"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "monitor.0.search_string", "test"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "monitor.0.transmitted_data", ""),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "live_record_state", "NOT_FORCED"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "live_record_description", "Active"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "pool_description", "SF Pool Resource of Type AAAA"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "region_failure_sensitivity", "HIGH"),
					// resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "backup_record.0.rdata", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "backup_record.0.description", "Type AAAA Backup record"),
				),
			},
			{
				Config: testAccResourceUpdateSFPoolAAAA(strings.ToUpper(zoneName), strings.ToUpper(ownerNameTypeAAAA)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_sfpool.aaaa", pool.SF),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "owner_name", ownerNameTypeAAAA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "ttl", "150"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "record_data.0", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "backup_record.#", "0"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "live_record_state", "NOT_FORCED"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "live_record_description", "Maintenance Activity"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "pool_description", "Update SF Pool Resource of Type AAAA"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "region_failure_sensitivity", "LOW"),
					// resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "monitor.0.method", "POST"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "monitor.0.search_string", "testing"),
					resource.TestCheckResourceAttr("ultradns_sfpool.aaaa", "monitor.0.transmitted_data", ""),
				),
			},
			{
				ResourceName:            "ultradns_sfpool.aaaa",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"live_record_state"},
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func testAccResourceSFPoolA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_sfpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_sfpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_sfpool.id}"
		record_type = "A"
		ttl = 120
		record_data = ["192.168.1.1"]
		region_failure_sensitivity = "HIGH"
		live_record_state = "FORCED_INACTIVE"
		live_record_description = "Maintenance Activity"
		pool_description = "SF Pool Resource of Type A"
		monitor{
			url = "%s"
			method = "POST"
			search_string = "test"
			transmitted_data = "foo=bar"
		}
		backup_record{
			rdata = "192.168.1.2"
			description = "Type A backup record"
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName, strings.TrimSuffix(acctest.TestHost, "/"))
}

func testAccResourceUpdateSFPoolA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_sfpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_sfpool.id}"
		owner_name = "%s"
		record_type = "1"
		ttl = 150
		record_data = ["192.168.1.2"]
		region_failure_sensitivity = "LOW"
		live_record_state = "NOT_FORCED"
		monitor{
			url = "%s"
			method = "GET"
		}
		backup_record{
			rdata = "192.168.1.1"
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName, acctest.TestHost)
}

func testAccResourceSFPoolAAAA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_sfpool" "aaaa" {
		zone_name = "${resource.ultradns_zone.primary_sfpool.id}"
		owner_name = "%s"
		record_type = "28"
		ttl = 120
		record_data = ["aaaa:bbbb:cccc:dddd:eeee:ffff:1111:2222"]
		region_failure_sensitivity = "HIGH"
		live_record_state = "NOT_FORCED"
		live_record_description = "Active"
		pool_description = "SF Pool Resource of Type AAAA"
		monitor{
			url = "%s"
			method = "GET"
			search_string = "test"
		}
		backup_record{
			rdata = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"
			description = "Type AAAA Backup record"
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName, acctest.TestHost)
}

func testAccResourceUpdateSFPoolAAAA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_sfpool" "aaaa" {
		zone_name = "${resource.ultradns_zone.primary_sfpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_sfpool.id}"
		record_type = "AAAA"
		ttl = 150
		record_data = ["aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"]
		region_failure_sensitivity = "LOW"
		live_record_state = "NOT_FORCED"
		live_record_description = "Maintenance Activity"
		pool_description = "Update SF Pool Resource of Type AAAA"
		monitor{
			url = "%s"
			method = "POST"
			search_string = "testing"
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName, strings.TrimSuffix(acctest.TestHost, "/"))
}
