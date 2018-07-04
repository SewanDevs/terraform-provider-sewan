package sewan_go_sdk

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
)

type FakeAirDrumResource_APIer struct{}

func (apier FakeAirDrumResource_APIer) Validate_status(api *API,
	resourceType string,
	client ClientTooler) error {

	var err error
	switch {
	case api.URL != RIGHT_API_URL:
		err = errors.New(WRONG_API_URL_ERROR)
	case api.Token != RIGHT_API_TOKEN:
		err = errors.New(WRONG_TOKEN_ERROR)
	default:
		err = nil
	}
	return err
}

func (apier FakeAirDrumResource_APIer) ValidateResourceType(resourceType string) error {
	return nil
}

func (apier FakeAirDrumResource_APIer) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	schemaTools *SchemaTooler,
	resourceType string,
	api *API) (error, interface{}) {

	return nil, ""
}

func (apier FakeAirDrumResource_APIer) Get_resource_creation_url(api *API,
	resourceType string) string {

	return ""
}

func (apier FakeAirDrumResource_APIer) Get_resource_url(api *API,
	resourceType string,
	id string) string {

	return ""
}

func (apier FakeAirDrumResource_APIer) Create_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	schemaTools *SchemaTooler,
	resourceType string,
	sewan *API) (error, map[string]interface{}) {

	return nil, nil
}
func (apier FakeAirDrumResource_APIer) Read_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	schemaTools *SchemaTooler,
	resourceType string,
	sewan *API) (error, map[string]interface{}, bool) {

	return nil, nil, true
}
func (apier FakeAirDrumResource_APIer) Update_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	schemaTools *SchemaTooler,
	resourceType string,
	sewan *API) error {

	return nil
}
func (apier FakeAirDrumResource_APIer) Delete_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	schemaTools *SchemaTooler,
	resourceType string,
	sewan *API) error {

	return nil
}
