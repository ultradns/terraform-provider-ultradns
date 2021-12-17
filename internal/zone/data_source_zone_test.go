package zone_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
)

func TestAccDataSourceZonePrimary(t *testing.T) {
	zoneName := fmt.Sprintf("test-acc-%s.com.", tfacctest.RandString(5))
	dataSourceName := "data.ultradns_zone.data_primary"

	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceZonePrimary(zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", zoneName),
					resource.TestCheckResourceAttr(dataSourceName, "account_name", acctest.TestUsername),
					resource.TestCheckResourceAttr(dataSourceName, "type", primaryZoneType),
					resource.TestCheckResourceAttr(dataSourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(dataSourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(dataSourceName, "owner", acctest.TestUsername),
					resource.TestCheckResourceAttr(dataSourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(dataSourceName, "inherit", "ALL"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func TestAccDataSourceZoneAlias(t *testing.T) {
	zoneName := fmt.Sprintf("test-acc-%s.com.", tfacctest.RandString(5))
	dataSourceName := "data.ultradns_zone.data_alias"

	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceZoneAlias(zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", zoneName),
					resource.TestCheckResourceAttr(dataSourceName, "account_name", acctest.TestUsername),
					resource.TestCheckResourceAttr(dataSourceName, "type", aliasZoneType),
					resource.TestCheckResourceAttr(dataSourceName, "dnssec_status", defaultDNSSECStatus),
					resource.TestCheckResourceAttr(dataSourceName, "status", defaultZoneStatus),
					resource.TestCheckResourceAttr(dataSourceName, "owner", acctest.TestUsername),
					resource.TestCheckResourceAttr(dataSourceName, "resource_record_count", defaultCount),
					resource.TestCheckResourceAttr(dataSourceName, "original_zone_name", testAccGetPrimaryZoneNameForAlias(zoneName)),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func testAccDataSourceZonePrimary(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "resource_primary" {
		name        = "%s"
		account_name = "%s"
		type        = "PRIMARY"
		primary_create_info {
			create_type = "NEW"
		}
	}

	data "ultradns_zone" "data_primary" {
		name = "${resource.ultradns_zone.resource_primary.id}"
	}
	`, zoneName, acctest.TestUsername)
}

func testAccDataSourceZoneAlias(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "resource_primary" {
		name        = "%s"
		account_name = "%s"
		type        = "PRIMARY"
		primary_create_info {
			create_type = "NEW"
		}
	}

	resource "ultradns_zone" "resource_alias" {
		name        = "%s"
		account_name = "%s"
		type        = "ALIAS"
		alias_create_info {
			  original_zone_name = "${resource.ultradns_zone.resource_primary.id}"
		}
	  }

	data "ultradns_zone" "data_alias" {
		name = "${resource.ultradns_zone.resource_alias.id}"
	}
	`, testAccGetPrimaryZoneNameForAlias(zoneName), acctest.TestUsername, zoneName, acctest.TestUsername)
}
