package probeping_test

import (
	"fmt"
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
)

func TestAccResourceProbePING(t *testing.T) {
	zoneNameSB := acctest.GetRandomZoneName()
	zoneNameTC := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckProbeResourceDestroy("ultradns_probe_ping", probe.PING),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProbePINGForSBPool(zoneNameSB, strings.ToUpper(ownerName)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_ping.ping_sb", probe.PING),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "agents.#", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "interval", "ONE_MINUTE"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "packets", "3"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "packet_size", "53"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "loss_percent_limit.0.fail", "1"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "total_limit.0.fail", "18"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "run_limit.0.fail", "5"),
				),
			},
			{
				Config: testAccResourceUpdateProbePINGForSBPool(strings.ToUpper(zoneNameSB), ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_ping.ping_sb", probe.PING),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "agents.#", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "interval", "TEN_MINUTES"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "packets", "6"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "packet_size", "106"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "loss_percent_limit.0.fail", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "total_limit.0.fail", "20"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_sb", "run_limit.0.fail", "6"),
				),
			},
			{
				ResourceName:      "ultradns_probe_ping.ping_sb",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceProbePINGForTCPool(strings.ToUpper(zoneNameTC), strings.ToUpper(ownerName)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_ping.ping_tc", probe.PING),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "agents.#", "3"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "interval", "HALF_MINUTE"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "packets", "6"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "packet_size", "106"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "loss_percent_limit.0.warning", "1"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "loss_percent_limit.0.critical", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "loss_percent_limit.0.fail", "3"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "total_limit.0.warning", "4"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "total_limit.0.critical", "5"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "total_limit.0.fail", "6"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "average_limit.0.warning", "7"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "average_limit.0.critical", "8"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "average_limit.0.fail", "9"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "run_limit.0.warning", "10"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "run_limit.0.critical", "11"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "run_limit.0.fail", "12"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "avg_run_limit.0.warning", "13"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "avg_run_limit.0.critical", "14"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "avg_run_limit.0.fail", "15"),
				),
			},
			{
				ResourceName:      "ultradns_probe_ping.ping_tc",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceUpdateProbePINGForTCPool(strings.ToUpper(zoneNameTC), strings.ToUpper(ownerName)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_ping.ping_tc", probe.PING),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "agents.#", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "interval", "FIFTEEN_MINUTES"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "packets", "3"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "packet_size", "53"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "loss_percent_limit.0.warning", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "loss_percent_limit.0.critical", "3"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "loss_percent_limit.0.fail", "4"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "total_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "total_limit.0.critical", "6"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "total_limit.0.fail", "7"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "average_limit.0.warning", "8"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "average_limit.0.critical", "9"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "average_limit.0.fail", "10"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "run_limit.0.warning", "11"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "run_limit.0.critical", "12"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "run_limit.0.fail", "13"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "avg_run_limit.0.warning", "14"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "avg_run_limit.0.critical", "15"),
					resource.TestCheckResourceAttr("ultradns_probe_ping.ping_tc", "avg_run_limit.0.fail", "16"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func testAccResourceProbePINGForSBPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_ping" "ping_sb" {
		zone_name = "${resource.ultradns_zone.primary_sbpool.id}"
		owner_name = "${resource.ultradns_sbpool.a.owner_name}"
		pool_record = "192.168.1.1"
		interval = "ONE_MINUTE"
		agents = ["NEW_YORK","DALLAS"]
		threshold = 2
		packets = 3
		packet_size = 53
		loss_percent_limit{
			fail = 1
		}
		total_limit{
			fail = 18
		}
		run_limit{
			fail = 5
		}
	}
	`, acctest.TestAccResourceSBPool(zoneName, ownerName))
}

func testAccResourceUpdateProbePINGForSBPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_ping" "ping_sb" {
		zone_name = "${resource.ultradns_zone.primary_sbpool.id}"
		owner_name = "${resource.ultradns_sbpool.a.owner_name}"
		pool_record = "192.168.1.1"
		interval = "TEN_MINUTES"
		agents = ["NEW_YORK","DALLAS"]
		threshold = 2
		packets = 6
		packet_size = 106
		loss_percent_limit{
			fail = 2
		}
		total_limit{
			fail = 20
		}
		run_limit{
			fail = 6
		}
	}
	`, acctest.TestAccResourceSBPool(zoneName, ownerName))
}

func testAccResourceProbePINGForTCPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_ping" "ping_tc" {
		zone_name = "${resource.ultradns_zone.primary_tcpool.id}"
		owner_name = "${resource.ultradns_tcpool.a.owner_name}"
		interval = "HALF_MINUTE"
		agents = ["NEW_YORK","DALLAS","PALO_ALTO"]
		threshold = 2
		packets = 6
		packet_size = 106
		loss_percent_limit{
			warning = 1 
			critical = 2
			fail = 3
		}
		total_limit{
			warning = 4 
			critical = 5
			fail = 6
		}
		average_limit{
			warning = 7 
			critical = 8
			fail = 9
		}
		run_limit{
			warning = 10 
			critical = 11
			fail = 12
		}
		avg_run_limit{
			warning = 13 
			critical = 14
			fail = 15
		}
	}
	`, acctest.TestAccResourceTCPool(zoneName, ownerName))
}

func testAccResourceUpdateProbePINGForTCPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_ping" "ping_tc" {
		zone_name = "${resource.ultradns_zone.primary_tcpool.id}"
		owner_name = "${resource.ultradns_tcpool.a.owner_name}"
		interval = "FIFTEEN_MINUTES"
		agents = ["NEW_YORK","DALLAS"]
		threshold = 2
		packets = 3
		packet_size = 53
		loss_percent_limit{
			warning = 2 
			critical = 3
			fail = 4
		}
		total_limit{
			warning = 5 
			critical = 6
			fail = 7
		}
		average_limit{
			warning = 8 
			critical = 9
			fail = 10
		}
		run_limit{
			warning = 11 
			critical = 12
			fail = 13
		}
		avg_run_limit{
			warning = 14 
			critical = 15
			fail = 16
		}
	}
	`, acctest.TestAccResourceTCPool(zoneName, ownerName))
}
