package sewan

import (
	"errors"
	sdk "gitlab.com/sewan_go_sdk"
	"testing"
)

func TestResourceVmCreate(t *testing.T) {
	testCases := []struct {
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
	vm_res := resourceVm()
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

	for _, testCase := range testCases {
		apiTooler.Api = testCase.TC_apier
		err = resourceVmCreate(d, m_struct)
		switch {
		case err == nil || testCase.Creation_Err == nil:
			if !(err == nil && testCase.Creation_Err == nil) {
				t.Errorf("\n\nTC %d : VM creation error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", testCase.Id, err, testCase.Creation_Err)
			}
		case err.Error() != testCase.Creation_Err.Error():
			t.Errorf("\n\nTC %d : VM creation error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				testCase.Id, err.Error(), testCase.Creation_Err.Error())
		}
	}
}

func TestResourceVmRead(t *testing.T) {
	testCases := []struct {
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
	vm_res := resourceVm()
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

	for _, testCase := range testCases {
		apiTooler.Api = testCase.TC_apier
		err = resourceVmRead(d, m_struct)
		switch {
		case err == nil || testCase.Read_Err == nil:
			if !(err == nil && testCase.Read_Err == nil) {
				t.Errorf("\n\nTC %d : VM update error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", testCase.Id, err, testCase.Read_Err)
			}
		case err.Error() != testCase.Read_Err.Error():
			t.Errorf("\n\nTC %d : VM update error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				testCase.Id, err.Error(), testCase.Read_Err.Error())
		}
	}
}

func TestResourceVmUpdate(t *testing.T) {
	testCases := []struct {
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
	vm_res := resourceVm()
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

	for _, testCase := range testCases {
		apiTooler.Api = testCase.TC_apier
		err = resourceVmUpdate(d, m_struct)
		switch {
		case err == nil || testCase.Update_Err == nil:
			if !(err == nil && testCase.Update_Err == nil) {
				t.Errorf("\n\nTC %d : VM update error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", testCase.Id, err, testCase.Update_Err)
			}
		case err.Error() != testCase.Update_Err.Error():
			t.Errorf("\n\nTC %d : VM update error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				testCase.Id, err.Error(), testCase.Update_Err.Error())
		}
	}
}

func TestResourceVmDelete(t *testing.T) {
	testCases := []struct {
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
	vm_res := resourceVm()
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

	for _, testCase := range testCases {
		apiTooler.Api = testCase.TC_apier
		err = resourceVmDelete(d, m_struct)
		switch {
		case err == nil || testCase.Delete_Err == nil:
			if !(err == nil && testCase.Delete_Err == nil) {
				t.Errorf("\n\nTC %d : VM deletion error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", testCase.Id, err, testCase.Delete_Err)
			}
		case err.Error() != testCase.Delete_Err.Error():
			t.Errorf("\n\nTC %d : VM deletion error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				testCase.Id, err.Error(), testCase.Delete_Err.Error())
		}
	}
}
