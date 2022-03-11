package sfpool_test

import (
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

func TestAccDataSourceSFPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerNameTypeA := tfacctest.RandString(3)
	ownerNameTypeAAAA := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_sfpool", pool.SF),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_sfpool",
					"a",
					zoneName,
					ownerNameTypeA+"."+zoneName,
					"1",
					testAccResourceSFPoolA(zoneName, ownerNameTypeA),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_sfpool.data_a", pool.SF),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "record_type", "A"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "record_data.0", "192.168.1.1"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "backup_record.0.rdata", "192.168.1.2"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "backup_record.0.description", "Type A backup record"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "monitor.0.method", "POST"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "monitor.0.search_string", "test"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "monitor.0.transmitted_data", "foo=bar"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "live_record_description", "Maintenance Activity"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "pool_description", "SF Pool Resource of Type A"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "region_failure_sensitivity", "HIGH"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_a", "status", "MANUAL"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_sfpool",
					"aaaa",
					strings.TrimSuffix(zoneName, "."),
					ownerNameTypeAAAA,
					"AAAA",
					testAccResourceSFPoolAAAA(zoneName, ownerNameTypeAAAA),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_sfpool.data_aaaa", pool.SF),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "owner_name", ownerNameTypeAAAA+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "record_data.0", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:2222"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "backup_record.0.rdata", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "backup_record.0.description", "Type AAAA Backup record"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "monitor.0.url", acctest.TestHost),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "monitor.0.method", "GET"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "monitor.0.search_string", "test"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "monitor.0.transmitted_data", ""),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "live_record_description", "Active"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "pool_description", "SF Pool Resource of Type AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "region_failure_sensitivity", "HIGH"),
					resource.TestCheckResourceAttr("data.ultradns_sfpool.data_aaaa", "status", "OK"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}
