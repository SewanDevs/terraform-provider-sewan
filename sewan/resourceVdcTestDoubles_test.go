package sewan

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	sdk "gitlab.com/sewan_go_sdk"
)

type VDC_successfull_CRUD_operations_AirDrumAPIer struct{}

func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return nil, TEST_VDC_MAP
}

func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return nil, TEST_VDC_MAP, true
}

func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}

func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}

func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) GetResourceCreationUrl(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) GetResourceUrl(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) ValidateStatus(api *sdk.API,
	resourceType string,
	client sdk.ClientTooler) error {

	return nil
}
func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	api *sdk.API) (error,
	interface{}) {

	return nil, nil
}

type VDC_failure_CRUD_operations_AirDrumAPIer struct{}

func (apier VDC_failure_CRUD_operations_AirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return errors.New(VDC_CREATION_FAILURE), nil
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return nil, nil, false
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return errors.New(VDC_UPDATE_FAILURE)
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return errors.New(VDC_DELETION_FAILURE)
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) GetResourceCreationUrl(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) GetResourceUrl(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) ValidateStatus(api *sdk.API,
	resourceType string,
	client sdk.ClientTooler) error {

	return nil
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	api *sdk.API) (error,
	interface{}) {

	return nil, nil
}

type VDC_readfailure_CRUD_operations_AirDrumAPIer struct{}

func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return nil, nil
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return errors.New(VDC_READ_FAILURE), nil, true
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) GetResourceCreationUrl(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) GetResourceUrl(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) ValidateStatus(api *sdk.API,
	resourceType string,
	client sdk.ClientTooler) error {

	return nil
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	api *sdk.API) (error,
	interface{}) {

	return nil, nil
}
