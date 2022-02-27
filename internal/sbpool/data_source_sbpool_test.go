package sbpool_test

import (
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
)

func TestAccDataSourceSBPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckSBPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_sbpool",
					"a",
					zoneName,
					ownerName+"."+zoneName,
					"1",
					testAccResourceSBPoolA(zoneName, ownerName),
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSBPoolExists("data.ultradns_sbpool.data_a"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "owner_name", ownerName+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "record_type", "A"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "status", "OK"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "pool_description", "SB Pool Resource of Type A"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "run_probes", "true"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "act_on_probes", "true"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "order", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "failure_threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "max_active", "1"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "max_served", "1"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "backup_record.0.rdata", "192.168.1.3"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "backup_record.0.failover_delay", "1"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "backup_record.0.available_to_serve", "true"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "rdata_info.#", "2"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}
