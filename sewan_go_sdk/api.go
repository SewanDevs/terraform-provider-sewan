package sewan_go_sdk

import (
	"errors"
	"net/http"
	"strings"
)

// API is the interface used to communicate with the  API
type API struct {
	Token  string
	URL    string
	Client *http.Client
}

// New creates a ready-to-use SDK client
func New(token string, url string) (*API, error) {
	api := &API{
		Token:  token,
		URL:    url,
		Client: &http.Client{},
	}
	return api, nil
}

type URL struct {
	S_url string
	Err   error
}

type URL_builder interface {
	Get_vm_creation_url(api_url string) string
	Get_vm_url(api_url string, id string) string
	Create_and_Validate_API_URL(url string) URL
}

type AirdrumURL_Builder struct{}

func (a AirdrumURL_Builder) Create_and_Validate_API_URL(url string) URL {
	valid_urls_slice := []string{
		"https://next.cloud-datacenter.fr/api/clouddc/",
	}
	var valid_urls_list strings.Builder

	var s_return_url string
	s_return_url = ""

	var url_valid bool
	url_valid = false

	var url_error error
	url_error = nil

	for _, valid_url := range valid_urls_slice {
		valid_urls_list.WriteString(valid_url)
		valid_urls_list.WriteString(", ")
		if url == valid_url {
			url_valid = true
		}
	}

	if url_valid == true {
		s_return_url = url
	} else {
		var error strings.Builder
		error.WriteString("Sewan's Airdrum API's URL is wrong, accepted urls : ")
		error.WriteString(valid_urls_list.String())
		url_error = errors.New(error.String())
	}

	return URL{s_return_url, url_error}
}

func (a AirdrumURL_Builder) Get_vm_creation_url(api_url URL) string {
	var vm_url strings.Builder
	s_api_url := api_url.S_url
	vm_url.WriteString(s_api_url)
	vm_url.WriteString("vm/")
	return vm_url.String()
}

func (a AirdrumURL_Builder) Get_vm_url(api_url URL, vm_id string) string {
	var vm_url strings.Builder
	s_create_url := a.Get_vm_creation_url(api_url)
	vm_url.WriteString(s_create_url)
	vm_url.WriteString(vm_id)
	vm_url.WriteString("/")
	return vm_url.String()
}
