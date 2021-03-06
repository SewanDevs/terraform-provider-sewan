package sewan

import (
	sdk "github.com/SewanDevs/sewan-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
)

//Provider returns a set up terraform provider
// doc : (https://godoc.org/github.com/hashicorp/terraform/helper/schema#Provider)
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: providerSchema(),
		ResourcesMap: map[string]*schema.Resource{
			"sewan_clouddc_vm":  resourceVM(),
			"sewan_clouddc_vdc": resourceVdc(),
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
		enterpriseField: &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Clouddc session enterprise",
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := configStruct{
		APIToken:   d.Get("api_token").(string),
		APIURL:     d.Get("api_url").(string),
		Enterprise: d.Get(enterpriseField).(string),
	}
	apiTooler := sdk.APITooler{
		Implementer: sdk.AirDrumResourcesAPI{},
		Initialyser: sdk.Initialyser{},
	}
	return config.clientStruct(&apiTooler)
}
