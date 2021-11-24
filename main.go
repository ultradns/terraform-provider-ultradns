package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/ultradns/terraform-provider-ultradns/ultradns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ultradns.Provider,
	})
}
