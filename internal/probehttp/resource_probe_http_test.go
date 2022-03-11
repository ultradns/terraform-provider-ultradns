package probehttp_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
)

func TestAccResourceProbeHTTP(t *testing.T) {
	zoneNameSB := acctest.GetRandomZoneName()
	zoneNameTC := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckProbeResourceDestroy("ultradns_probe_http", probe.HTTP),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProbeHTTPForSBPool(zoneNameSB, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_http.http_sb", probe.HTTP),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "agents.#", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "interval", "ONE_MINUTE"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "total_limit.0.fail", "15"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.method", "POST"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.protocol_version", "HTTP/1.0"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.url", "https://www.ultradns.com/"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.transmitted_data", "foo=bar"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.follow_redirects", "true"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.expected_response", "2XX"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.search_string.0.fail", "Failure"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.connect_limit.0.fail", "11"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.run_limit.0.fail", "12"),
				),
			},
			{
				ResourceName:      "ultradns_probe_http.http_sb",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceUpdateProbeHTTPForSBPool(zoneNameSB, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_http.http_sb", probe.HTTP),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "agents.#", "3"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "threshold", "3"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "interval", "FIVE_MINUTES"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "total_limit.0.fail", "16"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.method", "GET"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.protocol_version", "HTTP/1.0"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.url", "https://www.ultradns.com/"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.follow_redirects", "false"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.expected_response", "3XX"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.search_string.0.fail", "Fail"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.connect_limit.0.fail", "10"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_sb", "transaction.0.run_limit.0.fail", "15"),
				),
			},
			{
				Config: testAccResourceProbeHTTPForTCPool(zoneNameTC, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_http.http_tc", probe.HTTP),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "agents.#", "4"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "threshold", "4"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "interval", "HALF_MINUTE"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "total_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "total_limit.0.critical", "8"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "total_limit.0.fail", "10"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.#", "3"),
				),
			},
			{
				Config: testAccResourceUpdateProbeHTTPForTCPool(zoneNameTC, ownerName),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("ultradns_probe_http.http_tc", probe.HTTP),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "agents.#", "3"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "threshold", "2"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "interval", "TEN_MINUTES"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "total_limit.0.warning", "10"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "total_limit.0.critical", "12"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "total_limit.0.fail", "15"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.method", "GET"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.protocol_version", "HTTP/1.0"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.url", "https://www.ultradns.com/"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.follow_redirects", "false"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.expected_response", "2XX"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.connect_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.connect_limit.0.critical", "8"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.connect_limit.0.fail", "10"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.run_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.run_limit.0.critical", "8"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.run_limit.0.fail", "10"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.avg_connect_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.avg_connect_limit.0.critical", "8"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.avg_connect_limit.0.fail", "10"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.avg_run_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.avg_run_limit.0.critical", "8"),
					resource.TestCheckResourceAttr("ultradns_probe_http.http_tc", "transaction.0.avg_run_limit.0.fail", "10"),
				),
			},
			{
				ResourceName:      "ultradns_probe_http.http_tc",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func testAccResourceProbeHTTPForSBPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_http" "http_sb" {
		zone_name = "${resource.ultradns_zone.primary_sbpool.id}"
		owner_name = "${resource.ultradns_sbpool.a.owner_name}"
		interval = "ONE_MINUTE"
		agents = ["NEW_YORK","DALLAS"]
		threshold = 2
		total_limit{
			fail = 15
		}
		transaction{
			method = "POST"
			protocol_version = "HTTP/1.0"
			url = "https://www.ultradns.com/"
			transmitted_data = "foo=bar"
			follow_redirects = true
			expected_response = "2XX"
			search_string {
				fail = "Failure"
			}
			connect_limit{
				fail = 11
			}
			run_limit{
				fail = 12
			}
		}
	}
	`, acctest.TestAccResourceSBPool(zoneName, ownerName))
}

func testAccResourceUpdateProbeHTTPForSBPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_http" "http_sb" {
		zone_name = "${resource.ultradns_zone.primary_sbpool.id}"
		owner_name = "${resource.ultradns_sbpool.a.owner_name}"
		interval = "FIVE_MINUTES"
		agents = ["NEW_YORK","DALLAS","CHINA"]
		threshold = 3
		total_limit{
			fail = 16
		}
		transaction{
			method = "GET"
			protocol_version = "HTTP/1.0"
			url = "https://www.ultradns.com/"
			follow_redirects = false
			expected_response = "3XX"
			search_string {
				fail = "Fail"
			}
			connect_limit{
				fail = 10
			}
			run_limit{
				fail = 15
			}
		}
	}
	`, acctest.TestAccResourceSBPool(zoneName, ownerName))
}

func testAccResourceProbeHTTPForTCPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_http" "http_tc" {
		zone_name = "${resource.ultradns_zone.primary_tcpool.id}"
		owner_name = "${resource.ultradns_tcpool.a.owner_name}"
		pool_record = "192.168.1.1"
		interval = "HALF_MINUTE"
		agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
		threshold = 4
		total_limit{
			warning = "5"
			critical = "8"
			fail = "10"
		}
		transaction{
			method = "GET"
			protocol_version = "HTTP/1.0"
			url = "https://www.ultradns.com/"
			follow_redirects = false
			expected_response = "2XX"
			search_string {
				fail = "fail"
				warning = "warning"
				critical = "critical"
			}
			connect_limit{
				warning = "5"
				critical = "8"
				fail = "10"
			}
			run_limit{
				warning = "5"
				critical = "8"
				fail = "10"
			}
		}
		transaction{
			method = "POST"
			protocol_version = "HTTP/1.0"
			url = "https://www.ultradns.com/"
			transmitted_data = "foo=bar"
			follow_redirects = true
			expected_response = "2XX"
			search_string {
				fail = "Failure"
			}
			connect_limit{
				fail = 11
			}
			run_limit{
				fail = 12
			}
		}
		transaction{
			method = "POST"
			protocol_version = "HTTP/1.0"
			url = "https://www.ultradns.com/"
			transmitted_data = "foo=bar"
			follow_redirects = true
			expected_response = "2XX"
			search_string {
				fail = "Failure"
			}
			connect_limit{
				fail = 11
			}
			run_limit{
				fail = 12
			}
		}
	}
	`, acctest.TestAccResourceTCPool(zoneName, ownerName))
}

func testAccResourceUpdateProbeHTTPForTCPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_probe_http" "http_tc" {
		zone_name = "${resource.ultradns_zone.primary_tcpool.id}"
		owner_name = "${resource.ultradns_tcpool.a.owner_name}"
		pool_record = "192.168.1.1"
		interval = "TEN_MINUTES"
		agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO"]
		threshold = 2
		total_limit{
			warning = "10"
			critical = "12"
			fail = "15"
		}
		transaction{
			method = "GET"
			protocol_version = "HTTP/1.0"
			url = "https://www.ultradns.com/"
			follow_redirects = false
			expected_response = "2XX"
			connect_limit{
				warning = "5"
				critical = "8"
				fail = "10"
			}
			avg_connect_limit{
				warning = "5"
				critical = "8"
				fail = "10"
			}
			run_limit{
				warning = "5"
				critical = "8"
				fail = "10"
			}
			avg_run_limit{
				warning = "5"
				critical = "8"
				fail = "10"
			}
		}
	}
	`, acctest.TestAccResourceTCPool(zoneName, ownerName))
}
