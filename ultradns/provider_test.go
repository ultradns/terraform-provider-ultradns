package ultradns

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

var (
	testUsername  = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME")
	testPassword  = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD")
	testHost      = os.Getenv("ULTRADNS_UNIT_TEST_HOST_URL")
	testVersion   = os.Getenv("ULTRADNS_UNIT_TEST_API_VERSION")
	testUserAgent = os.Getenv("ULTRADNS_UNIT_TEST_USER_AGENT")
	testZoneName  = os.Getenv("ULTRADNS_UNIT_TEST_ZONE_NAME")
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProvider.ConfigureContextFunc = getTestAccProviderConfigureContextFunc
	testAccProviders = map[string]*schema.Provider{
		"ultradns": testAccProvider,
	}
}

func TestProvider(t *testing.T) {

	if err := Provider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}

func TestAccPreCheck(t *testing.T) {
	if testUsername == "" {
		t.Fatal("username required")
	}

	if testPassword == "" {
		t.Fatal("password required")
	}

	if testHost == "" {
		t.Fatal("host required")
	}

	if testVersion == "" {
		t.Fatal("version required")
	}

	if testUserAgent == "" {
		t.Fatal("user agent required")
	}

	if testZoneName == "" {
		t.Fatal("zone name required")
	}
}

func getTestAccProviderConfigureContextFunc(c context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	client, err := ultradns.NewClient(testUsername, testPassword, testHost, testVersion, testUserAgent)

	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, diags
}
