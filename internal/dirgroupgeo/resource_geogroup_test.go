package dirgroupgeo_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/geo"
)

const (
	resourceName = "ultradns_dirgroup_geo.ultradns_geo"
	resourceCode = "CA"
)

func TestAccResourceDirGroupGeo(t *testing.T) {
	geoData := &geo.DirGroupGeo{
		Name:        tfacctest.RandString(5),
		AccountName: acctest.TestAccount,
		Description: "Unit Test Geo Directional group description",
		Codes:       []string{"CA"},
	}

	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckDirGroupResourceDestroy(resourceName, geo.DirGroupType, geoData.DirGroupGeoID()),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDirGroupGeo(resourceName, geoData),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckDirGroupResourceExists(resourceName, geo.DirGroupType, geoData.DirGroupGeoID()),
					resource.TestCheckResourceAttr(resourceName, "name", geoData.Name),
					resource.TestCheckResourceAttr(resourceName, "account_name", geoData.AccountName),
					resource.TestCheckResourceAttr(resourceName, "codes.#", "1"),
				),
			},
			{
				Config: testAccResourceDirGroupGeo(resourceName, geoData),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckDirGroupResourceExists(resourceName, geo.DirGroupType, geoData.DirGroupGeoID()),
					resource.TestCheckResourceAttr(resourceName, "name", geoData.Name),
					resource.TestCheckResourceAttr(resourceName, "account_name", geoData.AccountName),
					resource.TestCheckResourceAttr(resourceName, "codes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", geoData.Description),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func testAccResourceDirGroupGeo(resourceName string, geoData *geo.DirGroupGeo) string {
	return fmt.Sprintf(`
	resource "ultradns_dirgroup_geo" "ultradns_geo" {
		name = "%s"
		account_name = "%s"
		description = "%s"
		codes = [ "CA" ]
	}`, geoData.Name, acctest.TestAccount, geoData.Description)
}
