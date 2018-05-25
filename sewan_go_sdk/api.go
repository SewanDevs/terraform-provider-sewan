package sewan_go_sdk

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// API is the interface used to communicate with the  API
type API struct {
	Token  string
	URL    API_URL
	Client *http.Client
}

type API_URL struct {
	S_url string
	Err   error
}

const API_UNAUTHORIZED_CODE = 401

// New creates a ready-to-use SDK client
func New(token string, url string) (*API, error) {
	var apiErr error
	api_url := API_URL{}

	api := &API{
		Token:  token,
		URL:    api_url,
		Client: &http.Client{},
	}

	api_url = api.validate_and_set_API_URL(url)
	api.URL = api_url
	apiErr = api.URL.Err
	if apiErr != nil {
		api = nil
	}
	return api, apiErr
}

func (api API) validateStatus() error {
	var apiErr error
	var responseBody string
	req, _ := http.NewRequest("GET", api.Get_vm_creation_url(), nil)
	req.Header.Add("authorization", "Token "+api.Token)
	resp, apiErr := api.Client.Do(req)

	if resp != nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		responseBody = string(bodyBytes)
		if resp.StatusCode == API_UNAUTHORIZED_CODE {
			apiErr = errors.New(resp.Status + responseBody)
		} else if resp.Header.Get("content-type") != "application/json" {
			apiErr = errors.New("Could not get a proper json response from \"" + api.URL.S_url + "\", the api is down or this url is wrong.")
		}
	} else {
		apiErr = errors.New("Could not get a response from \"" + api.URL.S_url + "\", the api is down or this url is wrong.")
	}

	return apiErr
}

type URL_builder interface {
	Get_vm_creation_url(api_url string) string
	Get_vm_url(api_url string, id string) string
	Validate_and_set_API_URL(url string) API_URL
	ValidateStatus() error
}

func (api API) validate_and_set_API_URL(url string) API_URL {
	api_url := API_URL{url, nil}
	api.URL = api_url
	api_url_err := api.validateStatus()
	if api_url_err != nil {
		api_url.S_url = ""
		api_url.Err = api_url_err
	}
	return api_url
}

func (api API) Get_vm_creation_url() string {
	var vm_url strings.Builder
	vm_url.WriteString(api.URL.S_url)
	vm_url.WriteString("vm/")
	return vm_url.String()
}

func (api API) Get_vm_url(vm_id string) string {
	var vm_url strings.Builder
	s_create_url := api.Get_vm_creation_url()
	vm_url.WriteString(s_create_url)
	vm_url.WriteString(vm_id)
	vm_url.WriteString("/")
	return vm_url.String()
}
