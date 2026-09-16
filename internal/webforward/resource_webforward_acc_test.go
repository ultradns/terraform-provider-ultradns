package webforward_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkwebforward "github.com/ultradns/ultradns-go-sdk/pkg/webforward"
)

const (
	webForwardZoneResource = "primary_webforward"
	webForwardResourceName = "ultradns_webforward.www"
)

func TestAccResourceWebForward(t *testing.T) {
	zoneName := acctest.GetRandomZoneName()
	requestTo := "www." + strings.TrimSuffix(zoneName, ".")

	testCase := resource.TestCase{
		PreCheck:     acctest.TestPreCheck(t),
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckWebForwardDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWebForward(zoneName, requestTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebForwardExists(webForwardResourceName),
					resource.TestCheckResourceAttr(webForwardResourceName, "zone_name", zoneName),
					resource.TestCheckResourceAttr(webForwardResourceName, "request_to", requestTo),
					resource.TestCheckResourceAttr(webForwardResourceName, "default_redirect_to", "https://example.com/"),
					resource.TestCheckResourceAttr(webForwardResourceName, "default_forward_type", sdkwebforward.HTTP301Redirect),
					resource.TestCheckResourceAttr(webForwardResourceName, "relative_forward_type", sdkwebforward.ParameterAndPath),
					resource.TestCheckResourceAttrSet(webForwardResourceName, "guid"),
				),
			},
			{
				Config: testAccResourceWebForwardUpdate(zoneName, requestTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebForwardExists(webForwardResourceName),
					resource.TestCheckResourceAttr(webForwardResourceName, "zone_name", zoneName),
					resource.TestCheckResourceAttr(webForwardResourceName, "request_to", requestTo),
					resource.TestCheckResourceAttr(webForwardResourceName, "default_redirect_to", "https://example.org/"),
					resource.TestCheckResourceAttr(webForwardResourceName, "default_forward_type", sdkwebforward.HTTP302Redirect),
					resource.TestCheckResourceAttr(webForwardResourceName, "relative_forward_type", sdkwebforward.Path),
				),
			},
			{
				ResourceName:      webForwardResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataSourceWebForward(zoneName, requestTo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.ultradns_webforward.by_guid", "request_to", webForwardResourceName, "request_to"),
					resource.TestCheckResourceAttrPair("data.ultradns_webforward.by_guid", "default_forward_type", webForwardResourceName, "default_forward_type"),
					resource.TestCheckResourceAttrPair("data.ultradns_webforward.by_request_to", "guid", webForwardResourceName, "guid"),
					resource.TestCheckResourceAttrPair("data.ultradns_webforward.by_request_to", "default_redirect_to", webForwardResourceName, "default_redirect_to"),
				),
			},
		},
	}

	resource.ParallelTest(t, testCase)
}

func testAccResourceWebForward(zoneName, requestTo string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_webforward" "www" {
		zone_name = "${resource.ultradns_zone.%s.id}"
		request_to = "%s"
		default_redirect_to = "https://example.com"
		default_forward_type = "HTTP_301_REDIRECT"
		relative_forward_type = "PARAMETER_AND_PATH"
	}
	`, acctest.TestAccResourceZonePrimary(webForwardZoneResource, zoneName), webForwardZoneResource, requestTo)
}

func testAccResourceWebForwardUpdate(zoneName, requestTo string) string {
	return fmt.Sprintf(`
	%s

	resource "ultradns_webforward" "www" {
		zone_name = "${resource.ultradns_zone.%s.id}"
		request_to = "%s"
		default_redirect_to = "https://example.org"
		default_forward_type = "HTTP_302_REDIRECT"
		relative_forward_type = "PATH"
	}
	`, acctest.TestAccResourceZonePrimary(webForwardZoneResource, zoneName), webForwardZoneResource, requestTo)
}

func testAccDataSourceWebForward(zoneName, requestTo string) string {
	return fmt.Sprintf(`
	%s

	data "ultradns_webforward" "by_guid" {
		zone_name = "${resource.ultradns_webforward.www.zone_name}"
		guid = "${resource.ultradns_webforward.www.guid}"
	}

	data "ultradns_webforward" "by_request_to" {
		zone_name = "${resource.ultradns_webforward.www.zone_name}"
		request_to = "%s"
	}
	`, testAccResourceWebForwardUpdate(zoneName, requestTo), requestTo)
}

func testAccCheckWebForwardExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return errors.ResourceNotFoundError(resourceName)
		}

		parts := strings.SplitN(rs.Primary.ID, ":", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid web forward resource id %q; expected zone_name:guid", rs.Primary.ID)
		}

		services := acctest.TestAccProvider.Meta().(*service.Service)
		if _, _, err := services.WebForwardService.Read(parts[0], parts[1]); err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckWebForwardDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ultradns_webforward" {
			continue
		}

		parts := strings.SplitN(rs.Primary.ID, ":", 2)
		if len(parts) != 2 {
			continue
		}

		services := acctest.TestAccProvider.Meta().(*service.Service)
		_, _, err := services.WebForwardService.Read(parts[0], parts[1])

		if err == nil {
			return errors.ResourceNotDestroyedError(rs.Primary.ID)
		}

		errMsg := err.Error()
		if strings.Contains(errMsg, "Resource not found") || strings.Contains(errMsg, "59002") || strings.Contains(errMsg, "Zone does not exist") {
			// The forward is gone, or the zone holding it is gone.
			continue
		}

		return err
	}

	return nil
}
