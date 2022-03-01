package record_test

import (
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
)

func TestAccDataSourceRecord(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	ownerNameTypeA := tfacctest.RandString(3)
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_record", ""),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"a",
					zoneName,
					ownerNameTypeA+"."+zoneName,
					"1",
					testAccResourceRecordA(zoneName, ownerNameTypeA),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_a", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "record_type", "A"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "record_data.0", "192.168.1.1"),
				),
			},

			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"cname",
					zoneName,
					tfacctest.RandString(3),
					"CNAME",
					testAccResourceRecordCNAME(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_cname", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "record_type", "CNAME"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "record_data.0", "example.com."),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"ptr",
					zoneName,
					tfacctest.RandString(3),
					"PTR",
					testAccResourceRecordPTR(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_ptr", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "record_type", "PTR"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "record_data.0", "example.com."),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"mx",
					zoneName,
					tfacctest.RandString(3),
					"MX",
					testAccResourceRecordMX(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_mx", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_mx", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_mx", "record_type", "MX"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_mx", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_mx", "record_data.0", "2 example.com."),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"txt",
					zoneName,
					tfacctest.RandString(3),
					"TXT",
					testAccResourceRecordTXT(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_txt", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "record_type", "TXT"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "record_data.0", "example.com."),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"aaaa",
					zoneName,
					tfacctest.RandString(3),
					"AAAA",
					testAccResourceRecordAAAA(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_aaaa", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_aaaa", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_aaaa", "record_data.0", "2001:db8:85a3:0:0:8a2e:370:7334"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"srv",
					zoneName,
					tfacctest.RandString(3),
					"SRV",
					testAccResourceRecordSRV(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_srv", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_srv", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_srv", "record_type", "SRV"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_srv", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_srv", "record_data.0", "5 6 7 example.com."),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"sshfp",
					zoneName,
					tfacctest.RandString(3),
					"44",
					testAccResourceRecordSSHFP(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_sshfp", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_sshfp", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_sshfp", "record_type", "SSHFP"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_sshfp", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_sshfp", "record_data.0", "1 2 54B5E539EAF593AEA410F80737530B71CCDE8B6C3D241184A1372E98BC7EDB37"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"apex",
					zoneName,
					tfacctest.RandString(3),
					"APEXALIAS",
					testAccResourceRecordAPEXALIAS(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_apex", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "record_type", "APEXALIAS"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "ttl", "120"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "record_data.0", "example.com."),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}
