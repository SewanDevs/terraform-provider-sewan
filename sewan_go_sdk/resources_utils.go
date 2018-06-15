package sewan_go_sdk

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
	"strings"
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
	Template          string        `json:"template"`
	State             string        `json:"state"`
	OS                string        `json:"os"`
	RAM               int           `json:"ram"`
	CPU               int           `json:"cpu"`
	Disks             []interface{} `json:"disks"`
	Nics              []interface{} `json:"nics"`
	Vdc               string        `json:"vdc"`
	Boot              string        `json:"boot"`
	Vdc_resource_disk string        `json:"vdc_resource_disk"`
	Slug              string        `json:"slug"`
	Token             string        `json:"token"`
	Backup            string        `json:"backup"`
	Disk_image        string        `json:"disk_image"`
	Platform_name     string        `json:"platform_name"`
	Backup_size       string        `json:"backup_size"`
	Comment           string        `json:"comment"`
	Outsourcing       string        `json:"outsourcing"`
	Dynamic_field     string        `json:"dynamic_field"`
}

func vdcInstanceCreate(d *schema.ResourceData,
	clientTooler *ClientTooler,
	api *API) (VDC, error) {
	return VDC{
		Name:          d.Get("name").(string),
		Enterprise:    d.Get("enterprise").(string),
		Datacenter:    d.Get("datacenter").(string),
		Vdc_resources: d.Get("vdc_resources").([]interface{}),
		Slug:          d.Get("slug").(string),
		Dynamic_field: d.Get("dynamic_field").(string),
	}, nil
}

func vmInstanceCreate(d *schema.ResourceData,
	clientTooler *ClientTooler,
	api *API) (VM, error) {

	var (
		vm               VM
		getTemplateError error
		enterprise_slug  string
	)
	getTemplateError = nil

	logger := loggerCreate("vmInstanceCreate.log")

	if d.Get("template") != nil {
		var templateList []map[string]interface{}

		//1 get template list
		// i : get enterprise from VDC field :
		//	enterprise_slug = d.Get("depends_on").enterprise
		// ii : get enterprise from a new vm field :
		//	enterprise_slug = d.Get("enterprise")
		enterprise_slug = "sewan-rd-cloud-beta"
		//clientToolerB := ClientTooler{
		//	Client: HttpClienter{},
		//}
		templateList,
		getTemplateError = clientTooler.Client.GetTemplatesList(clientTooler,
			enterprise_slug)
		logger.Println("templateList =", templateList)

		//2 validate option template is in the list

		//3 Create a VM instance with template paramaters
		//		for nics and disk, add paramater disks and nics to template's one
		//		for other parameter, replace template param with .tf file one's except for OS

		vm = VM{}

	} else {
		vm = VM{
			Name:              d.Get("name").(string),
			Template:          d.Get("template").(string),
			State:             d.Get("state").(string),
			OS:                d.Get("os").(string),
			RAM:               d.Get("ram").(int),
			CPU:               d.Get("cpu").(int),
			Disks:             d.Get("disks").([]interface{}),
			Nics:              d.Get("nics").([]interface{}),
			Vdc:               d.Get("vdc").(string),
			Boot:              d.Get("boot").(string),
			Vdc_resource_disk: d.Get("vdc_resource_disk").(string),
			Slug:              d.Get("slug").(string),
			Token:             d.Get("token").(string),
			Backup:            d.Get("backup").(string),
			Disk_image:        d.Get("disk_image").(string),
			Platform_name:     d.Get("platform_name").(string),
			Backup_size:       d.Get("backup_size").(string),
			Comment:           d.Get("comment").(string),
			Outsourcing:       d.Get("outsourcing").(string),
			Dynamic_field:     d.Get("dynamic_field").(string),
		}
	}
	return vm, getTemplateError
}

func (apier AirDrumResources_Apier) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceType string,
	api *API) (error, interface{}, string) {

	var (
		resourceInstance interface{}
		instanceName     string
		instanceError    error
	)

	switch resourceType {
	case "vdc":
		resourceInstance, instanceError = vdcInstanceCreate(d, clientTooler, api)
		instanceName = d.Get("name").(string)
	case "vm":
		resourceInstance, instanceError = vmInstanceCreate(d, clientTooler, api)
		instanceName = d.Get("name").(string)
	default:
		resourceInstance = nil
		instanceName = ""
		instanceError = apier.ValidateResourceType(resourceType)
	}

	return instanceError, resourceInstance, instanceName
}

func (apier AirDrumResources_Apier) ValidateResourceType(resourceType string) error {
	var err error

	switch resourceType {
	case "vdc":
		err = nil
	case "vm":
		err = nil
	default:
		err = errors.New("Resource of type \"" + resourceType + "\" not supported," +
			"list of accepted resource types :\n\r" +
			"- \"vdc\"\n\r" +
			"- \"vm\"")
	}

	return err
}

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
