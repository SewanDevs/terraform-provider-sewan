package sewansdk

import (
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
	"strconv"
	"strings"
)

// ResourceTooler contains implementation of Resourceer interface
type ResourceTooler struct {
	Resource Resourceer
}

// Resourceer interface is responsible of operations on terraform resource
type Resourceer interface {
	getResourceCreationURL(api *API,
		resourceType string) string
	getResourceURL(api *API,
		resourceType string,
		id string) string
	validateResourceType(resourceType string) error
	validateStatus(api *API,
		resourceType string,
		client ClientTooler) error
	resourceInstanceCreate(d *schema.ResourceData,
		clientTooler *ClientTooler,
		templatesTooler *TemplatesTooler,
		resourceType string,
		api *API) (interface{}, error)
}

// ResourceResourceer implements Resourceer interface
type ResourceResourceer struct{}

type dynamicFieldStruct struct {
	TerraformProvisioned    bool          `json:"terraform_provisioned"`
	CreationTemplate        string        `json:"creationTemplate"`
	TemplateDisksOnCreation []interface{} `json:"TemplateDisksOnCreation"`
}

type vdcResourceStruct struct {
	Resource string `json:"vdc_resources"`
	Used     int    `json:"used"`
	Total    int    `json:"total"`
	Slug     string `json:"slug"`
}

type vdcStruct struct {
	Name         string        `json:"name"`
	Enterprise   string        `json:"enterprise"`
	Datacenter   string        `json:"datacenter"`
	VdcResources []interface{} `json:"vdc_resources"`
	Slug         string        `json:"slug"`
	DynamicField string        `json:"dynamic_field"`
}

// This struct is not used in the code, however it remains here for dev doc purpose
//type vmDisk struct {
//	Name         string `json:"name"`
//	Size         int    `json:"size"`
//	StorageClass string `json:"storage_class"`
//	Slug         string `json:"slug"`
//	VDisk       string `json:"v_disk,omitempty"`
//}

// This struct is not used in the code, however it remains here for dev doc purpose
//type vmNic struct {
//	Vlan        string `json:"vlan"`
//	MacAddress string `json:"mac_address"`
//	Connected   bool   `json:"connected"`
//}

type vmStruct struct {
	Name         string        `json:"name"`
	Enterprise   string        `json:"enterprise"`
	Template     string        `json:"template,omitempty"`
	State        string        `json:"state"`
	OS           string        `json:"os,omitempty"`
	RAM          int           `json:"ram"`
	CPU          int           `json:"cpu"`
	Disks        []interface{} `json:"disks,omitempty"`
	Nics         []interface{} `json:"nics,omitempty"`
	Vdc          string        `json:"vdc"`
	Boot         string        `json:"boot"`
	StorageClass string        `json:"storage_class"`
	Slug         string        `json:"slug"`
	Token        string        `json:"token"`
	Backup       string        `json:"backup"`
	DiskImage    string        `json:"disk_image"`
	PlatformName string        `json:"platform_name"`
	BackupSize   int           `json:"backup_size"`
	Comment      string        `json:"comment,omitempty"`
	DynamicField string        `json:"dynamic_field"`
	Outsourcing  string        `json:"outsourcing"`
}

func vdcInstanceCreate(d *schema.ResourceData) (vdcStruct, error) {
	var (
		resourceName strings.Builder
	)
	vdc := vdcStruct{
		Name:         d.Get(NameField).(string),
		Enterprise:   d.Get(EnterpriseField).(string),
		Datacenter:   d.Get(DatacenterField).(string),
		VdcResources: d.Get(VdcResourceField).([]interface{}),
		Slug:         d.Get(SlugField).(string),
		DynamicField: d.Get(DynamicField).(string),
	}
	for index, resource := range vdc.VdcResources {
		resourceName.Reset()
		resourceName.WriteString(vdc.Enterprise)
		resourceName.WriteString(monoField)
		resourceName.WriteString(resource.(map[string]interface{})[ResourceField].(string))
		resource.(map[string]interface{})[ResourceField] = resourceName.String()
		vdc.VdcResources[index] = resource
	}
	return vdc, nil
}

func getTemplateAndUpdateSchema(templateName string,
	d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	api *API) (map[string]interface{}, error) {
	templateList, err1 := clientTooler.Client.getTemplatesList(clientTooler,
		d.Get(EnterpriseField).(string),
		api)
	if err1 != nil {
		return map[string]interface{}{}, err1
	}
	template, err2 := templatesTooler.TemplatesTools.FetchTemplateFromList(templateName,
		templateList)
	if err2 != nil {
		return map[string]interface{}{}, err2
	}
	err3 := templatesTooler.TemplatesTools.validateTemplate(template)
	if err3 != nil {
		return map[string]interface{}{}, err3
	}
	err4 := templatesTooler.TemplatesTools.updateSchemaFromTemplateOnResourceCreation(d,
		template)
	if err4 != nil {
		return map[string]interface{}{}, err4
	}
	return template, nil
}

func vmInstanceCreate(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	api *API) (vmStruct, error) {
	var (
		templateError error
		template      map[string]interface{}
		templateName  = d.Get(TemplateField).(string)
		vmName        strings.Builder
	)
	vmName.WriteString(d.Get(NameField).(string))
	if templateName != "" && d.Id() == "" {
		instanceNumber := d.Get(InstanceNumberField).(int)
		vmName.WriteString(resourceNameCountSeparator)
		vmName.WriteString(strconv.Itoa(instanceNumber))
		template,
			templateError = getTemplateAndUpdateSchema(templateName,
			d,
			clientTooler,
			templatesTooler,
			api)
	}
	if templateError != nil {
		return vmStruct{}, templateError
	}
	vm := vmStruct{
		Name:         vmName.String(),
		Enterprise:   d.Get(EnterpriseField).(string),
		State:        d.Get(StateField).(string),
		OS:           d.Get(OsField).(string),
		RAM:          d.Get(RAMField).(int),
		CPU:          d.Get(CPUField).(int),
		Disks:        d.Get(DisksField).([]interface{}),
		Nics:         d.Get(NicsField).([]interface{}),
		Vdc:          d.Get(VdcField).(string),
		Boot:         d.Get(BootField).(string),
		StorageClass: d.Get(StorageClassField).(string),
		Slug:         d.Get(SlugField).(string),
		Token:        d.Get(TokenField).(string),
		Backup:       d.Get(BackupField).(string),
		DiskImage:    d.Get(DiskImageField).(string),
		PlatformName: d.Get(PlatformNameField).(string),
		BackupSize:   d.Get(BackupSizeField).(int),
		DynamicField: d.Get(DynamicField).(string),
	}
	if d.Id() == "" {
		dynamicFieldStruct := dynamicFieldStruct{
			TerraformProvisioned:    true,
			CreationTemplate:        d.Get(TemplateField).(string),
			TemplateDisksOnCreation: nil,
		}
		if template != nil {
			dynamicFieldStruct.TemplateDisksOnCreation = template[DisksField].([]interface{})
			_, err := templatesTooler.TemplatesTools.createVMTemplateOverrideConfig(d,
				template)
			if err != nil {
				return vmStruct{}, err
			}
			vm.Template = d.Get(TemplateField).(string)
		}
		dynamicFieldJSON, err2 := json.Marshal(dynamicFieldStruct)
		if err2 != nil {
			return vmStruct{}, err2
		}
		vm.DynamicField = string(dynamicFieldJSON)
	}
	return vm, nil
}

// resourceInstanceCreate creates a resource structure initialized with
// fields values got from schema.
// Accepted resource types : "vm", "vdc"
func (resource ResourceResourceer) resourceInstanceCreate(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	resourceType string,
	api *API) (interface{}, error) {
	switch resourceType {
	case VdcResourceType:
		return vdcInstanceCreate(d)
	case VMResourceType:
		return vmInstanceCreate(d,
			clientTooler,
			templatesTooler,
			api)
	default:
		return nil, resource.validateResourceType(resourceType)
	}
}

//validateResourceType validates if resource type is in
// list of accepted resource types : "vm", "vdc"
func (resource ResourceResourceer) validateResourceType(resourceType string) error {
	switch resourceType {
	case VdcResourceType:
		return nil
	case VMResourceType:
		return nil
	default:
		return errWrongResourceTypeBuilder(resourceType)
	}
}

// getResourceCreationURL returns valid urls for resource creation :
// * https://cloud-datacenter.fr/api/clouddc/vm/
// * https://cloud-datacenter.fr/api/clouddc/vdc/
// * etc.
func (resource ResourceResourceer) getResourceCreationURL(api *API,
	resourceType string) string {
	var resourceURL strings.Builder
	resourceURL.WriteString(api.URL)
	resourceURL.WriteString(resourceType)
	resourceURL.WriteString("/")
	return resourceURL.String()
}

// getResourceURL returns valid resources urls :
// * https://cloud-datacenter.fr/api/clouddc/vm/<a resource id number>/
// * https://cloud-datacenter.fr/api/clouddc/vdc/<a resource id number>/
// * etc.
func (resource ResourceResourceer) getResourceURL(api *API,
	resourceType string,
	resourceID string) string {
	var resourceURL strings.Builder
	sCreateURL := resource.getResourceCreationURL(api, resourceType)
	resourceURL.WriteString(sCreateURL)
	resourceURL.WriteString(resourceID)
	resourceURL.WriteString("/")
	return resourceURL.String()
}

// validateStatus validates api status by sending a test GET request and validating response
func (resource ResourceResourceer) validateStatus(api *API,
	resourceType string,
	clientTooler ClientTooler) error {
	req, _ := http.NewRequest("GET",
		resource.getResourceCreationURL(api, resourceType),
		nil)
	req.Header.Add(httpAuthorization, httpTokenHeader+api.Token)
	resp, err1 := clientTooler.Client.do(api, req)
	if err1 != nil {
		return err1
	}
	_, err2 := clientTooler.Client.handleResponse(resp,
		http.StatusOK,
		httpJSONContentType)
	return err2
}
