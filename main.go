package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"terraform-provider-ultradns/ultradns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ultradns.Provider})
}
