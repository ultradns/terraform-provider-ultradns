package dirpool_test

import (
	"fmt"
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

const zoneResourceName = "primary_dirpool"

func TestAccResourceDirPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerNameA := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_dirpool", pool.DIR),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDIRPoolA(zoneName, strings.ToUpper(ownerNameA)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_dirpool.a", pool.DIR),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "owner_name", ownerNameA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "pool_description", "DIR Pool Resource of Type A"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "ignore_ecs", "false"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "conflict_resolve", "GEO"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "rdata_info.#", "2"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "no_response.0.geo_group_name", "geo_response_group"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "no_response.0.ip_group_name", "ip_response_group"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "no_response.0.geo_codes.#", "2"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "no_response.0.ip.#", "3"),
				),
			},
			{
				Config: testAccResourceUpdateDIRPoolA(zoneName, ownerNameA),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_dirpool.a", pool.DIR),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "owner_name", ownerNameA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "pool_description", ownerNameA+"."+zoneName),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "ignore_ecs", "true"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "conflict_resolve", "IP"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "rdata_info.0.rdata", "192.168.1.5"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "rdata_info.0.all_non_configured", "true"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "rdata_info.0.ttl", "800"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "no_response.0.geo_group_name", "geo_response_group"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "no_response.0.ip_group_name", "ip_response_group"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "no_response.0.geo_codes.0", "AG"),
					resource.TestCheckResourceAttr("ultradns_dirpool.a", "no_response.0.ip.0.address", "2.2.2.2"),
				),
			},
			{
				ResourceName:      "ultradns_dirpool.a",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceDIRPoolPTR(zoneName, tfacctest.RandString(3)),
			},
			{
				ResourceName:      "ultradns_dirpool.ptr",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceDIRPoolMX(zoneName, tfacctest.RandString(3)),
			},
			{
				ResourceName:      "ultradns_dirpool.mx",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceDIRPoolTXT(zoneName, tfacctest.RandString(3)),
			},
			{
				ResourceName:      "ultradns_dirpool.txt",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceDIRPoolAAAA(zoneName, tfacctest.RandString(3)),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_dirpool.aaaa", pool.DIR),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "pool_description", "DIR Pool Resource of Type AAAA"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "ignore_ecs", "true"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "conflict_resolve", "IP"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "rdata_info.0.rdata", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "rdata_info.0.geo_group_name", "geo_group"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "rdata_info.0.geo_codes.0", "EUR"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "rdata_info.0.ip_group_name", "ip_group"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "rdata_info.0.ip.0.start", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:4444"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "rdata_info.0.ip.0.end", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:6666"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "no_response.0.geo_group_name", "geo_response_group"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "no_response.0.ip_group_name", "ip_response_group"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "no_response.0.geo_codes.0", "AI"),
					resource.TestCheckResourceAttr("ultradns_dirpool.aaaa", "no_response.0.ip.0.address", "aaaa:bbbb:cccc:dddd:eeee:ffff:3333:5555"),
				),
			},
			{
				ResourceName:      "ultradns_dirpool.aaaa",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceDIRPoolSRV(zoneName, tfacctest.RandString(3)),
			},
			{
				ResourceName:      "ultradns_dirpool.srv",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func testAccResourceDIRPoolA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_dirpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_dirpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_dirpool.id}"
		record_type = "A"
		pool_description = "DIR Pool Resource of Type A"
    	rdata_info{
			rdata = "192.168.1.1"
			all_non_configured = true
			ttl = 500
		}
		rdata_info{
			rdata = "192.168.1.2"
			geo_group_name = "geo_group"
			geo_codes = ["NAM","EUR"]
			ip_group_name = "ip_group"
			ip{
				address = "200.1.1.1"
			}
			ip{
				start = "200.1.1.2"
				end = "200.1.1.5"
			}
			ip{
				cidr = "200.20.20.0/24"
			}
		}
		no_response{
			geo_group_name = "geo_response_group"
			geo_codes = ["AG","AI"]
			ip_group_name = "ip_response_group"
			ip{
				address = "1.1.1.1"
			}
			ip{
				start = "1.1.1.2"
				end = "1.1.1.5"
			}
			ip{
				cidr = "20.20.20.0/24"
			}
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccResourceUpdateDIRPoolA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_dirpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_dirpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_dirpool.id}"
		record_type = "A"
		ignore_ecs = true
    	conflict_resolve = "IP"
    	rdata_info{
			rdata = "192.168.1.5"
			all_non_configured = true
			ttl = 800
		}
		no_response{
			geo_group_name = "geo_response_group"
			geo_codes = ["AG"]
			ip_group_name = "ip_response_group"
			ip{
				address = "2.2.2.2"
			}
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccResourceDIRPoolPTR(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_dirpool" "ptr" {
		zone_name = "${resource.ultradns_zone.primary_dirpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_dirpool.id}"
		record_type = "PTR"
		rdata_info{
			rdata = "ns1.example.com."
			geo_group_name = "geo_group"
			geo_codes = ["NAM","EUR"]
		}
		no_response{
			all_non_configured = true
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccResourceDIRPoolMX(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_dirpool" "mx" {
		zone_name = "${resource.ultradns_zone.primary_dirpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_dirpool.id}"
		record_type = "MX"
		rdata_info{
			rdata = "2 example.com."
			geo_group_name = "geo_group"
			geo_codes = ["NAM","EUR"]
		}
		no_response{
			all_non_configured = true
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccResourceDIRPoolTXT(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_dirpool" "txt" {
		zone_name = "${resource.ultradns_zone.primary_dirpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_dirpool.id}"
		record_type = "TXT"
		rdata_info{
			rdata = "text data"
			geo_group_name = "geo_group"
			geo_codes = ["NAM","EUR"]
		}
		no_response{
			all_non_configured = true
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccResourceDIRPoolAAAA(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_dirpool" "aaaa" {
		zone_name = "${resource.ultradns_zone.primary_dirpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_dirpool.id}"
		record_type = "AAAA"
		pool_description = "DIR Pool Resource of Type AAAA"
		ignore_ecs = true
    	conflict_resolve = "IP"
		rdata_info{
			rdata = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"
			geo_group_name = "geo_group"
			geo_codes = ["EUR"]
			ip_group_name = "ip_group"
			ip{
				start = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:4444"
				end = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:6666"
			}
		}
		no_response{
			geo_group_name = "geo_response_group"
			geo_codes = ["AI"]
			ip_group_name = "ip_response_group"
			ip{
				address = "aaaa:bbbb:cccc:dddd:eeee:ffff:3333:5555"
			}
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}

func testAccResourceDIRPoolSRV(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_dirpool" "srv" {
		zone_name = "${resource.ultradns_zone.primary_dirpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_dirpool.id}"
		record_type = "SRV"
		rdata_info{
			rdata = "5 6 7 example.com."
			geo_group_name = "geo_group"
			geo_codes = ["NAM","EUR"]
		}
		no_response{
			all_non_configured = true
		}
	}
	`, acctest.TestAccResourceZonePrimary(zoneResourceName, zoneName), ownerName)
}
