package zone_test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
// 	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
// 	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
// )

// func TestAccDataSourceZonePrimary(t *testing.T) {
// 	zoneName := acctest.GetRandomZoneName()
// 	dataSourceName := "data.ultradns_zone.data_primary"

// 	testCase := resource.TestCase{
// 		PreCheck:     func() { acctest.TestPreCheck(t) },
// 		Providers:    acctest.TestAccProviders,
// 		CheckDestroy: testAccCheckZoneDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccDataSourceZonePrimary(zoneName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(dataSourceName, "name", zoneName),
// 					resource.TestCheckResourceAttr(dataSourceName, "account_name", acctest.TestUsername),
// 					resource.TestCheckResourceAttr(dataSourceName, "type", zone.Primary),
// 					resource.TestCheckResourceAttr(dataSourceName, "dnssec_status", defaultDNSSECStatus),
// 					resource.TestCheckResourceAttr(dataSourceName, "status", defaultZoneStatus),
// 					resource.TestCheckResourceAttr(dataSourceName, "owner", acctest.TestUsername),
// 					resource.TestCheckResourceAttr(dataSourceName, "resource_record_count", defaultCount),
// 					resource.TestCheckResourceAttr(dataSourceName, "inherit", "TSIG"),
// 				),
// 			},
// 		},
// 	}
// 	resource.ParallelTest(t, testCase)
// }

// func TestAccDataSourceZoneSecondary(t *testing.T) {
// 	zoneName := acctest.GetRandomSecondaryZoneName()
// 	dataSourceName := "data.ultradns_zone.data_secondary"

// 	testCase := resource.TestCase{
// 		PreCheck:     func() { acctest.TestPreCheck(t) },
// 		Providers:    acctest.TestAccProviders,
// 		CheckDestroy: testAccCheckZoneDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccDataSourceZoneSecondary(zoneName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(dataSourceName, "name", zoneName),
// 					resource.TestCheckResourceAttr(dataSourceName, "account_name", acctest.TestUsername),
// 					resource.TestCheckResourceAttr(dataSourceName, "type", zone.Secondary),
// 					resource.TestCheckResourceAttr(dataSourceName, "dnssec_status", defaultDNSSECStatus),
// 					resource.TestCheckResourceAttr(dataSourceName, "status", defaultZoneStatus),
// 					resource.TestCheckResourceAttr(dataSourceName, "owner", acctest.TestUsername),
// 					resource.TestCheckResourceAttr(dataSourceName, "resource_record_count", defaultCount),
// 				),
// 			},
// 		},
// 	}
// 	resource.ParallelTest(t, testCase)
// }

// func TestAccDataSourceZoneAlias(t *testing.T) {
// 	zoneName := acctest.GetRandomZoneName()
// 	dataSourceName := "data.ultradns_zone.data_alias"

// 	testCase := resource.TestCase{
// 		PreCheck:     func() { acctest.TestPreCheck(t) },
// 		Providers:    acctest.TestAccProviders,
// 		CheckDestroy: testAccCheckZoneDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccDataSourceZoneAlias(zoneName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(dataSourceName, "name", zoneName),
// 					resource.TestCheckResourceAttr(dataSourceName, "account_name", acctest.TestUsername),
// 					resource.TestCheckResourceAttr(dataSourceName, "type", zone.Alias),
// 					resource.TestCheckResourceAttr(dataSourceName, "dnssec_status", defaultDNSSECStatus),
// 					resource.TestCheckResourceAttr(dataSourceName, "status", defaultZoneStatus),
// 					resource.TestCheckResourceAttr(dataSourceName, "owner", acctest.TestUsername),
// 					resource.TestCheckResourceAttr(dataSourceName, "resource_record_count", defaultCount),
// 				),
// 			},
// 		},
// 	}
// 	resource.ParallelTest(t, testCase)
// }

// func testAccDataSourceZonePrimary(zoneName string) string {
// 	return fmt.Sprintf(`
// 	%s

// 	data "ultradns_zone" "data_primary" {
// 		name = "${resource.ultradns_zone.primary.id}"
// 	}
// 	`, testAccResourceZonePrimary(zoneName))
// }

// func testAccDataSourceZoneSecondary(zoneName string) string {
// 	return fmt.Sprintf(`
// 	%s

// 	data "ultradns_zone" "data_secondary" {
// 		name = "${resource.ultradns_zone.secondary.id}"
// 	}
// 	`, testAccResourceZoneSecondary(zoneName))
// }

// func testAccDataSourceZoneAlias(zoneName string) string {
// 	return fmt.Sprintf(`
// 	%s

// 	data "ultradns_zone" "data_alias" {
// 		name = "${resource.ultradns_zone.alias.id}"
// 	}
// 	`, testAccResourceZoneAlias(zoneName))
// }
