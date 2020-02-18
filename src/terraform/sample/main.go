package main

import (
	"log"
	"github.com/sairamaj/terraform-provider-sample/provider"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	log.Print("[TF-SAMPLE] in terraform-main")
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
