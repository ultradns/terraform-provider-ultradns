package ultradns

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func init() {
	resource.AddTestSweepers("ultradns_zone", &resource.Sweeper{
		Name: "ultradns_zone",
		F:    testAccZoneSweeper,
	})
}

func testAccZoneSweeper(r string) error {
	client := testAccProvider.Meta().(*ultradns.Client)
	offset := 0
	totalCount := 1
	for totalCount > offset {
		queryString := testAccGetZoneQueryString(offset)
		_, zoneList, err := client.ListZone(queryString)
		if err != nil {
			return err
		}
		for _, zone := range zoneList.Zones {
			if strings.HasPrefix(zone.Properties.Name, "test-acc") {
				_, er := client.DeleteZone(zone.Properties.Name)
				if er != nil {
					fmt.Errorf("error destroying %s during sweep: %s", zone.Properties.Name, er)
				}
			}
		}
		totalCount = zoneList.ResultInfo.TotalCount
		offset += 1000
	}

	return nil
}

func testAccGetZoneQueryString(offset int) string {
	return fmt.Sprintf("?&limit=%v&offset=%v", 1000, offset)
}
