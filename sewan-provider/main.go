package main

import (
        "github.com/hashicorp/terraform/plugin"
        "github.com/hashicorp/terraform/terraform"
        //"../../providers/sewan-provider"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return Provider()
                },
        })
}
