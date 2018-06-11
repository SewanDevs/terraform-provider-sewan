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

type VM struct {
	Name              string        `json:"name"`
	State             string        `json:"state"`
	OS                string        `json:"os"`
	RAM               int           `json:"ram"`
	CPU               int           `json:"cpu"`
	Disks             []interface{} `json:"disks"`
	Nics              []interface{} `json:"nics"`
	Vdc               string        `json:"vdc"`
	Boot              string        `json:"boot"`
	Vdc_resource_disk string        `json:"vdc_resource_disk"`
	//Template string `json:"template"`
	Slug          string `json:"slug"`
	Token         string `json:"token"`
	Backup        string `json:"backup"`
	Disk_image    string `json:"disk_image"`
	Platform_name string `json:"platform_name"`
	Backup_size   string `json:"backup_size"`
	Comment       string `json:"comment"`
	Outsourcing   string `json:"outsourcing"`
	Dynamic_field string `json:"dynamic_field"`
}

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

func vmInstanceCreate(d *schema.ResourceData) VM {
	return VM{
		Name:              d.Get("name").(string),
		State:             d.Get("state").(string),
		OS:                d.Get("os").(string),
		RAM:               d.Get("ram").(int),
		CPU:               d.Get("cpu").(int),
		Disks:             d.Get("disks").([]interface{}),
		Nics:              d.Get("nics").([]interface{}),
		Vdc:               d.Get("vdc").(string),
		Boot:              d.Get("boot").(string),
		Vdc_resource_disk: d.Get("vdc_resource_disk").(string),
		//Template:  d.Get("template").(string),
		Slug:          d.Get("slug").(string),
		Token:         d.Get("token").(string),
		Backup:        d.Get("backup").(string),
		Disk_image:    d.Get("disk_image").(string),
		Platform_name: d.Get("platform_name").(string),
		Backup_size:   d.Get("backup_size").(string),
		Comment:       d.Get("comment").(string),
		Outsourcing:   d.Get("outsourcing").(string),
		Dynamic_field: d.Get("dynamic_field").(string),
	}
}

func resourceInstanceCreate(d *schema.ResourceData,
	resourceType string) (interface{}, string, error) {

	var (
		resourceInstance interface{}
		instanceName     string
		err              error
	)

	switch resourceType {
	case "vdc":
		resourceInstance = vdcInstanceCreate(d)
		instanceName = d.Get("name").(string)
		err = nil
	case "vm":
		resourceInstance = vmInstanceCreate(d)
		instanceName = d.Get("name").(string)
		err = nil
	default:
		resourceInstance = nil
		instanceName = ""
		err = errors.New("Resource of type \"" + resourceType + "\" not supported," +
			"list of accepted resource types :\n\r" +
			"- \"vdc\"" +
			"- \"vm\"")
	}

	return resourceInstance, instanceName, err
}

const (
	VM_DESTROY_FAILURE  = "{\"detail\":\"Destroying the VM now\"}"
	VDC_DESTROY_FAILURE = "{\"detail\":\"Destroying the VDC now\"}"
)

//------------------------------------------------------------------------------
func (apier AirDrumResources_Apier) Create_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceType string,
	sewan *API) (error, map[string]interface{}) {

	var (
		resourceTypeErr            error
		create_req_err             error
		createError                error
		create_resp_body_err       error
		airDrumAPICreationResponse map[string]interface{}
		resourceInstance           interface{}
		responseBody               string
		instanceName               string
		resource_json              []byte
		resp_body_reader           interface{}
		bodyBytes                  []byte
	)
	req := &http.Request{}
	resp := &http.Response{}
	resourceInstance, instanceName, resourceTypeErr = resourceInstanceCreate(d,
		resourceType)
	createError = nil
	create_req_err = nil
	create_resp_body_err = nil
	airDrumAPICreationResponse = nil
	logger := loggerCreate("create_resource__\"")// +
		//instanceName + "\".log")
	api_tools := APITooler{
		Api: apier,
	}
	logger.Println("d =",d)
	logger.Println("clientTooler =",clientTooler)
	logger.Println("resourceType =",resourceType)
	logger.Println("sewanAPI =",sewan)

	if resourceTypeErr == nil {

		resource_json, create_req_err = json.Marshal(resourceInstance)
		if create_req_err == nil {
			req, create_req_err = http.NewRequest("POST",
				api_tools.Api.Get_resource_creation_url(sewan, resourceType),
				bytes.NewBuffer(resource_json))
			if create_req_err == nil {
				req.Header.Add("authorization", "Token "+sewan.Token)
				req.Header.Add("content-type", "application/json")
				resp, create_req_err = clientTooler.Client.Do(sewan, req)
			}
		}

		if resp != nil {
			if create_req_err != nil {
				createError = errors.New("Creation of \"" + instanceName +
					"\" failed, response reception error : " + create_req_err.Error())
			} else {
				defer resp.Body.Close()
				bodyBytes, create_resp_body_err = ioutil.ReadAll(resp.Body)
				responseBody = string(bodyBytes)
				resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
				switch {
				case create_resp_body_err != nil:
					createError = errors.New("Read of " + instanceName +
						" response body error " + create_resp_body_err.Error())
				case resp_body_json_err != nil:
					logger.Println("resp_body_json_err != nil\nresp.Body = ", resp.Body,
						"\nresp.StatusCode = ", resp.StatusCode,
						"\nresp.Status =", resp.Status,
						"\nresp.Header =", resp.Header,
						"\nresponseBody = ", responseBody)
					createError = errors.New("Creation of \"" + instanceName +
						"\" failed, " +
						"the response body is not a properly formated json :\n\r\"" +
						resp_body_json_err.Error() + "\"")
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

	} else {
		createError = resourceTypeErr
	}

	logger.Println("createError = ", createError,
		"\nairDrumAPICreationResponse = ", airDrumAPICreationResponse)
	return createError, airDrumAPICreationResponse
}

//------------------------------------------------------------------------------
func (apier AirDrumResources_Apier) Read_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceType string,
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
	logger := loggerCreate("read_resource_" + d.Get("name").(string) + ".log")
	logger.Println("--------------- ", d.Get("name").(string),
		" ( id= ", d.Id(), ") READ -----------------")
	api_tools := APITooler{
		Api: apier,
	}

	req, read_req_err = http.NewRequest("GET",
		api_tools.Api.Get_resource_url(sewan, resourceType, d.Id()), nil)
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
func (apier AirDrumResources_Apier) Update_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceType string,
	sewan *API) error {

	var (
		resourceTypeErr      error
		updateError          error
		update_req_err       error
		update_resp_body_err error
		resourceInstance     interface{}
		responseBody         string
		instanceName         string
		resource_json        []byte
		resp_body_reader     interface{}
		bodyBytes            []byte
	)
	req := &http.Request{}
	resp := &http.Response{}
	resourceInstance, instanceName, resourceTypeErr = resourceInstanceCreate(d,
		resourceType)
	updateError = nil
	update_req_err = nil
	update_resp_body_err = nil
	logger := loggerCreate("update_resource_" + d.Get("name").(string) + ".log")
	logger.Println("--------------- ", d.Get("name").(string), " ( id= ",
		d.Id(), ") UPDATE -----------------")
	api_tools := APITooler{
		Api: apier,
	}

	if resourceTypeErr == nil {

		resource_json, update_req_err = json.Marshal(resourceInstance)
		if update_req_err == nil {
			req, update_req_err = http.NewRequest("PUT",
				api_tools.Api.Get_resource_url(sewan, resourceType, d.Id()),
				bytes.NewBuffer(resource_json))
			if update_req_err == nil {
				req.Header.Add("authorization", "Token "+sewan.Token)
				req.Header.Add("content-type", "application/json")
				resp, update_req_err = clientTooler.Client.Do(sewan, req)
			}
		}

		if resp != nil {
			if update_req_err != nil {
				updateError = errors.New("Update of \"" + instanceName +
					"\" state failed, response reception error : " + update_req_err.Error())
			} else {
				defer resp.Body.Close()
				bodyBytes, update_resp_body_err = ioutil.ReadAll(resp.Body)
				responseBody = string(bodyBytes)
				switch {
				case update_resp_body_err != nil:
					updateError = errors.New("Read of \"" + d.Get("name").(string) +
						"\" state response body read error " + update_resp_body_err.Error())
				case resp.StatusCode == http.StatusOK:
					resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
					if resp_body_json_err != nil {
						updateError = errors.New("Read of \"" + d.Get("name").(string) +
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

	} else {
		updateError = resourceTypeErr
	}

	logger.Println("updateError = ", updateError)
	return updateError
}

//------------------------------------------------------------------------------
func (apier AirDrumResources_Apier) Delete_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceType string,
	sewan *API) error {

	var (
		resourceTypeErr          error
		deleteError              error
		delete_req_err           error
		delete_resp_body_err     error
		responseBody             string
		resp_body_reader         interface{}
		bodyBytes                []byte
		resource_destroy_failure string
	)
	switch resourceType {
	case "vdc":
		resource_destroy_failure = VDC_DESTROY_FAILURE
		resourceTypeErr = nil
	case "vm":
		resource_destroy_failure = VM_DESTROY_FAILURE
		resourceTypeErr = nil
	default:
		resource_destroy_failure = ""
		resourceTypeErr = errors.New("Resource of type \"" + resourceType +
			"\" not supported," +
			"list of accepted resource types :\n\r" +
			"- \"vdc\"" +
			"- \"vm\"")
	}
	req := &http.Request{}
	resp := &http.Response{}
	deleteError = nil
	delete_req_err = nil
	logger := loggerCreate("delete_resource_" + d.Get("name").(string) + ".log")
	logger.Println("--------------- ", d.Get("name").(string), " ( id= ", d.Id(),
		") DELETE -----------------")
	api_tools := APITooler{
		Api: apier,
	}

	if resourceTypeErr == nil {

		req, delete_req_err = http.NewRequest("DELETE",
			api_tools.Api.Get_resource_url(sewan, resourceType, d.Id()), nil)
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
					} else if responseBody != resource_destroy_failure {
						logger.Println("resource_destroy_failure", resource_destroy_failure)
						deleteError = errors.New(resp.Status + responseBody)
					}
				default:
					deleteError = errors.New(resp.Status + responseBody)
				}
			}
		} else {
			deleteError = delete_req_err
		}
	} else {
		deleteError = resourceTypeErr
	}

	logger.Println("deleteError = ", deleteError)
	return deleteError
}
