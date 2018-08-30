package sewansdk

import (
	"errors"
	"net/http"
	"strconv"
)

// Exported constants are resource field names
const (
	NameField                      = "name"
	EnterpriseField                = "enterprise"
	DatacenterField                = "datacenter"
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
	DiskImageField                 = "disk_image"
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
)

var (
	errDoRequest         = errors.New("do(request) error")
	errEmptyResp         = errors.New("empty API response")
	errEmptyRespBody     = errors.New("empty API response body")
	errEmptyTemplateList = errors.New("empty template list")
	//ErrResourceNotExist provide message for unexisting resource case
	ErrResourceNotExist                 = errors.New("resource does not exists")
	errUninitializedExpectedCode        = errors.New("expected code not initialized")
	errNilResponse                      = errors.New("response is nil")
	errZeroStatusCode                   = errors.New("response status code is zero")
	err500ServerError                   = errors.New("<h1>Server Error (500)</h1>")
	errHandleResponse                   = errors.New("handle response error")
	errUnexpectedvalidateStatusResponse = errors.New("unexpected response to validate status request")
	errCheckRedirectFailure             = errors.New("CheckRedirectReqFailure")
)

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
