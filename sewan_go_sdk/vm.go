package sewan_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
)

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

const (
	VM_DESTROY_FAILURE = "{\"detail\":\"Destroying the VM now\"}"
)

//------------------------------------------------------------------------------
func (apier AirDrumAPIer) Create_vm_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	sewan *API) (error, map[string]interface{}) {

	var (
		resp_body_reader           interface{}
		create_req_err             error
		createError                error
		create_resp_body_err       error
		airDrumAPICreationResponse map[string]interface{}
		responseBody               string
		vm_json                    []byte
		bodyBytes                  []byte
	)
	req := &http.Request{}
	resp := &http.Response{}
	vmInstance := vmInstanceCreate(d)
	createError = nil
	create_req_err = nil
	create_resp_body_err = nil
	airDrumAPICreationResponse = nil
	logger := loggerCreate("create_vm_" + vmInstance.Vdc + "_\"" +
		vmInstance.Name + "\".log")
	api_tools := APITooler{
		Api: apier,
	}

	vm_json, create_req_err = json.Marshal(vmInstance)
	if create_req_err == nil {
		req, create_req_err = http.NewRequest("POST",
			api_tools.Api.Get_vm_creation_url(sewan),
			bytes.NewBuffer(vm_json))
		if create_req_err == nil {
			req.Header.Add("authorization", "Token "+sewan.Token)
			req.Header.Add("content-type", "application/json")
			resp, create_req_err = clientTooler.Client.Do(sewan, req)
		}
	}

	if resp != nil {
		if create_req_err != nil {
			createError = errors.New("Creation of \"" + vmInstance.Name +
				"\" failed, response reception error : " + create_req_err.Error())
		} else {
			defer resp.Body.Close()
			bodyBytes, create_resp_body_err = ioutil.ReadAll(resp.Body)
			responseBody = string(bodyBytes)
			resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
			switch {
			case create_resp_body_err != nil:
				createError = errors.New("Read of " + vmInstance.Name +
					" response body error " + create_resp_body_err.Error())
			case resp_body_json_err != nil:
				logger.Println("resp_body_json_err != nil\nresp.Body = ", resp.Body,
					"\nresp.StatusCode = ", resp.StatusCode,
					"\nresp.Status =", resp.Status,
					"\nresp.Header =", resp.Header,
					"\nresponseBody = ", responseBody)
				createError = errors.New("Creation of \"" + vmInstance.Name +
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

	logger.Println("createError = ", createError,
		"\nairDrumAPICreationResponse = ", airDrumAPICreationResponse)
	return createError, airDrumAPICreationResponse
}

//------------------------------------------------------------------------------
func (apier AirDrumAPIer) Read_vm_resource(d *schema.ResourceData,
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
	logger := loggerCreate("read_vm_" + d.Get("name").(string) + ".log")
	logger.Println("--------------- ", d.Get("name").(string),
		" ( id= ", d.Id(), ") READ -----------------")
	api_tools := APITooler{
		Api: apier,
	}

	req, read_req_err = http.NewRequest("GET",
		api_tools.Api.Get_vm_url(sewan, d.Id()), nil)
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
func (apier AirDrumAPIer) Update_vm_resource(d *schema.ResourceData,
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
	vmInstance := vmInstanceCreate(d)
	updateError = nil
	update_req_err = nil
	update_resp_body_err = nil
	logger := loggerCreate("update_vm_" + d.Get("name").(string) + ".log")
	logger.Println("--------------- ", d.Get("name").(string), " ( id= ", d.Id(), ") UPDATE -----------------")
	api_tools := APITooler{
		Api: apier,
	}

	vm_json, update_req_err := json.Marshal(vmInstance)
	if update_req_err == nil {
		req, update_req_err = http.NewRequest("PUT",
			api_tools.Api.Get_vm_url(sewan, d.Id()),
			bytes.NewBuffer(vm_json))
		if update_req_err == nil {
			req.Header.Add("authorization", "Token "+sewan.Token)
			req.Header.Add("content-type", "application/json")
			resp, update_req_err = clientTooler.Client.Do(sewan, req)
		}
	}

	if resp != nil {
		if update_req_err != nil {
			updateError = errors.New("Update of \"" + vmInstance.Name +
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

	logger.Println("updateError = ", updateError)
	return updateError
}

//------------------------------------------------------------------------------
func (apier AirDrumAPIer) Delete_vm_resource(d *schema.ResourceData,
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
	logger := loggerCreate("delete_vm_" + d.Get("name").(string) + ".log")
	logger.Println("--------------- ", d.Get("name").(string), " ( id= ", d.Id(), ") DELETE -----------------")
	api_tools := APITooler{
		Api: apier,
	}

	req, delete_req_err = http.NewRequest("DELETE", api_tools.Api.Get_vm_url(sewan, d.Id()), nil)
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
				} else if responseBody != VM_DESTROY_FAILURE {
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
