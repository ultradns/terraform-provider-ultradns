package acctest

import (
	"context"
	"crypto/rand"
	"math/big"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
	"github.com/ultradns/terraform-provider-ultradns/internal/provider"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/geo"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/ip"
)

const (
	randZoneNamePrefix                = "terraform-plugin-acc-test-"
	randZoneNameSuffix                = ".com."
	randZoneNameWithSpecialCharSuffix = ".in-addr.arpa."
	randStringLength                  = 5
	randSecondaryZoneCount            = 50
)

var (
	TestHost      = os.Getenv("ULTRADNS_UNIT_TEST_HOST_URL")
	TestAccount   = os.Getenv("ULTRADNS_UNIT_TEST_ACCOUNT")
	TestUsername  = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME")
	testPassword  = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD")
	testUserAgent = os.Getenv("ULTRADNS_UNIT_TEST_USER_AGENT")
)

var (
	TestAccProviders map[string]*schema.Provider
	TestAccProvider  *schema.Provider
)

func init() {
	TestAccProvider = provider.Provider()
	TestAccProvider.ConfigureContextFunc = getTestAccProviderConfigureContextFunc
	TestAccProviders = map[string]*schema.Provider{
		"ultradns": TestAccProvider,
	}
}

func getTestAccProviderConfigureContextFunc(c context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	cnf := client.Config{
		Username:  TestUsername,
		Password:  testPassword,
		HostURL:   TestHost,
		UserAgent: testUserAgent,
	}

	client, err := client.NewClient(cnf)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	service, err := service.NewService(client)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return service, diags
}

func TestPreCheck(t *testing.T) func() {
	return func() {
		if TestUsername == "" {
			t.Fatal("username required for creating test client")
		}

		if testPassword == "" {
			t.Fatal("password required for creating test client")
		}

		if TestHost == "" {
			t.Fatal("host required for creating test client")
		}

		if TestAccount == "" {
			t.Fatal("account required for creating test client")
		}

		if testUserAgent == "" {
			t.Fatal("user agent required for creating test client")
		}
	}
}

func TestAccCheckRecordResourceExists(resourceName, pType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return errors.ResourceNotFoundError(resourceName)
		}

		services := TestAccProvider.Meta().(*service.Service)
		rrSetKey := rrset.GetRRSetKeyFromID(rs.Primary.ID)
		rrSetKey.PType = pType
		_, _, err := services.RecordService.Read(rrSetKey)
		if err != nil {
			return err
		}

		return nil
	}
}

func TestAccCheckProbeResourceExists(resourceName, pType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return errors.ResourceNotFoundError(resourceName)
		}

		services := TestAccProvider.Meta().(*service.Service)
		rrSetKey := probe.GetRRSetKeyFromID(rs.Primary.ID)
		rrSetKey.PType = pType
		_, _, err := services.ProbeService.Read(rrSetKey)
		if err != nil {
			return err
		}

		return nil
	}
}

func TestAccCheckDirGroupResourceExists(resourceName, resourceType, resourceID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return errors.ResourceNotFoundError(resourceName)
		}

		services := TestAccProvider.Meta().(*service.Service)
		switch {
		case resourceType == ip.DirGroupType:
			_, dirGroupResponse, _, err := services.DirGroupIPService.Read(resourceID)
			if err != nil || dirGroupResponse.Name != rs.Primary.ID {
				return err
			}

		case resourceType == geo.DirGroupType:
			_, dirGroupResponse, _, err := services.DirGroupGeoService.Read(resourceID)
			if err != nil || dirGroupResponse.Name != rs.Primary.ID {
				return err
			}
		}

		return nil
	}
}
func TestAccCheckRecordResourceDestroy(resourceName, pType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceName {
				continue
			}

			services := TestAccProvider.Meta().(*service.Service)
			rrSetKey := rrset.GetRRSetKeyFromID(rs.Primary.ID)
			rrSetKey.PType = pType
			_, response, err := services.RecordService.Read(rrSetKey)

			if err == nil {
				if len(response.RRSets) > 0 && response.RRSets[0].OwnerName == rrSetKey.Owner {
					return errors.ResourceNotDestroyedError(rs.Primary.ID)
				}
			}
		}

		return nil
	}
}

func TestAccCheckProbeResourceDestroy(resourceName, pType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceName {
				continue
			}

			services := TestAccProvider.Meta().(*service.Service)
			rrSetKey := rrset.GetRRSetKeyFromID(rs.Primary.ID)
			rrSetKey.PType = pType
			_, response, err := services.ProbeService.Read(rrSetKey)

			if err == nil {
				if response.Type == pType {
					return errors.ResourceNotDestroyedError(rs.Primary.ID)
				}
			}
		}

		return nil
	}
}

func TestAccCheckDirGroupResourceDestroy(resourceName, resourceType, resourceID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceName {
				continue
			}

			services := TestAccProvider.Meta().(*service.Service)
			_, dirGroupResponse, _, err := services.DirGroupGeoService.Read(resourceID)
			if err == nil {
				if dirGroupResponse.Name == rs.Primary.ID {
					return errors.ResourceNotDestroyedError(rs.Primary.ID)
				}
			}
		}
		return nil
	}
}
func GetRandomZoneName() string {
	return randZoneNamePrefix + acctest.RandString(randStringLength) + randZoneNameSuffix
}

func GetRandomZoneNameWithSpecialChar() string {
	return randZoneNamePrefix + "/" + acctest.RandString(randStringLength) + "/" + acctest.RandString(randStringLength) + randZoneNameWithSpecialCharSuffix
}

func GetRandomSecondaryZoneName() string {
	if num, err := rand.Int(rand.Reader, big.NewInt(randSecondaryZoneCount)); err == nil {
		return randZoneNamePrefix + num.String() + randZoneNameSuffix
	}

	return randZoneNamePrefix + "0" + randZoneNameSuffix
}
