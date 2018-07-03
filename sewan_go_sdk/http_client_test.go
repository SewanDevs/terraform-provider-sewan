package sewan_go_sdk

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

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
			GetTemplatesList_Success_HttpClienterFake{},
			"unit test enterprise",
			TEMPLATES_LIST,
			nil,
		},
		{
			2,
			GetTemplatesList_Failure_HttpClienterFake{},
			"unit test enterprise",
			nil,
			errors.New("HandleResponse() error"),
		},
		{
			3,
			ErrorResponse_HttpClienterFake{},
			"unit test enterprise",
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
				t.Errorf("\n\nTC %d : GetTemplatesList() error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Error)
			} else {
				switch {
				case !reflect.DeepEqual(test_case.TemplateList, templates_list):
					t.Errorf("\n\nTC %d : Wrong template list,"+
						"\n\rgot: \"%s\"\n\rwant: \"%s\"",
						test_case.Id, templates_list, test_case.TemplateList)
				}
			}
		case err != nil && test_case.Error != nil:
			if templates_list != nil {
				t.Errorf("\n\nTC %d : Wrong response read element,"+
					" it should be nil as error is not nil,"+
					"\n\rgot map: \n\r\"%s\"\n\rwant map: \n\r\"%s\"\n\r",
					test_case.Id, templates_list, test_case.TemplateList)
			}
			if err.Error() != test_case.Error.Error() {
				t.Errorf("\n\nTC %d : Wrong response handle error,"+
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
				t.Errorf("\n\nTC %d : HandleResponse() error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Error)
			} else {
				switch {
				case !reflect.DeepEqual(test_case.ResponseBody, responseBody):
					t.Errorf("\n\nTC %d : Wrong response read element,"+
						"\n\rgot: \"%s\"\n\rwant: \"%s\"",
						test_case.Id, responseBody, test_case.ResponseBody)
				}
			}
		case err != nil && test_case.Error != nil:
			if responseBody != nil {
				t.Errorf("\n\nTC %d : Wrong response read element,"+
					" it should be nil as error is not nil,"+
					"\n\rgot map: \n\r\"%s\"\n\rwant map: \n\r\"%s\"\n\r",
					test_case.Id, responseBody, test_case.ResponseBody)
			}
			if err.Error() != test_case.Error.Error() {
				t.Errorf("\n\nTC %d : Wrong response handle error,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, err.Error(), test_case.Error.Error())
			}
		}
	}
}
