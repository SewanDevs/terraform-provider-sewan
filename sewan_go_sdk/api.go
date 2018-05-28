package sewan_go_sdk

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const API_UNAUTHORIZED_CODE = 401

var (
	api_tooler API_tooler
)

type Tooler struct {
	ApiTooler API_tooler
}

type API struct {
	Token  string
	URL    string
	Client *http.Client
}

type API_tooler interface {
	Get_vm_creation_url() string
	Get_vm_url(id string) string
	Validate_status() error
}

func New(token string, url string) (*API, error) {
	var apiErr error

	api := &API{
		Token:  token,
		URL:    url,
		Client: &http.Client{},
	}

	apiErr = api.Validate_status()
	if apiErr != nil {
		api = nil
	}
	return api, apiErr
}

func (api API) Get_vm_creation_url() string {
	var vm_url strings.Builder
	vm_url.WriteString(api.URL)
	vm_url.WriteString("vm/")
	return vm_url.String()
}

func (api API) Get_vm_url(vm_id string) string {
	var vm_url strings.Builder
	api_tools := Tooler{
		ApiTooler: api,
	}
	s_create_url := api_tools.ApiTooler.Get_vm_creation_url()
	vm_url.WriteString(s_create_url)
	vm_url.WriteString(vm_id)
	vm_url.WriteString("/")
	return vm_url.String()
}

func (api API) Validate_status() error {
	var apiErr error
	var responseBody string
	api_tools := Tooler{
		ApiTooler: api,
	}
	req, _ := http.NewRequest("GET", api_tools.ApiTooler.Get_vm_creation_url(), nil)
	req.Header.Add("authorization", "Token "+api.Token)
	resp, apiErr := api.Client.Do(req)

	if resp != nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		responseBody = string(bodyBytes)
		if resp.StatusCode == API_UNAUTHORIZED_CODE {
			apiErr = errors.New(resp.Status + responseBody)
		} else if resp.Header.Get("content-type") != "application/json" {
			apiErr = errors.New("Could not get a proper json response from \"" + api.URL + "\", the api is down or this url is wrong.")
		}
	} else {
		apiErr = errors.New("Could not get a response from \"" + api.URL + "\", the api is down or this url is wrong.")
	}

	return apiErr
}
