package sewan_go_sdk

import (
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
)

const (
	DEFAULT_RESOURCE_TYPE = "vm"
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
	Get_resource_creation_url(api *API,
		resourceType string) string
	Get_resource_url(api *API,
		resourceType string,
		id string) string
	Validate_status(api *API,
		resourceType string,
		client ClientTooler) error
	Create_resource(d *schema.ResourceData,
		clientTooler *ClientTooler,
		resourceType string,
		sewan *API) (error, map[string]interface{})
	Read_resource(d *schema.ResourceData,
		clientTooler *ClientTooler,
		resourceType string,
		sewan *API) (error, map[string]interface{}, bool)
	Update_resource(d *schema.ResourceData,
		clientTooler *ClientTooler,
		resourceType string,
		sewan *API) error
	Delete_resource(d *schema.ResourceData,
		clientTooler *ClientTooler,
		resourceType string,
		sewan *API) error
}
type AirDrumResources_Apier struct{}

type ClientTooler struct {
	Client Clienter
}
type Clienter interface {
	Do(api *API, req *http.Request) (*http.Response, error)
}
type HttpClienter struct{}

func (client HttpClienter) Do(api *API, req *http.Request) (*http.Response, error) {
	resp, err := api.Client.Do(req)
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
	apiClientErr = api_tools.Api.Validate_status(api, DEFAULT_RESOURCE_TYPE, clientTooler)
	return apiClientErr
}
