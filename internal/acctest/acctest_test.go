package acctest_test

import (
	"testing"

	"github.com/ultradns/terraform-provider-ultradns/internal/provider"
)

func TestProvider(t *testing.T) {
	if err := provider.Provider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}
