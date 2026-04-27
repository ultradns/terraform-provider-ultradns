package acctest

import (
	"context"
	"os"
	"strings"
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
	cdnresource "github.com/ultradns/ultradns-go-sdk/pkg/cdn/resource"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/geo"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/ip"
)

const (
	randZoneNamePrefix                = "terraform-plugin-acc-test-"
	randCDNZoneNamePrefix             = "terraform-plugin-cdn-acc-test-"
	randZoneNameSuffix                = ".com."
	randZoneNameWithSpecialCharSuffix = ".in-addr.arpa."
	randStringLength                  = 5
	randCDNStringLength               = 10
	randSecondaryZoneCount            = 50
)

var (
	TestHost          = os.Getenv("ULTRADNS_UNIT_TEST_HOST_URL")
	TestAccount       = os.Getenv("ULTRADNS_UNIT_TEST_ACCOUNT")
	TestAccountCDN    = os.Getenv("ULTRADNS_UNIT_TEST_ACCOUNT_CDN")
	TestUsername      = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME")
	TestUsernameCDN   = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME_CDN")
	TestNameServer    = os.Getenv("ULTRADNS_UNIT_TEST_NAME_SERVER")
	TestSecondaryZone = os.Getenv("ULTRADNS_UNIT_TEST_SECONDARY_ZONE_NAME")
	testPassword      = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD")
	testPasswordCDN   = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD_CDN")
	testUserAgent     = os.Getenv("ULTRADNS_UNIT_TEST_USER_AGENT")
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

func NewTestAccProvidersCDN() map[string]*schema.Provider {
	providerCDN := provider.Provider()
	providerCDN.ConfigureContextFunc = getTestAccProviderConfigureContextFuncCDN

	return map[string]*schema.Provider{
		"ultradns": providerCDN,
	}
}

func getTestAccProviderConfigureContextFunc(c context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return configureTestAccProviderContext(TestUsername, testPassword)
}

func getTestAccProviderConfigureContextFuncCDN(c context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return configureTestAccProviderContext(TestUsernameCDN, testPasswordCDN)
}

func configureTestAccProviderContext(username, password string) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	cnf := client.Config{
		Username:  username,
		Password:  password,
		HostURL:   TestHost,
		UserAgent: testUserAgent,
	}

	client, err := client.NewClient(cnf)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	if os.Getenv("ULTRADNS_UNIT_TEST_DEBUG_HTTP") == "1" {
		client.EnableDefaultDebugLogger()
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

func TestPreCheckCDN(t *testing.T) func() {
	return func() {
		TestPreCheck(t)()

		if TestUsernameCDN == "" {
			t.Fatal("cdn username required for creating test client")
		}

		if testPasswordCDN == "" {
			t.Fatal("cdn password required for creating test client")
		}

		if TestAccountCDN == "" {
			t.Fatal("cdn account required for creating test client")
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
		switch resourceType {
		case ip.DirGroupType:
			_, dirGroupResponse, _, err := services.DirGroupIPService.Read(resourceID)
			if err != nil || dirGroupResponse.Name != rs.Primary.ID {
				return err
			}

		case geo.DirGroupType:
			_, dirGroupResponse, _, err := services.DirGroupGeoService.Read(resourceID)
			if err != nil || dirGroupResponse.Name != rs.Primary.ID {
				return err
			}
		}

		return nil
	}
}

func TestAccCheckCDNResourceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return errors.ResourceNotFoundError(resourceName)
		}

		accountName := rs.Primary.Attributes["account_name"]
		fqdn := rs.Primary.Attributes["fqdn"]

		services, err := getCDNService()
		if err != nil {
			return err
		}
		_, payload, err := services.CDNResourceService.Read(accountName, fqdn)
		if err != nil {
			return err
		}
		if payload == nil || !strings.EqualFold(payload.FQDN, fqdn) {
			return errors.ResourceNotFoundError(rs.Primary.ID)
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
			} else {
				errMsg := err.Error()
				if strings.Contains(errMsg, "70002") || strings.Contains(errMsg, "Data not found") || strings.Contains(errMsg, "1801") || strings.Contains(errMsg, "Zone does not exist") {
					// Ignore 'Data not found' and 'Zone does not exist' errors
					continue
				} else {
					return err
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

func TestAccCheckCDNResourceDestroy(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceName {
				continue
			}

			accountName := rs.Primary.Attributes["account_name"]
			fqdn := rs.Primary.Attributes["fqdn"]

			services, err := getCDNService()
			if err != nil {
				return err
			}
			_, listPayload, err := services.CDNResourceService.List(accountName, &cdnresource.ListOptions{Page: 1, Size: 1000})
			if err != nil {
				return err
			}

			if listPayload == nil {
				continue
			}

			for _, item := range listPayload.Content {
				if item != nil && strings.EqualFold(item.FQDN, fqdn) {
					return errors.ResourceNotDestroyedError(rs.Primary.ID)
				}
			}
		}

		return nil
	}
}

func getCDNService() (*service.Service, error) {
	cnf := client.Config{
		Username:  TestUsernameCDN,
		Password:  testPasswordCDN,
		HostURL:   TestHost,
		UserAgent: testUserAgent,
	}

	client, err := client.NewClient(cnf)
	if err != nil {
		return nil, err
	}

	return service.NewService(client)
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

func GetRandomCDNZoneName() string {
	return randCDNZoneNamePrefix + acctest.RandString(randCDNStringLength) + randZoneNameSuffix
}

func GetRandomCDNName() string {
	return "cdn-" + acctest.RandString(randCDNStringLength)
}

func GetRandomZoneNameWithSpecialChar() string {
	return randZoneNamePrefix + "/" + acctest.RandString(randStringLength) + "/" + acctest.RandString(randStringLength) + randZoneNameWithSpecialCharSuffix
}

// func GetRandomSecondaryZoneName() string {
// 	if num, err := rand.Int(rand.Reader, big.NewInt(randSecondaryZoneCount)); err == nil {
// 		return randZoneNamePrefix + num.String() + randZoneNameSuffix
// 	}

// 	return randZoneNamePrefix + "0" + randZoneNameSuffix
// }
