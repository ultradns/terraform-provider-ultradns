package tcpool_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

const zoneResourceName = "primary_tcpool"

func TestAccResourceTCPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_tcpool", pool.TC),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccResourceTCPool(zoneName, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_tcpool.a", pool.TC),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "owner_name", ownerName+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "pool_description", ownerName+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "run_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "act_on_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "failure_threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "max_to_lb", "2"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "backup_record.0.rdata", "192.168.1.3"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "backup_record.0.failover_delay", "1"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "backup_record.0.available_to_serve", "true"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "rdata_info.#", "2"),
				),
			},
			{
				Config: testAccResourceUpdateTCPool(zoneName, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_tcpool.a", pool.TC),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "owner_name", ownerName+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "ttl", "150"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "status", "OK"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "pool_description", "TC Pool Resource of Type A"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "run_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "act_on_probes", "true"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "failure_threshold", "1"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "max_to_lb", "1"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "rdata_info.0.rdata", "192.168.1.5"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "rdata_info.0.priority", "2"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "rdata_info.0.threshold", "1"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "rdata_info.0.failover_delay", "2"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "rdata_info.0.run_probes", "false"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "rdata_info.0.state", "NORMAL"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "backup_record.0.rdata", "192.168.1.6"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "backup_record.0.failover_delay", "2"),
					resource.TestCheckResourceAttr("ultradns_tcpool.a", "backup_record.0.available_to_serve", "true"),
				),
			},
			{
				ResourceName:      "ultradns_tcpool.a",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func testAccResourceUpdateTCPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_tcpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_tcpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_tcpool.id}"
		record_type = "1"
		ttl = 150
		pool_description = "TC Pool Resource of Type A"
    	run_probes = true
    	act_on_probes = true
    	failure_threshold = 1
    	max_to_lb = 1
    	rdata_info{
			priority = 2
			threshold = 1
			rdata = "192.168.1.5"
			failover_delay = 2
			run_probes = false
			state = "NORMAL"
		}
		backup_record{
			rdata = "192.168.1.6"
			failover_delay = 2
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}
