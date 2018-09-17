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
	Iso          string        `json:"disk_image"`
	PlatformName string        `json:"platform_name"`
	BackupSize   int           `json:"backup_size"`
	Comment      string        `json:"comment,omitempty"`
	DynamicField string        `json:"dynamic_field"`
	Outsourcing  string        `json:"outsourcing"`
}

// getDataCenterCos returns the cos field of a data center,
// the cos field represents the resource type of DatCenter :
// * mono or HA (high availability)
func getDataCenterCos(dataCenter string, api *API) string {
	for _, listDataCenter := range api.Meta.DataCenterList {
		listDataCenterSlug := listDataCenter.(map[string]interface{})[SlugField].(string)
		if dataCenter == listDataCenterSlug {
			return listDataCenter.(map[string]interface{})[ResourceCosField].(string)
		}
	}
	return ""
}

// validateDatacenter validates datacenter is in available dataCenter list in api.
func validateDatacenter(dataCenter string, api *API) error {
	var (
		isInSlice     bool
		sliceElements strings.Builder
	)
	for _, sliceElem := range api.Meta.DataCenterList {
		sliceElemSlug := sliceElem.(map[string]interface{})[SlugField].(string)
		sliceElements.WriteString(" \"")
		sliceElements.WriteString(sliceElemSlug)
		sliceElements.WriteString("\"")
		if dataCenter == sliceElemSlug {
			isInSlice = true
		}
	}
	if isInSlice {
		return nil
	}
	return errNotInList(dataCenter, sliceElements.String())
}

// validateVdcResources validates VDC resources exists in clouddc environment
// resources list.
func validateVdcResources(d *schema.ResourceData,
	api *API, cos string) error {
	for _, resource := range d.Get(VdcResourceField).([]interface{}) {
		var (
			resourceExists        bool
			availableResourceList strings.Builder
		)
		for _, apiResource := range api.Meta.EnterpriseResourceList {
			isRightCos := (apiResource.(map[string]interface{})[ResourceCosField] == cos)
			isResourceExistingInMeta := (apiResource.(map[string]interface{})[NameField] == resource.(map[string]interface{})[ResourceField])
			if isRightCos {
				availableResourceList.WriteString(" \"")
				availableResourceList.WriteString(apiResource.(map[string]interface{})[NameField].(string))
				availableResourceList.WriteString("\"")
				if isResourceExistingInMeta {
					resourceExists = true
				}
			}
		}
		if !resourceExists {
			return errResourceNotExist(resource.(map[string]interface{})[ResourceField].(string),
				availableResourceList.String())
		}
	}
	return nil
}

// validateTemplate validates template is in available template list in api.
func validateTemplate(d *schema.ResourceData,
	api *API) error {
	var (
		isInSlice     bool
		sliceElements strings.Builder
		template      = d.Get(TemplateField).(string)
	)
	if template == "" {
		return nil
	}
	for _, sliceElem := range api.Meta.TemplateList {
		sliceElemSlug := sliceElem.(map[string]interface{})[SlugField].(string)
		sliceElements.WriteString(" \"")
		sliceElements.WriteString(sliceElemSlug)
		sliceElements.WriteString("\"")
		if template == sliceElemSlug {
			isInSlice = true
		}
	}
	if isInSlice {
		return nil
	}
	return errNotInList(template, sliceElements.String())
}

// validateVMsVDC validates the vm's vdc is in available vdc list in api.
func validateVMsVDC(d *schema.ResourceData,
	api *API) error {
	var (
		isInSlice     bool
		sliceElements strings.Builder
		vdc           = d.Get(VdcField).(string)
	)
	for _, sliceElem := range api.Meta.EnterpriseVdcList {
		sliceElemSlug := sliceElem.(map[string]interface{})[SlugField].(string)
		sliceElements.WriteString(" \"")
		sliceElements.WriteString(sliceElemSlug)
		sliceElements.WriteString("\"")
		if vdc == sliceElemSlug {
			isInSlice = true
		}
	}
	if isInSlice {
		return nil
	}
	return errNotInList(vdc, sliceElements.String())
}

// validateNics validates the vm's nics are in available vlan list in api.
func validateNics(d *schema.ResourceData,
	api *API) error {
	var (
		isInSlice     bool
		sliceElements strings.Builder
		nicList       = d.Get(NicsField).([]interface{})
	)
	for _, nic := range nicList {
		isInSlice = false
		nicVlan := nic.(map[string]interface{})[VlanNameField].(string)
		for _, sliceElem := range api.Meta.VlanList {
			sliceElemSlug := sliceElem.(map[string]interface{})[SlugField].(string)
			sliceElements.WriteString(" \"")
			sliceElements.WriteString(sliceElemSlug)
			sliceElements.WriteString("\"")
			if nicVlan == sliceElemSlug {
				isInSlice = true
			}
		}
		if !isInSlice {
			return errNotInList(nicVlan, sliceElements.String())
		}
	}
	return nil
}

// validateResourceFieldsValue validates all resources values match on of
// of the available value in clouddc environment resources list
func validateResourceFieldsValue(d *schema.ResourceData,
	api *API,
	resourceType string) error {
	switch resourceType {
	case VdcResourceType:
		dataCenter := d.Get(DataCenterField).(string)
		err1 := validateDatacenter(dataCenter, api)
		if err1 != nil {
			return err1
		}
		err2 := validateVdcResources(d, api, getDataCenterCos(dataCenter, api))
		if err2 != nil {
			return err2
		}
	case VMResourceType:
		err3 := validateTemplate(d, api)
		if err3 != nil {
			return err3
		}
		err4 := validateNics(d, api)
		if err4 != nil {
			return err4
		}
		err5 := validateVMsVDC(d, api)
		if err5 != nil {
			return err5
		}
	}
	return nil
}

// resourceInstanceCreate creates a resource structure initialized with
// fields values got from schema.
// Accepted resource types : "vm", "vdc"
func (resource ResourceResourceer) resourceInstanceCreate(d *schema.ResourceData,
	templatesTooler *TemplatesTooler,
	resourceType string,
	api *API) (interface{}, error) {
	switch resourceType {
	case VdcResourceType:
		err1 := validateResourceFieldsValue(d, api, VdcResourceType)
		if err1 != nil {
			return vdcStruct{}, err1
		}
		return vdcInstanceCreate(d, api)
	case VMResourceType:
		err2 := validateResourceFieldsValue(d, api, VMResourceType)
		if err2 != nil {
			return vmStruct{}, err2
		}
		return vmInstanceCreate(d,
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

func vdcInstanceCreate(d *schema.ResourceData, api *API) (vdcStruct, error) {
	vdc := vdcStruct{
		Name:         d.Get(NameField).(string),
		Enterprise:   api.Enterprise,
		Datacenter:   d.Get(DataCenterField).(string),
		VdcResources: d.Get(VdcResourceField).([]interface{}),
		Slug:         d.Get(SlugField).(string),
		DynamicField: d.Get(DynamicField).(string),
	}
	for index, resource := range vdc.VdcResources {
		resourceSlug, err := getResourceSlug(resource.(map[string]interface{})[ResourceField].(string),
			*api.Meta)
		if err != nil {
			return vdcStruct{}, err
		}
		vdc.VdcResources[index].(map[string]interface{})[ResourceField] = resourceSlug
	}
	return vdc, nil
}

func getResourceSlug(resourceName string, meta APIMeta) (string, error) {
	for _, resource := range meta.EnterpriseResourceList {
		resourceExistsInMeta := (resource.(map[string]interface{})[NameField] == resourceName)
		isResourceMonoTyped := resource.(map[string]interface{})[ResourceCosField] == MonoResourceType
		if resourceExistsInMeta && isResourceMonoTyped {
			return resource.(map[string]interface{})[SlugField].(string), nil
		}
	}
	return "", errResourceNotExist(resourceName, "")
}

func getTemplateAndUpdateSchema(templateSlug string,
	d *schema.ResourceData,
	templatesTooler *TemplatesTooler,
	api *API) (map[string]interface{}, error) {
	template, err1 := templatesTooler.TemplatesTools.FetchTemplateFromList(templateSlug,
		api.Meta.TemplateList)
	if err1 != nil {
		return map[string]interface{}{}, err1
	}
	err2 := templatesTooler.TemplatesTools.validateTemplate(template)
	if err2 != nil {
		return map[string]interface{}{}, err2
	}
	err3 := templatesTooler.TemplatesTools.updateSchemaFromTemplateOnResourceCreation(d,
		template)
	if err3 != nil {
		return map[string]interface{}{}, err3
	}
	return template, nil
}

func vmInstanceCreate(d *schema.ResourceData,
	templatesTooler *TemplatesTooler,
	api *API) (vmStruct, error) {
	var (
		templateError error
		template      map[string]interface{}
		templateSlug  = d.Get(TemplateField).(string)
		vmName        strings.Builder
	)
	vmName.WriteString(d.Get(NameField).(string))
	if templateSlug != "" && d.Id() == "" {
		instanceNumber := d.Get(InstanceNumberField).(int)
		vmName.WriteString(resourceNameCountSeparator)
		vmName.WriteString(strconv.Itoa(instanceNumber))
		template,
			templateError = getTemplateAndUpdateSchema(templateSlug,
			d,
			templatesTooler,
			api)
	}
	if templateError != nil {
		return vmStruct{}, templateError
	}
	vm := vmStruct{
		Name:         vmName.String(),
		Enterprise:   api.Enterprise,
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
		Iso:          d.Get(IsoField).(string),
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
