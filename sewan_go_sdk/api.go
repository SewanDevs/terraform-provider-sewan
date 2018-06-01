package sewan_go_sdk

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/hashicorp/terraform/helper/schema"
)

type API struct {
	Token  string
	URL    string
	Client *http.Client
}
type APITooler struct {
	Api APIer
}
type APIer interface {
	Get_vm_creation_url(api *API) string
	Get_vm_url(api *API, id string) string
	Validate_status(api *API, client ClientTooler) error
	Create_vm_resource(d *schema.ResourceData, clientTooler *ClientTooler, sewan *API) (error, map[string]interface{})
	Read_vm_resource(d *schema.ResourceData, clientTooler *ClientTooler, sewan *API) (error, map[string]interface{},bool)
	Update_vm_resource(d *schema.ResourceData, clientTooler *ClientTooler, sewan *API) (error)
	Delete_vm_resource(d *schema.ResourceData, clientTooler *ClientTooler, sewan *API) (error)
}
type AirDrumAPIer struct{}

type ClientTooler struct {
	Client Clienter
}
type Clienter interface{
	Do(api *API,req *http.Request) (*http.Response,error)
}
type HttpClienter struct{}

func (client HttpClienter) Do(api *API,req *http.Request) (*http.Response,error){
	resp,err := api.Client.Do(req)
	return resp, err
}

func (api_tools *APITooler) New(token string, url string) *API {
	return &API{
		Token:  token,
		URL:    url,
		Client: &http.Client{},
	}
}

func (api_tools *APITooler) CheckStatus(api *API) error {
	var apiClientErr error
	clientTooler := ClientTooler{Client: HttpClienter{}}
	apiClientErr = api_tools.Api.Validate_status(api, clientTooler)
	return apiClientErr
}

func (apier AirDrumAPIer) Get_vm_creation_url(api *API) string {
	var vm_url strings.Builder
	vm_url.WriteString(api.URL)
	vm_url.WriteString("vm/")
	return vm_url.String()
}

func (apier AirDrumAPIer) Get_vm_url(api *API,vm_id string) string {
	var vm_url strings.Builder
	api_tools := APITooler{
		Api: apier,
	}
	s_create_url := api_tools.Api.Get_vm_creation_url(api)
	vm_url.WriteString(s_create_url)
	vm_url.WriteString(vm_id)
	vm_url.WriteString("/")
	return vm_url.String()
}

func (apier AirDrumAPIer) Validate_status(api *API, clientTooler ClientTooler) error {
	var apiErr error
	var responseBody string
	api_tools := APITooler{
		Api: apier,
	}
	req,_ := http.NewRequest("GET", api_tools.Api.Get_vm_creation_url(api), nil)
	req.Header.Add("authorization", "Token "+api.Token)
	resp, apiErr := clientTooler.Client.Do(api,req)

	if apiErr == nil {
		if resp.Body!=nil{
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			responseBody = string(bodyBytes)
			switch  {
			case resp.StatusCode == http.StatusUnauthorized:
				apiErr = errors.New(resp.Status + responseBody)
			case resp.Header.Get("content-type") != "application/json":
				apiErr = errors.New("Could not get a proper json response from \"" + api.URL + "\", the api is down or this url is wrong.")
			}
		} else {
			apiErr = errors.New("Could not get a response body from \"" + api.URL + "\", the api is down or this url is wrong.")
		}
	} else {
		apiErr = errors.New("Could not get a response from \"" + api.URL + "\", the api is down or this url is wrong.")
	}

	return apiErr
}
