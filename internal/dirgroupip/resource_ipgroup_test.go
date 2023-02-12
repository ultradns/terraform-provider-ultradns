package dirgroupip_test

import (
	"fmt"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/ip"
)

const (
	resourceName = "ultradns_dirgroup_ip.ultradns_ip"
	resourceIP   = "192.168.1.1"
)

func TestAccResourceDirGroupIP(t *testing.T) {
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
				Config: testAccResourceDirGroupIP(resourceName, ipData),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckDirGroupResourceExists(resourceName, ip.DirGroupType, ipData.DirGroupIPID()),
					resource.TestCheckResourceAttr(resourceName, "name", ipData.Name),
					resource.TestCheckResourceAttr(resourceName, "account_name", ipData.AccountName),
					resource.TestCheckResourceAttr(resourceName, "ip.#", "3"),
				),
			},
			{
				Config: testAccResourceDirGroupIP(resourceName, ipData),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckDirGroupResourceExists(resourceName, ip.DirGroupType, ipData.DirGroupIPID()),
					resource.TestCheckResourceAttr(resourceName, "name", ipData.Name),
					resource.TestCheckResourceAttr(resourceName, "account_name", ipData.AccountName),
					resource.TestCheckResourceAttr(resourceName, "ip.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "description", ipData.Description),
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

func testAccResourceDirGroupIP(resourceName string, ipData *ip.DirGroupIP) string {
	return fmt.Sprintf(`
	resource "ultradns_dirgroup_ip" "ultradns_ip" {
		name = "%s"
		account_name = "%s"
		description = "%s"
		ip{
			address = "192.168.3.4"
		}
		ip{
			start = "192.168.1.0"
			end   = "192.168.1.10"
		}
		ip{
			cidr  = "192.168.2.0/24"
		}
	}`, ipData.Name, acctest.TestAccount, ipData.Description)
}
