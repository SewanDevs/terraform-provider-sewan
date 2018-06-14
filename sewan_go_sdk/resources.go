package sewan_go_sdk

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func (apier AirDrumResources_Apier) Get_resource_creation_url(api *API,
	resourceType string) string {

	var resource_url strings.Builder
	resource_url.WriteString(api.URL)
	resource_url.WriteString(resourceType)
	resource_url.WriteString("/")
	return resource_url.String()
}

func (apier AirDrumResources_Apier) Get_resource_url(api *API,
	resourceType string,
	resource_id string) string {

	var resource_url strings.Builder
	api_tools := APITooler{
		Api: apier,
	}
	s_create_url := api_tools.Api.Get_resource_creation_url(api, resourceType)
	resource_url.WriteString(s_create_url)
	resource_url.WriteString(resource_id)
	resource_url.WriteString("/")
	return resource_url.String()
}

func (apier AirDrumResources_Apier) Validate_status(api *API,
	resourceType string,
	clientTooler ClientTooler) error {

	var apiErr error
	var responseBody string
	api_tools := APITooler{
		Api: apier,
	}
	req, _ := http.NewRequest("GET",
		api_tools.Api.Get_resource_creation_url(api, resourceType),
		nil)
	req.Header.Add("authorization", "Token "+api.Token)
	resp, apiErr := clientTooler.Client.Do(api, req)

	if apiErr == nil {
		if resp.Body != nil {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			responseBody = string(bodyBytes)
			switch {
			case resp.StatusCode == http.StatusUnauthorized:
				apiErr = errors.New(resp.Status + responseBody)
			case resp.Header.Get("content-type") != "application/json":
				apiErr = errors.New("Could not get a proper json response from \"" +
					api.URL + "\", the api is down or this url is wrong.")
			}
		} else {
			apiErr = errors.New("Could not get a response body from \"" + api.URL +
				"\", the api is down or this url is wrong.")
		}
	} else {
		apiErr = errors.New("Could not get a response from \"" + api.URL +
			"\", the api is down or this url is wrong.")
	}

	return apiErr
}
