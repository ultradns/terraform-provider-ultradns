package probetcp_test

import (
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	sdkprobe "github.com/ultradns/ultradns-go-sdk/pkg/probe"
)

func TestAccDataSourceProbeTCP(t *testing.T) {
	zoneNameSB := acctest.GetRandomZoneName()
	zoneNameTC := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckProbeResourceDestroy("ultradns_probe_tcp", sdkprobe.TCP),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceProbe(
					"ultradns_probe_tcp",
					"tcp_sb",
					zoneNameSB,
					ownerName,
					testAccResourceProbeTCPForSBPool(zoneNameSB, ownerName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("data.ultradns_probe_tcp.data_tcp_sb", sdkprobe.TCP),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_sb", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_sb", "agents.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_sb", "threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_sb", "interval", "ONE_MINUTE"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_sb", "port", "443"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_sb", "control_ip", ""),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_sb", "connect_limit.0.fail", "5"),
				),
			},
			{
				Config: acctest.TestAccDataSourceProbeWithOptions(
					"ultradns_probe_tcp",
					"tcp_tc",
					zoneNameTC,
					ownerName,
					"FIFTEEN_MINUTES",
					"",
					testAccResourceUpdateProbeTCPForTCPool(zoneNameTC, ownerName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("data.ultradns_probe_tcp.data_tcp_tc", sdkprobe.TCP),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "agents.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "interval", "FIFTEEN_MINUTES"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "port", "443"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "connect_limit.0.warning", "11"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "connect_limit.0.critical", "12"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "connect_limit.0.fail", "13"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "avg_connect_limit.0.warning", "14"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "avg_connect_limit.0.critical", "15"),
					resource.TestCheckResourceAttr("data.ultradns_probe_tcp.data_tcp_tc", "avg_connect_limit.0.fail", "16"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}
