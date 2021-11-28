package ultradns

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

var (
	testUsername  = os.Getenv("ULTRADNS_UNIT_TEST_USERNAME")
	testPassword  = os.Getenv("ULTRADNS_UNIT_TEST_PASSWORD")
	testHost      = os.Getenv("ULTRADNS_UNIT_TEST_HOST_URL")
	testVersion   = os.Getenv("ULTRADNS_UNIT_TEST_API_VERSION")
	testUserAgent = os.Getenv("ULTRADNS_UNIT_TEST_USER_AGENT")
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

func TestMain(m *testing.M) {
	testAccProvider.Configure(context.TODO(), terraform.NewResourceConfigRaw(make(map[string]interface{})))
	resource.TestMain(m)
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}

func TestAccPreCheck(t *testing.T) {
	if testUsername == "" {
		t.Fatal("username required for creating test client")
	}

	if testPassword == "" {
		t.Fatal("password required for creating test client")
	}

	if testHost == "" {
		t.Fatal("host required for creating test client")
	}

	if testVersion == "" {
		t.Fatal("version required for creating test client")
	}

	if testUserAgent == "" {
		t.Fatal("user agent required for creating test client")
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
