package dirgroupip_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/ip"
)

func TestAccDataSourceDirGroupIP(t *testing.T) {
	dataSourceName := "data.ultradns_dirgroup_ip.data_ultradns_ip"
	ipData := &ip.DirGroupIP{
		Name:        tfacctest.RandString(5),
		AccountName: acctest.TestAccount,
		Description: "Unit Test IP Directional group description",
		IPs: []*ip.IPAddress{
			&ip.IPAddress{
				Address: "192.168.3.4",
			},
			&ip.IPAddress{
				Start: "192.168.1.0",
				End:   "192.168.1.10",
			},
			&ip.IPAddress{
				Cidr: "192.168.2.0/24",
			},
		},
	}

	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckDirGroupResourceDestroy(resourceName, ip.DirGroupType, ipData.DirGroupIPID()),

		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDirGroupIP("ultradns_ip", testAccResourceDirGroupIP("ultradns_ip", ipData)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", ipData.Name),
					resource.TestCheckResourceAttr(dataSourceName, "account_name", ipData.AccountName),
					resource.TestCheckResourceAttr(dataSourceName, "description", ipData.Description),
					resource.TestCheckResourceAttr(dataSourceName, "ip.#", "3"),
				),
			},
		},
	}
	resource.ParallelTest(t, testCase)
}

func testAccDataSourceDirGroupIP(datasourceName, resource string) string {
	return fmt.Sprintf(`
	%[1]s
	data "ultradns_dirgroup_ip" "data_%[2]s" {
		name = "${resource.ultradns_dirgroup_ip.%[2]s.id}"
	}
`, resource, datasourceName)
}
