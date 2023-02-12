package dirgroupgeo_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/geo"
)

func TestAccDataSourceDirGroupGeo(t *testing.T) {
	geoData := &geo.DirGroupGeo{
		Name:        tfacctest.RandString(5),
		AccountName: acctest.TestAccount,
		Description: "Unit Test Geo Directional group description",
		Codes:       []string{"CA"},
	}
	dataSourceName := "data.ultradns_dirgroup_geo.data_ultradns_geo"
	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckDirGroupResourceDestroy(resourceName, geo.DirGroupType, geoData.DirGroupGeoID()),

		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDirGroupGeo("ultradns_geo", testAccResourceDirGroupGeo("ultradns_geo", geoData)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", geoData.Name),
					resource.TestCheckResourceAttr(dataSourceName, "account_name", geoData.AccountName),
					resource.TestCheckResourceAttr(dataSourceName, "description", geoData.Description),
					resource.TestCheckResourceAttr(dataSourceName, "codes.#", "1"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func testAccDataSourceDirGroupGeo(datasourceName, resource string) string {
	return fmt.Sprintf(`
	%[1]s
	data "ultradns_dirgroup_geo" "data_%[2]s" {
		name = "${resource.ultradns_dirgroup_geo.%[2]s.id}"
	}
`, resource, datasourceName)
}
