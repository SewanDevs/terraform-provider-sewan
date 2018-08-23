package main

import (
	"github.com/SewanDevs/terraform-provider-sewan/sewan"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return sewan.Provider()
		},
	})
}
