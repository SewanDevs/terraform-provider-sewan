package sewan_go_sdk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type ClientTooler struct {
	Client Clienter
}
type Clienter interface {
	Do(api *API, req *http.Request) (*http.Response, error)
	GetTemplatesList(clientTooler *ClientTooler,
		enterprise_slug string, api *API) ([]interface{}, error)
	HandleResponse(resp *http.Response,
		expectedCode int,
		expectedBodyFormat string) (interface{}, error)
	//UpdateDynamic_Field()
}
type HttpClienter struct{}

func (client HttpClienter) Do(api *API, req *http.Request) (*http.Response, error) {
	resp, err := api.Client.Do(req)
	return resp, err
}

func (client HttpClienter) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	var (
		req_err              error
		resp_err             error
		handler_resp_err     error
		return_err           error         = nil
		template_list        interface{}   = nil
		return_template_list []interface{} = nil
		templates_list_url   strings.Builder
	)
	resp := &http.Response{}
	req := &http.Request{}
	templates_list_url.WriteString(api.URL)
	templates_list_url.WriteString("template/?enterprise__slug=")
	templates_list_url.WriteString(enterprise_slug)
	logger := loggerCreate("getTemplatesList.log")

	logger.Println("templates_list_url.String() = ", templates_list_url.String())
	req, req_err = http.NewRequest("GET",
		templates_list_url.String(),
		nil)
	req.Header.Add("authorization", "Token "+api.Token)
	logger.Println("req = ", req)
	logger.Println("req_err = ", req_err)
	if req_err == nil {
		resp, resp_err = clientTooler.Client.Do(api, req)
		logger.Println("resp = ", resp)
		if resp_err == nil {
			template_list, handler_resp_err = clientTooler.Client.HandleResponse(resp,
				http.StatusOK,
				"application/json")
			if template_list != nil {
				return_template_list = template_list.([]interface{})
			}
			if handler_resp_err != nil {
				return_err = handler_resp_err
			}
		} else {
			return_err = resp_err
		}
	} else {
		return_err = req_err
	}
	logger.Println("return_template_list = ", return_template_list)
	return return_template_list, return_err
}

func (client HttpClienter) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	var (
		resp_err         error       = nil
		responseBody     interface{} = nil
		contentType      string
		bodyBytes        []byte
		resp_body_reader interface{}
		read_body_err    error = nil
		read_json_err    error = nil
	)

	if resp.StatusCode == expectedCode {
		contentType = resp.Header.Get("Content-Type")

		if contentType == expectedBodyFormat {
			switch contentType {
			case "application/json":
				bodyBytes, read_body_err = ioutil.ReadAll(resp.Body)
				read_json_err = json.Unmarshal(bodyBytes, &resp_body_reader)
				switch {
				case read_body_err != nil:
					resp_err = errors.New("Read of response body error " +
						read_body_err.Error())
				case read_json_err != nil:
					resp_err = errors.New("Response body is not a properly formated json :" +
						read_json_err.Error())
				default:
					responseBody = resp_body_reader.(interface{})
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
			resp_err = errors.New("Wrong content type, \n\r expected :" +
				expectedBodyFormat + "\n\r got :" + contentType)
		}
	} else {
		resp_err = errors.New("Wrong response status code, \n\r expected :" +
			strconv.Itoa(expectedCode) + "\n\r got :" + strconv.Itoa(resp.StatusCode) +
			"\n\rFull response status : " + resp.Status)
	}
	return responseBody, resp_err
}
