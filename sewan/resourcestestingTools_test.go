package sewan

import (
	sdk "github.com/SewanDevs/sewan_go_sdk"
)

func ResourceCRUDTestInit() *Client {
	config := Config{
		Api_token: "4242",
		Api_url:   unitTestApiUrl,
	}
	apiTooler := sdk.APITooler{}
	clientTooler := sdk.ClientTooler{
		Client: sdk.HttpClienter{},
	}
	templatesTooler := sdk.TemplatesTooler{
		TemplatesTools: sdk.Template_Templater{},
	}
	resourceTooler := sdk.ResourceTooler{
		Resource: sdk.ResourceResourceer{},
	}
	schemaTooler := sdk.SchemaTooler{
		SchemaTools: sdk.SchemaSchemaer{},
	}
	api := apiTooler.New(
		config.Api_token,
		config.Api_url,
	)
	metaStruct := &Client{api,
		&apiTooler,
		&clientTooler,
		&templatesTooler,
		&resourceTooler,
		&schemaTooler}
	return metaStruct
}
