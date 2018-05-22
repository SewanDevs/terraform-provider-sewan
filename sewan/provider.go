package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: providerSchema(),
		ResourcesMap: map[string]*schema.Resource{
			"sewan_clouddc_vm":  resource_vm(),
			"sewan_clouddc_vdc": resource_vdc(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_token": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Airdrum session token",
		},
		"api_url": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Airdrum API's URL",
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Api_token: d.Get("api_token").(string),
		Api_url:   d.Get("api_url").(string),
	}
	return config.Client()
}
