package sewan

import (
	sdk "github.com/SewanDevs/sewan-sdk-go"
)

func resourceCRUDTestInit() *clientStruct {
	config := configStruct{
		APIToken:   unitTestToken,
		APIURL:     unitTestAPIURL,
		Enterprise: unitTestEnterprise,
	}
	apiTooler := sdk.APITooler{
		Implementer: sdk.AirDrumResourcesAPI{},
		Initialyser: sdk.Initialyser{},
	}
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
	api := apiTooler.Initialyser.New(
		config.APIToken,
		config.APIURL,
		config.Enterprise,
	)
	api.Meta = sdk.APIMeta{
		EnterpriseResourceList: enterpriseResourceMetaDataList,
		DatacenterList:         dataCenterMetaDataList,
		TemplatesList:          []interface{}{},
		VlansList:              []interface{}{},
		SnapshotsList:          []interface{}{},
		DiskImageList:          []interface{}{},
		OvaList:                []interface{}{},
	}
	return &clientStruct{api,
		clientToolerStruct{
			&apiTooler,
			&clientTooler,
			&templatesTooler,
			&resourceTooler,
			&schemaTooler},
	}
}

func resourceTestInit() *clientStruct {
	config := configStruct{
		APIToken:   unitTestToken,
		APIURL:     unitTestAPIURL,
		Enterprise: unitTestEnterprise,
	}
	apiTooler := sdk.APITooler{
		Implementer: nil,
		Initialyser: initSuccess{},
	}
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
	api := apiTooler.Initialyser.New(
		config.APIToken,
		config.APIURL,
		config.Enterprise,
	)
	api.Meta = sdk.APIMeta{
		EnterpriseResourceList: enterpriseResourceMetaDataList,
		DatacenterList:         dataCenterMetaDataList,
		TemplatesList:          []interface{}{},
		VlansList:              []interface{}{},
		SnapshotsList:          []interface{}{},
		DiskImageList:          []interface{}{},
		OvaList:                []interface{}{},
	}
	return &clientStruct{api,
		clientToolerStruct{
			&apiTooler,
			&clientTooler,
			&templatesTooler,
			&resourceTooler,
			&schemaTooler},
	}
}
