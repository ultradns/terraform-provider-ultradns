package cdn_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	cdnresource "github.com/ultradns/ultradns-go-sdk/pkg/cdn/resource"
)

func TestAccDataSourceCDN(t *testing.T) {
	fqdn := acctest.GetRandomCDNZoneName()
	cdnFQDN := "www." + fqdn
	dataSourceName := "data.ultradns_cdn.data_single"
	resourceName := "single"
	name := acctest.GetRandomCDNName()

	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheckCDN(t),
		Providers:    acctest.NewTestAccProvidersCDN(),
		CheckDestroy: acctest.TestAccCheckCDNResourceDestroy("ultradns_cdn"),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccDataSourceCDN(
					"single",
					resourceName,
					acctest.TestAccResourceCDN(resourceName, fqdn, cdnresource.TypeSynthetic, name, "Synthetic data source test"),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "account_name", acctest.TestAccountCDN),
					resource.TestCheckResourceAttr(dataSourceName, "fqdn", cdnFQDN),
					resource.TestCheckResourceAttr(dataSourceName, "type", cdnresource.TypeSynthetic),
					resource.TestCheckResourceAttr(dataSourceName, "name", name),
					resource.TestCheckResourceAttrSet(dataSourceName, "resource_id"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func TestAccDataSourceCDNs(t *testing.T) {
	byodFQDN := acctest.GetRandomCDNZoneName()
	syntheticFQDN := acctest.GetRandomCDNZoneName()
	dataSourceName := "data.ultradns_cdns.data_all"
	byodName := acctest.GetRandomCDNName()
	syntheticName := acctest.GetRandomCDNName()

	config := fmt.Sprintf(`
		%s
		%s
		%s
	`,
		acctest.TestAccResourceCDN("byod", byodFQDN, cdnresource.TypeBYOD, byodName, "BYOD list test"),
		acctest.TestAccResourceCDN("synthetic", syntheticFQDN, cdnresource.TypeSynthetic, syntheticName, "Synthetic list test"),
		acctest.TestAccDataSourceCDNs("all", "byod", 1, 100, ""),
	)

	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheckCDN(t),
		Providers:    acctest.NewTestAccProvidersCDN(),
		CheckDestroy: acctest.TestAccCheckCDNResourceDestroy("ultradns_cdn"),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "account_name", acctest.TestAccountCDN),
					resource.TestCheckResourceAttrSet(dataSourceName, "total_pages"),
					resource.TestCheckResourceAttrSet(dataSourceName, "total_elements"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}


