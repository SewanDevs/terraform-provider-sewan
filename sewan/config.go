package sewan

import (
	"net/http"
	sdk "terraform-provider-sewan/sewan_go_sdk"
)

const (
	VM_RESOURCE_TYPE  = "vm"
	VDC_RESOURCE_TYPE = "vdc"
)

type Config struct {
	Api_token string
	Api_url   string
}

type API struct {
	Token  string
	URL    string
	Client *http.Client
}

type Client struct {
	sewan                 *sdk.API
	sewan_apiTooler       *sdk.APITooler
	sewan_clientTooler    *sdk.ClientTooler
	sewan_templatesTooler *sdk.TemplatesTooler
	sewan_schemaTooler    *sdk.SchemaTooler
}

func (c *Config) Client() (*Client, error) {
	apiTooler := sdk.APITooler{
		Api: sdk.AirDrumResources_Apier{},
	}
	clientTooler := sdk.ClientTooler{
		Client: sdk.HttpClienter{},
	}
	templatesTooler := sdk.TemplatesTooler{
		TemplatesTools: sdk.Template_Templater{},
	}
	schemaTooler := sdk.SchemaTooler{
		SchemaTools: sdk.Schema_Schemaer{},
	}
	api := apiTooler.New(
		c.Api_token,
		c.Api_url,
	)
	err := apiTooler.CheckStatus(api)

	return &Client{api,
			&apiTooler,
			&clientTooler,
			&templatesTooler,
			&schemaTooler},
		err
}
