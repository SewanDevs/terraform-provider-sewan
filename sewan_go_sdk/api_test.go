package sewan_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
	"testing"
)

//------------------------------------------------------------------------------
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

type FakeAirDrumResource_APIer struct{}

func (apier FakeAirDrumResource_APIer) Validate_status(api *API,
	resourceType string,
	client ClientTooler) error {

	var err error
	switch {
	case api.URL != RIGHT_API_URL:
		err = errors.New(WRONG_API_URL_ERROR)
	case api.Token != RIGHT_API_TOKEN:
		err = errors.New(WRONG_TOKEN_ERROR)
	default:
		err = nil
	}
	return err
}

func (apier FakeAirDrumResource_APIer) ValidateResourceType(resourceType string) error {
	return nil
}

func (apier FakeAirDrumResource_APIer) ResourceInstanceCreate(d *schema.ResourceData,
	resourceType string) (error, interface{}, string) {

	return nil, nil, ""
}

func (apier FakeAirDrumResource_APIer) Get_resource_creation_url(api *API,
	resourceType string) string {

	return ""
}

func (apier FakeAirDrumResource_APIer) Get_resource_url(api *API,
	resourceType string,
	id string) string {

	return ""
}

func (apier FakeAirDrumResource_APIer) Create_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceType string,
	sewan *API) (error, map[string]interface{}) {

	return nil, nil
}
func (apier FakeAirDrumResource_APIer) Read_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceType string,
	sewan *API) (error, map[string]interface{}, bool) {

	return nil, nil, true
}
func (apier FakeAirDrumResource_APIer) Update_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceType string,
	sewan *API) error {

	return nil
}
func (apier FakeAirDrumResource_APIer) Delete_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceType string,
	sewan *API) error {

	return nil
}

type FakeHttpClienter struct{}

func (client FakeHttpClienter) Do(api *API, req *http.Request) (*http.Response, error) {
	var err error
	err = nil
	type body struct {
		detail string `json:"detail"`
	}
	resp := http.Response{}

	if api.URL != NO_RESP_API_URL {
		resp.Status = "200 OK"
		resp.StatusCode = http.StatusOK
		switch {
		case api.URL == WRONG_API_URL || api.URL == NOT_JSON_RESP_API_URL:
			resp.Header = map[string][]string{"Content-Type": {"text/plain; charset=utf-8"}}
			resp.Body = ioutil.NopCloser(bytes.NewBufferString("A plain text."))
		case api.URL == RIGHT_API_URL:
			if api.Token != RIGHT_API_TOKEN {
				resp.Status = "401 Unauthorized"
				resp.StatusCode = http.StatusUnauthorized
				resp.Body = ioutil.NopCloser(bytes.NewBufferString("{\"detail\":\"Invalid token.\"}"))
			} else {
				resp.Header = map[string][]string{"Content-Type": {"application/json"}}
				body_json, _ := json.Marshal(body{detail: ""})
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(body_json))
			}
		}
	} else {
		err = errors.New("No response error.")
	}
	return &resp, err
}

//------------------------------------------------------------------------------
func TestDo(t *testing.T) {
	//Not tested, ref=TD-35489-UT-35737-1
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
