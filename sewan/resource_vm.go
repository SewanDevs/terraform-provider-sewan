package sewan

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const RESOURCE_CREATE_HTTP_SUCCESS_CODE = 201
const RESOURCE_DELETE_HTTP_SUCCESS_CODE = 200
const RESOURCE_GET_HTTP_SUCCESS_CODE = 200

// NB : The following 2 vars will be deleted when the provider config will be handled
const DEST_REQ_URL = "https://next.cloud-datacenter.fr/api/clouddc/vm/"
const CONN_TOKEN = "17f061821bac9e12f9a2ded3928e624ae7c28448"

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
			//"slug": &schema.Schema{
			//	Type:     schema.TypeString,
			//	Optional: true,
			//},
			"vdc": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdc_resource_disk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ram": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			//"os": &schema.Schema{
			//	Type:     schema.TypeString,
			//	Optional: true,
			//},
			//"nic": &schema.Schema{
			//	Type:     schema.TypeString,
			//	Optional: true,
			//},
			//"token": &schema.Schema{
			//	Type:     schema.TypeString,
			//	Optional: true,
			//},
			//"plateform_name": &schema.Schema{
			//	Type:     schema.TypeString,
			//	Optional: true,
			//},
			"backup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disk_image": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"boot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			//"backup_size": &schema.Schema{
			//	Type:     schema.TypeString,
			//	Optional: true,
			//},
			//"comment": &schema.Schema{
			//	Type:     schema.TypeString,
			//	Optional: true,
			//},
			//"outsourcing": &schema.Schema{
			//	Type:     schema.TypeString,
			//	Optional: true,
			//},
			//"dynamic_field": &schema.Schema{
			//	Type:     schema.TypeString,
			//	Optional: true,
			//},
		},
	}
}

type VM struct {
	Name              string
	Slug              string
	Vdc               string
	Vdc_resource_disk string
	RAM               string
	CPU               string
	Template          string
	State             string
	Disks             string
	//Os                string
	//Nics              string
	//Token             string
	//Plateform_name    string
	Backup            string
	Disk_image        string
	Boot              string
	//Backup_size       string
	//Comment           string
	//Outsourcing       string
	//Dynamic_field     string
}

func vmInstanceCreate(d *schema.ResourceData) VM {
	return VM{
		Name:              d.Get("name").(string),

		//Slug:"",
		//State:"",
		//Disks:"",
		//Token:"",
		//Plateform_name:"",
		//Comment:"",
		//Outsourcing:"",
		//Dynamic_field:"",

		Vdc:               d.Get("vdc").(string),
		CPU:               d.Get("cpu").(string),
		RAM:               d.Get("ram").(string),
		Boot:              d.Get("boot").(string),
		//Nics:              d.Get("nics").(string),
		Template:          d.Get("template").(string),
		//Os:                d.Get("os").(string),
		Backup:            d.Get("backup").(string),
		//Backup_size:       d.Get("backup_size").(string),
		Vdc_resource_disk: d.Get("vdc_resource_disk").(string),
		Disk_image:        d.Get("disk_image").(string),
	}
}

func resourceVMCreate(d *schema.ResourceData, m interface{}) error {
	vmInstance := vmInstanceCreate(d)
	var requestBody bytes.Buffer
	var responseBody string
	var resp_body_reader interface{}
	logger := loggerCreate("resourceVMCreate_" + vmInstance.Vdc + "_" + vmInstance.Name + ".log")
	var returnError error
	returnError = nil
	client := &http.Client{}

	logger.Println("--------------- ",vmInstance.Name," CREATION -----------------")

	requestBody.WriteString("{\"name\":\"" + vmInstance.Name + "\",")
	requestBody.WriteString("\"vdc\":\"" + vmInstance.Vdc + "\",")
	requestBody.WriteString("\"ram\":" + vmInstance.RAM + ",")
	requestBody.WriteString("\"cpu\":" + vmInstance.RAM + ",")
	requestBody.WriteString("\"disk_image\":\"" + vmInstance.Disk_image + "\",")
	requestBody.WriteString("\"boot\":\"" + vmInstance.Boot + "\",")
	//requestBody.WriteString("\"nics\":" + vmInstance.Nic + ",")
	requestBody.WriteString("\"template\":\"" + vmInstance.Template + "\",")
	requestBody.WriteString("\"vdc_resource_disk\":\"" + vmInstance.Vdc_resource_disk + "\",")
	requestBody.WriteString("\"backup\":\"" + vmInstance.Backup + "\"}")

	req, _ := http.NewRequest("POST", DEST_REQ_URL, &requestBody)

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
	}

	logger.Println("Creation of ", vmInstance.Name, " response status = ", resp.Status)
	logger.Println("Creation of ", vmInstance.Name, " response body = ", responseBody)

	if resp.StatusCode != RESOURCE_CREATE_HTTP_SUCCESS_CODE {
		logger.Println("Creation of ", vmInstance.Name, " resource failed : ", resp.Status, responseBody)
		returnError = errors.New(resp.Status + responseBody)
	} else {
		resp_body_json_err := json.Unmarshal(bodyBytes, &resp_body_reader)
		if resp_body_json_err != nil{
			returnError = resp_body_json_err
		}
		resp_body_map := resp_body_reader.(map[string]interface{})
		for key, value := range resp_body_map {
			switch value_type := value.(type) {
			case string:
				logger.Println(key, "(type string) = ", value_type)
			case float64:
				logger.Println(key, "(type float64) = ", value_type)
			case []interface{}:
				logger.Println(key, "(type array) = ")
				for i, u := range value_type {
					logger.Println(i, u)
				}
			default:
				logger.Println(key, "(unknow type) : ", value_type)
			}
			if key == "id" {
				s_id := strconv.FormatFloat(value.(float64), 'f', -1, 64)
				defer d.SetId(s_id)
			}
		}
	}

	return returnError
}

func resourceVMRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceVMUpdate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceVMDelete(d *schema.ResourceData, m interface{}) error {

	var responseBody string
	var returnError error
	returnError = nil
	client := &http.Client{}
	var destDELETE_URL strings.Builder
	vmName := d.Get("name").(string)
	vmId := d.Id()
	logger := loggerCreate("resourceVMDelete_" + vmName + ".log")
	destDELETE_URL.WriteString("https://next.cloud-datacenter.fr/api/clouddc/vm/")
	destDELETE_URL.WriteString(vmId)
	destDELETE_URL.WriteString("/")
	s_destDELETE_URL := destDELETE_URL.String()

	logger.Println("--------------- ",vmName," DELETION, id = ",vmId," -----------------")

	req, _ := http.NewRequest("DELETE", s_destDELETE_URL,nil)

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
	}

	if resp.StatusCode != RESOURCE_DELETE_HTTP_SUCCESS_CODE && responseBody!="{\"detail\":\"Destroying the VM now\"}"{
		logger.Println("Deletion of ", vmName, " resource failed : ", resp.Status, responseBody)
		returnError = errors.New(resp.Status + responseBody)
	} else {
		defer d.SetId("")
	}

	return returnError
}
