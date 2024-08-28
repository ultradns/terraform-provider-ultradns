package record_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
)

const zoneResource = "record_handling"

func TestAccRecordHandling(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckRecordResourceDestroy("ultradns_record", ""),
		Steps: []resource.TestStep{
			{
				Config: getRecordHandlingConfig(
					acctest.TestAccResourceZonePrimary(zoneResource, zoneName),
					getDataSourceConfig("soa", "${resource.ultradns_zone.record_handling.name}", zoneName, "SOA"),
					getDataSourceConfig("ns", "${resource.ultradns_zone.record_handling.name}", zoneName, "NS"),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_soa", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "owner_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "record_type", "SOA"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "record_data.0", "udns1.ultradns.net. antonyrohith.akash@vercara.com. 86400 86400 86400 86400"),
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_ns", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "owner_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "record_type", "NS"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "record_data.#", "2"),
				),
			},
			{
				Config: getRecordHandlingConfig(
					acctest.TestAccResourceZonePrimary(zoneResource, zoneName),
					getResourceConfig("soa", "${resource.ultradns_zone.record_handling.name}", zoneName, "SOA", "60", "udns1.ultradns.net. antonyrohith.akash@ultradns.com. 60 60 60 60"),
					getDataSourceConfig("soa", "${resource.ultradns_record.soa.zone_name}", zoneName, "SOA"),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.soa", ""),
					resource.TestCheckResourceAttr("ultradns_record.soa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.soa", "owner_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.soa", "record_type", "SOA"),
					resource.TestCheckResourceAttr("ultradns_record.soa", "ttl", "60"),
					resource.TestCheckResourceAttr("ultradns_record.soa", "record_data.0", "udns1.ultradns.net. antonyrohith.akash@ultradns.com. 60 60 60 60"),
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_soa", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "owner_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "record_type", "SOA"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "ttl", "60"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "record_data.0", "udns1.ultradns.net. antonyrohith.akash@ultradns.com. 60 60 60 60"),
				),
			},
			{
				ResourceName:      "ultradns_record.soa",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: getRecordHandlingConfig(
					acctest.TestAccResourceZonePrimary(zoneResource, zoneName),
					getResourceConfig("ns", "${resource.ultradns_zone.record_handling.name}", zoneName, "NS", "60", "udns3.ultradns.net."),
					getResourceConfig("caa", "${resource.ultradns_record.caa_one.zone_name}", zoneName, "CAA", "60", "0 issue letsencrypt"),
					getResourceConfig("caa_one", "${resource.ultradns_zone.record_handling.name}", zoneName, "CAA", "60", "0 issue a b c d"),
					getResourceConfig("https", "${resource.ultradns_zone.record_handling.name}", zoneName, "HTTPS", "60", "1 www.ultradns.com. ech=dGVzdA== no-default-alpn ipv4hint=1.2.5.4,9.8.7.6 key65444=privatekeyTesting ipv6hint=2001:db8:3333:4444:5555:6666:7777:8888,2001:db8:3333:4444:cccc:dddd:eeee:ffff port=8080 alpn=h3,h3-29,h2 mandatory=alpn,key65444"),
					getDataSourceConfig("soa", "${resource.ultradns_zone.record_handling.name}", zoneName, "SOA"),
					getDataSourceConfig("ns", "${resource.ultradns_record.ns.zone_name}", zoneName, "NS"),
					getDataSourceConfig("caa", "${resource.ultradns_record.caa.zone_name}", zoneName, "CAA"),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.ns", ""),
					resource.TestCheckResourceAttr("ultradns_record.ns", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ns", "owner_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.ns", "record_type", "NS"),
					resource.TestCheckResourceAttr("ultradns_record.ns", "ttl", "60"),
					resource.TestCheckResourceAttr("ultradns_record.ns", "record_data.0", "udns3.ultradns.net."),
					acctest.TestAccCheckRecordResourceExists("ultradns_record.caa", ""),
					resource.TestCheckResourceAttr("ultradns_record.caa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.caa", "owner_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.caa", "record_type", "CAA"),
					resource.TestCheckResourceAttr("ultradns_record.caa", "ttl", "60"),
					resource.TestCheckResourceAttr("ultradns_record.caa", "record_data.0", "0 issue letsencrypt"),
					acctest.TestAccCheckRecordResourceExists("ultradns_record.caa_one", ""),
					resource.TestCheckResourceAttr("ultradns_record.caa_one", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.caa_one", "owner_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.caa_one", "record_type", "CAA"),
					resource.TestCheckResourceAttr("ultradns_record.caa_one", "ttl", "60"),
					resource.TestCheckResourceAttr("ultradns_record.caa_one", "record_data.0", "0 issue a b c d"),
					acctest.TestAccCheckRecordResourceExists("ultradns_record.https", ""),
					resource.TestCheckResourceAttr("ultradns_record.https", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.https", "owner_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.https", "record_type", "HTTPS"),
					resource.TestCheckResourceAttr("ultradns_record.https", "ttl", "60"),
					resource.TestCheckResourceAttr("ultradns_record.https", "record_data.0", "1 www.ultradns.com. ech=dGVzdA== no-default-alpn ipv4hint=1.2.5.4,9.8.7.6 key65444=privatekeyTesting ipv6hint=2001:db8:3333:4444:5555:6666:7777:8888,2001:db8:3333:4444:cccc:dddd:eeee:ffff port=8080 alpn=h3,h3-29,h2 mandatory=alpn,key65444"),
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_soa", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "owner_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "record_type", "SOA"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "ttl", "60"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_soa", "record_data.0", "udns1.ultradns.net. antonyrohith.akash@ultradns.com. 60 60 60 60"),
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_ns", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "owner_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "record_type", "NS"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "ttl", "60"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "record_data.#", "3"),
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_caa", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "owner_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "record_type", "CAA"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "ttl", "60"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "record_data.#", "2"),
				),
			},
			{
				ResourceName:      "ultradns_record.https",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: getRecordHandlingConfig(
					acctest.TestAccResourceZonePrimary(zoneResource, zoneName),
					getResourceConfig("caa", "${resource.ultradns_zone.record_handling.name}", zoneName, "CAA", "60", "0 issue ultradns"),
					getResourceConfig("https", "${resource.ultradns_zone.record_handling.name}", zoneName, "HTTPS", "60", "1 www.ultradns.com. alpn=h3,h3-29,h2 mandatory=alpn"),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("ultradns_record.caa", ""),
					resource.TestCheckResourceAttr("ultradns_record.caa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.caa", "owner_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.caa", "record_type", "CAA"),
					resource.TestCheckResourceAttr("ultradns_record.caa", "ttl", "60"),
					resource.TestCheckResourceAttr("ultradns_record.caa", "record_data.0", "0 issue ultradns"),
					acctest.TestAccCheckRecordResourceExists("ultradns_record.https", ""),
					resource.TestCheckResourceAttr("ultradns_record.https", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.https", "owner_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_record.https", "record_type", "HTTPS"),
					resource.TestCheckResourceAttr("ultradns_record.https", "ttl", "60"),
					resource.TestCheckResourceAttr("ultradns_record.https", "record_data.0", "1 www.ultradns.com. alpn=h3,h3-29,h2 mandatory=alpn"),
				),
			},
			{
				ResourceName:      "ultradns_record.https",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "ultradns_record.caa",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: getRecordHandlingConfig(
					acctest.TestAccResourceZonePrimary(zoneResource, zoneName),
					getResourceConfig("caa", "${resource.ultradns_zone.record_handling.name}", zoneName, "CAA", "60", "0 issue ultradns"),
					getDataSourceConfig("ns", "${resource.ultradns_zone.record_handling.name}", zoneName, "NS"),
					getDataSourceConfig("caa", "${resource.ultradns_zone.record_handling.name}", zoneName, "CAA"),
				),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_ns", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "owner_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "record_type", "NS"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "ttl", "60"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_ns", "record_data.#", "2"),
					acctest.TestAccCheckRecordResourceExists("data.ultradns_record.data_caa", ""),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "owner_name", zoneName),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "record_type", "CAA"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "ttl", "60"),
					resource.TestCheckResourceAttr("data.ultradns_record.data_caa", "record_data.0", "0 issue \"ultradns\""),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func getRecordHandlingConfig(v ...any) string {
	strFmt := ""
	for _, _ = range v {
		strFmt += "%s\n"
	}
	return fmt.Sprintf(strFmt, v...)
}

func getResourceConfig(rsName, zoneName, ownerName, rrType, ttl, rdata string) string {
	return fmt.Sprintf(`
	resource "ultradns_record" "%s" {
		zone_name = "%s"
		owner_name = "%s"
		record_type = "%s"
		ttl = %s
		record_data = ["%s"]
	}
	`, rsName, zoneName, ownerName, rrType, ttl, rdata)
}

func getDataSourceConfig(dsName, zoneName, ownerName, rrType string) string {
	return fmt.Sprintf(`
	data "ultradns_record" "data_%s" {
		zone_name = "%s"
		owner_name = "%s"
		record_type = "%s"
	}
	`, dsName, zoneName, ownerName, rrType)
}
