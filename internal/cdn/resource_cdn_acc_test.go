package cdn_test

import (
	"regexp"
	"testing"

	tfacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	cdnresource "github.com/ultradns/ultradns-go-sdk/pkg/cdn/resource"
)

func TestAccResourceCDNBYOD(t *testing.T) {
	fqdn := acctest.GetRandomZoneName()
	cdnFQDN := "www." + fqdn
	resourceName := "ultradns_cdn.byod"
	name := "cdn-" + tfacctest.RandString(6)
	updatedName := "cdn-" + tfacctest.RandString(6)

	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckCDNResourceDestroy("ultradns_cdn"),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccResourceCDN("byod", fqdn, cdnresource.TypeBYOD, name, "BYOD acceptance test"),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckCDNResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "account_name", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "fqdn", cdnFQDN),
					resource.TestCheckResourceAttr(resourceName, "type", cdnresource.TypeBYOD),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "BYOD acceptance test"),
					resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
					// Verify config_properties keys are present in state.
					resource.TestCheckResourceAttrSet(resourceName, "config_properties.cdnEnablementMap"),
					resource.TestCheckResourceAttrSet(resourceName, "config_properties.trafficDistribution"),
					// Verify preference_properties keys are present in state.
					resource.TestCheckResourceAttrSet(resourceName, "preference_properties.availabilityThresholds"),
					resource.TestCheckResourceAttrSet(resourceName, "preference_properties.performanceFiltering"),
					resource.TestCheckResourceAttrSet(resourceName, "preference_properties.enabledSubdivisionCountries"),
				),
			},
			{
				Config: acctest.TestAccResourceCDN("byod", fqdn, cdnresource.TypeBYOD, updatedName, "BYOD acceptance update"),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckCDNResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "fqdn", cdnFQDN),
					resource.TestCheckResourceAttr(resourceName, "type", cdnresource.TypeBYOD),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "BYOD acceptance update"),
					// Re-verify config/preference keys survive an update round-trip.
					resource.TestCheckResourceAttrSet(resourceName, "config_properties.cdnEnablementMap"),
					resource.TestCheckResourceAttrSet(resourceName, "config_properties.trafficDistribution"),
					resource.TestCheckResourceAttrSet(resourceName, "preference_properties.availabilityThresholds"),
					resource.TestCheckResourceAttrSet(resourceName, "preference_properties.performanceFiltering"),
					resource.TestCheckResourceAttrSet(resourceName, "preference_properties.enabledSubdivisionCountries"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cdn_providers", "config_properties", "preference_properties"},
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func TestAccResourceCDNSynthetic(t *testing.T) {
	fqdn := acctest.GetRandomZoneName()
	cdnFQDN := "www." + fqdn
	resourceName := "ultradns_cdn.synthetic"
	name := "cdn-" + tfacctest.RandString(6)
	updatedName := "cdn-" + tfacctest.RandString(6)

	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: acctest.TestAccCheckCDNResourceDestroy("ultradns_cdn"),
		Steps: []resource.TestStep{
			{
				Config: acctest.TestAccResourceCDN("synthetic", fqdn, cdnresource.TypeSynthetic, name, "Synthetic acceptance test"),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckCDNResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "account_name", acctest.TestAccount),
					resource.TestCheckResourceAttr(resourceName, "fqdn", cdnFQDN),
					resource.TestCheckResourceAttr(resourceName, "type", cdnresource.TypeSynthetic),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				),
			},
			{
				Config: acctest.TestAccResourceCDN("synthetic", fqdn, cdnresource.TypeSynthetic, updatedName, "Synthetic acceptance update"),
				Check: resource.ComposeTestCheckFunc(
					acctest.TestAccCheckCDNResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "fqdn", cdnFQDN),
					resource.TestCheckResourceAttr(resourceName, "type", cdnresource.TypeSynthetic),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cdn_providers", "config_properties", "preference_properties"},
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

// TestAccResourceCDNInvalidClientCdnID verifies that the schema-level
// ValidateFunc rejects a clientCdnId that does not match ^[A-Za-z0-9\-_]{1,64}$
// before any API call is made.
func TestAccResourceCDNInvalidClientCdnID(t *testing.T) {
	fqdn := acctest.GetRandomZoneName()
	name := "cdn-" + tfacctest.RandString(6)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  acctest.TestPreCheck(t),
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      acctest.TestAccResourceCDNWithClientCdnID("neg", fqdn, cdnresource.TypeBYOD, name, "invalid cdn id!"),
				ExpectError: regexp.MustCompile("must match"),
			},
		},
	})
}

// TestAccResourceCDNInvalidType verifies that the schema-level ValidateFunc
// rejects an unknown CDN type before any API call is made.
func TestAccResourceCDNInvalidType(t *testing.T) {
	fqdn := acctest.GetRandomZoneName()
	name := "cdn-" + tfacctest.RandString(6)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  acctest.TestPreCheck(t),
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      acctest.TestAccResourceCDN("neg", fqdn, "INVALID_TYPE", name, "type validation test"),
				ExpectError: regexp.MustCompile("expected type to be one of"),
			},
		},
	})
}