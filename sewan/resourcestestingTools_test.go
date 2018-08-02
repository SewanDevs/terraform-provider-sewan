package sewan

import (
	sdk "gitlab.com/sewan_go_sdk"
)

func ResourceCRUDTestInit() *Client {
	config := Config{
		Api_token: "4242",
		Api_url:   UNIT_TEST_API_URL,
	}
	apiTooler := sdk.APITooler{}
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
		config.Api_token,
		config.Api_url,
	)
	metaStruct := &Client{api,
		&apiTooler,
		&clientTooler,
		&templatesTooler,
		&schemaTooler}
	return metaStruct
}
