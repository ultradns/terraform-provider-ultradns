package probeping_test

import (
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	sdkprobe "github.com/ultradns/ultradns-go-sdk/pkg/probe"
)

func TestAccDataSourceProbePING(t *testing.T) {
	zoneNameSB := acctest.GetRandomZoneName()
	zoneNameTC := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckProbeResourceDestroy("ultradns_probe_ping", sdkprobe.PING),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceProbe(
					"ultradns_probe_ping",
					"ping_sb",
					strings.ToUpper(zoneNameSB),
					ownerName,
					testAccResourceProbePINGForSBPool(zoneNameSB, strings.ToUpper(ownerName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("data.ultradns_probe_ping.data_ping_sb", sdkprobe.PING),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_sb", "zone_name", zoneNameSB),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_sb", "owner_name", ownerName+"."+zoneNameSB),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_sb", "pool_record", "192.168.1.1"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_sb", "agents.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_sb", "threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_sb", "interval", "ONE_MINUTE"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_sb", "packets", "3"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_sb", "packet_size", "53"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_sb", "loss_percent_limit.0.fail", "1"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_sb", "total_limit.0.fail", "18"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_sb", "run_limit.0.fail", "5"),
				),
			},
			{
				Config: acctest.TestAccDataSourceProbeWithOptions(
					"ultradns_probe_ping",
					"ping_tc",
					zoneNameTC,
					strings.ToUpper(ownerName),
					"FIFTEEN_MINUTES",
					"",
					testAccResourceUpdateProbePINGForTCPool(zoneNameTC, ownerName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckProbeResourceExists("data.ultradns_probe_ping.data_ping_tc", sdkprobe.PING),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "zone_name", zoneNameTC),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "owner_name", ownerName+"."+zoneNameTC),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "agents.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "interval", "FIFTEEN_MINUTES"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "packets", "3"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "packet_size", "53"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "loss_percent_limit.0.warning", "2"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "loss_percent_limit.0.critical", "3"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "loss_percent_limit.0.fail", "4"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "total_limit.0.warning", "5"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "total_limit.0.critical", "6"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "total_limit.0.fail", "7"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "average_limit.0.warning", "8"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "average_limit.0.critical", "9"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "average_limit.0.fail", "10"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "run_limit.0.warning", "11"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "run_limit.0.critical", "12"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "run_limit.0.fail", "13"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "avg_run_limit.0.warning", "14"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "avg_run_limit.0.critical", "15"),
					resource.TestCheckResourceAttr("data.ultradns_probe_ping.data_ping_tc", "avg_run_limit.0.fail", "16"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}
