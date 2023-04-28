package probedns_test

import (
	"fmt"
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
)

func TestAccResourceProbeDNS(t *testing.T) {
	zoneNameSB := acctest.GetRandomZoneName()
	zoneNameTC := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckProbeResourceDestroy("ultradns_probe_dns", probe.DNS),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProbeDNSForSBPool(zoneNameSB, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_dns.dns_sb", probe.DNS),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "agents.#", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "interval", "ONE_MINUTE"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "port", "55"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "tcp_only", "true"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "type", "SOA"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "response.0.fail", "fail"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "run_limit.0.fail", "5"),
				),
			},
			{
				Config: testAccResourceUpdateProbeDNSForSBPool(zoneNameSB, strings.ToUpper(ownerName)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_dns.dns_sb", probe.DNS),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "agents.#", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "interval", "TEN_MINUTES"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "port", "53"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "tcp_only", "false"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "type", "NULL"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "query_name", "www.ultradns.com"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "response.0.fail", "failure"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_sb", "run_limit.0.fail", "8"),
				),
			},
			{
				ResourceName:      "ultradns_probe_dns.dns_sb",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceProbeDNSForTCPool(strings.ToUpper(zoneNameTC), ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_dns.dns_tc", probe.DNS),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "agents.#", "3"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "interval", "HALF_MINUTE"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "port", "53"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "tcp_only", "false"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "type", "NULL"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "query_name", "www.ultradns.com"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "response.0.warning", "warning"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "response.0.critical", "critical"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "response.0.fail", "fail"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "run_limit.0.warning", "10"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "run_limit.0.critical", "11"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "run_limit.0.fail", "12"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "avg_run_limit.0.warning", "13"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "avg_run_limit.0.critical", "14"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "avg_run_limit.0.fail", "15"),
				),
			},
			{
				ResourceName:      "ultradns_probe_dns.dns_tc",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceUpdateProbeDNSForTCPool(zoneNameTC, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_dns.dns_tc", probe.DNS),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "agents.#", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "interval", "FIFTEEN_MINUTES"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "port", "56"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "tcp_only", "true"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "type", "SOA"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "response.0.warning", "warn"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "response.0.critical", "critical_warning"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "response.0.fail", "failure"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "run_limit.0.warning", "11"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "run_limit.0.critical", "12"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "run_limit.0.fail", "13"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "avg_run_limit.0.warning", "14"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "avg_run_limit.0.critical", "15"),
					resource.TestCheckResourceAttr("ultradns_probe_dns.dns_tc", "avg_run_limit.0.fail", "16"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func testAccResourceProbeDNSForSBPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_dns" "dns_sb" {
		zone_name = "${resource.ultradns_zone.primary_sbpool.id}"
		owner_name = "${resource.ultradns_sbpool.a.owner_name}"
		pool_record = "192.168.1.1"
		interval = "ONE_MINUTE"
		agents = ["NEW_YORK","DALLAS"]
		threshold = 2
		port = 55
		tcp_only = true
		type = "SOA"
		response{
			fail = "fail"
		}
		run_limit{
			fail = 5
		}
	}
	`, acctest.TestAccResourceSBPool(zoneName, ownerName))
}

func testAccResourceUpdateProbeDNSForSBPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_dns" "dns_sb" {
		zone_name = "${resource.ultradns_zone.primary_sbpool.id}"
		owner_name = "${resource.ultradns_sbpool.a.owner_name}"
		pool_record = "192.168.1.1"
		interval = "TEN_MINUTES"
		agents = ["NEW_YORK","DALLAS"]
		threshold = 2
		query_name = "www.ultradns.com"
		response{
			fail = "failure"
		}
		run_limit{
			fail = 8
		}
	}
	`, acctest.TestAccResourceSBPool(zoneName, ownerName))
}

func testAccResourceProbeDNSForTCPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_dns" "dns_tc" {
		zone_name = "${resource.ultradns_zone.primary_tcpool.id}"
		owner_name = "${resource.ultradns_tcpool.a.owner_name}"
		interval = "HALF_MINUTE"
		agents = ["NEW_YORK","DALLAS","PALO_ALTO"]
		threshold = 2
		query_name = "www.ultradns.com"
		response{
			warning = "warning" 
			critical = "critical"
			fail = "fail"
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

func testAccResourceUpdateProbeDNSForTCPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_dns" "dns_tc" {
		zone_name = "${resource.ultradns_zone.primary_tcpool.id}"
		owner_name = "${resource.ultradns_tcpool.a.owner_name}"
		interval = "FIFTEEN_MINUTES"
		agents = ["NEW_YORK","DALLAS"]
		threshold = 2
		port = 56
		tcp_only = true
		type = "SOA"
		response{
			warning = "warn" 
			critical = "critical_warning"
			fail = "failure"
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
