package acctest

import (
	"context"
	"crypto/rand"
	"math/big"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/provider"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
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
	TestUsername  = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME")
	testPassword  = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD")
	testUserAgent = os.Getenv("ULTRADNS_UNIT_TEST_USER_AGENT")
)

var TestAccProviders map[string]*schema.Provider
var TestAccProvider *schema.Provider

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

func TestPreCheck(t *testing.T) {
	if TestUsername == "" {
		t.Fatal("username required for creating test client")
	}

	if testPassword == "" {
		t.Fatal("password required for creating test client")
	}

	if TestHost == "" {
		t.Fatal("host required for creating test client")
	}

	if testUserAgent == "" {
		t.Fatal("user agent required for creating test client")
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
