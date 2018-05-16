package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: providerSchema(),
		ResourcesMap:	map[string]*schema.Resource{
			"sewan_serverVM": resourceVM(),
			"sewan_clientVM": resourceVM(),
		},
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"air_drum_token": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Airdrum session token",
		},
		"air_drum_url": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Airdrum API's URL",
		},
	}
}
