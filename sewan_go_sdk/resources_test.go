package sewan_go_sdk

import (
	"errors"
	"net/http"
	"testing"
)

//------------------------------------------------------------------------------
func TestGet_resource_url(t *testing.T) {
	test_cases := []struct {
		Id     int
		api    API
		vm_id  string
		vm_url string
	}{
		{1,
			API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			"42",
			RIGHT_VM_URL_42,
		},
		{2,
			API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			"PATATE",
			RIGHT_VM_URL_PATATE,
		},
	}

	api_tools := APITooler{
		Api: AirDrumResources_Apier{},
	}

	for _, test_case := range test_cases {
		s_vm_url := api_tools.Api.Get_resource_url(&test_case.api,
			VM_RESOURCE_TYPE,
			test_case.vm_id)

		switch {
		case s_vm_url != test_case.vm_url:
			t.Errorf("VM url was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"",
				s_vm_url, test_case.vm_url)
		}
	}
}

//------------------------------------------------------------------------------
func TestGet_resource_creation_url(t *testing.T) {
	test_cases := []struct {
		Id                    int
		api                   API
		resource_creation_url string
	}{
		{1,
			API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			RIGHT_VM_CREATION_API_URL,
		},
	}

	api_tools := APITooler{
		Api: AirDrumResources_Apier{},
	}

	for _, test_case := range test_cases {
		s_resource_creation_url := api_tools.Api.Get_resource_creation_url(&test_case.api,
			VM_RESOURCE_TYPE)

		switch {
		case s_resource_creation_url != test_case.resource_creation_url:
			t.Errorf("resource api creation url was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				s_resource_creation_url, test_case.resource_creation_url)
		}
	}
}

//------------------------------------------------------------------------------
func TestValidate_status(t *testing.T) {
	test_cases := []struct {
		Id           int
		Api          API
		Err          error
		ResourceType string
	}{
		{1,
			API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			nil,
			VM_RESOURCE_TYPE,
		},
		{2,
			API{
				WRONG_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			errors.New("401 Unauthorized{\"detail\":\"Invalid token.\"}"),
			VM_RESOURCE_TYPE,
		},
		{3,
			API{
				RIGHT_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a proper json response from \"" +
				WRONG_API_URL + "\", the api is down or this url is wrong."),
			VM_RESOURCE_TYPE,
		},
		{4,
			API{
				WRONG_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a proper json response from \"" +
				WRONG_API_URL + "\", the api is down or this url is wrong."),
			VM_RESOURCE_TYPE,
		},
		{5,
			API{
				RIGHT_API_TOKEN,
				NO_RESP_BODY_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a response body from \"" +
				NO_RESP_BODY_API_URL + "\", the api is down or this url is wrong."),
			VM_RESOURCE_TYPE,
		},
		{6,
			API{
				RIGHT_API_TOKEN,
				NO_RESP_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a response from \"" +
				NO_RESP_API_URL + "\", the api is down or this url is wrong."),
			VM_RESOURCE_TYPE,
		},
	}

	apiTooler := APITooler{
		Api: AirDrumResources_Apier{},
	}
	clientTooler := ClientTooler{
		Client: FakeHttpClienter{},
	}
	var apiClientErr error

	for _, test_case := range test_cases {
		apiClientErr = apiTooler.Api.Validate_status(&test_case.Api,
			test_case.ResourceType,
			clientTooler)

		switch {
		case apiClientErr == nil || test_case.Err == nil:
			if !(apiClientErr == nil && test_case.Err == nil) {
				t.Errorf("TC %d : Validate_status() error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, apiClientErr, test_case.Err)
			}
		case apiClientErr.Error() != test_case.Err.Error():
			t.Errorf("TC %d : Validate_status() error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, apiClientErr.Error(), test_case.Err.Error())
		}
	}
}
