package sewansdk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// ClientTooler contains implementation of Clienter interface
type ClientTooler struct {
	Client Clienter
}

// Clienter interface is responsible of HTTP operations
type Clienter interface {
	do(api *API, req *http.Request) (*http.Response, error)
	getEnvResourceList(clientTooler *ClientTooler,
		api *API, resourceType string) ([]interface{}, error)
	handleResponse(resp *http.Response,
		expectedCode int,
		expectedBodyFormat string) (interface{}, error)
}

// HTTPClienter implements Clienter interface
type HTTPClienter struct{}

// Do is a wrapper of package net/http/Do method for dependency injection
func (client HTTPClienter) do(api *API, req *http.Request) (*http.Response, error) {
	return api.Client.Do(req)
}

// getPhysicalResourcesMeta returns enterprise physical resource name prefix and suffix,
// as they can be not consistent from an enterprise environment to another.
// accepted resourceType :
// * template (returns list of template available within api.Meta)
// * resource (returns physical resources list available within api.Meta :
//		resources for critical datacenter and resources for non-critical datacenters)
// * datacenter (returns list of available datacenter within api.Meta)
// * ova
// * vlan
// * disk image
// * snapshot
func (client HTTPClienter) getEnvResourceList(clientTooler *ClientTooler,
	api *API, resourceType string) ([]interface{}, error) {
	err1 := stringSliceContains(resourceType, resourceSlice)
	if err1 != nil {
		return nil, errNotInList(resourceType, strings.Join(resourceSlice, ", "))
	}
	resourceList, err2 := getJSONList(clientTooler, api, resourceType)
	if err2 == errEmptyJSON {
		return nil, errEmptyResourcesList(resourceType)
	}
	return resourceList, err2
}

// getJSONList returns a list of "listType" extracted from an Json got from
// environment api instance
func getJSONList(clientTooler *ClientTooler,
	api *API,
	listType string) ([]interface{}, error) {
	var listURL strings.Builder
	listURL.WriteString(api.URL)
	listURL.WriteString(listType)
	listURL.WriteString(entrepriseSlugHTTPReqParam)
	listURL.WriteString(api.Enterprise)
	if listType == clouddcEnvironmentTemplate {
		listURL.WriteString(clouddcGenericTemplateEnterprise)
	}
	req, err1 := http.NewRequest("GET",
		listURL.String(),
		nil)
	if err1 != nil {
		return nil, err1
	}
	req.Header.Add(httpAuthorization, httpTokenHeader+api.Token)
	resp, err2 := clientTooler.Client.do(api, req)
	if err2 != nil {
		return nil, err2
	}
	jsonList, err3 := clientTooler.Client.handleResponse(resp,
		http.StatusOK,
		httpJSONContentType)
	if err3 != nil {
		return nil, err3
	}
	if jsonList == nil {
		return nil, errEmptyJSON
	}
	return jsonList.([]interface{}), nil
}

// handleResponse formats a response body to the "expectedBodyFormat", tests the
// response status code against the "expectedCode".
// Handled response body format : "application/json", "text/html", ""
func (client HTTPClienter) handleResponse(resp *http.Response,
	expectedCode int,
	expectedBodyFormat string) (interface{}, error) {
	if resp == nil {
		return "", errEmptyResp
	}
	if resp.Body == nil {
		return "", errEmptyRespBody
	}
	defer resp.Body.Close()
	contentType := resp.Header.Get(httpRespContentType)
	if contentType != expectedBodyFormat {
		return nil, errRespStatusCodeBuilder(resp, expectedCode,
			"Wrong response content type,\n"+
				"expected :"+expectedBodyFormat+"\ngot :"+contentType)
	}
	switch contentType {
	case httpJSONContentType:
		return handleJSONContentType(resp, expectedCode)
	case httpHTMLTextContentType:
		return handleHTMLContentType(resp, expectedCode)
	case "":
		return nil, errRespStatusCodeBuilder(resp, expectedCode, "")
	default:
		return nil, errRespStatusCodeBuilder(resp, expectedCode,
			errAPIUnhandledRespType+
				resp.Header.Get(httpRespContentType)+
				errValidateAPIURL)
	}
}

func handleJSONContentType(resp *http.Response,
	expectedCode int) (interface{}, error) {
	var (
		respBodyReader interface{}
	)
	bodyBytes, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return nil, errRespStatusCodeBuilder(resp, expectedCode,
			"\nRead of response body error "+err1.Error())
	}
	if string(bodyBytes) == "" {
		return nil, errRespStatusCodeBuilder(resp, expectedCode, "")
	}
	err2 := json.Unmarshal(bodyBytes, &respBodyReader)
	if err2 != nil {
		return nil, errRespStatusCodeBuilder(resp, expectedCode,
			errJSONFormat+err2.Error()+
				"\nJson :"+string(bodyBytes))
	}
	err3 := errRespStatusCodeBuilder(resp, expectedCode, "")
	if err3 != nil {
		return nil, errors.New(err3.Error() +
			"\nResponse body error :" + string(bodyBytes))
	}
	return respBodyReader.(interface{}), nil
}

func handleHTMLContentType(resp *http.Response,
	expectedCode int) (interface{}, error) {
	bodyBytes, err4 := ioutil.ReadAll(resp.Body)
	if err4 != nil {
		return nil, errRespStatusCodeBuilder(resp, expectedCode, err4.Error())
	}
	err5 := errRespStatusCodeBuilder(resp, expectedCode, "")
	if err5 != nil {
		return nil, errors.New(err5.Error() +
			"\nResponse body error :" + string(bodyBytes))
	}
	return string(bodyBytes), nil
}
