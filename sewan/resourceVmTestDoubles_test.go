package sewan

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	sdk "gitlab.com/sewan_go_sdk"
)

type VM_successfull_CRUD_operations_AirDrumAPIer struct{}

func (apier VM_successfull_CRUD_operations_AirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return nil, NO_TEMPLATE_VM_MAP
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return nil, NO_TEMPLATE_VM_MAP, true
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) GetResourceCreationUrl(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) GetResourceUrl(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) ValidateStatus(api *sdk.API,
	resourceType string,
	client sdk.ClientTooler) error {

	return nil
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceType string,
	api *sdk.API) (error,
	interface{}) {

	return nil, nil
}

type VM_failure_CRUD_operations_AirDrumAPIer struct{}

func (apier VM_failure_CRUD_operations_AirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return errors.New(VM_CREATION_FAILURE), nil
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return nil, nil, false
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return errors.New(VM_UPDATE_FAILURE)
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return errors.New(VM_DELETION_FAILURE)
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) GetResourceCreationUrl(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) GetResourceUrl(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) ValidateStatus(api *sdk.API,
	resourceType string,
	client sdk.ClientTooler) error {

	return nil
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceType string,
	api *sdk.API) (error,
	interface{}) {

	return nil, nil
}

type VM_readfailure_CRUD_operations_AirDrumAPIer struct{}

func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return nil, nil
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return errors.New(VM_READ_FAILURE), nil, true
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) GetResourceCreationUrl(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) GetResourceUrl(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) ValidateStatus(api *sdk.API,
	resourceType string,
	client sdk.ClientTooler) error {

	return nil
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceType string,
	api *sdk.API) (error,
	interface{}) {

	return nil, nil
}
