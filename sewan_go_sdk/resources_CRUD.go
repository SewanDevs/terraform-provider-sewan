package sewan_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
)

//------------------------------------------------------------------------------
func (apier AirDrumResources_Apier) Create_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	schemaTools *SchemaTooler,
	resourceType string,
	sewan *API) (error, map[string]interface{}) {
	var (
		resource_instance_creation_err error = nil
		create_req_err                 error = nil
		createError                    error = nil
		create_resp_body_err           error = nil
		created_resource               map[string]interface{}
		resourceInstance               interface{}
		responseBody                   string
		instanceName                   string = d.Get("name").(string)
		resource_json                  []byte
		resp_body_reader               interface{}
		bodyBytes                      []byte
	)
	api_tools := APITooler{
		Api: apier,
	}
	req := &http.Request{}
	resp := &http.Response{}
	resource_instance_creation_err,
		resourceInstance = api_tools.Api.ResourceInstanceCreate(d,
		clientTooler,
		templatesTooler,
		schemaTools,
		resourceType,
		sewan)
	logger := loggerCreate("create_resource_" + instanceName + ".log")

	if resource_instance_creation_err == nil {
		logger.Println("resourceInstance = ", resourceInstance)
		resource_json, create_req_err = json.Marshal(resourceInstance)
		if create_req_err == nil {
			req, create_req_err = http.NewRequest("POST",
				api_tools.Api.Get_resource_creation_url(sewan, resourceType),
				bytes.NewBuffer(resource_json))
			logger.Println("req.Body = ", req.Body)
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
				if resp.Body != nil {
					defer resp.Body.Close()
					bodyBytes, create_resp_body_err = ioutil.ReadAll(resp.Body)
					responseBody = string(bodyBytes)
				} else {
					bodyBytes = []byte{}
					create_resp_body_err = nil
					responseBody = ""
				}
				switch resp.Header.Get("Content-Type") {
				case "application/json":
					resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
					switch {
					case create_resp_body_err != nil:
						createError = errors.New("Read of " + instanceName +
							" response body error " + create_resp_body_err.Error())
					case resp_body_json_err != nil:
						createError = errors.New("Creation of \"" + instanceName +
							"\" failed, " +
							"the response body is not a properly formated json :\n\r\"" +
							resp_body_json_err.Error() + "\"")
					default:
						if resp.StatusCode == http.StatusCreated {
							created_resource = resp_body_reader.(map[string]interface{})
							for key, value := range created_resource {
								read_value,
									updateError := schemaTools.SchemaTools.Read_element(key,
									value,
									logger)
								if updateError == nil {
									created_resource[key] = read_value
								}
							}
						} else {
							createError = errors.New(resp.Status + responseBody)
						}
					}
				case "text/html":
					createError = errors.New(resp.Status + responseBody)
				default:
					createError = errors.New("Unhandled api response type : " +
						resp.Header.Get("Content-Type") +
						"\nPlease validate the configuration api url.")
				}
			}
		} else {
			createError = create_req_err
		}
	} else {
		createError = resource_instance_creation_err
	}
	logger.Println("createError = ", createError,
		"\ncreated_resource = ", created_resource)
	return createError, created_resource
}

//------------------------------------------------------------------------------
func (apier AirDrumResources_Apier) Read_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	schemaTools *SchemaTooler,
	resourceType string,
	sewan *API) (error, map[string]interface{}, bool) {

	var (
		readError                      error = nil
		read_req_err                   error = nil
		resource_instance_creation_err error = nil
		read_resp_body_err             error = nil
		read_resource                  map[string]interface{}
		responseBody                   string
		resp_body_reader               interface{}
		resource_exists                bool   = true
		instanceName                   string = d.Get("name").(string)
		bodyBytes                      []byte
	)
	req := &http.Request{}
	resp := &http.Response{}
	logger := loggerCreate("read_resource_" + instanceName + ".log")
	api_tools := APITooler{
		Api: apier,
	}
	resource_instance_creation_err = api_tools.Api.ValidateResourceType(resourceType)

	if resource_instance_creation_err == nil {
		req, read_req_err = http.NewRequest("GET",
			api_tools.Api.Get_resource_url(sewan, resourceType, d.Id()), nil)
		if read_req_err == nil {
			req.Header.Add("authorization", "Token "+sewan.Token)
			resp, read_req_err = clientTooler.Client.Do(sewan, req)
		}

		if resp != nil {
			if read_req_err != nil {
				readError = errors.New("Read of \"" + instanceName +
					"\" state failed, response reception error : " + read_req_err.Error())
			} else {
				if resp.Body != nil {
					defer resp.Body.Close()
					bodyBytes, read_resp_body_err = ioutil.ReadAll(resp.Body)
					responseBody = string(bodyBytes)
				} else {
					bodyBytes = []byte{}
					read_resp_body_err = nil
					responseBody = ""
				}
				switch resp.Header.Get("Content-Type") {
				case "application/json":
					switch {
					case read_resp_body_err != nil:
						readError = errors.New("Read of " + instanceName +
							" state response body read error " + read_resp_body_err.Error())
					case resp.StatusCode == http.StatusOK:
						resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
						if resp_body_json_err != nil {
							readError = errors.New("Read of \"" + instanceName +
								"\" failed, response body json error :\n\r\"" +
								resp_body_json_err.Error() + "\"")
						} else {
							read_resource = resp_body_reader.(map[string]interface{})

							for key, value := range read_resource {
								read_value,
									updateError := schemaTools.SchemaTools.Read_element(key,
									value,
									logger)
								if updateError == nil {
									read_resource[key] = read_value
								}
							}
						}
					case resp.StatusCode == http.StatusNotFound:
						resource_exists = false
					default:
						readError = errors.New(resp.Status + responseBody)
					}
				case "text/html":
					readError = errors.New(resp.Status + responseBody)
				default:
					readError = errors.New("Unhandled api response type : " +
						resp.Header.Get("Content-Type") +
						"\nPlease validate the configuration api url.")
				}
			}
		} else {
			readError = read_req_err
		}
	} else {
		readError = resource_instance_creation_err
	}

	logger.Println("readError =", readError,
		"\nread_resource =", read_resource,
		"\nresource_exists =", resource_exists,
	)
	return readError, read_resource, resource_exists
}

//------------------------------------------------------------------------------
func (apier AirDrumResources_Apier) Update_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	schemaTools *SchemaTooler,
	resourceType string,
	sewan *API) error {

	var (
		resource_instance_creation_err error = nil
		updateError                    error = nil
		update_req_err                 error = nil
		update_resp_body_err           error = nil
		resourceInstance               interface{}
		responseBody                   string
		instanceName                   string = d.Get("name").(string)
		resource_json                  []byte
		resp_body_reader               interface{}
		bodyBytes                      []byte
	)
	req := &http.Request{}
	resp := &http.Response{}
	api_tools := APITooler{
		Api: apier,
	}
	logger := loggerCreate("update_resource_" + instanceName + ".log")
	resource_instance_creation_err,
		resourceInstance = api_tools.Api.ResourceInstanceCreate(d,
		clientTooler,
		templatesTooler,
		schemaTools,
		resourceType,
		sewan)

	if resource_instance_creation_err == nil {

		resource_json, update_req_err = json.Marshal(resourceInstance)
		if update_req_err == nil {
			req, update_req_err = http.NewRequest("PUT",
				api_tools.Api.Get_resource_url(sewan, resourceType, d.Id()),
				bytes.NewBuffer(resource_json))
			logger.Println("req.Body = ", req.Body)
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
				if resp.Body != nil {
					defer resp.Body.Close()
					bodyBytes, update_resp_body_err = ioutil.ReadAll(resp.Body)
					responseBody = string(bodyBytes)
				} else {
					bodyBytes = []byte{}
					update_resp_body_err = nil
					responseBody = ""
				}
				switch resp.Header.Get("Content-Type") {
				case "application/json":
					switch {
					case update_resp_body_err != nil:
						updateError = errors.New("Read of \"" + instanceName +
							"\" state response body read error " + update_resp_body_err.Error())
					case resp.StatusCode == http.StatusOK:
						resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
						if resp_body_json_err != nil {
							updateError = errors.New("Read of \"" + instanceName +
								"\" failed, response body json error :\n\r\"" +
								resp_body_json_err.Error())
						}
					default:
						updateError = errors.New(resp.Status + responseBody)
					}
				case "text/html":
					updateError = errors.New(resp.Status + responseBody)
				default:
					updateError = errors.New("Unhandled api response type : " +
						resp.Header.Get("Content-Type") +
						"\nPlease validate the configuration api url.")
				}
			}
		} else {
			updateError = update_req_err
		}

	} else {
		updateError = resource_instance_creation_err
	}

	logger.Println("updateError = ", updateError)
	return updateError
}

//------------------------------------------------------------------------------
func (apier AirDrumResources_Apier) Delete_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	schemaTools *SchemaTooler,
	resourceType string,
	sewan *API) error {

	var (
		resource_instance_creation_err error = nil
		deleteError                    error = nil
		delete_req_err                 error = nil
		delete_resp_body_err           error = nil
		responseBody                   string
		resp_body_reader               interface{}
		bodyBytes                      []byte
		instanceName                   string = d.Get("name").(string)
	)
	api_tools := APITooler{
		Api: apier,
	}
	resource_instance_creation_err = api_tools.Api.ValidateResourceType(resourceType)
	req := &http.Request{}
	resp := &http.Response{}
	logger := loggerCreate("delete_resource_" + instanceName + ".log")
	logger.Println("--------------- ", instanceName, " ( id= ", d.Id(),
		") DELETE -----------------")

	if resource_instance_creation_err == nil {

		req, delete_req_err = http.NewRequest("DELETE",
			api_tools.Api.Get_resource_url(sewan, resourceType, d.Id()), nil)
		if delete_req_err == nil {
			req.Header.Add("authorization", "Token "+sewan.Token)
			resp, delete_req_err = clientTooler.Client.Do(sewan, req)
		}

		if resp != nil {
			if delete_req_err != nil {
				deleteError = errors.New("Deletion of \"" + instanceName +
					"\" state failed, response reception error : " + delete_req_err.Error())
			} else {
				if resp.Body != nil {
					defer resp.Body.Close()
					bodyBytes, delete_resp_body_err = ioutil.ReadAll(resp.Body)
					responseBody = string(bodyBytes)
				} else {
					bodyBytes = []byte{}
					delete_resp_body_err = nil
					responseBody = ""
				}
				if resp.StatusCode != http.StatusNoContent {
					switch resp.Header.Get("Content-Type") {
					case "application/json":
						switch {
						case delete_resp_body_err != nil:
							deleteError = errors.New("Deletion of " + instanceName +
								" response reception error : " + delete_resp_body_err.Error())
						default:
							resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
							switch {
							case resp_body_json_err != nil:
								deleteError = errors.New("Read of \"" + instanceName +
									"\" failed, response body json error :\n\r\"" +
									resp_body_json_err.Error())
							default:
								deleteError = errors.New(resp.Status + responseBody)
							}
						}
					case "text/html":
						deleteError = errors.New(resp.Status + responseBody)
					default:
						deleteError = errors.New("Unhandled api response type : " +
							resp.Header.Get("Content-Type") +
							"\nPlease validate the configuration api url.")
					}
				}
			}
		} else {
			deleteError = delete_req_err
		}
	} else {
		deleteError = resource_instance_creation_err
	}

	logger.Println("deleteError = ", deleteError)
	return deleteError
}
