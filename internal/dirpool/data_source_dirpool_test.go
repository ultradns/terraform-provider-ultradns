package dirpool_test

import (
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
)

func TestAccDataSourceDIRPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerNameTypeA := tfacctest.RandString(3)
	ownerNameTypePTR := tfacctest.RandString(3)
	ownerNameTypeMX := tfacctest.RandString(3)
	ownerNameTypeTXT := tfacctest.RandString(3)
	ownerNameTypeAAAA := tfacctest.RandString(3)
	ownerNameTypeSRV := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_dirpool", pool.DIR),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_dirpool",
					"a",
					zoneName,
					ownerNameTypeA+"."+zoneName,
					"1",
					testAccResourceDIRPoolA(zoneName, ownerNameTypeA),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_dirpool.data_a", pool.DIR),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_a", "record_type", "A"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_a", "pool_description", "DIR Pool Resource of Type A"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_a", "ignore_ecs", "false"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_a", "conflict_resolve", "GEO"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_a", "rdata_info.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_a", "no_response.0.geo_group_name", "geo_response_group"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_a", "no_response.0.ip_group_name", "ip_response_group"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_a", "no_response.0.geo_codes.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_a", "no_response.0.ip.#", "3"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_dirpool",
					"ptr",
					zoneName,
					ownerNameTypePTR,
					"PTR",
					testAccResourceDIRPoolPTR(zoneName, ownerNameTypePTR),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_dirpool.data_ptr", pool.DIR),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_ptr", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_ptr", "owner_name", ownerNameTypePTR+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_ptr", "record_type", "PTR"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_ptr", "pool_description", ownerNameTypePTR+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_ptr", "ignore_ecs", "false"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_ptr", "conflict_resolve", "GEO"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_ptr", "rdata_info.0.rdata", "ns1.example.com."),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_ptr", "rdata_info.0.geo_group_name", "geo_group"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_ptr", "rdata_info.0.geo_codes.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_ptr", "no_response.0.all_non_configured", "true"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_dirpool",
					"mx",
					zoneName,
					ownerNameTypeMX,
					"MX",
					testAccResourceDIRPoolMX(zoneName, ownerNameTypeMX),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_dirpool.data_mx", pool.DIR),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_mx", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_mx", "owner_name", ownerNameTypeMX+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_mx", "record_type", "MX"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_mx", "pool_description", ownerNameTypeMX+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_mx", "ignore_ecs", "false"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_mx", "conflict_resolve", "GEO"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_mx", "rdata_info.0.rdata", "2 example.com."),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_mx", "rdata_info.0.geo_group_name", "geo_group"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_mx", "rdata_info.0.geo_codes.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_mx", "no_response.0.all_non_configured", "true"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_dirpool",
					"txt",
					zoneName,
					ownerNameTypeTXT,
					"TXT",
					testAccResourceDIRPoolTXT(zoneName, ownerNameTypeTXT),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_dirpool.data_txt", pool.DIR),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_txt", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_txt", "owner_name", ownerNameTypeTXT+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_txt", "record_type", "TXT"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_txt", "pool_description", ownerNameTypeTXT+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_txt", "ignore_ecs", "false"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_txt", "conflict_resolve", "GEO"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_txt", "rdata_info.0.rdata", "text data"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_txt", "rdata_info.0.geo_group_name", "geo_group"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_txt", "rdata_info.0.geo_codes.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_txt", "no_response.0.all_non_configured", "true"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_dirpool",
					"aaaa",
					zoneName,
					ownerNameTypeAAAA,
					"AAAA",
					testAccResourceDIRPoolAAAA(zoneName, ownerNameTypeAAAA),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_dirpool.data_aaaa", pool.DIR),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "pool_description", "DIR Pool Resource of Type AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "ignore_ecs", "true"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "conflict_resolve", "IP"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "rdata_info.0.rdata", "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "rdata_info.0.geo_group_name", "geo_group"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "rdata_info.0.geo_codes.0", "EUR"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "rdata_info.0.ip_group_name", "ip_group"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "rdata_info.0.ip.0.start", "AAAA:BBBB:CCCC:DDDD:EEEE:FFFF:1111:4444"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "rdata_info.0.ip.0.end", "AAAA:BBBB:CCCC:DDDD:EEEE:FFFF:1111:6666"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "no_response.0.geo_group_name", "geo_response_group"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "no_response.0.ip_group_name", "ip_response_group"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "no_response.0.geo_codes.0", "AI"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_aaaa", "no_response.0.ip.0.address", "AAAA:BBBB:CCCC:DDDD:EEEE:FFFF:3333:5555"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_dirpool",
					"srv",
					zoneName,
					ownerNameTypeSRV,
					"SRV",
					testAccResourceDIRPoolSRV(zoneName, ownerNameTypeSRV),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_dirpool.data_srv", pool.DIR),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_srv", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_srv", "owner_name", ownerNameTypeSRV+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_srv", "record_type", "SRV"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_srv", "pool_description", ownerNameTypeSRV+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_srv", "ignore_ecs", "false"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_srv", "conflict_resolve", "GEO"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_srv", "rdata_info.0.rdata", "5 6 7 example.com."),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_srv", "rdata_info.0.geo_group_name", "geo_group"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_srv", "rdata_info.0.geo_codes.#", "2"),
					resource.TestCheckResourceAttr("data.ultradns_dirpool.data_srv", "no_response.0.all_non_configured", "true"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}
