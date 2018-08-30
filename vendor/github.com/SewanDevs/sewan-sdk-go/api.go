package sewansdk

import (
	"bytes"
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
	"strings"
)

const (
	defaultResourceType = VMResourceType
)

// API struct represents distant Sewan clouddc API
type API struct {
	Token  string
	URL    string
	Client *http.Client
}

// APITooler contains implementation of APIer interface
type APITooler struct {
	APIImplementer APIer
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

// New creates an API instance
func (apiTools *APITooler) New(token string, url string) *API {
	return &API{
		Token:  token,
		URL:    url,
		Client: &http.Client{},
	}
}

// CheckCloudDcStatus checks availability of clouddc through its API
func (apiTools *APITooler) CheckCloudDcStatus(api *API,
	clientTooler *ClientTooler,
	resourceTooler *ResourceTooler) error {
	var apiClientErr error
	apiClientErr = resourceTooler.Resource.validateStatus(api,
		defaultResourceType,
		*clientTooler)
	return apiClientErr
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
		clientTooler,
		templatesTooler,
		resourceType,
		sewan)
	if err1 != nil {
		return map[string]interface{}{}, err1.(error)
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
				readResource.(map[string]interface{}))
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
	readResource map[string]interface{}) error {
	var (
		resourceNamePrefix strings.Builder
		resourcesList      []interface{}
	)
	resourceNamePrefix.WriteString(readResource[EnterpriseField].(string))
	resourceNamePrefix.WriteString(monoField)
	for _, resource := range readResource[VdcResourceField].([]interface{}) {
		resource.(map[string]interface{})[ResourceField] = strings.TrimPrefix(resource.(map[string]interface{})[ResourceField].(string),
			resourceNamePrefix.String())
		resourcesList = append(resourcesList, resource)
	}
	return d.Set(VdcResourceField, resourcesList)
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
		clientTooler,
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
