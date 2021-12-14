package zone_test

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

func init() {
	resource.AddTestSweepers("ultradns_zone", &resource.Sweeper{
		Name: "ultradns_zone",
		F:    testAccZoneSweeper,
	})
}

func testAccZoneSweeper(r string) error {
	services := acctest.TestAccProvider.Meta().(*service.Service)
	cursor := ""
	initial := true

	for {
		query := testAccGetZoneQueryString(cursor)
		_, zoneList, err := services.ZoneService.ListZone(query)

		if err != nil {
			return err
		}

		for _, zone := range zoneList.Zones {
			if strings.HasPrefix(zone.Properties.Name, "test-acc") {
				_, er := services.ZoneService.DeleteZone(zone.Properties.Name)
				if er != nil {
					log.Printf("error destroying %s during sweep: %s", zone.Properties.Name, er)
				}
			}
		}

		if zoneList.CursorInfo.Next == "" && !initial {
			return nil
		}

		initial = false
		cursor = zoneList.CursorInfo.Next
	}
}

func testAccGetZoneQueryString(cursor string) *helper.QueryInfo {
	return &helper.QueryInfo{
		Limit:  1000,
		Cursor: cursor,
	}
}
