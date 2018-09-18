package sewansdk

import (
	"bytes"
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
)

const (
	defaultResourceType = VMResourceType
)

// API struct represents distant Sewan clouddc API
type API struct {
	Token      string
	URL        string
	Enterprise string
	Meta       *APIMeta
	Client     *http.Client
}

// APITooler contains implementation of APIer and APIInitialyser interfaces
type APITooler struct {
	Implementer APIer
	Initialyser APIInitialyser
}

// APIMeta stores specific meta data about a clouddc environment
type APIMeta struct {
	EnterpriseResourceList []interface{}
	EnterpriseVdcList      []interface{}
	DataCenterList         []interface{}
	TemplateList           []interface{}
	VlanList               []interface{}
	SnapshotList           []interface{}
	IsoList                []interface{}
	OvaList                []interface{}
	BackupPlanList         []interface{}
}

// APIer interface is responsible of CRUD operations on Sewan's clouddc resources,
// they are done through AirDrumAPI consumption.
type APIer interface {
	CreateResource(d *schema.ResourceData,
		clientTooler *ClientTooler,
		templatesTooler *TemplatesTooler,
		resourceTooler *ResourceTooler,
		resourceType string,
		sewan *API) (map[string]interface{}, error)
	ReadResource(d *schema.ResourceData,
		clientTooler *ClientTooler,
		resourceTooler *ResourceTooler,
		resourceType string,
		sewan *API) (map[string]interface{}, error)
	UpdateResource(d *schema.ResourceData,
		clientTooler *ClientTooler,
		templatesTooler *TemplatesTooler,
		resourceTooler *ResourceTooler,
		resourceType string,
		sewan *API) error
	DeleteResource(d *schema.ResourceData,
		clientTooler *ClientTooler,
		resourceTooler *ResourceTooler,
		resourceType string,
		sewan *API) error
}

// AirDrumResourcesAPI implements APIer interface
type AirDrumResourcesAPI struct{}

// APIInitialyser interface is responsible of initializing an API client
type APIInitialyser interface {
	New(token string, url string, enterprise string) *API
	CheckCloudDcStatus(api *API,
		clientTooler *ClientTooler,
		resourceTooler *ResourceTooler) error
	GetClouddcEnvMeta(api *API,
		clientTooler *ClientTooler) (*APIMeta, error)
	// Add here fields validation
	//ValidateResourceFieldsValue
}

// Initialyser implements APIInitialyser interface
type Initialyser struct{}

// New creates an API instance
func (initialyser Initialyser) New(token string, url string, enterprise string) *API {
	return &API{
		Token:      token,
		URL:        url,
		Enterprise: enterprise,
		Client:     &http.Client{},
	}
}

// CheckCloudDcStatus checks availability of clouddc through its API
func (initialyser Initialyser) CheckCloudDcStatus(api *API,
	clientTooler *ClientTooler,
	resourceTooler *ResourceTooler) error {
	return resourceTooler.Resource.validateStatus(api,
		defaultResourceType,
		*clientTooler)
}

// GetClouddcEnvMeta gets Clouddc environnements meta data :
// * physical clouddc resource lists :
//		- non critical resource list
//		- critical resource list (redondant resource for critic uses)
//		- other resource list (Windows server license, RedHat licenses, etc.)
func (initialyser Initialyser) GetClouddcEnvMeta(api *API,
	clientTooler *ClientTooler) (*APIMeta, error) {
	var (
		apiMeta APIMeta
	)
	resourceMetaDataList,
		err1 := clientTooler.Client.getEnvResourceList(clientTooler,
		api, clouddcEnvironmentResource)
	if err1 != nil {
		return nil, err1
	}
	templateList, err2 := clientTooler.Client.getEnvResourceList(clientTooler,
		api, clouddcEnvironmentTemplate)
	if err2 != nil {
		return nil, err2
	}
	dataCenterList, err3 := clientTooler.Client.getEnvResourceList(clientTooler,
		api, clouddcEnvironmentDatacenter)
	if err3 != nil {
		return nil, err3
	}
	vlanList, err4 := clientTooler.Client.getEnvResourceList(clientTooler,
		api, clouddcEnvironmentVlan)
	if err4 != nil {
		return nil, err4
	}
	snapshotList, err5 := clientTooler.Client.getEnvResourceList(clientTooler,
		api, clouddcEnvironmentSnapshot)
	if err5 != nil {
		return nil, err5
	}
	isoList, err6 := clientTooler.Client.getEnvResourceList(clientTooler,
		api, clouddcEnvironmentIso)
	if err6 != nil {
		return nil, err6
	}
	ovaList, err7 := clientTooler.Client.getEnvResourceList(clientTooler,
		api, clouddcEnvironmentOva)
	if err7 != nil {
		return nil, err7
	}
	backupPlanList, err8 := clientTooler.Client.getEnvResourceList(clientTooler,
		api, clouddcEnvironmentBackupPlan)
	if err8 != nil {
		return nil, err8
	}
	vdcList,
		err9 := clientTooler.Client.getEnvResourceList(clientTooler,
		api, clouddcEnvironmentVdc)
	if err9 != nil {
		return nil, err9
	}
	logger := LoggerCreate("GetClouddcEnvMeta.log")
	apiMeta.EnterpriseResourceList = resourceMetaDataList
	logger.Println("resourceMetaDataList =", resourceMetaDataList)
	apiMeta.EnterpriseVdcList = vdcList
	logger.Println("vdcList =", vdcList)
	apiMeta.DataCenterList = dataCenterList
	logger.Println(" dataCenterList =", dataCenterList)
	apiMeta.TemplateList = templateList
	logger.Println("templateList =", templateList)
	apiMeta.VlanList = vlanList
	logger.Println("vlanList =", vlanList)
	apiMeta.SnapshotList = snapshotList
	logger.Println("snapshotList =", snapshotList)
	apiMeta.IsoList = isoList
	logger.Println("isoList =", isoList)
	apiMeta.OvaList = ovaList
	logger.Println("ovaList =", ovaList)
	apiMeta.BackupPlanList = backupPlanList
	logger.Println("backupPlanList =", backupPlanList)
	// redmine ticket #37823 : resource's lists validation lack
	return &apiMeta, nil
}

// CreateResource creates Sewan clouddc resource
//
// * input data : schema.ResourceData (godoc.org/github.com/hashicorp/terraform/helper/schema#ResourceData)
//
// * return map of created resource or eventual creation error
func (apier AirDrumResourcesAPI) CreateResource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	resourceTooler *ResourceTooler,
	resourceType string,
	sewan *API) (map[string]interface{}, error) {
	var (
		instanceName = d.Get(NameField).(string)
	)
	resourceInstance, err1 := resourceTooler.Resource.resourceInstanceCreate(d,
		templatesTooler,
		resourceType,
		sewan)
	if err1 != nil {
		return map[string]interface{}{}, err1
	}
	resourceJSON, err2 := json.Marshal(resourceInstance)
	if err2 != nil {
		return map[string]interface{}{}, err2
	}
	req, err3 := http.NewRequest("POST",
		resourceTooler.Resource.getResourceCreationURL(sewan, resourceType),
		bytes.NewBuffer(resourceJSON))
	if err3 != nil {
		return map[string]interface{}{}, err3
	}
	req.Header.Add(httpAuthorization, httpTokenHeader+sewan.Token)
	req.Header.Add(httpReqContentType, httpJSONContentType)
	resp, err4 := clientTooler.Client.do(sewan, req)
	switch {
	case err4 != nil:
		return map[string]interface{}{}, errDoCrudRequestsBuilder(creationOperation,
			instanceName, err4)
	default:
		createdResource, err5 := clientTooler.Client.handleResponse(resp,
			http.StatusCreated,
			httpJSONContentType)
		if createdResource == nil {
			return map[string]interface{}{}, err5
		}
		return createdResource.(map[string]interface{}), err5
	}
}

// ReadResource reads Sewan clouddc resource's state and update terraform's resource's state
//
// * input data : schema.ResourceData (with existing id) (godoc.org/github.com/hashicorp/terraform/helper/schema#ResourceData)
//
// * return map of read resource or eventual read error
func (apier AirDrumResourcesAPI) ReadResource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceTooler *ResourceTooler,
	resourceType string,
	sewan *API) (map[string]interface{}, error) {
	err1 := resourceTooler.Resource.validateResourceType(resourceType)
	if err1 != nil {
		return map[string]interface{}{}, err1
	}
	req, err2 := http.NewRequest("GET",
		resourceTooler.Resource.getResourceURL(sewan, resourceType, d.Id()), nil)
	if err2 != nil {
		return map[string]interface{}{}, err2
	}
	req.Header.Add(httpAuthorization, httpTokenHeader+sewan.Token)
	resp, err3 := clientTooler.Client.do(sewan, req)
	switch {
	case err3 != nil:
		return map[string]interface{}{}, errDoCrudRequestsBuilder(readOperation,
			d.Get(NameField).(string),
			err3)
	default:
		if (resp != nil) && (resp.StatusCode == http.StatusNotFound) {
			return map[string]interface{}{}, ErrResourceNotExist
		}
		readResource, err4 := clientTooler.Client.handleResponse(resp,
			http.StatusOK,
			httpJSONContentType)
		if readResource == nil {
			return map[string]interface{}{}, err4
		}
		if resourceType == VdcResourceType {
			err5 := updateSchemaReadVdcResource(d,
				readResource.(map[string]interface{}), sewan)
			if err5 != nil {
				return map[string]interface{}{}, err5
			}
		}
		return readResource.(map[string]interface{}), err4
	}
}

// updateSchemaReadVdcResource handle VDC resource names
//
// AirDrum use complex vdc resource names (example : <enterprise name>-mono-<resource name>),
// the aim of this function is to rm "<enterprise name>-mono-" name part to
// simplify terraform user experience
func updateSchemaReadVdcResource(d *schema.ResourceData,
	readResource map[string]interface{}, api *API) error {
	var (
		resourcesList []interface{}
	)
	for _, resource := range readResource[VdcResourceField].([]interface{}) {
		resourceName,
			err := getResourceName(resource.(map[string]interface{})[ResourceField].(string),
			*api.Meta)
		if err != nil {
			return err
		}
		resource.(map[string]interface{})[ResourceField] = resourceName
		resourcesList = append(resourcesList, resource)
	}
	return d.Set(VdcResourceField, resourcesList)
}

// getResourceName extracts from APIMeta and returns corresponding resource's name
func getResourceName(resourceSlug string, meta APIMeta) (string, error) {
	for _, resource := range meta.EnterpriseResourceList {
		resourceExistsInMeta := (resource.(map[string]interface{})[SlugField] == resourceSlug)
		isResourceMonoTyped := resource.(map[string]interface{})[ResourceCosField] == MonoResourceType
		if resourceExistsInMeta && isResourceMonoTyped {
			return resource.(map[string]interface{})[NameField].(string), nil
		}
	}
	return "", errResourceNotExist(resourceSlug, "")
}

// UpdateResource update Sewan clouddc resource's
//
// * input data : schema.ResourceData (with existing id) (godoc.org/github.com/hashicorp/terraform/helper/schema#ResourceData)
//
// * return map of updated resource or eventual update error
func (apier AirDrumResourcesAPI) UpdateResource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	templatesTooler *TemplatesTooler,
	resourceTooler *ResourceTooler,
	resourceType string,
	sewan *API) error {
	resourceInstance,
		err1 := resourceTooler.Resource.resourceInstanceCreate(d,
		templatesTooler,
		resourceType,
		sewan)
	if err1 != nil {
		return err1
	}
	resourceJSON, err2 := json.Marshal(resourceInstance)
	if err2 != nil {
		return err2
	}
	req, err3 := http.NewRequest("PUT",
		resourceTooler.Resource.getResourceURL(sewan, resourceType, d.Id()),
		bytes.NewBuffer(resourceJSON))
	if err3 != nil {
		return err3
	}
	req.Header.Add(httpAuthorization, httpTokenHeader+sewan.Token)
	req.Header.Add(httpReqContentType, httpJSONContentType)
	resp, err4 := clientTooler.Client.do(sewan, req)
	switch {
	case err4 != nil:
		return errDoCrudRequestsBuilder(updateOperation,
			d.Get(NameField).(string),
			err4)
	default:
		_, err5 := clientTooler.Client.handleResponse(resp,
			http.StatusOK,
			httpJSONContentType)
		return err5
	}
}

// DeleteResource deletes a Sewan clouddc resource
//
// * input data : schema.ResourceData (with existing id) (godoc.org/github.com/hashicorp/terraform/helper/schema#ResourceData)
//
// * return eventual deletion error
func (apier AirDrumResourcesAPI) DeleteResource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	resourceTooler *ResourceTooler,
	resourceType string,
	sewan *API) error {
	err1 := resourceTooler.Resource.validateResourceType(resourceType)
	if err1 != nil {
		return err1
	}
	req, err2 := http.NewRequest("DELETE",
		resourceTooler.Resource.getResourceURL(sewan, resourceType, d.Id()), nil)
	if err2 != nil {
		return err2
	}
	req.Header.Add(httpAuthorization, httpTokenHeader+sewan.Token)
	resp, err3 := clientTooler.Client.do(sewan, req)
	switch {
	case err3 != nil:
		return errDoCrudRequestsBuilder(deleteOperation,
			d.Get(NameField).(string),
			err3)
	default:
		_, err4 := clientTooler.Client.handleResponse(resp,
			http.StatusNoContent,
			httpJSONContentType)
		return err4
	}
}
