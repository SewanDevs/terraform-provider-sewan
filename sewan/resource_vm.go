package sewan

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

// NB : The following 2 vars will be deleted when the provider config will be handled
const DEST_REQ_URL = "https://next.cloud-datacenter.fr/api/clouddc/vm/"
const CONN_TOKEN = "26e9cba39d1f66bfb916f10a0815158ed55d24d7"

func resourceVM() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMCreate,
		Read:   resourceVMRead,
		Update: resourceVMUpdate,
		Delete: resourceVMDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdc": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdc_resource_disk": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ram": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			//"template": &schema.Schema{
			//	Type:     schema.TypeString,
			//	Optional: true,
			//},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"disks": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"size": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"v_disk": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"slug": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"nics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"mac_adress": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"connected": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Sensitive: true,
			},
			"platform_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"disk_image": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"boot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"outsourcing": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Sensitive: true,
			},
			"dynamic_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

type VM struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Vdc string `json:"vdc"`
	Vdc_resource_disk string `json:"vdc_resource_disk"`
	RAM string `json:"ram"`
	CPU string `json:"cpu"`
	//Template string `json:"template"`
	State string `json:"state"`
	Disks []interface{} `json:"disks"`
	OS string `json:"os"`
	Nics []interface{} `json:"nics"`
	Token string `json:"token"`
	Platform_name string `json:"platform_name"`
	Backup string `json:"backup"`
	Disk_image string `json:"disk_image"`
	Boot string `json:"boot"`
	Backup_size string `json:"backup_size"`
	Comment string `json:"comment"`
	Outsourcing string `json:"outsourcing"`
	Dynamic_field string `json:"dynamic_field"`
}

func vmInstanceCreate(d *schema.ResourceData) VM {
	return VM{
		Name: d.Get("name").(string),
		Comment:d.Get("comment").(string),
		Dynamic_field:d.Get("dynamic_field").(string),
		Vdc: d.Get("vdc").(string),
		CPU: d.Get("cpu").(string),
		RAM: d.Get("ram").(string),
		Boot: d.Get("boot").(string),
		Disks: d.Get("disks").([]interface{}),
		Nics: d.Get("nics").([]interface{}),
		//Template:	d.Get("template").(string),
		OS: d.Get("os").(string),
		Backup: d.Get("backup").(string),
		Vdc_resource_disk: d.Get("vdc_resource_disk").(string),
		Disk_image: d.Get("disk_image").(string),
		Slug: d.Get("slug").(string),
		State: d.Get("state").(string),
		Backup_size: d.Get("backup_size").(string),
		Token: d.Get("token").(string),
		Platform_name: d.Get("platform_name").(string),
		Outsourcing: d.Get("outsourcing").(string),
	}
}

func resourceVMCreate(d *schema.ResourceData, m interface{}) error {
	vmInstance := vmInstanceCreate(d)
	//var requestBody bytes.Buffer
	var responseBody string
	var resp_body_reader interface{}
	logger := loggerCreate("resourceVM_Create_" + vmInstance.Vdc + "_" + vmInstance.Name + ".log")
	var returnError error
	returnError = nil
	client := &http.Client{}

	logger.Println("--------------- ", vmInstance.Name, " CREATE -----------------")

	vm_struct := VM{
		Name: vmInstance.Name,
		Slug: vmInstance.Slug,
		Vdc:  vmInstance.Vdc,
		Vdc_resource_disk: vmInstance.Vdc_resource_disk,
		RAM: vmInstance.RAM,
		CPU: vmInstance.CPU,
		//Template: vmInstance.Template,
		State: vmInstance.State,
		Disks: vmInstance.Disks,
		OS: vmInstance.OS,
		Nics: vmInstance.Nics,
		Token: vmInstance.Token,
		Platform_name: vmInstance.Platform_name,
		Backup: vmInstance.Backup,
		Disk_image: vmInstance.Disk_image,
		Boot: vmInstance.Boot,
		Backup_size: vmInstance.Backup_size,
		Comment: vmInstance.Comment,
		Outsourcing: vmInstance.Outsourcing,
		Dynamic_field: vmInstance.Dynamic_field,
	}

	vm_json, err_json := json.Marshal(vm_struct)
	logger.Println("vm_struct =", vm_struct)
	logger.Println("err_json =", err_json)
	logger.Println("vm_json =", vm_json)

	req, _ := http.NewRequest("POST", DEST_REQ_URL, bytes.NewBuffer(vm_json))

	req.Header.Add("authorization", "Token "+CONN_TOKEN)
	req.Header.Add("content-type", "application/json")

	logger.Println("Creation of ", vmInstance.Name, "request Header = ", req.Header)
	logger.Println("Creation of ", vmInstance.Name, "request body = ", req.Body)

	resp, create_err := client.Do(req)
	defer resp.Body.Close()
	bodyBytes, create_resp_body_read_err := ioutil.ReadAll(resp.Body)
	responseBody = string(bodyBytes)

	if create_err != nil {
		logger.Println("Creation of ", vmInstance.Name, " response reception error : ", create_err)
		returnError = errors.New("Creation of " + vmInstance.Name + " response reception error : " + create_err.Error())
	}

	if create_resp_body_read_err != nil {
		logger.Println("Creation of ", vmInstance.Name, " response body read error ", create_resp_body_read_err)
		returnError = errors.New("Creation of " + vmInstance.Name + " response body read error " + create_resp_body_read_err.Error())
	}

	logger.Println("Creation of ", vmInstance.Name, " response status = ", resp.Status)
	logger.Println("Creation of ", vmInstance.Name, " response body = ", responseBody)

	if resp.StatusCode != RESOURCE_CREATE_HTTP_SUCCESS_CODE {
		logger.Println("Creation of ", vmInstance.Name, " resource failed : ", resp.Status, responseBody)
		returnError = errors.New(resp.Status + responseBody)
	} else {
		resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
		if resp_body_json_err != nil {
			returnError = resp_body_json_err
		}
		resp_body_map := resp_body_reader.(map[string]interface{})
		_ = update_local_resource_state(resp_body_map, logger, d)

		//for key, value := range resp_body_map {
		//	if key == "id" {
		//		s_id := strconv.FormatFloat(value.(float64), 'f', -1, 64)
		//		defer d.SetId(s_id)
		//	}else{
		//		defer d.Set(key,value)
		//	}
		//	//if key == "nics" {
		//	//	for key_map, value_map := range value {
		//	//		d.Set(key, read_value)
		//	//		read_value = nil
		//	//	}
		//	//}
		//}
	}

	return returnError
}

func resourceVMRead(d *schema.ResourceData, m interface{}) error {
	var returnError error
	var responseBody string
	returnError = nil
	var destREAD_URL strings.Builder
	var resp_body_reader interface{}
	client := &http.Client{}
	vmName := d.Get("name").(string)
	vmId := d.Id()
	logger := loggerCreate("resourceVM_Read_" + vmName + ".log")
	logger.Println("--------------- ", vmName, " ( id= ", vmId, ") READ -----------------")

	destREAD_URL.WriteString("https://next.cloud-datacenter.fr/api/clouddc/vm/")
	destREAD_URL.WriteString(vmId)
	destREAD_URL.WriteString("/")
	s_destREAD_URL := destREAD_URL.String()

	req, _ := http.NewRequest("GET", s_destREAD_URL, nil)

	req.Header.Add("authorization", "Token "+CONN_TOKEN)

	resp, read_req_err := client.Do(req)
	defer resp.Body.Close()

	bodyBytes, read_resp_body_read_err := ioutil.ReadAll(resp.Body)
	responseBody = string(bodyBytes)

	if read_req_err != nil {
		logger.Println("Read of ", vmName, " state reception error : ", read_req_err)
		returnError = errors.New("Read of " + vmName + " state reception error : " + read_req_err.Error())
	}

	if read_resp_body_read_err != nil {
		logger.Println("Read of ", vmName, " state response body read error ", read_resp_body_read_err)
		returnError = errors.New("Read of " + vmName + " state response body read error " + read_resp_body_read_err.Error())
	}

	if resp.StatusCode == RESOURCE_GET_HTTP_SUCCESS_CODE {
		resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
		if resp_body_json_err != nil {
			returnError = resp_body_json_err
		} else {
		resp_body_map := resp_body_reader.(map[string]interface{})
		returnError = update_local_resource_state(resp_body_map, logger, d)
		}
	} else if resp.StatusCode == RESOURCE_GET_HTTP_NOT_FOUND_CODE {
		logger.Println(vmName, " not found, The resource may have been deleted by an other Airdrum API client.")
		logger.Println("Airdrum api response : ", resp.Status, responseBody)
		logger.Println("Deletion of ", vmName, "in terraform resources list.")
		defer d.SetId("")
	} else {
		logger.Println("Unknow error : ")
		returnError = errors.New("Unknow error : " + resp.Status + responseBody)
	}

	return returnError
}

func resourceVMUpdate(d *schema.ResourceData, m interface{}) error {
	vmName := d.Get("name").(string)
	vmId := d.Id()
	var responseBody string
	var returnError error
	var destREAD_URL strings.Builder
	destREAD_URL.WriteString("https://next.cloud-datacenter.fr/api/clouddc/vm/")
	destREAD_URL.WriteString(vmId)
	destREAD_URL.WriteString("/")
	s_destREAD_URL := destREAD_URL.String()
	client := &http.Client{}
	logger := loggerCreate("resourceVM_Update_" + vmName + ".log")

	logger.Println("--------------- ", vmName, " ( id= ", vmId, ") UPDATE -----------------")

	vm_struct := VM{
		Name: d.Get("name").(string),
		Slug: d.Get("slug").(string),
		Vdc: d.Get("vdc").(string),
		Vdc_resource_disk: d.Get("vdc_resource_disk").(string),
		RAM: d.Get("ram").(string),
		CPU: d.Get("cpu").(string),
		//Template: vmInstance.Template,
		State: d.Get("state").(string),
		Disks: d.Get("disks").([]interface{}),
		OS: d.Get("os").(string),
		Nics: d.Get("nics").([]interface{}),
		Token: d.Get("token").(string),
		Platform_name: d.Get("platform_name").(string),
		Backup: d.Get("backup").(string),
		Disk_image: d.Get("disk_image").(string),
		Boot: d.Get("boot").(string),
		Backup_size: d.Get("backup_size").(string),
		Comment: d.Get("comment").(string),
		Outsourcing: d.Get("outsourcing").(string),
		Dynamic_field: d.Get("dynamic_field").(string),
	}

	vm_json, err_json := json.Marshal(vm_struct)
	logger.Println("vm_struct =", vm_struct)
	logger.Println("err_json =", err_json)
	logger.Println("vm_json =", vm_json)

	req, _ := http.NewRequest("PUT", s_destREAD_URL, bytes.NewBuffer(vm_json))

	req.Header.Add("authorization", "Token "+CONN_TOKEN)
	req.Header.Add("content-type", "application/json")

	logger.Println("Update of ", vmName, "request Header = ", req.Header)
	logger.Println("Update of ", vmName, "request body = ", req.Body)

	resp, create_err := client.Do(req)
	defer resp.Body.Close()
	bodyBytes, update_resp_body_read_err := ioutil.ReadAll(resp.Body)
	responseBody = string(bodyBytes)

	if create_err != nil {
		logger.Println("Update of ", vmName, " response reception error : ", create_err)
		returnError = errors.New("Update of " + vmName + " response reception error : " + create_err.Error())
	}

	if update_resp_body_read_err != nil {
		logger.Println("Update of ", vmName, " response body read error ", update_resp_body_read_err)
		returnError = errors.New("Update of " + vmName + " response body read error " + update_resp_body_read_err.Error())
	}

	logger.Println("Update of ", vmName, " response status = ", resp.Status)
	logger.Println("Update of ", vmName, " response body = ", responseBody)

	if resp.StatusCode != RESOURCE_UPDATE_HTTP_SUCCESS_CODE {
		logger.Println("Update of ", vmName, " resource failed : ", resp.Status, responseBody)
		returnError = errors.New(resp.Status + responseBody)
	}

	return returnError
}

func resourceVMDelete(d *schema.ResourceData, m interface{}) error {

	var responseBody string
	var returnError error
	returnError = nil
	client := &http.Client{}
	var destDELETE_URL strings.Builder
	vmName := d.Get("name").(string)
	vmId := d.Id()
	logger := loggerCreate("resourceVM_Delete_" + vmName + ".log")
	destDELETE_URL.WriteString("https://next.cloud-datacenter.fr/api/clouddc/vm/")
	destDELETE_URL.WriteString(vmId)
	destDELETE_URL.WriteString("/")
	s_destDELETE_URL := destDELETE_URL.String()

	logger.Println("--------------- ", vmName, " ( id= ", vmId, ") DELETE -----------------")

	req, _ := http.NewRequest("DELETE", s_destDELETE_URL, nil)

	req.Header.Add("authorization", "Token "+CONN_TOKEN)

	resp, delete_err := client.Do(req)
	defer resp.Body.Close()

	bodyBytes, delete_resp_body_read_err := ioutil.ReadAll(resp.Body)
	responseBody = string(bodyBytes)

	if delete_err != nil {
		logger.Println("Deletion of ", vmName, " response reception error : ", delete_err)
		returnError = errors.New("Deletion of " + vmName + " response reception error : " + delete_err.Error())
	}

	if delete_resp_body_read_err != nil {
		logger.Println("Deletion of ", vmName, " response body read error ", delete_resp_body_read_err)
		returnError = errors.New("Deletion of " + vmName + " response reception error : " + delete_err.Error())
	}

	if resp.StatusCode != RESOURCE_DELETE_HTTP_SUCCESS_CODE && responseBody != "{\"detail\":\"Destroying the VM now\"}" {
		logger.Println("Deletion of ", vmName, " resource failed : ", resp.Status, responseBody)
		returnError = errors.New(resp.Status + responseBody)
	} else {
		defer d.SetId("")
	}

	return returnError
}
