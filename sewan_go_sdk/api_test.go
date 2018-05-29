package sewan_go_sdk

import (
	"errors"
	"net/http"
	"testing"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

//------------------------------------------------------------------------------
const (
	RIGHT_API_URL = "https://next.cloud-datacenter.fr/api/clouddc/"
	RIGHT_VM_CREATION_API_URL = "https://next.cloud-datacenter.fr/api/clouddc/vm/"
	RIGHT_VM_URL_PATATE = "https://next.cloud-datacenter.fr/api/clouddc/vm/PATATE/"
	RIGHT_VM_URL_42 = "https://next.cloud-datacenter.fr/api/clouddc/vm/42/"
	WRONG_API_URL = "a wrong url"
	WRONG_API_URL_ERROR = "Wrong api url msg"
	NO_RESP_API_URL = "https://NO_RESP_API_URL.fr"
	NO_RESP_BODY_API_URL = "https://NO_BODY_API_URL.org"
	NOT_JSON_RESP_API_URL = "https://next.cloud-datacenter.fr"
	RIGHT_API_TOKEN = "42424242424242424242424242424242"
	WRONG_API_TOKEN = "a wrong token"
	WRONG_TOKEN_ERROR = "Wrong api token msg"
)

type FakeAirDrumAPIer struct {}

func (apier FakeAirDrumAPIer) Validate_status(api API, client ClientTooler) error {
	var err error
	switch{
	case api.URL!=RIGHT_API_URL:
		err=errors.New(WRONG_API_URL_ERROR)
	case api.Token!=RIGHT_API_TOKEN:
		err=errors.New(WRONG_TOKEN_ERROR)
	default:
		err=nil
	}
	return err
}

func (apier FakeAirDrumAPIer) Get_vm_creation_url(api API) string {
	return ""
}

func (apier FakeAirDrumAPIer) Get_vm_url(api API,id string) string {
	return ""
}

type FakeHttpClienter struct{}

func (client FakeHttpClienter) Do(api API,req *http.Request) (*http.Response,error){
	var err error
	err = nil
	type body struct{detail string `json:"detail"`}
	resp := http.Response{}

	if api.URL!=NO_RESP_API_URL{
		resp.Status = "200 OK"
		resp.StatusCode = http.StatusOK
		switch{
		case api.URL==WRONG_API_URL || api.URL==NOT_JSON_RESP_API_URL:
			resp.Header = map[string][]string{"Content-Type":{"text/plain; charset=utf-8"}}
			resp.Body = ioutil.NopCloser(bytes.NewBufferString("A plain text."))
		case api.URL==RIGHT_API_URL:
			if api.Token!=RIGHT_API_TOKEN {
				resp.Status = "401 Unauthorized"
				resp.StatusCode = http.StatusUnauthorized
				resp.Body = ioutil.NopCloser(bytes.NewBufferString("{\"detail\":\"Invalid token.\"}"))
			} else {
				resp.Header = map[string][]string{"Content-Type":{"application/json"}}
				body_json,_ := json.Marshal(body{detail:""})
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(body_json))
			}
		}
	} else {
		err = errors.New("No response error.")
	}
	return &resp, err
}

//------------------------------------------------------------------------------
func TestNew(t *testing.T) {

	test_cases := []struct {
		Id int
		Input_token string
		Input_url string
		Output_api API
	}{
		{	1,
			WRONG_API_TOKEN,
			RIGHT_API_URL,
			API{WRONG_API_TOKEN,RIGHT_API_URL,nil},
		},
		{	2,
			RIGHT_API_TOKEN,
			WRONG_API_URL,
			API{RIGHT_API_TOKEN,WRONG_API_URL,nil},
		},
		{	3,
			WRONG_API_TOKEN,
			WRONG_API_URL,
			API{WRONG_API_TOKEN,WRONG_API_URL,nil},
		},
		{	4,
			RIGHT_API_TOKEN,
			RIGHT_API_URL,
			API{RIGHT_API_TOKEN,RIGHT_API_URL,nil},
		},
	}

	fake_api_tools := APITooler{
		Api: FakeAirDrumAPIer{},
	}

	for _, test_case := range test_cases {
		api := fake_api_tools.New(
			test_case.Input_token,
			test_case.Input_url,
		)

		if api.Token!=test_case.Output_api.Token {
			t.Errorf("TC %d : API token error was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"",test_case.Id, api.Token, test_case.Output_api.Token)
		}
		if api.URL!=test_case.Output_api.URL {
			t.Errorf("TC %d : API token error was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"",test_case.Id, api.URL, test_case.Output_api.URL)
		}
	}
}

//------------------------------------------------------------------------------
func TestCheckStatus(t *testing.T) {

	test_cases := []struct {
		Id int
		Input_api *API
		Err error
	}{
		{	1,
			&API{
				WRONG_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			errors.New(WRONG_TOKEN_ERROR),
		},
		{	2,
			&API{
				RIGHT_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New(WRONG_API_URL_ERROR),
		},
		{	3,
			&API{
				WRONG_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New(WRONG_API_URL_ERROR),
		},
		{	4,
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
		fake_api_tools.Api= FakeAirDrumAPIer{}
		err := fake_api_tools.CheckStatus(test_case.Input_api)
		switch{
		case err==nil || test_case.Err==nil:
			if !(err==nil && test_case.Err==nil){
				t.Errorf("TC %d : Check API error was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"",test_case.Id, err, test_case.Err)
			}
		case err.Error()!=test_case.Err.Error():
			t.Errorf("TC %d : Check API error was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"",test_case.Id, err.Error(), test_case.Err.Error())
		}
	}
}

//------------------------------------------------------------------------------
func TestGet_vm_creation_url(t *testing.T) {
	test_cases := []struct {
		api API
		vm_creation_url string
	}{
		{	API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			RIGHT_VM_CREATION_API_URL,
		},
	}

	api_tools := APITooler{
		Api: AirDrumAPIer{},
	}

	for _, test_case := range test_cases {
		s_vm_creation_url := api_tools.Api.Get_vm_creation_url(test_case.api)
		if s_vm_creation_url != test_case.vm_creation_url {
			t.Errorf("VM api creation url was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"", s_vm_creation_url, test_case.vm_creation_url)
		}
	}
}

//------------------------------------------------------------------------------
func TestGet_vm_url(t *testing.T) {
	test_cases := []struct {
		api API
		vm_id	string
		vm_url string
	}{
		{	API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			"42",
			RIGHT_VM_URL_42,
		},
		{	API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			"PATATE",
			RIGHT_VM_URL_PATATE,
		},
	}

	api_tools := APITooler{
		Api: AirDrumAPIer{},
	}

	for _, test_case := range test_cases {
		s_vm_url := api_tools.Api.Get_vm_url(test_case.api,test_case.vm_id)
		if s_vm_url != test_case.vm_url {
			t.Errorf("VM url was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"", s_vm_url, test_case.vm_url)
		}
	}
}

//------------------------------------------------------------------------------
func TestValidate_status(t *testing.T) {
	test_cases := []struct {
		Id	int
		Api API
		Err error
	}{
		{	1,
			API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			nil,
		},
		{	2,
			API{
				WRONG_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			errors.New("401 Unauthorized{\"detail\":\"Invalid token.\"}"),
		},
		{	3,
			API{
				RIGHT_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a proper json response from \""+WRONG_API_URL+"\", the api is down or this url is wrong."),
		},
		{	4,
			API{
				WRONG_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a proper json response from \""+WRONG_API_URL+"\", the api is down or this url is wrong."),
		},
		{	5,
			API{
				RIGHT_API_TOKEN,
				NO_RESP_BODY_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a response body from \""+NO_RESP_BODY_API_URL+"\", the api is down or this url is wrong."),
		},
		{	6,
			API{
				RIGHT_API_TOKEN,
				NO_RESP_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a response from \""+NO_RESP_API_URL+"\", the api is down or this url is wrong."),
		},
	}

	apiTooler := APITooler{
		Api: AirDrumAPIer{},
	}
	clientTooler := ClientTooler{
		Client: FakeHttpClienter{},
	}
	var apiClientErr error

	for _, test_case := range test_cases {
		apiClientErr = apiTooler.Api.Validate_status(test_case.Api, clientTooler)
		switch{
		case apiClientErr==nil || test_case.Err==nil:
			if !(apiClientErr==nil && test_case.Err==nil){
				t.Errorf("TC %d : Validate_status() error was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"",test_case.Id, apiClientErr, test_case.Err)
			}
		case apiClientErr.Error()!=test_case.Err.Error():
			t.Errorf("TC %d : Validate_status() error was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"",test_case.Id, apiClientErr.Error(), test_case.Err.Error())
		}
	}
}
