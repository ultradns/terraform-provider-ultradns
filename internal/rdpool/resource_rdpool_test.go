package rdpool_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
)

func TestAccResourceRDPool(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()

	testCase := resource.TestCase{
		PreCheck:     func() { acctest.TestPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckRDPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRDPoolA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRDPoolExists("ultradns_rdpool.a"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "record_type", "A"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "record_data.0", "192.168.1.1"),
					resource.TestCheckResourceAttr("ultradns_rdpool.a", "order", "FIXED"),
				),
			},
			{
				ResourceName:      "ultradns_rdpool.a",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceRDPoolAAAA(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRDPoolExists("ultradns_rdpool.aaaa"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "zone_name", zoneName),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "record_type", "AAAA"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "ttl", "120"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "record_data.0", "2001:db8:85a3:0:0:8a2e:370:7334"),
					resource.TestCheckResourceAttr("ultradns_rdpool.aaaa", "order", "ROUND_ROBIN"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func testAccCheckRDPoolExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return errors.ResourceNotFoundError(resourceName)
		}

		services := acctest.TestAccProvider.Meta().(*service.Service)
		rrSetKey := rrset.GetRRSetKeyFromID(rs.Primary.ID)
		_, _, err := services.RDPoolService.ReadRDPool(rrSetKey)

		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckRDPoolDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ultradns_rdpool" {
			continue
		}

		services := acctest.TestAccProvider.Meta().(*service.Service)
		rrSetKey := rrset.GetRRSetKeyFromID(rs.Primary.ID)
		_, rdPoolResponse, err := services.RDPoolService.ReadRDPool(rrSetKey)

		if err == nil {
			if len(rdPoolResponse.RRSets) > 0 && rdPoolResponse.RRSets[0].OwnerName == rrSetKey.Name {
				return errors.ResourceNotDestroyedError(rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccResourceZonePrimary(zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "primary_rdpool" {
		name        = "%s"
		account_name = "%s"
		type        = "PRIMARY"
		primary_create_info {
			create_type = "NEW"
		}
	}
	`, zoneName, acctest.TestUsername)
}

func testAccResourceRDPoolA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_rdpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_rdpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_rdpool.id}"
		record_type = "A"
		ttl = 120
		record_data = ["192.168.1.1"]
		order = "FIXED"
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}

func testAccResourceRDPoolAAAA(zoneName string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_rdpool" "aaaa" {
		zone_name = "${resource.ultradns_zone.primary_rdpool.id}"
		owner_name = "%s"
		record_type = "AAAA"
		ttl = 120
		record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
		order = "ROUND_ROBIN"
	}
	`, testAccResourceZonePrimary(zoneName), tfacctest.RandString(3))
}
