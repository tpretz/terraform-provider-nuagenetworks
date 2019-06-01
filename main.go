package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/tpretz/terraform-provider-nuagenetworks/nuagenetworks"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nuagenetworks.Provider})
}
