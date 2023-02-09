package probetcp_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
)

func TestAccResourceProbeTCP(t *testing.T) {
	zoneNameSB := acctest.GetRandomZoneName()
	zoneNameTC := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckProbeResourceDestroy("ultradns_probe_tcp", probe.TCP),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProbeTCPForSBPool(zoneNameSB, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_tcp.tcp_sb", probe.TCP),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "agents.#", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "interval", "ONE_MINUTE"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "port", "443"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "connect_limit.0.fail", "5"),
				),
			},
			{
				Config: testAccResourceUpdateProbeTCPForSBPool(zoneNameSB, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_tcp.tcp_sb", probe.TCP),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "agents.#", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "interval", "TEN_MINUTES"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "port", "443"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "control_ip", ""),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_sb", "connect_limit.0.fail", "8"),
				),
			},
			{
				ResourceName:      "ultradns_probe_tcp.tcp_sb",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceProbeTCPForTCPool(zoneNameTC, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_tcp.tcp_tc", probe.TCP),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "agents.#", "3"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "interval", "HALF_MINUTE"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "port", "443"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "connect_limit.0.warning", "10"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "connect_limit.0.critical", "11"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "connect_limit.0.fail", "12"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "avg_connect_limit.0.warning", "13"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "avg_connect_limit.0.critical", "14"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "avg_connect_limit.0.fail", "15"),
				),
			},
			{
				ResourceName:      "ultradns_probe_tcp.tcp_tc",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceUpdateProbeTCPForTCPool(zoneNameTC, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_tcp.tcp_tc", probe.TCP),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "agents.#", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "interval", "FIFTEEN_MINUTES"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "port", "443"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "connect_limit.0.warning", "11"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "connect_limit.0.critical", "12"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "connect_limit.0.fail", "13"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "avg_connect_limit.0.warning", "14"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "avg_connect_limit.0.critical", "15"),
					resource.TestCheckResourceAttr("ultradns_probe_tcp.tcp_tc", "avg_connect_limit.0.fail", "16"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func testAccResourceProbeTCPForSBPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_tcp" "tcp_sb" {
		zone_name = "${resource.ultradns_zone.primary_sbpool.id}"
		owner_name = "${resource.ultradns_sbpool.a.owner_name}"
		pool_record = "192.168.1.1"
		interval = "ONE_MINUTE"
		agents = ["NEW_YORK","DALLAS"]
		threshold = 2
		port = 443
		connect_limit{
			fail = 5
		}
	}
	`, acctest.TestAccResourceSBPool(zoneName, ownerName))
}

func testAccResourceUpdateProbeTCPForSBPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_tcp" "tcp_sb" {
		zone_name = "${resource.ultradns_zone.primary_sbpool.id}"
		owner_name = "${resource.ultradns_sbpool.a.owner_name}"
		pool_record = "192.168.1.1"
		interval = "TEN_MINUTES"
		agents = ["NEW_YORK","DALLAS"]
		threshold = 2
		connect_limit{
			fail = 8
		}
	}
	`, acctest.TestAccResourceSBPool(zoneName, ownerName))
}

func testAccResourceProbeTCPForTCPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_tcp" "tcp_tc" {
		zone_name = "${resource.ultradns_zone.primary_tcpool.id}"
		owner_name = "${resource.ultradns_tcpool.a.owner_name}"
		interval = "HALF_MINUTE"
		agents = ["NEW_YORK","DALLAS","PALO_ALTO"]
		threshold = 2
		port = 443
		connect_limit{
			warning = 10
			critical = 11
			fail = 12
		}
		avg_connect_limit{
			warning = 13
			critical = 14
			fail = 15
		}
	}
	`, acctest.TestAccResourceTCPool(zoneName, ownerName))
}

func testAccResourceUpdateProbeTCPForTCPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_tcp" "tcp_tc" {
		zone_name = "${resource.ultradns_zone.primary_tcpool.id}"
		owner_name = "${resource.ultradns_tcpool.a.owner_name}"
		interval = "FIFTEEN_MINUTES"
		agents = ["NEW_YORK","DALLAS"]
		threshold = 2
		port = 443
		connect_limit{
			warning = 11
			critical = 12
			fail = 13
		}
		avg_connect_limit{
			warning = 14 
			critical = 15
			fail = 16
		}
	}
	`, acctest.TestAccResourceTCPool(zoneName, ownerName))
}
