package sewan_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
)

type VDC struct {
	Name          string        `json:"name"`
	Enterprise    string        `json:"enterprise"`
	Datacenter    string        `json:"datacenter"`
	Vdc_resources []interface{} `json:"vdc_resources"`
	Slug          string        `json:"slug"`
	Dynamic_field string        `json:"dynamic_field"`
}

const (
	VDC_DESTROY_FAILURE = "{\"detail\":\"Destroying the VM now\"}"
)


//------------------------------------------------------------------------------
func vdcInstanceCreate(d *schema.ResourceData) VDC {
	return VDC{
		Name:          d.Get("name").(string),
		Enterprise:    d.Get("enterprise").(string),
		Datacenter:    d.Get("datacenter").(string),
		Vdc_resources: d.Get("vdc_resources").([]interface{}),
		Slug:          d.Get("slug").(string),
		Dynamic_field: d.Get("dynamic_field").(string),
	}
}

//------------------------------------------------------------------------------
func (apier AirDrumAPIer) Create_vdc_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	sewan *API) (error, map[string]interface{}) {
	var (
		resp_body_reader           interface{}
		create_req_err             error
		createError                error
		create_resp_body_err       error
		airDrumAPICreationResponse map[string]interface{}
		responseBody               string
		vdc_json                   []byte
		bodyBytes                  []byte
	)
	req := &http.Request{}
	resp := &http.Response{}
	vdcInstance := vdcInstanceCreate(d)
	createError = nil
	create_req_err = nil
	create_resp_body_err = nil
	airDrumAPICreationResponse = nil
	logger := loggerCreate("create_vdc_" + vdcInstance.Name + "\".log")
	//api_tools := APITooler{
	//	Api: apier,
	//}

	vdc_json, create_req_err = json.Marshal(vdcInstance)
	if create_req_err == nil {
		req, create_req_err = http.NewRequest("POST",
			"https://next.cloud-datacenter.fr/api/clouddc/vdc/", //api_tools.Api.Get_vdc_creation_url(sewan),
			bytes.NewBuffer(vdc_json))
		if create_req_err == nil {
			req.Header.Add("authorization", "Token "+sewan.Token)
			req.Header.Add("content-type", "application/json")
			resp, create_req_err = clientTooler.Client.Do(sewan, req)
		}
	}

	if resp != nil {
		if create_req_err != nil {
			createError = errors.New("Creation of \"" + vdcInstance.Name +
				"\" failed, response reception error : " + create_req_err.Error())
		} else {
			defer resp.Body.Close()
			bodyBytes, create_resp_body_err = ioutil.ReadAll(resp.Body)
			responseBody = string(bodyBytes)
			resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
			switch {
			case create_resp_body_err != nil:
				createError = errors.New("Read of " + vdcInstance.Name +
					" response body error " + create_resp_body_err.Error())
			case resp_body_json_err != nil:
				createError = errors.New("Creation of \"" + vdcInstance.Name + "\" failed, " +
					"response body json error :\n\r\"" + resp_body_json_err.Error() + "\"")
			default:
				if resp.StatusCode != http.StatusCreated {
					createError = errors.New(resp.Status + responseBody)
				} else {
					airDrumAPICreationResponse = resp_body_reader.(map[string]interface{})
				}
			}
		}
	} else {
		createError = create_req_err
	}

	logger.Println("createError = ", createError,
		"\nairDrumAPICreationResponse = ", airDrumAPICreationResponse)
	return createError, airDrumAPICreationResponse
}

//------------------------------------------------------------------------------
func (apier AirDrumAPIer) Read_vdc_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	sewan *API) (error, map[string]interface{}, bool) {

	var (
		readError                  error
		read_req_err               error
		airDrumAPICreationResponse map[string]interface{}
		responseBody               string
		resp_body_reader           interface{}
		resource_exists            bool
	)
	req := &http.Request{}
	resp := &http.Response{}
	readError = nil
	read_req_err = nil
	airDrumAPICreationResponse = nil
	resource_exists = true
	logger := loggerCreate("read_vdc_" + d.Get("name").(string) + ".log")
	logger.Println("--------------- ", d.Get("name").(string),
		" ( id= ", d.Id(), ") READ -----------------")
	//api_tools := APITooler{
	//	Api: apier,
	//}

	req, read_req_err = http.NewRequest("GET",
		"https://next.cloud-datacenter.fr/api/clouddc/vdc/"+d.Id(), //api_tools.Api.Get_vdc_url(sewan, d.Id()),
		nil)
	if read_req_err == nil {
		req.Header.Add("authorization", "Token "+sewan.Token)
		resp, read_req_err = clientTooler.Client.Do(sewan, req)
	}

	if resp != nil {
		if read_req_err != nil {
			readError = errors.New("Read of \"" + d.Get("name").(string) +
				"\" state failed, response reception error : " + read_req_err.Error())
		} else {
			defer resp.Body.Close()
			bodyBytes, read_resp_body_err := ioutil.ReadAll(resp.Body)
			responseBody = string(bodyBytes)
			switch {
			case read_resp_body_err != nil:
				readError = errors.New("Read of " + d.Get("name").(string) +
					" state response body read error " + read_resp_body_err.Error())
			case resp.StatusCode == http.StatusOK:
				resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
				if resp_body_json_err != nil {
					readError = errors.New("Read of \"" + d.Get("name").(string) +
						"\" failed, response body json error :\n\r\"" +
						resp_body_json_err.Error() + "\"")
				} else {
					airDrumAPICreationResponse = resp_body_reader.(map[string]interface{})
				}
			case resp.StatusCode == http.StatusNotFound:
				resource_exists = false
			default:
				readError = errors.New(resp.Status + responseBody)
			}
		}
	} else {
		readError = read_req_err
	}
	logger.Println("readError =", readError,
		"\nairDrumAPICreationResponse =", airDrumAPICreationResponse,
		"\nresource_exists =", resource_exists)
	return readError, airDrumAPICreationResponse, resource_exists

}

//------------------------------------------------------------------------------
func (apier AirDrumAPIer) Update_vdc_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	sewan *API) error {
		var (
			updateError          error
			update_req_err       error
			update_resp_body_err error
			responseBody         string
			resp_body_reader     interface{}
			bodyBytes            []byte
		)
		req := &http.Request{}
		resp := &http.Response{}
		vdcInstance := vdcInstanceCreate(d)
		updateError = nil
		update_req_err = nil
		update_resp_body_err = nil
		logger := loggerCreate("update_vdc_" + d.Get("name").(string) + ".log")
		logger.Println("--------------- ", d.Get("name").(string), " ( id= ", d.Id(), ") UPDATE -----------------")
		//api_tools := APITooler{
		//	Api: apier,
		//}

		vdc_json, update_req_err := json.Marshal(vdcInstance)
		if update_req_err == nil {
			req, update_req_err = http.NewRequest("PUT",
				"https://next.cloud-datacenter.fr/api/clouddc/vdc/"+d.Id(),//api_tools.Api.Get_vdc_url(sewan, d.Id()),
				bytes.NewBuffer(vdc_json))
			if update_req_err == nil {
				req.Header.Add("authorization", "Token "+sewan.Token)
				req.Header.Add("content-type", "application/json")
				resp, update_req_err = clientTooler.Client.Do(sewan, req)
			}
		}

		if resp != nil {
			if update_req_err != nil {
				updateError = errors.New("Update of \"" + vdcInstance.Name +
					"\" state failed, response reception error : " + update_req_err.Error())
			} else {
				defer resp.Body.Close()
				bodyBytes, update_resp_body_err = ioutil.ReadAll(resp.Body)
				responseBody = string(bodyBytes)
				switch {
				case update_resp_body_err != nil:
					updateError = errors.New("Update of \"" + d.Get("name").(string) +
						"\" state response body read error " + update_resp_body_err.Error())
				case resp.StatusCode == http.StatusOK:
					resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
					if resp_body_json_err != nil {
						updateError = errors.New("Update of \"" + d.Get("name").(string) +
							"\" failed, response body json error :\n\r\"" +
							resp_body_json_err.Error())
					}
				case resp.StatusCode != http.StatusOK:
					updateError = errors.New("" + resp.Status + responseBody)
				default:
					updateError = errors.New(resp.Status + responseBody)
				}
			}
		} else {
			updateError = update_req_err
		}

		logger.Println("updateError = ", updateError)
		return updateError
}

//------------------------------------------------------------------------------
func (apier AirDrumAPIer) Delete_vdc_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	sewan *API) error {

		var (
			deleteError          error
			delete_req_err       error
			delete_resp_body_err error
			responseBody         string
			resp_body_reader     interface{}
			bodyBytes            []byte
		)
		req := &http.Request{}
		resp := &http.Response{}
		deleteError = nil
		delete_req_err = nil
		logger := loggerCreate("delete_vdc_" + d.Get("name").(string) + ".log")
		logger.Println("--------------- ", d.Get("name").(string), " ( id= ", d.Id(), ") DELETE -----------------")
		//api_tools := APITooler{
		//	Api: apier,
		//}

		req, delete_req_err = http.NewRequest("DELETE",
			"https://next.cloud-datacenter.fr/api/clouddc/vdc/"+d.Id(),//api_tools.Api.Get_vdc_url(sewan, d.Id()),
			nil)
		if delete_req_err == nil {
			req.Header.Add("authorization", "Token "+sewan.Token)
			resp, delete_req_err = clientTooler.Client.Do(sewan, req)
		}

		if resp != nil {
			if delete_req_err != nil {
				deleteError = errors.New("Deletion of \"" + d.Get("name").(string) +
					"\" state failed, response reception error : " + delete_req_err.Error())
			} else {
				defer resp.Body.Close()
				bodyBytes, delete_resp_body_err = ioutil.ReadAll(resp.Body)
				responseBody = string(bodyBytes)
				switch {
				case delete_resp_body_err != nil:
					deleteError = errors.New("Deletion of " + d.Get("name").(string) +
						" response reception error : " + delete_resp_body_err.Error())
				case resp.StatusCode == http.StatusOK:
					resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
					if resp_body_json_err != nil {
						deleteError = errors.New("Read of \"" + d.Get("name").(string) +
							"\" failed, response body json error :\n\r\"" +
							resp_body_json_err.Error())
					} else if responseBody != VDC_DESTROY_FAILURE {
						deleteError = errors.New(resp.Status + responseBody)
					}
				default:
					deleteError = errors.New(resp.Status + responseBody)
				}
			}
		} else {
			deleteError = delete_req_err
		}

		logger.Println("deleteError = ", deleteError)
		return deleteError
}
