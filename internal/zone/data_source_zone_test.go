package zone_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func TestAccDataSourceZonePrimary(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	dataSourceName := "data.ultradns_zone.data_primary"

	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceZone(
					"primary",
					acctest.TestAccResourceZonePrimary("primary", zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", zoneName),
					resource.TestCheckResourceAttr(dataSourceName, "account_name", acctest.TestAccount),
					resource.TestCheckResourceAttr(dataSourceName, "type", zone.Primary),
					resource.TestCheckResourceAttr(dataSourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(dataSourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(dataSourceName, "owner", acctest.TestUsername),
					resource.TestCheckResourceAttr(dataSourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(dataSourceName, "notify_addresses.#", defaultCount),
					resource.TestCheckResourceAttr(dataSourceName, "restrict_ip.#", defaultCount),
				),
			},
			{
				Config: testAccDataSourceZone(
					"primary",
					acctest.TestAccResourceZonePrimary("primary", strings.ToUpper(zoneName)),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", zoneName),
					resource.TestCheckResourceAttr(dataSourceName, "account_name", acctest.TestAccount),
					resource.TestCheckResourceAttr(dataSourceName, "type", zone.Primary),
					resource.TestCheckResourceAttr(dataSourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(dataSourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(dataSourceName, "owner", acctest.TestUsername),
					resource.TestCheckResourceAttr(dataSourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(dataSourceName, "notify_addresses.#", defaultCount),
					resource.TestCheckResourceAttr(dataSourceName, "restrict_ip.#", defaultCount),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func TestAccDataSourceZoneSecondary(t *testing.T) {
	zoneName := acctest.TestSecondaryZone
	dataSourceName := "data.ultradns_zone.data_secondary"

	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceZone(
					"secondary",
					testAccResourceZoneSecondary(zoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", zoneName),
					resource.TestCheckResourceAttr(dataSourceName, "account_name", acctest.TestAccount),
					resource.TestCheckResourceAttr(dataSourceName, "type", zone.Secondary),
					resource.TestCheckResourceAttr(dataSourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(dataSourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(dataSourceName, "owner", acctest.TestUsername),
					resource.TestCheckResourceAttr(dataSourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(dataSourceName, "primary_name_server_1.0.ip", acctest.TestNameServer),
					resource.TestCheckResourceAttr(dataSourceName, "notification_email_address", "test@ultradns.com"),
					resource.TestCheckResourceAttr(dataSourceName, "transfer_status_details.0.last_refresh_status", "SUCCESSFUL"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func TestAccDataSourceZoneAlias(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	primaryZoneName := acctest.GetRandomZoneNameWithSpecialChar()
	dataSourceName := "data.ultradns_zone.data_alias"

	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceZone(
					"alias",
					testAccResourceZoneAlias(zoneName, primaryZoneName),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", zoneName),
					resource.TestCheckResourceAttr(dataSourceName, "account_name", acctest.TestAccount),
					resource.TestCheckResourceAttr(dataSourceName, "type", zone.Alias),
					resource.TestCheckResourceAttr(dataSourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(dataSourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(dataSourceName, "owner", acctest.TestUsername),
					resource.TestCheckResourceAttr(dataSourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(dataSourceName, "original_zone_name", primaryZoneName),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func testAccDataSourceZone(datasourceName, resource string) string {
	return fmt.Sprintf(`
	%[1]s
	data "ultradns_zone" "data_%[2]s" {
		name = "${resource.ultradns_zone.%[2]s.id}"
	}
	`, resource, datasourceName)
}
