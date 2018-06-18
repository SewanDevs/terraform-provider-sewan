package sewan_go_sdk

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

const (
	RIGHT_API_URL             = "https://next.cloud-datacenter.fr/api/clouddc/"
	RIGHT_VM_CREATION_API_URL = "https://next.cloud-datacenter.fr/api/clouddc/vm/"
	RIGHT_VM_URL_PATATE       = "https://next.cloud-datacenter.fr/api/clouddc/vm/PATATE/"
	RIGHT_VM_URL_42           = "https://next.cloud-datacenter.fr/api/clouddc/vm/42/"
	WRONG_API_URL             = "a wrong url"
	WRONG_API_URL_ERROR       = "Wrong api url msg"
	NO_RESP_API_URL           = "https://NO_RESP_API_URL.fr"
	NO_RESP_BODY_API_URL      = "https://NO_BODY_API_URL.org"
	NOT_JSON_RESP_API_URL     = "https://next.cloud-datacenter.fr"
	RIGHT_API_TOKEN           = "42424242424242424242424242424242"
	WRONG_API_TOKEN           = "a wrong token"
	WRONG_TOKEN_ERROR         = "Wrong api token msg"
)

//------------------------------------------------------------------------------
//-------------Units tests------------------------------------------------------
//------------------------------------------------------------------------------
func TestDo(t *testing.T) {
	//Not tested, ref=TD-35489-UT-35737-1
}

//------------------------------------------------------------------------------
func TestGetTemplatesList(t *testing.T) {
	test_cases := []struct {
		Id              int
		TC_clienter     Clienter
		Enterprise_slug string
		TemplateList    []interface{}
		Error           error
	}{
		{
			1,
			GetTemplatesList_Sucess_HttpClienterFake{},
			"sewan-rd-cloud-beta",
			TEMPLATES_LIST,
			nil,
		},
		{
			2,
			GetTemplatesList_Failure_HttpClienterFake{},
			"sewan-rd-cloud-beta",
			nil,
			errors.New("HandleResponse() error"),
		},
		{
			3,
			ErrorResponse_HttpClienterFake{},
			"sewan-rd-cloud-beta",
			nil,
			errors.New(REQ_ERR),
		},
	}
	var (
		templates_list []interface{}
		err            error
	)
	client_tooler := ClientTooler{}
	client_tooler.Client = HttpClienter{}
	fake_client_tooler := ClientTooler{}
	apiTooler := APITooler{}
	api := apiTooler.New("token", "url")

	for _, test_case := range test_cases {
		fake_client_tooler.Client = test_case.TC_clienter
		templates_list, err = client_tooler.Client.GetTemplatesList(&fake_client_tooler,
			test_case.Enterprise_slug, api)

		switch {
		case err == nil || test_case.Error == nil:
			if !(err == nil && test_case.Error == nil) {
				t.Errorf("TC %d : GetTemplatesList() error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Error)
			} else {
				switch {
				case !reflect.DeepEqual(test_case.TemplateList, templates_list):
					t.Errorf("TC %d : Wrong template list,"+
						"\n\rgot: \"%s\"\n\rwant: \"%s\"",
						test_case.Id, templates_list, test_case.TemplateList)
				}
			}
		case err != nil && test_case.Error != nil:
			if templates_list != nil {
				t.Errorf("TC %d : Wrong response read element,"+
					" it should be nil as error is not nil,"+
					"\n\rgot map: \n\r\"%s\"\n\rwant map: \n\r\"%s\"\n\r",
					test_case.Id, templates_list, test_case.TemplateList)
			}
			if err.Error() != test_case.Error.Error() {
				t.Errorf("TC %d : Wrong response handle error,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, err.Error(), test_case.Error.Error())
			}
		}
	}
}

//------------------------------------------------------------------------------
func TestHandleResponse(t *testing.T) {
	test_cases := []struct {
		Id                 int
		Response           *http.Response
		ExpectedCode       int
		ExpectedBodyFormat string
		ResponseBody       interface{}
		Error              error
	}{
		{
			1,
			HttpResponseFake_OK_json(),
			http.StatusOK,
			"application/json",
			JsonStub(),
			nil,
		},
		{
			2,
			HttpResponseFake_OK_template_list_json(),
			http.StatusOK,
			"application/json",
			JsonTemplateListFake(),
			nil,
		},
		{
			3,
			HttpResponseFake_500_texthtml(),
			http.StatusInternalServerError,
			"text/html",
			"<h1>Server Error (500)</h1>",
			nil,
		},
		{
			4,
			HttpResponseFake_500_json(),
			http.StatusInternalServerError,
			"text/html",
			nil,
			errors.New("Wrong content type, \n\r expected :text/html\n\r got :application/json"),
		},
		{
			5,
			HttpResponseFake_OK_json(),
			http.StatusInternalServerError,
			"text/html",
			nil,
			errors.New("Wrong response status code, \n\r expected :500\n\r got :200" +
				"\n\rFull response status : 200 OK"),
		},
		{
			6,
			HttpResponseFake_OK_json(),
			http.StatusInternalServerError,
			"text/html",
			nil,
			errors.New("Wrong response status code, \n\r expected :500\n\r got :200" +
				"\n\rFull response status : 200 OK"),
		},
		{
			7,
			HttpResponseFake_OK_no_content(),
			http.StatusOK,
			"",
			nil,
			nil,
		},
		{
			8,
			HttpResponseFake_OK_wrongjson(),
			http.StatusOK,
			"application/json",
			nil,
			errors.New("Response body is not a properly formated json :" +
				"invalid character 'a' looking for beginning of value"),
		},
		{
			9,
			HttpResponseFake_OK_image(),
			http.StatusOK,
			"image",
			nil,
			errors.New("Unhandled api response type : image" +
				"\nPlease validate the configuration api url."),
		},
	}

	var (
		responseBody interface{}
		err          error
	)
	clienter := HttpClienter{}

	for _, test_case := range test_cases {
		responseBody, err = clienter.HandleResponse(test_case.Response,
			test_case.ExpectedCode, test_case.ExpectedBodyFormat)

		switch {
		case err == nil || test_case.Error == nil:
			if !(err == nil && test_case.Error == nil) {
				t.Errorf("TC %d : HandleResponse() error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Error)
			} else {
				switch {
				case !reflect.DeepEqual(test_case.ResponseBody, responseBody):
					t.Errorf("TC %d : Wrong response read element,"+
						"\n\rgot: \"%s\"\n\rwant: \"%s\"",
						test_case.Id, responseBody, test_case.ResponseBody)
				}
			}
		case err != nil && test_case.Error != nil:
			if responseBody != nil {
				t.Errorf("TC %d : Wrong response read element,"+
					" it should be nil as error is not nil,"+
					"\n\rgot map: \n\r\"%s\"\n\rwant map: \n\r\"%s\"\n\r",
					test_case.Id, responseBody, test_case.ResponseBody)
			}
			if err.Error() != test_case.Error.Error() {
				t.Errorf("TC %d : Wrong response handle error,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, err.Error(), test_case.Error.Error())
			}
		}
	}
}

//------------------------------------------------------------------------------
func TestNew(t *testing.T) {

	test_cases := []struct {
		Id          int
		Input_token string
		Input_url   string
		Output_api  API
	}{
		{1,
			WRONG_API_TOKEN,
			RIGHT_API_URL,
			API{WRONG_API_TOKEN, RIGHT_API_URL, nil},
		},
		{2,
			RIGHT_API_TOKEN,
			WRONG_API_URL,
			API{RIGHT_API_TOKEN, WRONG_API_URL, nil},
		},
		{3,
			WRONG_API_TOKEN,
			WRONG_API_URL,
			API{WRONG_API_TOKEN, WRONG_API_URL, nil},
		},
		{4,
			RIGHT_API_TOKEN,
			RIGHT_API_URL,
			API{RIGHT_API_TOKEN, RIGHT_API_URL, nil},
		},
	}

	fake_api_tools := APITooler{
		Api: FakeAirDrumResource_APIer{},
	}

	for _, test_case := range test_cases {
		api := fake_api_tools.New(
			test_case.Input_token,
			test_case.Input_url,
		)

		switch {
		case api.Token != test_case.Output_api.Token:
			t.Errorf("TC %d : API token error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, api.Token, test_case.Output_api.Token)
		case api.URL != test_case.Output_api.URL:
			t.Errorf("TC %d : API token error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, api.URL, test_case.Output_api.URL)
		}
	}
}

//------------------------------------------------------------------------------
func TestCheckStatus(t *testing.T) {

	test_cases := []struct {
		Id        int
		Input_api *API
		Err       error
	}{
		{1,
			&API{
				WRONG_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			errors.New(WRONG_TOKEN_ERROR),
		},
		{2,
			&API{
				RIGHT_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New(WRONG_API_URL_ERROR),
		},
		{3,
			&API{
				WRONG_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New(WRONG_API_URL_ERROR),
		},
		{4,
			&API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			nil,
		},
	}

	fake_api_tools := APITooler{}

	for _, test_case := range test_cases {
		fake_api_tools.Api = FakeAirDrumResource_APIer{}
		err := fake_api_tools.CheckStatus(test_case.Input_api)
		switch {
		case err == nil || test_case.Err == nil:
			if !(err == nil && test_case.Err == nil) {
				t.Errorf("TC %d : Check API error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, err, test_case.Err)
			}
		case err.Error() != test_case.Err.Error():
			t.Errorf("TC %d : Check API error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Err.Error())
		}
	}
}
