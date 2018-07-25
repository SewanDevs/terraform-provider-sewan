package sewan

import (
	"errors"
	sdk "gitlab.com/sewan_go_sdk"
	"testing"
)

func TestResourceVdcCreate(t *testing.T) {
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
	vdc_res := resourceVdc()
	d := vdc_res.TestResourceData()
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
		err = resourceVdcCreate(d, m_struct)
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

func TestResourceVdcRead(t *testing.T) {
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
	vdc_res := resourceVdc()
	d := vdc_res.TestResourceData()
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
		err = resourceVdcRead(d, m_struct)
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

func TestResourceVdcUpdate(t *testing.T) {
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
	vdc_res := resourceVdc()
	d := vdc_res.TestResourceData()
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
		err = resourceVdcUpdate(d, m_struct)
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

func TestResourceVdcDelete(t *testing.T) {
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
	vdc_res := resourceVdc()
	d := vdc_res.TestResourceData()
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
		err = resourceVdcDelete(d, m_struct)
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
