package sewan

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	sdk "gitlab.com/sewan_go_sdk"
	"testing"
)

//------------------------------------------------------------------------------
//--Structures init, interface implementation fakes, various test items etc.----
//------------------------------------------------------------------------------
var (
	NO_TEMPLATE_VM_MAP = map[string]interface{}{
		NAME_FIELD:       "Unit test vm",
		ENTERPRISE_FIELD: "unit test enterprise",
		STATE_FIELD:      "UP",
		OS_FIELD:         "Debian",
		RAM_FIELD:        8,
		CPU_FIELD:        4,
		DISKS_FIELD: []interface{}{
			map[string]interface{}{
				NAME_FIELD: "disk 1",
				SIZE_FIELD: 24,
				SLUG_FIELD: SLUG_FIELD,
			},
		},
		NICS_FIELD: []interface{}{
			map[string]interface{}{
				VLAN_NAME_FIELD:  "vlan 1 update",
				MAC_ADRESS_FIELD: "24",
				CONNECTED_FIELD:  true,
			},
			map[string]interface{}{
				VLAN_NAME_FIELD:  "vlan 2",
				MAC_ADRESS_FIELD: "24",
				CONNECTED_FIELD:  true,
			},
		},
		VDC_FIELD:           VDC_FIELD,
		BOOT_FIELD:          "on disk",
		STORAGE_CLASS_FIELD: "storage_enterprise",
		SLUG_FIELD:          "42",
		TOKEN_FIELD:         "424242",
		BACKUP_FIELD:        "backup-no_backup",
		DISK_IMAGE_FIELD:    "",
		PLATFORM_NAME_FIELD: "42",
		BACKUP_SIZE_FIELD:   42,
		COMMENT_FIELD:       "42",
		DYNAMIC_FIELD:       "42",
	}
)

const (
	VM_CREATION_FAILURE = "VM creation failed."
	VM_READ_FAILURE     = "VM read failed."
	VM_UPDATE_FAILURE   = "VM update failed."
	VM_DELETION_FAILURE = "VM deletion failed."
)

type VM_successfull_CRUD_operations_AirDrumAPIer struct{}

func (apier VM_successfull_CRUD_operations_AirDrumAPIer) Create_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return nil, NO_TEMPLATE_VM_MAP
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) Read_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return nil, NO_TEMPLATE_VM_MAP, true
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) Update_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) Delete_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) Get_resource_creation_url(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) Get_resource_url(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) Validate_status(api *sdk.API,
	resourceType string,
	client sdk.ClientTooler) error {

	return nil
}
func (apier VM_successfull_CRUD_operations_AirDrumAPIer) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	api *sdk.API) (error,
	interface{}) {

	return nil, nil
}

type VM_failure_CRUD_operations_AirDrumAPIer struct{}

func (apier VM_failure_CRUD_operations_AirDrumAPIer) Create_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return errors.New(VM_CREATION_FAILURE), nil
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) Read_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return nil, nil, false
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) Update_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return errors.New(VM_UPDATE_FAILURE)
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) Delete_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return errors.New(VM_DELETION_FAILURE)
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) Get_resource_creation_url(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) Get_resource_url(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) Validate_status(api *sdk.API,
	resourceType string,
	client sdk.ClientTooler) error {

	return nil
}
func (apier VM_failure_CRUD_operations_AirDrumAPIer) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	api *sdk.API) (error,
	interface{}) {

	return nil, nil
}

type VM_readfailure_CRUD_operations_AirDrumAPIer struct{}

func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) Create_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}) {

	return nil, nil
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) Read_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) (error, map[string]interface{}, bool) {

	return errors.New(VM_READ_FAILURE), nil, true
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) Update_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) Delete_resource(d *schema.ResourceData,
	clientTooler *sdk.ClientTooler,
	templatesTooler *sdk.TemplatesTooler,
	schemaTools *sdk.SchemaTooler,
	resourceType string,
	sewan *sdk.API) error {

	return nil
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) Get_resource_creation_url(api *sdk.API,
	resourceType string) string {
	return ""
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) Get_resource_url(api *sdk.API,
	resourceType string,
	id string) string {

	return ""
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) ValidateResourceType(resourceType string) error {
	return nil
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) Validate_status(api *sdk.API,
	resourceType string,
	client sdk.ClientTooler) error {

	return nil
}
func (apier VM_readfailure_CRUD_operations_AirDrumAPIer) ResourceInstanceCreate(d *schema.ResourceData,
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
func TestResource_vm_create(t *testing.T) {
	test_cases := []struct {
		Id           int
		TC_apier     sdk.APIer
		Creation_Err error
	}{
		{
			1,
			VM_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VM_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VM_CREATION_FAILURE),
		},
	}
	vm_res := resource_vm()
	d := vm_res.TestResourceData()
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
		err = resource_vm_create(d, m_struct)
		switch {
		case err == nil || test_case.Creation_Err == nil:
			if !(err == nil && test_case.Creation_Err == nil) {
				t.Errorf("\n\nTC %d : VM creation error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Creation_Err)
			}
		case err.Error() != test_case.Creation_Err.Error():
			t.Errorf("\n\nTC %d : VM creation error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Creation_Err.Error())
		}
	}
}

func TestResource_vm_read(t *testing.T) {
	test_cases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Read_Err   error
		Res_exists bool
	}{
		{
			1,
			VM_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
			true,
		},
		{
			2,
			VM_failure_CRUD_operations_AirDrumAPIer{},
			nil,
			false,
		},
		{
			3,
			VM_readfailure_CRUD_operations_AirDrumAPIer{},
			errors.New(VM_READ_FAILURE),
			false,
		},
	}
	vm_res := resource_vm()
	d := vm_res.TestResourceData()
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
		err = resource_vm_read(d, m_struct)
		switch {
		case err == nil || test_case.Read_Err == nil:
			if !(err == nil && test_case.Read_Err == nil) {
				t.Errorf("\n\nTC %d : VM update error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Read_Err)
			}
		case err.Error() != test_case.Read_Err.Error():
			t.Errorf("\n\nTC %d : VM update error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Read_Err.Error())
		}
	}
}

func TestResource_vm_update(t *testing.T) {
	test_cases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Update_Err error
	}{
		{
			1,
			VM_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VM_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VM_UPDATE_FAILURE),
		},
	}
	vm_res := resource_vm()
	d := vm_res.TestResourceData()
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
		err = resource_vm_update(d, m_struct)
		switch {
		case err == nil || test_case.Update_Err == nil:
			if !(err == nil && test_case.Update_Err == nil) {
				t.Errorf("\n\nTC %d : VM update error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Update_Err)
			}
		case err.Error() != test_case.Update_Err.Error():
			t.Errorf("\n\nTC %d : VM update error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Update_Err.Error())
		}
	}
}

func TestResource_vm_delete(t *testing.T) {
	test_cases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Delete_Err error
	}{
		{
			1,
			VM_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VM_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VM_DELETION_FAILURE),
		},
	}
	vm_res := resource_vm()
	d := vm_res.TestResourceData()
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
		err = resource_vm_delete(d, m_struct)
		switch {
		case err == nil || test_case.Delete_Err == nil:
			if !(err == nil && test_case.Delete_Err == nil) {
				t.Errorf("\n\nTC %d : VM deletion error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Delete_Err)
			}
		case err.Error() != test_case.Delete_Err.Error():
			t.Errorf("\n\nTC %d : VM deletion error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Delete_Err.Error())
		}
	}
}
