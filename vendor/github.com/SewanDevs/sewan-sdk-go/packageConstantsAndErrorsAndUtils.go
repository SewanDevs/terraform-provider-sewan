package sewansdk

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// Exported constants are resource field names
const (
	ResourceCosField               = "cos"
	MonoResourceType               = "Mono"
	NameField                      = "name"
	EnterpriseField                = "enterprise"
	DataCenterField                = "datacenter"
	VdcResourceField               = "vdc_resources"
	ResourceField                  = "resource"
	TotalField                     = "total"
	UsedField                      = "used"
	SlugField                      = "slug"
	StateField                     = "state"
	OsField                        = "os"
	RAMField                       = "ram"
	CPUField                       = "cpu"
	DisksField                     = "disks"
	VDiskField                     = "v_disk"
	SizeField                      = "size"
	StorageClassField              = "storage_class"
	NicsField                      = "nics"
	VlanNameField                  = "vlan"
	MacAdressField                 = "mac_address"
	ConnectedField                 = "connected"
	VdcField                       = "vdc"
	BootField                      = "boot"
	TokenField                     = "token"
	BackupField                    = "backup"
	IsoField                       = "disk_image"
	PlatformNameField              = "platform_name"
	BackupSizeField                = "backup_size"
	CommentField                   = "comment"
	TemplateField                  = "template"
	IDField                        = "id"
	DynamicField                   = "dynamic_field"
	OutsourcingField               = "outsourcing"
	monoField                      = "-mono-"
	InstanceNumberField            = "instance_number"
	VMResourceType                 = "vm"
	VdcResourceType                = VdcField
	resourceNameCountSeparator     = "-"
	resourceDynamicInstanceNumber  = "${count.index + 1}"
	httpReqContentType             = "content-type"
	httpRespContentType            = "Content-Type"
	httpJSONContentType            = "application/json"
	httpHTMLTextContentType        = "text/html"
	httpAuthorization              = "authorization"
	httpTokenHeader                = "Token "
	errTestResultDiffs             = "\n\rGot: \"%s\"\n\rWant: \"%s\""
	errAPIUnhandledRespType        = "Unhandled api response type : "
	errValidateAPIURL              = "\nPlease validate the configuration api url."
	errReadOf                      = "Read of \""
	errUpdateStateFailedAndRespErr = "\" state failed, response reception error : "
	errJSONRespFailedAndJSONErr    = "\" failed, response body json error :\n\r\""
	errAPIDownOrwrongAPIURL        = "\", the api is down or this url is wrong."
	errEmptyResponse               = "Empty response error."
	errJSONFormat                  = "Response body is not a properly formated json :"
	creationOperation              = "Creation"
	readOperation                  = "Read"
	updateOperation                = "Update"
	deleteOperation                = "Delete"
	entrepriseSlugHTTPReqParam     = "/?enterprise__slug="
)

var (
	errDoRequest     = errors.New("do(request) error")
	errEmptyResp     = errors.New("empty API response")
	errEmptyRespBody = errors.New("empty API response body")
	errEmptyJSON     = errors.New("empty json")
	//ErrResourceNotExist provide message for unexisting resource case
	ErrResourceNotExist                 = errResourceNotExist("", "")
	errUninitializedExpectedCode        = errors.New("expected code not initialized")
	errNilResponse                      = errors.New("response is nil")
	errZeroStatusCode                   = errors.New("response status code is zero")
	err500ServerError                   = errors.New("<h1>Server Error (500)</h1>")
	errHandleResponse                   = errors.New("handle response error")
	errUnexpectedvalidateStatusResponse = errors.New("unexpected response to validate status request")
	errCheckRedirectFailure             = errors.New("CheckRedirectReqFailure")
	clouddcEnvironmentResource          = ResourceField
	clouddcEnvironmentVdc               = VdcField
	clouddcEnvironmentDatacenter        = DataCenterField
	clouddcEnvironmentTemplate          = TemplateField
	clouddcGenericTemplateEnterprise    = ",sewanadmin"
	clouddcEnvironmentVlan              = "vlan"
	clouddcEnvironmentSnapshot          = "snapshot"
	clouddcEnvironmentIso               = "disk-image"
	clouddcEnvironmentOva               = "ova"
	clouddcEnvironmentBackupPlan        = "backup-plan"
	resourceSlice                       = []string{
		clouddcEnvironmentResource,
		clouddcEnvironmentVdc,
		clouddcEnvironmentDatacenter,
		clouddcEnvironmentTemplate,
		clouddcEnvironmentVlan,
		clouddcEnvironmentSnapshot,
		clouddcEnvironmentIso,
		clouddcEnvironmentOva,
		clouddcEnvironmentBackupPlan,
	}
)

func errEmptyResourcesList(resourceType string) error {
	return errors.New("empty " + resourceType + " list")
}

func errNotInList(elem string, list string) error {
	return errors.New("\"" + elem + "\"" + " is not in :" + list)
}

func errResourceNotExist(resourceName string, availableResources string) error {
	if availableResources == "" {
		return errors.New("\"" + resourceName + "\" resource does not exists")
	}
	return errors.New("\"" + resourceName +
		"\" resource does not exists, available resources : " + availableResources)
}

func errRespStatusCodeBuilder(resp *http.Response,
	expectedCode int,
	additionalErrMsg string) error {
	if expectedCode == 0 {
		return errUninitializedExpectedCode
	}
	if resp == nil {
		return errNilResponse
	}
	if resp.StatusCode == 0 {
		return errZeroStatusCode
	}
	if expectedCode == resp.StatusCode {
		if additionalErrMsg == "" {
			return nil
		}
		return errors.New(additionalErrMsg)
	}
	return errors.New("Wrong response status code," +
		"\nexpected :" + strconv.Itoa(expectedCode) +
		"\ngot :" + strconv.Itoa(resp.StatusCode) +
		"\nFull response status : " + resp.Status + "\n" + additionalErrMsg)
}

func errDoCrudRequestsBuilder(crudOperation string,
	instanceName string,
	err error) error {
	of := " of \""
	postMsg := "\" failed, POST response reception error : "
	getMsg := "\" failed, GET response reception error : "
	deleteMsg := "\" failed, DELETE response reception error : "
	if instanceName == "" {
		return errors.New("instanceName is empty string")
	}
	if err == nil {
		return errors.New("request execution error is nil")
	}
	switch crudOperation {
	case creationOperation:
		return errors.New(creationOperation + of + instanceName +
			postMsg + err.Error())
	case readOperation:
		return errors.New(readOperation + of + instanceName +
			getMsg + err.Error())
	case updateOperation:
		return errors.New(updateOperation + of + instanceName +
			postMsg + err.Error())
	case deleteOperation:
		return errors.New(deleteOperation + of + instanceName +
			deleteMsg + err.Error())
	default:
		return errors.New(crudOperation + "is not a crudOperation from list :" +
			creationOperation + readOperation + updateOperation + deleteOperation)
	}
}

func errWrongResourceTypeBuilder(resourceType string) error {
	if resourceType == "" {
		return errors.New("no resource type provided")
	}
	return errors.New("Resource of type \"" + resourceType + "\" not supported," +
		"list of accepted resource types :\n\r" +
		"- \"" + VdcResourceType + "\"\n\r" +
		"- \"" + VMResourceType + "\"")
}

func stringSliceContains(elem string, slice []string) error {
	var (
		isInSlice bool
	)
	for _, listElem := range slice {
		if elem == listElem {
			isInSlice = true
		}
	}
	if isInSlice {
		return nil
	}
	return errNotInList(elem, strings.Join(slice, ", "))
}
