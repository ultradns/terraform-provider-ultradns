package record_test

import (
	"strings"
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
					strings.ToUpper(ownerNameTypeA)+"."+zoneName,
					"1",
					testAccResourceRecordA(zoneName, strings.ToUpper(ownerNameTypeA)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_a", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "owner_name", ownerNameTypeA+"."+zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "record_type", "A"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_a", "record_data.0", "192.168.1.1"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"ns",
					strings.ToUpper(zoneName),
					zoneName,
					"2",
					testAccResourceRecordNS(strings.ToUpper(zoneName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_ns", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "owner_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "record_type", "NS"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "record_data.#", "3"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"cname",
					strings.ToUpper(zoneName),
					strings.ToUpper(tfacctest.RandString(3)),
					"CNAME",
					testAccResourceRecordCNAME(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_cname", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "record_type", "CNAME"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_cname", "record_data.0", "example.com."),
				),
			},
			// {
			// 	Config: acctest.TestAccDataSourceRRSet(
			// 		"ultradns_record",
			// 		"soa",
			// 		zoneName,
			// 		strings.ToUpper(zoneName),
			// 		"SOA",
			// 		testAccResourceRecordSOA(zoneName),
			// 	),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_soa", ""),
			// 		resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "zone_name", zoneName),
			// 		resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "owner_name", zoneName),
			// 		resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "record_type", "SOA"),
			// 		resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "ttl", "800"),
			// 		resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "record_data.0", "udns1.ultradns.net. sample@example.com. 10800 3600 2592000 10800"),
			// 	),
			// },
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"ptr",
					zoneName,
					strings.ToUpper(tfacctest.RandString(3)),
					"PTR",
					testAccResourceRecordPTR(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_ptr", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "record_type", "PTR"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ptr", "record_data.0", "example.com."),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"mx",
					strings.ToUpper(zoneName),
					tfacctest.RandString(3),
					"MX",
					testAccResourceRecordMX(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_mx", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_mx", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_mx", "record_type", "MX"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_mx", "ttl", "800"),
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
					testAccResourceRecordTXT(strings.ToUpper(zoneName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_txt", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "record_type", "TXT"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_txt", "record_data.0", "example.com."),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"aaaa",
					strings.ToUpper(zoneName),
					strings.ToUpper(tfacctest.RandString(3)),
					"AAAA",
					testAccResourceRecordAAAA(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_aaaa", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_aaaa", "ttl", "800"),
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
					resource.TestCheckResourceAttr("data.ultradns_record.data_srv", "ttl", "800"),
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
					resource.TestCheckResourceAttr("data.ultradns_record.data_sshfp", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_sshfp", "record_data.0", "1 2 54B5E539EAF593AEA410F80737530B71CCDE8B6C3D241184A1372E98BC7EDB37"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"apex",
					strings.ToUpper(zoneName),
					tfacctest.RandString(3),
					"APEXALIAS",
					testAccResourceRecordAPEXALIAS(strings.ToUpper(zoneName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_apex", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "record_type", "APEXALIAS"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_apex", "record_data.0", "example.com."),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"ds",
					strings.ToUpper(zoneName),
					zoneName,
					"DS",
					testAccResourceRecordDS(strings.ToUpper(zoneName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_ds", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ds", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ds", "record_type", "DS"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ds", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ds", "record_data.0", "25286 1 1 340437DC66C3DFAD0B3E849740D2CF1A4151671D"),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"caa",
					strings.ToUpper(zoneName),
					zoneName,
					"CAA",
					testAccResourceRecordCAA(strings.ToUpper(zoneName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_caa", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "record_type", "CAA"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "record_data.0", "0 issue \"ultradns\""),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"svcb",
					strings.ToUpper(zoneName),
					tfacctest.RandString(3),
					"SVCB",
					testAccResourceRecordSVCB(strings.ToUpper(zoneName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_svcb", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_svcb", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_svcb", "record_type", "SVCB"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_svcb", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_svcb", "record_data.0", "0 www.ultradns.com."),
				),
			},
			{
				Config: acctest.TestAccDataSourceRRSet(
					"ultradns_record",
					"https",
					strings.ToUpper(zoneName),
					tfacctest.RandString(3),
					"HTTPS",
					testAccResourceRecordHTTPS(strings.ToUpper(zoneName)),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_https", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_https", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_https", "record_type", "HTTPS"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_https", "ttl", "800"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_https", "record_data.0", "1 www.ultradns.com. mandatory=alpn,key65444 alpn=\"h3,h3-29,h2\" no-default-alpn port=8080 ipv4hint=1.2.3.4,9.8.7.6  ech=dGVzdA== ipv6hint=2001:db8:3333:4444:5555:6666:7777:8888,2001:db8:3333:4444:cccc:dddd:eeee:ffff key65444=\"privateKeyTesting\""),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}
