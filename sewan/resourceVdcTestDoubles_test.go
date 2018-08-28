package sewan

import (
	"errors"
	sdk "github.com/SewanDevs/sewan-sdk-go"
	sdk "github.com/SewanDevs/sewan_go_sdk"
	"github.com/hashicorp/terraform/helper/schema"
)

type VdcSuccessfullCrudOperationsAirDrumAPIer struct{}

func (apier VdcSuccessfullCrudOperationsAirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTooler *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {

	return testVdcMap, nil
}
func (apier VdcSuccessfullCrudOperationsAirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTooler *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return testVdcMap, nil
}
func (apier VdcSuccessfullCrudOperationsAirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTooler *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}
func (apier VdcSuccessfullCrudOperationsAirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTooler *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}

type VdcFailureCrudOperationsAirDrumAPIer struct{}

func (apier VdcFailureCrudOperationsAirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, errors.New(vdcCreationFailure)
}
func (apier VdcFailureCrudOperationsAirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
func (apier VdcFailureCrudOperationsAirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return errors.New(vdcUpdateFailure)
}
func (apier VdcFailureCrudOperationsAirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return errors.New(vdcDeletionFailure)
}

type VdcReadFailureCrudOperationsAirDrumAPIer struct{}

func (apier VdcReadFailureCrudOperationsAirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
func (apier VdcReadFailureCrudOperationsAirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, errors.New(vdcReadFailure)
}
func (apier VdcReadFailureCrudOperationsAirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}
func (apier VdcReadFailureCrudOperationsAirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}

type VdcNotFoundErrorOnReadOperationsAirDrumAPIer struct{}

func (apier VdcNotFoundErrorOnReadOperationsAirDrumAPIer) CreateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
func (apier VdcNotFoundErrorOnReadOperationsAirDrumAPIer) ReadResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) (map[string]interface{}, error) {
	return map[string]interface{}{}, sdk.ErrResourceNotExist
}
func (apier VdcNotFoundErrorOnReadOperationsAirDrumAPIer) UpdateResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}
func (apier VdcNotFoundErrorOnReadOperationsAirDrumAPIer) DeleteResource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	resourceTools *sdk.ResourceTooler,
	resourceType string,
	sewan *sdk.API) error {
	return nil
}
