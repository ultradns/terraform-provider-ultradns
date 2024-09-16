package sbpool_test

import (
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

func TestAccDataSourceSBPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerName := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_sbpool", pool.SB),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_sbpool",
					"a",
					strings.ToUpper(zoneName),
					strings.ToUpper(ownerName+"."+zoneName),
					"1",
					acctest.TestAccResourceSBPool(strings.ToUpper(zoneName), strings.ToUpper(ownerName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_sbpool.data_a", pool.SB),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "owner_name", ownerName+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "record_type", "A"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "ttl", "120"),
					// resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "status", "OK"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "pool_description", "SB Pool Resource of Type A"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "run_probes", "true"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "act_on_probes", "true"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "order", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "failure_threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "max_active", "1"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "max_served", "1"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "backup_record.0.rdata", "192.168.1.3"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "backup_record.0.failover_delay", "1"),
					// resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "backup_record.0.available_to_serve", "true"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_a", "rdata_info.#", "2"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_sbpool",
					"aaaa",
					strings.ToUpper(zoneName),
					strings.ToUpper(ownerName+"."+zoneName),
					"AAAA",
					acctest.TestAccResourceSBPoolAAAA(strings.ToUpper(zoneName), strings.ToUpper(ownerName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_sbpool.data_aaaa", pool.SB),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "owner_name", ownerName+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "ttl", "120"),
					// resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "status", "OK"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "pool_description", "SB Pool Resource of Type AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "run_probes", "true"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "act_on_probes", "true"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "order", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "failure_threshold", "2"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "max_active", "1"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "max_served", "1"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "backup_record.0.rdata", "2001:db8:85a3:0:0:8a2e:370:7337"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "backup_record.0.failover_delay", "1"),
					// resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "backup_record.0.available_to_serve", "true"),
					resource.TestCheckResourceAttr("data.ultradns_sbpool.data_aaaa", "rdata_info.#", "2"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}
