package sewan

import (
	sdk "gitlab.com/rd/sewan_go_sdk"
)

func ResourceCRUDTestInit() *Client {
	config := Config{
		APIToken: "4242",
		APIURL:   unitTestAPIURL,
	}
	apiTooler := sdk.APITooler{}
	clientTooler := sdk.ClientTooler{
		Client: sdk.HTTPClienter{},
	}
	templatesTooler := sdk.TemplatesTooler{
		TemplatesTools: sdk.TemplateTemplater{},
	}
	resourceTooler := sdk.ResourceTooler{
		Resource: sdk.ResourceResourceer{},
	}
	schemaTooler := sdk.SchemaTooler{
		SchemaTools: sdk.SchemaSchemaer{},
	}
	api := apiTooler.New(
		config.APIToken,
		config.APIURL,
	)
	metaStruct := &Client{api,
		&apiTooler,
		&clientTooler,
		&templatesTooler,
		&resourceTooler,
		&schemaTooler}
	return metaStruct
}
