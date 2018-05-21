package sewan_go_sdk

import (
  "bytes"
  "encoding/json"
  "errors"
  "github.com/hashicorp/terraform/helper/schema"
  "io/ioutil"
  "net/http"
  "strings"
)

const RESOURCE_CREATE_HTTP_SUCCESS_CODE = 201
const RESOURCE_UPDATE_HTTP_SUCCESS_CODE = 200
const RESOURCE_DELETE_HTTP_SUCCESS_CODE = 200
const RESOURCE_GET_HTTP_SUCCESS_CODE = 200
const RESOURCE_GET_HTTP_NOT_FOUND_CODE = 404

func get_vm_url(api_url string, id string) string{
  var vm_url strings.Builder
  vm_url.WriteString(api_url)
  vm_url.WriteString(id)
  vm_url.WriteString("/")
  return vm_url.String()
}

type vm struct {
  Name              string        `json:"name"`
  State             string        `json:"state"`
  OS                string        `json:"os"`
  RAM               string        `json:"ram"`
  CPU               string        `json:"cpu"`
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

func vmInstanceCreate(d *schema.ResourceData) vm {
  return vm{
    Name:              d.Get("name").(string),
    State:             d.Get("state").(string),
    OS:                d.Get("os").(string),
    RAM:               d.Get("ram").(string),
    CPU:               d.Get("cpu").(string),
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

func (client Http_client) Create_vm_resource(d *schema.ResourceData) (error, map[string]interface{}) {
  vmInstance := vmInstanceCreate(d)
  var responseBody string
  var resp_body_reader interface{}
  var createError error
  createError = nil
  var airDrumAPICreationResponse map[string]interface{}
  logger := loggerCreate("create_vm_" + vmInstance.Vdc + "_" + vmInstance.Name + ".log")

  vm_json, err_json := json.Marshal(vmInstance)
  logger.Println("vmInstance =", vmInstance)
  logger.Println("err_json =", err_json)
  logger.Println("vm_json =", vm_json)

  req, _ := http.NewRequest("POST", client.Conf.Api_url, bytes.NewBuffer(vm_json))

  req.Header.Add("authorization", "Token "+client.Conf.Api_token)
  req.Header.Add("content-type", "application/json")

  logger.Println("Creation of ", vmInstance.Name, "request Header = ", req.Header)
  logger.Println("Creation of ", vmInstance.Name, "request body = ", req.Body)

  resp, create_err := client.Net_http_client.Do(req)
  defer resp.Body.Close()
  bodyBytes, create_resp_body_read_err := ioutil.ReadAll(resp.Body)
  responseBody = string(bodyBytes)

  if create_err != nil {
    logger.Println("Creation of ", vmInstance.Name, " response reception error : ", create_err)
    createError = errors.New("Creation of " + vmInstance.Name + " response reception error : " + create_err.Error())
  }

  if create_resp_body_read_err != nil {
    logger.Println("Creation of ", vmInstance.Name, " response body read error ", create_resp_body_read_err)
    createError = errors.New("Creation of " + vmInstance.Name + " response body read error " + create_resp_body_read_err.Error())
  }

  logger.Println("Creation of ", vmInstance.Name, " response status = ", resp.Status)
  logger.Println("Creation of ", vmInstance.Name, " response body = ", responseBody)

  if resp.StatusCode != RESOURCE_CREATE_HTTP_SUCCESS_CODE {
    logger.Println("Creation of ", vmInstance.Name, " resource failed : ", resp.Status, responseBody)
    createError = errors.New(resp.Status + responseBody)
  } else {
    resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
    if resp_body_json_err != nil {
      createError = resp_body_json_err
    }
    airDrumAPICreationResponse = resp_body_reader.(map[string]interface{})
  }
  return createError, airDrumAPICreationResponse
}

func Read_vm_resource(d *schema.ResourceData,client Http_client) (error, map[string]interface{}, bool) {
  var readError error
  readError = nil
  var resource_exists bool
  resource_exists = true
  var airDrumAPICreationResponse map[string]interface{}
  var responseBody string
  var resp_body_reader interface{}
  s_vm_url := get_vm_url(client.Conf.Api_url,d.Id())
  logger := loggerCreate("read_vm_" + d.Get("name").(string) + ".log")
  logger.Println("--------------- ", d.Get("name").(string), " ( id= ", d.Id(), ") READ -----------------")

  req, _ := http.NewRequest("GET", s_vm_url, nil)

  req.Header.Add("authorization", "Token "+client.Conf.Api_token)

  resp, read_req_err := client.Net_http_client.Do(req)
  defer resp.Body.Close()

  bodyBytes, read_resp_body_read_err := ioutil.ReadAll(resp.Body)
  responseBody = string(bodyBytes)

  if read_req_err != nil {
    logger.Println("Read of ", d.Get("name").(string), " state reception error : ", read_req_err)
    readError = errors.New("Read of " + d.Get("name").(string) + " state reception error : " + read_req_err.Error())
  }

  if read_resp_body_read_err != nil {
    logger.Println("Read of ", d.Get("name").(string), " state response body read error ", read_resp_body_read_err)
    readError = errors.New("Read of " + d.Get("name").(string) + " state response body read error " + read_resp_body_read_err.Error())
  }

  if resp.StatusCode == RESOURCE_GET_HTTP_SUCCESS_CODE {
    resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
    if resp_body_json_err != nil {
      readError = resp_body_json_err
    } else {
      airDrumAPICreationResponse = resp_body_reader.(map[string]interface{})
    }
  } else if resp.StatusCode == RESOURCE_GET_HTTP_NOT_FOUND_CODE {
    logger.Println(d.Get("name").(string), " not found, The resource may have been deleted by an other Airdrum API client.")
    resource_exists = false
  } else {
    logger.Println("Unknow error : ")
    readError = errors.New("Unknow error : " + resp.Status + responseBody)
  }
  return readError,airDrumAPICreationResponse,resource_exists
}

func Update_vm_resource(d *schema.ResourceData,client Http_client) (error) {
  var responseBody string
  var updateError error
  updateError = nil
  vmInstance := vmInstanceCreate(d)
  s_vm_url := get_vm_url(client.Conf.Api_url,d.Id())
  logger := loggerCreate("update_vm_" + d.Get("name").(string) + ".log")

  logger.Println("--------------- ", d.Get("name").(string), " ( id= ", d.Id(), ") UPDATE -----------------")

  vm_json, err_json := json.Marshal(vmInstance)
  logger.Println("vmInstance =", vmInstance)
  logger.Println("err_json =", err_json)
  logger.Println("vm_json =", vm_json)

  req, _ := http.NewRequest("PUT", s_vm_url, bytes.NewBuffer(vm_json))

  req.Header.Add("authorization", "Token "+client.Conf.Api_token)
  req.Header.Add("content-type", "application/json")

  logger.Println("Update of ", d.Get("name").(string), "request Header = ", req.Header)
  logger.Println("Update of ", d.Get("name").(string), "request body = ", req.Body)

  resp, create_err := client.Net_http_client.Do(req)
  defer resp.Body.Close()
  bodyBytes, update_resp_body_read_err := ioutil.ReadAll(resp.Body)
  responseBody = string(bodyBytes)

  if create_err != nil {
    logger.Println("Update of ", d.Get("name").(string), " response reception error : ", create_err)
    updateError = errors.New("Update of " + d.Get("name").(string) + " response reception error : " + create_err.Error())
  }

  if update_resp_body_read_err != nil {
    logger.Println("Update of ", d.Get("name").(string), " response body read error ", update_resp_body_read_err)
    updateError = errors.New("Update of " + d.Get("name").(string) + " response body read error " + update_resp_body_read_err.Error())
  }

  logger.Println("Update of ", d.Get("name").(string), " response status = ", resp.Status)
  logger.Println("Update of ", d.Get("name").(string), " response body = ", responseBody)

  if resp.StatusCode != RESOURCE_UPDATE_HTTP_SUCCESS_CODE {
    logger.Println("Update of ", d.Get("name").(string), " resource failed : ", resp.Status, responseBody)
    updateError = errors.New(resp.Status + responseBody)
  }

  return updateError
}

func Delete_vm_resource(d *schema.ResourceData,client Http_client) (error) {
  var responseBody string
  var deleteError error
  deleteError = nil
  s_vm_url := get_vm_url(client.Conf.Api_url,d.Id())
  logger := loggerCreate("update_vm_" + d.Get("name").(string) + ".log")

  logger.Println("--------------- ", d.Get("name").(string), " ( id= ", d.Id(), ") DELETE -----------------")

  req, _ := http.NewRequest("DELETE", s_vm_url, nil)

  req.Header.Add("authorization", "Token "+client.Conf.Api_token)

  resp, delete_err := client.Net_http_client.Do(req)
  defer resp.Body.Close()

  bodyBytes, delete_resp_body_read_err := ioutil.ReadAll(resp.Body)
  responseBody = string(bodyBytes)

  if delete_err != nil {
    logger.Println("Deletion of ", d.Get("name").(string), " response reception error : ", delete_err)
    deleteError = errors.New("Deletion of " + d.Get("name").(string) + " response reception error : " + delete_err.Error())
  }

  if delete_resp_body_read_err != nil {
    logger.Println("Deletion of ", d.Get("name").(string), " response body read error ", delete_resp_body_read_err)
    deleteError = errors.New("Deletion of " + d.Get("name").(string) + " response reception error : " + delete_err.Error())
  }

  if resp.StatusCode != RESOURCE_DELETE_HTTP_SUCCESS_CODE && responseBody != "{\"detail\":\"Destroying the VM now\"}" {
    logger.Println("Deletion of ", d.Get("name").(string), " resource failed : ", resp.Status, responseBody)
    deleteError = errors.New(resp.Status + responseBody)
  }

  return deleteError
}
