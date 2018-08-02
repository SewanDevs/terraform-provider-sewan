package sewan

import (
	"errors"
	sdk "gitlab.com/sewan_go_sdk"
	"testing"
)

func TestResourceVdcCreate(t *testing.T) {
	testCases := []struct {
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

	for _, testCase := range testCases {
		apiTooler.Api = testCase.TC_apier
		err = resourceVdcCreate(d, m_struct)
		switch {
		case err == nil || testCase.Creation_Err == nil:
			if !(err == nil && testCase.Creation_Err == nil) {
				t.Errorf("\n\nTC %d : VDC creation error was incorrect,"+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Creation_Err)
			}
		case err.Error() != testCase.Creation_Err.Error():
			t.Errorf("\n\nTC %d : VDC creation error was incorrect,"+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Creation_Err.Error())
		}
	}
}

func TestResourceVdcRead(t *testing.T) {
	testCases := []struct {
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

	for _, testCase := range testCases {
		apiTooler.Api = testCase.TC_apier
		err = resourceVdcRead(d, m_struct)
		switch {
		case err == nil || testCase.Read_Err == nil:
			if !(err == nil && testCase.Read_Err == nil) {
				t.Errorf("\n\nTC %d : VDC update error was incorrect,"+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Read_Err)
			}
		case err.Error() != testCase.Read_Err.Error():
			t.Errorf("\n\nTC %d : VDC update error was incorrect,"+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Read_Err.Error())
		}
	}
}

func TestResourceVdcUpdate(t *testing.T) {
	testCases := []struct {
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

	for _, testCase := range testCases {
		apiTooler.Api = testCase.TC_apier
		err = resourceVdcUpdate(d, m_struct)
		switch {
		case err == nil || testCase.Update_Err == nil:
			if !(err == nil && testCase.Update_Err == nil) {
				t.Errorf("\n\nTC %d : VDC update error was incorrect,"+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Update_Err)
			}
		case err.Error() != testCase.Update_Err.Error():
			t.Errorf("\n\nTC %d : VDC update error was incorrect,"+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Update_Err.Error())
		}
	}
}

func TestResourceVdcDelete(t *testing.T) {
	testCases := []struct {
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

	for _, testCase := range testCases {
		apiTooler.Api = testCase.TC_apier
		err = resourceVdcDelete(d, m_struct)
		switch {
		case err == nil || testCase.Delete_Err == nil:
			if !(err == nil && testCase.Delete_Err == nil) {
				t.Errorf("\n\nTC %d : VDC deletion error was incorrect,"+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Delete_Err)
			}
		case err.Error() != testCase.Delete_Err.Error():
			t.Errorf("\n\nTC %d : VDC deletion error was incorrect,"+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Delete_Err.Error())
		}
	}
}
