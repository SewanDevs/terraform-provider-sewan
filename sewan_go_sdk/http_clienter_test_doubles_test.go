package sewan_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//------------------------------------------------------------------------------
type GetTemplatesList_Success_HttpClienterFake struct{}

func (client GetTemplatesList_Success_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	//return a resp with TEMPLATES_LIST
	return nil, nil
}

func (client GetTemplatesList_Success_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return TEMPLATES_LIST, nil
}

func (client GetTemplatesList_Success_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return TEMPLATES_LIST, nil
}

//------------------------------------------------------------------------------
// Error response *ClientTooler
type GetTemplatesList_Failure_HttpClienterFake struct{}

func (client GetTemplatesList_Failure_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	return nil, nil
}

func (client GetTemplatesList_Failure_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, errors.New("GetTemplatesList() error")
}

func (client GetTemplatesList_Failure_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, errors.New("HandleResponse() error")
}

//------------------------------------------------------------------------------
// Error response *ClientTooler
type ErrorResponse_HttpClienterFake struct{}

func (client ErrorResponse_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	return nil, errors.New(REQ_ERR)
}

func (client ErrorResponse_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client ErrorResponse_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// Response body error *ClientTooler
type BadBodyResponseContentType_HttpClienterFake struct{}

func (client BadBodyResponseContentType_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "image")
	resp.Body = ioutil.NopCloser(bytes.NewBufferString("Internal Server Error"))
	resp.StatusCode = http.StatusInternalServerError
	return &resp, nil
}

func (client BadBodyResponseContentType_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client BadBodyResponseContentType_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// Response body error *ClientTooler
type StatusInternalServerError_HttpClienterFake struct{}

func (client StatusInternalServerError_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "text/html")
	resp.Body = ioutil.NopCloser(bytes.NewBufferString("<h1>Server Error (500)</h1>"))
	resp.StatusCode = http.StatusInternalServerError
	return &resp, nil
}

func (client StatusInternalServerError_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client StatusInternalServerError_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// Response body error *ClientTooler
type BadBodyResponse_StatusCreated_HttpClienterFake struct{}

func (client BadBodyResponse_StatusCreated_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.Body = ioutil.NopCloser(bytes.NewBufferString("{\"detail\"\"Invalid json string}}.\"}"))
	resp.StatusCode = http.StatusCreated
	return &resp, nil
}

func (client BadBodyResponse_StatusCreated_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client BadBodyResponse_StatusCreated_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// Response body error *ClientTooler
type BadBodyResponse_StatusOK_HttpClienterFake struct{}

func (client BadBodyResponse_StatusOK_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.Body = ioutil.NopCloser(bytes.NewBufferString("{\"detail\"\"Invalid json string}}.\"}"))
	resp.StatusCode = http.StatusOK
	return &resp, nil
}

func (client BadBodyResponse_StatusOK_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client BadBodyResponse_StatusOK_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// 401 Reponse code error *ClientTooler
type Error401_HttpClienterFake struct{}

func (client Error401_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusUnauthorized
	resp.Status = UNAUTHORIZED_STATUS
	body := Resp_Body{"Token non valide."}
	js, _ := json.Marshal(body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

func (client Error401_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client Error401_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// 404 Reponse code error *ClientTooler
type Error404_HttpClienterFake struct{}

func (client Error404_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusNotFound
	resp.Status = NOT_FOUND_STATUS
	body := Resp_Body{"Not found."}
	js, _ := json.Marshal(body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

func (client Error404_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client Error404_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// Creation success *ClientTooler
type VDC_CreationSuccess_HttpClienterFake struct{}

func (client VDC_CreationSuccess_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusCreated
	js, _ := json.Marshal(VDC_READ_RESPONSE_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

func (client VDC_CreationSuccess_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client VDC_CreationSuccess_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// Read success *ClientTooler
type VDC_ReadSuccess_HttpClienterFake struct{}

func (client VDC_ReadSuccess_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	js, _ := json.Marshal(VDC_READ_RESPONSE_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

func (client VDC_ReadSuccess_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client VDC_ReadSuccess_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// Update success *ClientTooler
type VDC_UpdateSuccess_HttpClienterFake struct{}

func (client VDC_UpdateSuccess_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	js, _ := json.Marshal(VDC_CREATION_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

func (client VDC_UpdateSuccess_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client VDC_UpdateSuccess_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
type VDC_DeleteSuccess_HttpClienterFake struct{}

func (client VDC_DeleteSuccess_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusNoContent
	return &resp, nil
}

func (client VDC_DeleteSuccess_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client VDC_DeleteSuccess_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// Creation success *ClientTooler
type VM_CreationSuccess_HttpClienterFake struct{}

func (client VM_CreationSuccess_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusCreated
	js, _ := json.Marshal(NO_TEMPLATE_VM_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

func (client VM_CreationSuccess_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client VM_CreationSuccess_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// Read success *ClientTooler
type VM_ReadSuccess_HttpClienterFake struct{}

func (client VM_ReadSuccess_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	js, _ := json.Marshal(NO_TEMPLATE_VM_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

func (client VM_ReadSuccess_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client VM_ReadSuccess_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// Update success *ClientTooler
type VM_UpdateSuccess_HttpClienterFake struct{}

func (client VM_UpdateSuccess_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	js, _ := json.Marshal(NO_TEMPLATE_VM_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

func (client VM_UpdateSuccess_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client VM_UpdateSuccess_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
type VM_DeleteSuccess_HttpClienterFake struct{}

func (client VM_DeleteSuccess_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusNoContent
	return &resp, nil
}

func (client VM_DeleteSuccess_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client VM_DeleteSuccess_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
type DeleteWRONGResponseBody_HttpClienterFake struct{}

func (client DeleteWRONGResponseBody_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	resp.Body = ioutil.NopCloser(bytes.NewBufferString(DESTROY_WRONG_MSG))
	return &resp, nil
}

func (client DeleteWRONGResponseBody_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client DeleteWRONGResponseBody_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// req failure *ClientTooler
type CheckRedirectReqFailure_HttpClienterFake struct{}

func (client CheckRedirectReqFailure_HttpClienterFake) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	return &resp, errors.New(CHECK_REDIRECT_FAILURE)
}

func (client CheckRedirectReqFailure_HttpClienterFake) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client CheckRedirectReqFailure_HttpClienterFake) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}

//------------------------------------------------------------------------------
// Fake for api_test.go
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

func (client FakeHttpClienter) GetTemplatesList(clientTooler *ClientTooler,
	enterprise_slug string, api *API) ([]interface{}, error) {

	return nil, nil
}

func (client FakeHttpClienter) HandleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {

	return nil, nil
}
