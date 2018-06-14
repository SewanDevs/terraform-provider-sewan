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
	ValidateResourceType(resourceType string) error
	Validate_status(api *API,
		resourceType string,
		client ClientTooler) error
	ResourceInstanceCreate(d *schema.ResourceData,
		clientTooler *ClientTooler,
		resourceType string,
		api *API) (error,
		interface{},
		string)
	Create_resource(d *schema.ResourceData,
		clientTooler *ClientTooler,
		resourceType string,
		sewan *API) (error,
		map[string]interface{})
	Read_resource(d *schema.ResourceData,
		clientTooler *ClientTooler,
		resourceType string,
		sewan *API) (error,
		map[string]interface{},
		bool)
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
	GetTemplatesList(enterprise_slug string) ([]map[string]interface{},
		error)
	HandleResponse(resp *http.Response,
		expectedCode int,
		expectedBodyFormat string) (interface{}, error)
}
type HttpClienter struct{}

func (client HttpClienter) Do(api *API, req *http.Request) (*http.Response, error) {
	resp, err := api.Client.Do(req)
	return resp, err
}

func (client HttpClienter) GetTemplatesList(enterprise_slug string) ([]map[string]interface{},
	error) {

	//i Create GET req

	// ii Exec req
	// resp, err := client.Do(req)

	// iii handle resp, param : resp + wanted return code + wanted resp type
	// templateList,err := api.Client.HandleResp(resp,200,[]map[string]interface{})

	return []map[string]interface{}{}, nil
}

func (client HttpClienter) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	//var (
	//
	//)

	//if resp.StatusCode == expectedCode{
	//
	//} else {
	//
	//}
	return "", nil
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
