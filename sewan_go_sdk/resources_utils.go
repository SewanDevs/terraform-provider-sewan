package sewan_go_sdk

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type VDC_resource struct {
	Resource string `json:"resource"`
	Used     int    `json:"used"`
	Total    int    `json:"total"`
	Slug     string `json:"slug"`
}

type VDC struct {
	Name          string        `json:"name"`
	Enterprise    string        `json:"enterprise"`
	Datacenter    string        `json:"datacenter"`
	Vdc_resources []interface{} `json:"vdc_resources"`
	Slug          string        `json:"slug"`
	Dynamic_field string        `json:"dynamic_field"`
}

type VM_DISK struct {
	Name          string `json:"name"`
	Size          int    `json:"size"`
	Storage_class string `json:"storage_class"`
	Slug          string `json:"slug"`
	V_disk        string `json:"v_disk"`
}

type VM_NIC struct {
	Vlan        string `json:"vlan"`
	Mac_address string `json:"mac_address"`
	Connected   bool   `json:"connected"`
}

type VM struct {
	Name          string        `json:"name"`
	Enterprise    string        `json:"enterprise"`
	Template      string        `json:"template,omitempty"`
	State         string        `json:"state"`
	OS            string        `json:"os,omitempty"`
	RAM           int           `json:"ram"`
	CPU           int           `json:"cpu"`
	Disks         []interface{} `json:"disks,omitempty"`
	Nics          []interface{} `json:"nics,omitempty"`
	Vdc           string        `json:"vdc"`
	Boot          string        `json:"boot"`
	Storage_class string        `json:"storage_class"`
	Slug          string        `json:"slug"`
	Token         string        `json:"token"`
	Backup        string        `json:"backup"`
	Disk_image    string        `json:"disk_image"`
	Platform_name string        `json:"platform_name"`
	Backup_size   int           `json:"backup_size"`
	Comment       string        `json:"comment",omitempty`
	Outsourcing   string        `json:"outsourcing,omitempty"`
	Dynamic_field string        `json:"dynamic_field,omitempty"`
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
		vm                    VM
		getTemplateError      error  = nil
		template_list_Error   error  = nil
		instanceCreationError error  = nil
		template_exists       bool   = false
		template              string = d.Get("template").(string)
		enterprise            string = d.Get("enterprise").(string)
	)
	// @TODO : log to delete
	logger := loggerCreate("vmInstanceCreate" + d.Get("name").(string) + ".log")
	logger.Println("template =", template)

	if template != "" {
		vm = VM{}
		var templateList []interface{}

		templateList,
			getTemplateError = clientTooler.Client.GetTemplatesList(clientTooler,
			enterprise, api)
		// @TODO : log to delete
		logger.Println("templateList =", templateList)
		logger.Println("getTemplateError =", getTemplateError)

		if getTemplateError == nil {
			template_exists, template_list_Error = Handle_templates_list(d, templateList)
			switch {
			case template_list_Error != nil:
				instanceCreationError = template_list_Error
			}
			if template_exists == false {
				instanceCreationError = errors.New("Unavailable template : " +
					template)
			}
		} else {
			instanceCreationError = getTemplateError
		}
	} else {
		// @TODO : log to delete
		logger.Println("template = nil")
	}

	logger.Println("instanceCreationError =", instanceCreationError)

	if instanceCreationError == nil {

		// @TODO : log to delete
		logger.Println("vm = VM{} set")

		vm = VM{
			Name:          d.Get("name").(string),
			Enterprise:    d.Get("enterprise").(string),
			State:         d.Get("state").(string),
			OS:            d.Get("os").(string),
			RAM:           d.Get("ram").(int),
			CPU:           d.Get("cpu").(int),
			Disks:         d.Get("disks").([]interface{}),
			Nics:          d.Get("nics").([]interface{}),
			Vdc:           d.Get("vdc").(string),
			Boot:          d.Get("boot").(string),
			Storage_class: d.Get("storage_class").(string),
			Slug:          d.Get("slug").(string),
			Token:         d.Get("token").(string),
			Backup:        d.Get("backup").(string),
			Disk_image:    d.Get("disk_image").(string),
			Platform_name: d.Get("platform_name").(string),
			Backup_size:   d.Get("backup_size").(int),
			Comment:       d.Get("comment").(string),
			Outsourcing:   d.Get("outsourcing").(string),
			Dynamic_field: d.Get("dynamic_field").(string),
		}

		if d.Id() == "" {
			vm.Template = d.Get("template").(string)
		} else {
			vm.Template = ""
		}

	}
	logger.Println("vm = ", vm)

	return vm, instanceCreationError
}

func (apier AirDrumResources_Apier) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceType string,
	api *API) (error, interface{}) {

	var (
		resourceInstance interface{} = nil
		instanceError    error       = nil
	)

	logger := loggerCreate("ResourceInstanceCreate.log")
	logger.Println("resourceType = ", resourceType)
	switch resourceType {
	case "vdc":
		logger.Println("vdc case")
		resourceInstance, instanceError = vdcInstanceCreate(d, clientTooler, api)
	case "vm":
		logger.Println("vm case")
		resourceInstance, instanceError = vmInstanceCreate(d, clientTooler, api)
	default:
		logger.Println("default case")
		instanceError = apier.ValidateResourceType(resourceType)
	}

	logger.Println("instanceError, resourceInstance = ", instanceError, resourceInstance)
	return instanceError, resourceInstance
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

func Delete_terraform_resource(d *schema.ResourceData) {
	d.SetId("")
}

func Update_local_resource_state(resource_state map[string]interface{},
	d *schema.ResourceData) error {

	var (
		updateError error = nil
		read_value  interface{}
	)
	logger := LoggerCreate("update_local_resource_state_" +
		d.Get("name").(string) + ".log")
	for key, value := range resource_state {
		read_value, updateError = read_element(key, value, logger)
		logger.Println("Set \"", key, "\" to \"", read_value, "\"")
		if key == "id" {
			var s_id string = ""
			switch {
			case reflect.TypeOf(value).Kind() == reflect.Float64:
				s_id = strconv.FormatFloat(value.(float64), 'f', -1, 64)
			case reflect.TypeOf(value).Kind() == reflect.Int:
				s_id = strconv.Itoa(value.(int))
			case reflect.TypeOf(value).Kind() == reflect.String:
				s_id = value.(string)
			default:
				updateError = errors.New("Format of " + key + "(" +
					reflect.TypeOf(value).Kind().String() + ") not handled.")
			}
			d.SetId(s_id)
		} else {
			updateError = d.Set(key, read_value)
		}
		read_value = nil
	}
	return updateError
}

func read_element(key interface{}, value interface{},
	logger *log.Logger) (interface{}, error) {

	var (
		readError  error = nil
		read_value interface{}
	)
	switch value_type := value.(type) {
	case string:
		read_value = value.(string)
	case bool:
		read_value = value.(bool)
	case float64:
		read_value = int(value.(float64))
	case int:
		read_value = value.(int)
	case map[string]interface{}:
		var read_map_value map[string]interface{}
		read_map_value = make(map[string]interface{})
		var map_item interface{}
		for map_key, map_value := range value_type {
			map_item, readError = read_element(map_key, map_value, logger)
			read_map_value[map_key] = map_item
		}
		read_value = read_map_value
	case []interface{}:
		var read_list_value []interface{}
		var list_item interface{}
		for list_key, list_value := range value_type {
			list_item, readError = read_element(list_key, list_value, logger)
			read_list_value = append(read_list_value, list_item)
		}
		read_value = read_list_value
	default:
		if value == nil {
			read_value = nil
		} else {
			readError = errors.New("Format " +
				reflect.TypeOf(value_type).Kind().String() + " not handled.")
		}
	}
	return read_value, readError
}
