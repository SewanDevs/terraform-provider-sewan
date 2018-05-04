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
		"api_token": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Airdrum session token",
		},
		"endpoint": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Airdrum API's URL",
		},
		"timeout": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			Description: "session's timeout",
		},
		"max_retries": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			Description: "Max times a connection will be retry.",
		},
	}
}
