package probedns_test

import (
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	sdkprobe "github.com/ultradns/ultradns-go-sdk/pkg/probe"
)

func TestAccDataSourceProbeDNS(t *testing.T) {
	zoneNameSB := acctest.GetRandomZoneName()
	zoneNameTC := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckProbeResourceDestroy("ultradns_probe_dns", sdkprobe.DNS),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceProbe(
					"ultradns_probe_dns",
					"dns_sb",
					zoneNameSB,
					ownerName,
					"A",
					testAccResourceProbeDNSForSBPool(strings.ToUpper(zoneNameSB), strings.ToUpper(ownerName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("data.ultradns_probe_dns.data_dns_sb", sdkprobe.DNS),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb", "agents.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb", "threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb", "interval", "ONE_MINUTE"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb", "port", "55"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb", "tcp_only", "true"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb", "type", "SOA"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb", "response.0.fail", "fail"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb", "run_limit.0.fail", "5"),
				),
			},
			{
				Config: acctest.TestAccDataSourceProbeWithOptions(
					"ultradns_probe_dns",
					"dns_tc",
					zoneNameTC,
					ownerName,
					"A",
					"FIFTEEN_MINUTES",
					"",
					testAccResourceUpdateProbeDNSForTCPool(zoneNameTC, ownerName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("data.ultradns_probe_dns.data_dns_tc", sdkprobe.DNS),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "agents.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "interval", "FIFTEEN_MINUTES"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "port", "56"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "tcp_only", "true"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "type", "SOA"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "response.0.warning", "warn"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "response.0.critical", "critical_warning"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "response.0.fail", "failure"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "run_limit.0.warning", "11"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "run_limit.0.critical", "12"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "run_limit.0.fail", "13"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "avg_run_limit.0.warning", "14"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "avg_run_limit.0.critical", "15"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_tc", "avg_run_limit.0.fail", "16"),
				),
			},
			{
				Config: acctest.TestAccDataSourceProbe(
					"ultradns_probe_dns",
					"dns_sb_aaaa",
					zoneNameSB,
					ownerName,
					"AAAA",
					testAccResourceProbeDNSForSBPoolAAAA(strings.ToUpper(zoneNameSB), strings.ToUpper(ownerName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("data.ultradns_probe_dns.data_dns_sb_aaaa", sdkprobe.DNS),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "pool_type", "AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "pool_record", "2001:db8:85a3:0:0:8a2e:370:7335"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "agents.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "interval", "ONE_MINUTE"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "port", "55"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "tcp_only", "true"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "type", "SOA"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "response.0.fail", "fail"),
					resource.TestCheckResourceAttr("data.ultradns_probe_dns.data_dns_sb_aaaa", "run_limit.0.fail", "5"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}
