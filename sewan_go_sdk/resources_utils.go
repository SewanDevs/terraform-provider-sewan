package sewan_go_sdk

import (
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	DELETION_FIELD = "deletion"
)

type Dynamic_field_struct struct {
	Terraform_provisioned       bool          `json:"terraform_provisioned"`
	Creation_template           string        `json:"creation_template"`
	Disks_created_from_template []interface{} `json:"disks_created_from_template"`
}

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
	V_disk        string `json:"v_disk,omitempty"`
	Deletion      bool   `json:"deletion,omitempty"`
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
	Dynamic_field string        `json:"dynamic_field"`
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
	templatesTooler *TemplatesTooler,
	schemaTools *SchemaTooler,
	api *API) (VM, error) {

	var (
		vm                             VM
		get_templates_list_error       error                  = nil
		fetch_template_from_list_error error                  = nil
		instance_creation_error        error                  = nil
		template                       map[string]interface{} = nil
		template_name                  string                 = d.Get("template").(string)
		enterprise                     string                 = d.Get("enterprise").(string)
	)
	logger := LoggerCreate("vminstanceCreate" + d.Id() + ".log")

	if template_name != "" {
		vm = VM{}
		var templateList []interface{}

		templateList,
			get_templates_list_error = clientTooler.Client.GetTemplatesList(clientTooler,
			enterprise, api)
		if get_templates_list_error == nil {
			template,
				fetch_template_from_list_error = templatesTooler.TemplatesTools.FetchTemplateFromList(template_name,
				templateList)
			switch {
			case fetch_template_from_list_error != nil:
				instance_creation_error = fetch_template_from_list_error
			default:
				instance_creation_error = templatesTooler.TemplatesTools.UpdateSchemaFromTemplate(d,
					template,
					templatesTooler,
					schemaTools)
			}
		} else {
			instance_creation_error = get_templates_list_error
		}
	}

	logger.Println("instance_creation_error = ", instance_creation_error)
	if instance_creation_error == nil {
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
			Outsourcing:   d.Get("outsourcing").(string),
			Dynamic_field: d.Get("dynamic_field").(string),
		}
		if d.Id() == "" {
			vm.Template = d.Get("template").(string)
			dynamic_field_struct := Dynamic_field_struct{
				Terraform_provisioned:       true,
				Creation_template:           vm.Template,
				Disks_created_from_template: nil,
			}
			if template != nil {
				dynamic_field_struct.Disks_created_from_template = template["disks"].([]interface{})
			}
			dynamic_field_json, _ := json.Marshal(dynamic_field_struct)
			vm.Dynamic_field = string(dynamic_field_json)

		} else {
			vm.Template = ""
			var update_disks_slices []interface{} = []interface{}{}
			for _, disk := range d.Get(DISKS_PARAM).([]interface{}) {
				if disk.(map[string]interface{})[DELETION_FIELD] != true {
					update_disks_slices = append(update_disks_slices, disk)
				}
			}
			vm.Disks = update_disks_slices
		}
	}
	logger.Println("vm = ", vm)
	logger.Println("instance_creation_error = ", instance_creation_error)
	return vm, instance_creation_error
}

func (apier AirDrumResources_Apier) ResourceInstanceCreate(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	schemaTools *SchemaTooler,
	resourceType string,
	api *API) (error, interface{}) {

	var (
		resourceInstance interface{} = nil
		instanceError    error       = nil
	)

	switch resourceType {
	case "vdc":
		resourceInstance, instanceError = vdcInstanceCreate(d,
			clientTooler,
			api)
	case "vm":
		resourceInstance, instanceError = vmInstanceCreate(d,
			clientTooler,
			templatesTooler,
			schemaTools,
			api)
	default:
		instanceError = apier.ValidateResourceType(resourceType)
	}

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
