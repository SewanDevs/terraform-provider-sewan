package sewan

import (
	"errors"
	sdk "gitlab.com/sewan_go_sdk"
	"github.com/hashicorp/terraform/helper/schema"
	"testing"
)

//------------------------------------------------------------------------------
//--Structures init, interface implementation fakes, various test items etc.----
//------------------------------------------------------------------------------
var (
	TEST_VDC_MAP = map[string]interface{}{
		NAME_FIELD:       "Unit test vdc resource",
		ENTERPRISE_FIELD: "unit test enterprise",
		DATACENTER_FIELD: "dc1",
		VDC_RESOURCE_FIELD: []interface{}{
			map[string]interface{}{
				RESOURCE_FIELD: RAM_FIELD,
				TOTAL_FIELD:    20,
			},
			map[string]interface{}{
				RESOURCE_FIELD: CPU_FIELD,
				TOTAL_FIELD:    1,
			},
			map[string]interface{}{
				RESOURCE_FIELD: "storage_enterprise",
				TOTAL_FIELD:    10,
			},
			map[string]interface{}{
				RESOURCE_FIELD: "storage_performance",
				TOTAL_FIELD:    10,
			},
			map[string]interface{}{
				RESOURCE_FIELD: "storage_high_performance",
				TOTAL_FIELD:    10,
			},
		},
	}
)

const (
	VDC_CREATION_FAILURE = "VDC creation failed."
	VDC_READ_FAILURE     = "VDC read failed."
	VDC_UPDATE_FAILURE   = "VDC update failed."
	VDC_DELETION_FAILURE = "VDC deletion failed."
)

type VDC_successfull_CRUD_operations_AirDrumAPIer struct{}

func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) Create_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return nil, TEST_VDC_MAP
}

func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) Read_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return nil, TEST_VDC_MAP, true
}

func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) Update_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}

func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) Delete_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}

func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) Get_resource_creation_url(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) Get_resource_url(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VDC_successfull_CRUD_operations_AirDrumAPIer) Validate_status(api *sdk.API,
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

func (apier VDC_failure_CRUD_operations_AirDrumAPIer) Create_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return errors.New(VDC_CREATION_FAILURE), nil
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) Read_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return nil, nil, false
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) Update_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return errors.New(VDC_UPDATE_FAILURE)
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) Delete_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return errors.New(VDC_DELETION_FAILURE)
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) Get_resource_creation_url(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) Get_resource_url(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VDC_failure_CRUD_operations_AirDrumAPIer) Validate_status(api *sdk.API,
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

func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) Create_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return nil, nil
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) Read_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return errors.New(VDC_READ_FAILURE), nil, true
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) Update_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) Delete_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) Get_resource_creation_url(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) Get_resource_url(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VDC_readfailure_CRUD_operations_AirDrumAPIer) Validate_status(api *sdk.API,
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

//------------------------------------------------------------------------------
//-------------Units tests------------------------------------------------------
//------------------------------------------------------------------------------
func TestResource_vdc_create(t *testing.T) {
	test_cases := []struct {
		Id           int
		TC_apier     sdk.APIer
		Creation_Err error
	}{
		{
			1,
			VDC_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VDC_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VDC_CREATION_FAILURE),
		},
	}
	vdc_res := resource_vdc()
	d := vdc_res.TestResourceData()
	config := Config{
		Api_token: "4242",
		Api_url:   "https://42.org",
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
	m_struct := &Client{api,
		&apiTooler,
		&clientTooler,
		&templatesTooler,
		&schemaTooler}
	var err error

	for _, test_case := range test_cases {
		apiTooler.Api = test_case.TC_apier
		err = resource_vdc_create(d, m_struct)
		switch {
		case err == nil || test_case.Creation_Err == nil:
			if !(err == nil && test_case.Creation_Err == nil) {
				t.Errorf("\n\nTC %d : VDC creation error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Creation_Err)
			}
		case err.Error() != test_case.Creation_Err.Error():
			t.Errorf("\n\nTC %d : VDC creation error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Creation_Err.Error())
		}
	}
}

func TestResource_vdc_read(t *testing.T) {
	test_cases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Read_Err   error
		Res_exists bool
	}{
		{
			1,
			VDC_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
			true,
		},
		{
			2,
			VDC_failure_CRUD_operations_AirDrumAPIer{},
			nil,
			false,
		},
		{
			3,
			VDC_readfailure_CRUD_operations_AirDrumAPIer{},
			errors.New(VDC_READ_FAILURE),
			false,
		},
	}
	vdc_res := resource_vdc()
	d := vdc_res.TestResourceData()
	config := Config{
		Api_token: "4242",
		Api_url:   "https://42.org",
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
	m_struct := &Client{api,
		&apiTooler,
		&clientTooler,
		&templatesTooler,
		&schemaTooler}
	var err error

	for _, test_case := range test_cases {
		apiTooler.Api = test_case.TC_apier
		err = resource_vdc_read(d, m_struct)
		switch {
		case err == nil || test_case.Read_Err == nil:
			if !(err == nil && test_case.Read_Err == nil) {
				t.Errorf("\n\nTC %d : VDC update error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Read_Err)
			}
		case err.Error() != test_case.Read_Err.Error():
			t.Errorf("\n\nTC %d : VDC update error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Read_Err.Error())
		}
	}
}

func TestResource_vdc_update(t *testing.T) {
	test_cases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Update_Err error
	}{
		{
			1,
			VDC_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VDC_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VDC_UPDATE_FAILURE),
		},
	}
	vdc_res := resource_vdc()
	d := vdc_res.TestResourceData()
	config := Config{
		Api_token: "4242",
		Api_url:   "https://42.org",
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
	m_struct := &Client{api,
		&apiTooler,
		&clientTooler,
		&templatesTooler,
		&schemaTooler}
	var err error

	for _, test_case := range test_cases {
		apiTooler.Api = test_case.TC_apier
		err = resource_vdc_update(d, m_struct)
		switch {
		case err == nil || test_case.Update_Err == nil:
			if !(err == nil && test_case.Update_Err == nil) {
				t.Errorf("\n\nTC %d : VDC update error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Update_Err)
			}
		case err.Error() != test_case.Update_Err.Error():
			t.Errorf("\n\nTC %d : VDC update error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Update_Err.Error())
		}
	}
}

func TestResource_vdc_delete(t *testing.T) {
	test_cases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Delete_Err error
	}{
		{
			1,
			VDC_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VDC_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VDC_DELETION_FAILURE),
		},
	}
	vdc_res := resource_vdc()
	d := vdc_res.TestResourceData()
	config := Config{
		Api_token: "4242",
		Api_url:   "https://42.org",
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
	m_struct := &Client{api,
		&apiTooler,
		&clientTooler,
		&templatesTooler,
		&schemaTooler}
	var err error

	for _, test_case := range test_cases {
		apiTooler.Api = test_case.TC_apier
		err = resource_vdc_delete(d, m_struct)
		switch {
		case err == nil || test_case.Delete_Err == nil:
			if !(err == nil && test_case.Delete_Err == nil) {
				t.Errorf("\n\nTC %d : VDC deletion error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Delete_Err)
			}
		case err.Error() != test_case.Delete_Err.Error():
			t.Errorf("\n\nTC %d : VDC deletion error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Delete_Err.Error())
		}
	}
}
