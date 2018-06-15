package sewan_go_sdk

import (
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
	"strconv"
	"errors"
	"io/ioutil"
	"encoding/json"
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
	GetTemplatesList(clientTooler *ClientTooler,
		enterprise_slug string) ([]map[string]interface{},error)
	HandleResponse(resp *http.Response,
		expectedCode int,
		expectedBodyFormat string) (interface{}, error)
}
type HttpClienter struct{}

func (client HttpClienter) Do(api *API, req *http.Request) (*http.Response, error) {
	resp, err := api.Client.Do(req)
	return resp, err
}

func (client HttpClienter) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string) ([]map[string]interface{},error) {

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

	var (
		resp_err error
		responseBody interface{}
		contentType string
		bodyBytes                  []byte
		resp_body_reader           interface{}
		read_body_err error
		read_json_err error
	)
	responseBody = nil
	resp_err = nil

	if resp.StatusCode == expectedCode{
		contentType = resp.Header.Get("Content-Type")

		if contentType == expectedBodyFormat{
			switch contentType {
			case "application/json":
				bodyBytes, read_body_err = ioutil.ReadAll(resp.Body)
				read_json_err = json.Unmarshal(bodyBytes, &resp_body_reader)
				switch {
				case read_body_err != nil:
					resp_err = errors.New("Read of response body error " +
						read_body_err.Error())
				case read_json_err != nil:
					resp_err = errors.New("Response body is not a properly formated json :"+
						read_json_err.Error())
				default:
					responseBody = resp_body_reader.(map[string]interface{})
				}
			case "text/html":
				bodyBytes, read_body_err = ioutil.ReadAll(resp.Body)
				responseBody = string(bodyBytes)
			case "":
				responseBody = nil
			default:
				resp_err = errors.New("Unhandled api response type : " +
					resp.Header.Get("Content-Type") +
					"\nPlease validate the configuration api url.")
			}
		} else {
			resp_err = errors.New("Wrong content type, \n\r expected :"+
			expectedBodyFormat+"\n\r got :"+contentType)
		}
	} else {
		resp_err = errors.New("Wrong response status code, \n\r expected :"+
		strconv.Itoa(expectedCode)+"\n\r got :"+strconv.Itoa(resp.StatusCode))
	}
	return responseBody, resp_err
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
