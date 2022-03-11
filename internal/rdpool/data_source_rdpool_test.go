package rdpool_test

import (
	"strings"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

func TestAccDataSourceRDPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerNameTypeA := tfacctest.RandString(3)
	ownerNameTypeAAAA := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_rdpool", pool.RD),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_rdpool",
					"a",
					strings.TrimSuffix(zoneName, "."),
					ownerNameTypeA,
					"A",
					testAccResourceRDPoolA(zoneName, ownerNameTypeA),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_rdpool.data_a", pool.RD),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_a", "record_type", "A"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_a", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_a", "record_data.0", "192.168.1.1"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_a", "order", "FIXED"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_a", "description", "RD Pool Resource of Type A"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_rdpool",
					"aaaa",
					zoneName,
					ownerNameTypeA+"."+zoneName,
					"28",
					testAccResourceRDPoolAAAA(zoneName, ownerNameTypeAAAA),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_rdpool.data_aaaa", pool.RD),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_aaaa", "owner_name", ownerNameTypeAAAA+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_aaaa", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_aaaa", "record_data.0", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:2222"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_aaaa", "order", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr("data.ultradns_rdpool.data_aaaa", "description", ownerNameTypeAAAA+"."+zoneName),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}
