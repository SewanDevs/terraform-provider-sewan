package sewan

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	sdk "gitlab.com/rd/sewan_go_sdk"
)

type VMSuccessfullCrudOperationsAirDrumAPIer struct{}

func (apier VMSuccessfullCrudOperationsAirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return noTemplateVmMap, nil
}
func (apier VMSuccessfullCrudOperationsAirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return noTemplateVmMap, nil
}
func (apier VMSuccessfullCrudOperationsAirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}
func (apier VMSuccessfullCrudOperationsAirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}

type VMFailureCrudOperationsAirDrumAPIer struct{}

func (apier VMFailureCrudOperationsAirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, errors.New(vdcCreationFailure)
}
func (apier VMFailureCrudOperationsAirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
func (apier VMFailureCrudOperationsAirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return errors.New(vdcUpdateFailure)
}
func (apier VMFailureCrudOperationsAirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return errors.New(vmDeletionFailure)
}

type VMReadFailureCrudOperationsAirDrumAPIer struct{}

func (apier VMReadFailureCrudOperationsAirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
func (apier VMReadFailureCrudOperationsAirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, errors.New(vdcReadFailure)
}
func (apier VMReadFailureCrudOperationsAirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}
func (apier VMReadFailureCrudOperationsAirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}

type VMNotFoundErrorOnReadOperationsAirDrumAPIer struct{}

func (apier VMNotFoundErrorOnReadOperationsAirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
func (apier VMNotFoundErrorOnReadOperationsAirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, sdk.ErrResourceNotExist
}
func (apier VMNotFoundErrorOnReadOperationsAirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}
func (apier VMNotFoundErrorOnReadOperationsAirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}
