package sbpool_test

import (
	"fmt"
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

const zoneResourceName = "primary_sbpool"

func TestAccResourceSBPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_sbpool", pool.SB),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccResourceSBPool(strings.ToUpper(zoneName), ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_sbpool.a", pool.SB),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "owner_name", ownerName+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "pool_description", "SB Pool Resource of Type A"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "run_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "act_on_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "order", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "failure_threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "max_active", "1"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "max_served", "1"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "backup_record.0.rdata", "192.168.1.3"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "backup_record.0.failover_delay", "1"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "backup_record.0.available_to_serve", "true"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "rdata_info.#", "2"),
				),
			},
			{
				Config: testAccResourceUpdateSBPool(zoneName, strings.ToUpper(ownerName)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_sbpool.a", pool.SB),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "owner_name", ownerName+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "ttl", "150"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "pool_description", ownerName+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "run_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "act_on_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "order", "FIXED"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "failure_threshold", "1"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "max_active", "1"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "max_served", "1"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "rdata_info.0.rdata", "192.168.1.7"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "rdata_info.0.priority", "2"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "rdata_info.0.threshold", "1"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "rdata_info.0.failover_delay", "2"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "rdata_info.0.run_probes", "false"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "rdata_info.0.state", "NORMAL"),
					resource.TestCheckResourceAttr("ultradns_sbpool.a", "backup_record.#", "3"),
				),
			},
			{
				ResourceName:      "ultradns_sbpool.a",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: acctest.TestAccResourceSBPoolAAAA(strings.ToUpper(zoneName), ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_sbpool.aaaa", pool.SB),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "owner_name", ownerName+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "pool_description", "SB Pool Resource of Type AAAA"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "run_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "act_on_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "order", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "failure_threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "max_active", "1"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "max_served", "1"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "backup_record.0.rdata", "2001:db8:85a3:0:0:8a2e:370:7337"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "backup_record.0.failover_delay", "1"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "backup_record.0.available_to_serve", "true"),
					resource.TestCheckResourceAttr("ultradns_sbpool.aaaa", "rdata_info.#", "2"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func testAccResourceUpdateSBPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_sbpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_sbpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_sbpool.id}"
		record_type = "1"
		ttl = 150
    	run_probes = true
    	act_on_probes = true
    	order = "FIXED"
    	failure_threshold = 1
    	max_active = 1
    	rdata_info{
			priority = 2
			threshold = 1
			rdata = "192.168.1.7"
			failover_delay = 2
			run_probes = false
			state = "NORMAL"
		}
		backup_record{
			rdata = "192.168.1.4"
			failover_delay = 1
		}
		backup_record{
			rdata = "192.168.1.5"
			failover_delay = 1
		}
		backup_record{
			rdata = "192.168.1.6"
			failover_delay = 1
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}
