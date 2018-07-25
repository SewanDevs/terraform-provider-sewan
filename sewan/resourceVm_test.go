package sewan

import (
	"errors"
	sdk "gitlab.com/sewan_go_sdk"
	"testing"
)

func TestResourceVmCreate(t *testing.T) {
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

func TestResourceVmRead(t *testing.T) {
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

func TestResourceVmUpdate(t *testing.T) {
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

func TestResourceVmDelete(t *testing.T) {
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
