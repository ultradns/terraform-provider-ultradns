package probehttp_test

import (
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	sdkprobe "github.com/ultradns/ultradns-go-sdk/pkg/probe"
)

func TestAccDataSourceProbeHTTP(t *testing.T) {
	zoneNameSB := acctest.GetRandomZoneName()
	zoneNameTC := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckProbeResourceDestroy("ultradns_probe_http", sdkprobe.HTTP),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceProbe(
					"ultradns_probe_http",
					"http_sb",
					strings.ToUpper(zoneNameSB),
					strings.ToUpper(ownerName),
					"A",
					testAccResourceProbeHTTPForSBPool(zoneNameSB, ownerName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("data.ultradns_probe_http.data_http_sb", sdkprobe.HTTP),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "agents.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "interval", "ONE_MINUTE"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "total_limit.0.fail", "15"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "transaction.0.method", "POST"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "transaction.0.protocol_version", "HTTP/1.0"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "transaction.0.url", "https://www.ultradns.com/"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "transaction.0.transmitted_data", "foo=bar"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "transaction.0.follow_redirects", "true"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "transaction.0.expected_response", "2XX"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "transaction.0.search_string.0.fail", "Failure"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "transaction.0.connect_limit.0.fail", "11"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_sb", "transaction.0.run_limit.0.fail", "12"),
				),
			},
			{
				Config: acctest.TestAccDataSourceProbeWithOptions(
					"ultradns_probe_http",
					"http_tc",
					strings.ToUpper(zoneNameTC),
					strings.ToUpper(ownerName),
					"A",
					"TEN_MINUTES",
					"192.168.1.1",
					testAccResourceUpdateProbeHTTPForTCPool(strings.ToUpper(zoneNameTC), ownerName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("data.ultradns_probe_http.data_http_tc", sdkprobe.HTTP),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "agents.#", "3"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "interval", "TEN_MINUTES"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "total_limit.0.warning", "10"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "total_limit.0.critical", "12"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "total_limit.0.fail", "15"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.method", "GET"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.protocol_version", "HTTP/1.0"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.url", "https://www.ultradns.com/"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.follow_redirects", "false"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.expected_response", "2XX"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.connect_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.connect_limit.0.critical", "8"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.connect_limit.0.fail", "10"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.run_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.run_limit.0.critical", "8"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.run_limit.0.fail", "10"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.avg_connect_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.avg_connect_limit.0.critical", "8"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.avg_connect_limit.0.fail", "10"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.avg_run_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.avg_run_limit.0.critical", "8"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc", "transaction.0.avg_run_limit.0.fail", "10"),
				),
			},
			{
				Config: acctest.TestAccDataSourceProbeWithOptions(
					"ultradns_probe_http",
					"http_tc_aaaa",
					strings.ToUpper(zoneNameTC),
					strings.ToUpper(ownerName),
					"AAAA",
					"HALF_MINUTE",
					"2001:db8:85a3:0:0:8a2e:370:7335",
					testAccResourceProbeHTTPForTCPoolAAAA(strings.ToUpper(zoneNameTC), ownerName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("data.ultradns_probe_http.data_http_tc_aaaa", sdkprobe.HTTP),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc_aaaa", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc_aaaa", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc_aaaa", "pool_record", "2001:db8:85a3:0:0:8a2e:370:7335"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc_aaaa", "agents.#", "4"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc_aaaa", "threshold", "4"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc_aaaa", "interval", "HALF_MINUTE"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc_aaaa", "total_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc_aaaa", "total_limit.0.critical", "8"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc_aaaa", "total_limit.0.fail", "10"),
					resource.TestCheckResourceAttr("data.ultradns_probe_http.data_http_tc_aaaa", "transaction.#", "3"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}
