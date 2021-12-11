package acctest_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ultradns/terraform-provider-ultradns/internal/acctest"
	"github.com/ultradns/terraform-provider-ultradns/internal/provider"
)

func TestMain(m *testing.M) {
	acctest.TestAccProvider.Configure(context.TODO(), terraform.NewResourceConfigRaw(make(map[string]interface{})))
	resource.TestMain(m)
}

func TestProvider(t *testing.T) {
	if err := provider.Provider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}
