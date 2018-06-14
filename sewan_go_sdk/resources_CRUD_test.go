package sewan_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

//------------------------------------------------------------------------------
//-------------Units tests------------------------------------------------------
//------------------------------------------------------------------------------
func TestCreate_resource(t *testing.T) {
	test_cases := []struct {
		Id               int
		TC_clienter      Clienter
		ResourceType     string
		Creation_Err     error
		Created_resource map[string]interface{}
	}{
		{
			1,
			ErrorResponse_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New(REQ_ERR),
			nil,
		},
		{
			2,
			BadBodyResponse_StatusCreated_HttpClienter{},
			VDC_RESOURCE_TYPE,
			errors.New("Creation of \"Unit test resource\" failed, " +
				"the response body is not a properly formated json :" +
				"\n\r\"invalid character '\"' after object key\""),
			nil,
		},
		{
			3,
			Error401_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New(UNAUTHORIZED_MSG),
			nil,
		},
		{
			4,
			VDC_CreationSuccess_HttpClienter{},
			VDC_RESOURCE_TYPE,
			nil,
			TEST_VDC_READ_RESPONSE_MAP,
		},
		{
			5,
			VM_CreationSuccess_HttpClienter{},
			VM_RESOURCE_TYPE,
			nil,
			TEST_VM_MAP,
		},
		{
			6,
			CheckRedirectReqFailure_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New("Creation of \"Unit test resource\" failed, response reception " +
				"error : CheckRedirectReqFailure"),
			nil,
		},
		{
			7,
			BadBodyResponseContentType_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New("Unhandled api response type : image" +
				"\nPlease validate the configuration api url."),
			nil,
		},
		{
			8,
			StatusInternalServerError_HttpClienter{},
			VDC_RESOURCE_TYPE,
			errors.New("<h1>Server Error (500)</h1>"),
			nil,
		},
		{
			8,
			StatusInternalServerError_HttpClienter{},
			VDC_RESOURCE_TYPE,
			errors.New("<h1>Server Error (500)</h1>"),
			nil,
		},
		{
			9,
			VM_CreationSuccess_HttpClienter{},
			WRONG_RESOURCE_TYPE,
			errors.New("Resource of type \"a_non_supported_resource_type\"" +
				" not supported,list of accepted resource types :\n\r" +
				"- \"vdc\"\n\r" +
				"- \"vm\""),
			nil,
		},
	}
	var (
		sewan             *API
		err               error
		resp_creation_map map[string]interface{}
		resource_res      *schema.Resource
		d                 *schema.ResourceData
	)
	Apier := AirDrumResources_Apier{}

	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		resource_res = resource(test_case.ResourceType)
		d = resource_res.TestResourceData()
		d.SetId("UnitTest resource1")
		d.Set("name", "Unit test resource")
		fake_client_tooler.Client = test_case.TC_clienter
		err, resp_creation_map = Apier.Create_resource(d,
			&fake_client_tooler,
			test_case.ResourceType,
			sewan)

		switch {
		case err == nil || test_case.Creation_Err == nil:
			if !(err == nil && test_case.Creation_Err == nil) {
				t.Errorf("TC %d : resource creation error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Creation_Err)
			} else {
				switch {
				case !reflect.DeepEqual(test_case.Created_resource, resp_creation_map):
					t.Errorf("TC %d : Wrong created resource map,"+
						"\n\rgot: \"%s\"\n\rwant: \"%s\"",
						test_case.Id, resp_creation_map, test_case.Created_resource)
				}
			}
		case err != nil && test_case.Creation_Err != nil:
			if resp_creation_map != nil {
				t.Errorf("TC %d : Wrong created resource map,"+
					" it should be nil as error is not nil,"+
					"\n\rgot map: \n\r\"%s\"\n\rwant map: \n\r\"%s\"\n\r",
					test_case.Id, resp_creation_map, test_case.Created_resource)
			}
			if err.Error() != test_case.Creation_Err.Error() {
				t.Errorf("TC %d : resource creation error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, err.Error(), test_case.Creation_Err.Error())
			}
		case !reflect.DeepEqual(test_case.Created_resource, resp_creation_map):
			t.Errorf("TC %d : Wrong created resource map,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, resp_creation_map, test_case.Created_resource)
		}
	}
}

//------------------------------------------------------------------------------
func TestRead_resource(t *testing.T) {
	test_cases := []struct {
		Id              int
		TC_clienter     Clienter
		ResourceType    string
		Read_Err        error
		Read_resource   map[string]interface{}
		Resource_exists bool
	}{
		{
			1,
			ErrorResponse_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New(REQ_ERR),
			nil,
			true,
		},
		{
			2,
			BadBodyResponse_StatusOK_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New("Read of \"Unit test resource\" failed, response body json " +
				"error :\n\r\"invalid character '\"' after object key\""),
			nil,
			true,
		},
		{
			3,
			Error401_HttpClienter{},
			VDC_RESOURCE_TYPE,
			errors.New(UNAUTHORIZED_MSG),
			nil,
			true,
		},
		{
			4,
			Error404_HttpClienter{},
			VM_RESOURCE_TYPE,
			nil,
			nil,
			false,
		},
		{
			5,
			CheckRedirectReqFailure_HttpClienter{},
			VDC_RESOURCE_TYPE,
			errors.New("Read of \"Unit test resource\" state failed, response reception " +
				"error : CheckRedirectReqFailure"),
			nil,
			true,
		},
		{
			6,
			VDC_ReadSuccess_HttpClienter{},
			VDC_RESOURCE_TYPE,
			nil,
			TEST_VDC_READ_RESPONSE_MAP,
			true,
		},
		{
			7,
			VM_ReadSuccess_HttpClienter{},
			VM_RESOURCE_TYPE,
			nil,
			TEST_VM_MAP,
			true,
		},
		{
			8,
			BadBodyResponseContentType_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New("Unhandled api response type : image" +
				"\nPlease validate the configuration api url."),
			nil,
			true,
		},
		{
			9,
			StatusInternalServerError_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New("<h1>Server Error (500)</h1>"),
			nil,
			true,
		},
		{
			10,
			VDC_ReadSuccess_HttpClienter{},
			WRONG_RESOURCE_TYPE,
			errors.New("Resource of type \"a_non_supported_resource_type\"" +
				" not supported,list of accepted resource types :\n\r" +
				"- \"vdc\"\n\r" +
				"- \"vm\""),
			nil,
			true,
		},
	}
	var (
		sewan             *API
		err               error
		resp_creation_map map[string]interface{}
		res_exists        bool
		resource_res      *schema.Resource
		d                 *schema.ResourceData
	)
	Apier := AirDrumResources_Apier{}
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		resource_res = resource(test_case.ResourceType)
		d = resource_res.TestResourceData()
		d.SetId("UnitTest resource1")
		d.Set("name", "Unit test resource")
		fake_client_tooler.Client = test_case.TC_clienter
		err, resp_creation_map, res_exists = Apier.Read_resource(d,
			&fake_client_tooler,
			test_case.ResourceType,
			sewan)

		switch {
		case err == nil || test_case.Read_Err == nil:
			if !((err == nil) && (test_case.Read_Err == nil)) {
				t.Errorf("TC %d : resource read error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Read_Err)
			} else {
				switch {
				case res_exists != test_case.Resource_exists:
					t.Errorf("TC %d : Wrong read resource exists value"+
						"\n\rgot: \"%v\"\n\rwant: \"%v\"",
						test_case.Id, res_exists, test_case.Resource_exists)
				case !reflect.DeepEqual(test_case.Read_resource, resp_creation_map):
					t.Errorf("TC %d : Wrong resource read resource map,"+
						"\n\rgot: \"%s\"\n\rwant: \"%s\"",
						test_case.Id, resp_creation_map, test_case.Read_resource)
				}
			}
		case err != nil && test_case.Read_Err != nil:
			if resp_creation_map != nil {
				t.Errorf("TC %d : Wrong created resource map,"+
					" it should be nil as error is not nil,"+
					"\n\rgot map: \n\r\"%s\"\n\rwant map: \n\r\"%s\"\n\r",
					test_case.Id, resp_creation_map, test_case.Read_resource)
			}
			if err.Error() != test_case.Read_Err.Error() {
				t.Errorf("TC %d : resource read error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, err.Error(), test_case.Read_Err.Error())
			}
		case res_exists != test_case.Resource_exists:
			t.Errorf("TC %d : Wrong read resource exists value"+
				"\n\rgot: \"%v\"\n\rwant: \"%v\"",
				test_case.Id, res_exists, test_case.Resource_exists)
		case !reflect.DeepEqual(test_case.Read_resource, resp_creation_map):
			t.Errorf("TC %d : Wrong resource read resource map,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, resp_creation_map, test_case.Read_resource)
		}
	}
}

//------------------------------------------------------------------------------
func TestUpdate_resource(t *testing.T) {
	test_cases := []struct {
		Id           int
		TC_clienter  Clienter
		ResourceType string
		Update_Err   error
	}{
		{
			1,
			ErrorResponse_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New(REQ_ERR),
		},
		{
			2,
			BadBodyResponse_StatusOK_HttpClienter{},
			VDC_RESOURCE_TYPE,
			errors.New("Read of \"Unit test resource\" failed, response body json " +
				"error :\n\r\"invalid character '\"' after object key"),
		},
		{
			3,
			Error401_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New(UNAUTHORIZED_MSG),
		},
		{
			4,
			Error404_HttpClienter{},
			VDC_RESOURCE_TYPE,
			errors.New(NOT_FOUND_MSG),
		},
		{
			5,
			CheckRedirectReqFailure_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New("Update of \"Unit test resource\" state failed, response reception " +
				"error : CheckRedirectReqFailure"),
		},
		{
			6,
			VDC_UpdateSuccess_HttpClienter{},
			VDC_RESOURCE_TYPE,
			nil,
		},
		{
			7,
			VM_UpdateSuccess_HttpClienter{},
			VM_RESOURCE_TYPE,
			nil,
		},
		{
			8,
			BadBodyResponseContentType_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New("Unhandled api response type : image" +
				"\nPlease validate the configuration api url."),
		},
		{
			9,
			StatusInternalServerError_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New("<h1>Server Error (500)</h1>"),
		},
		{
			10,
			VDC_UpdateSuccess_HttpClienter{},
			WRONG_RESOURCE_TYPE,
			errors.New("Resource of type \"a_non_supported_resource_type\"" +
				" not supported,list of accepted resource types :\n\r" +
				"- \"vdc\"\n\r" +
				"- \"vm\""),
		},
	}
	Apier := AirDrumResources_Apier{}
	var (
		sewan        *API
		err          error
		resource_res *schema.Resource
		d            *schema.ResourceData
	)
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		resource_res = resource(test_case.ResourceType)
		d = resource_res.TestResourceData()
		d.SetId("UnitTest resource1")
		d.Set("name", "Unit test resource")
		fake_client_tooler.Client = test_case.TC_clienter
		err = Apier.Update_resource(d,
			&fake_client_tooler,
			test_case.ResourceType,
			sewan)

		switch {
		case err == nil || test_case.Update_Err == nil:
			if !(err == nil && test_case.Update_Err == nil) {
				t.Errorf("TC %d : resource read error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Update_Err)
			}
		case err.Error() != test_case.Update_Err.Error():
			t.Errorf("TC %d : resource read error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Update_Err.Error())
		}
	}
}

//------------------------------------------------------------------------------
func TestDelete_resource(t *testing.T) {
	test_cases := []struct {
		Id           int
		TC_clienter  Clienter
		ResourceType string
		Delete_Err   error
	}{
		{
			1,
			ErrorResponse_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New(REQ_ERR),
		},
		{
			2,
			BadBodyResponse_StatusOK_HttpClienter{},
			VDC_RESOURCE_TYPE,
			errors.New("Read of \"Unit test resource\" failed, response body json " +
				"error :\n\r\"invalid character '\"' after object key"),
		},
		{
			3,
			Error401_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New(UNAUTHORIZED_MSG),
		},
		{
			4,
			Error404_HttpClienter{},
			VDC_RESOURCE_TYPE,
			errors.New(NOT_FOUND_MSG),
		},
		{
			5,
			CheckRedirectReqFailure_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New("Deletion of \"Unit test resource\" state failed, response reception " +
				"error : CheckRedirectReqFailure"),
		},
		{
			6,
			VDC_DeleteSuccess_HttpClienter{},
			VDC_RESOURCE_TYPE,
			nil,
		},
		{
			7,
			VM_DeleteSuccess_HttpClienter{},
			VM_RESOURCE_TYPE,
			nil,
		},
		{
			8,
			DeleteWRONGResponseBody_HttpClienter{},
			VDC_RESOURCE_TYPE,
			errors.New(DESTROY_WRONG_MSG),
		},
		{
			9,
			BadBodyResponseContentType_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New("Unhandled api response type : image" +
				"\nPlease validate the configuration api url."),
		},
		{
			10,
			StatusInternalServerError_HttpClienter{},
			VM_RESOURCE_TYPE,
			errors.New("<h1>Server Error (500)</h1>"),
		},
		{
			7,
			VM_DeleteSuccess_HttpClienter{},
			WRONG_RESOURCE_TYPE,
			errors.New("Resource of type \"a_non_supported_resource_type\"" +
				" not supported,list of accepted resource types :\n\r" +
				"- \"vdc\"\n\r" +
				"- \"vm\""),
		},
	}
	var (
		sewan        *API
		err          error
		resource_res *schema.Resource
		d            *schema.ResourceData
	)
	Apier := AirDrumResources_Apier{}
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		resource_res = resource(test_case.ResourceType)
		d = resource_res.TestResourceData()
		d.SetId("UnitTest resource1")
		d.Set("name", "Unit test resource")
		fake_client_tooler.Client = test_case.TC_clienter
		err = Apier.Delete_resource(d,
			&fake_client_tooler,
			test_case.ResourceType,
			sewan)

		switch {
		case err == nil || test_case.Delete_Err == nil:
			if !(err == nil && test_case.Delete_Err == nil) {
				t.Errorf("TC %d : resource read error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Delete_Err)
			}
		case err.Error() != test_case.Delete_Err.Error():
			t.Errorf("TC %d : resource read error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Delete_Err.Error())
		}
	}
}

//------------------------------------------------------------------------------
//--Structures init, interface implementation fakes, various test items etc.----
//------------------------------------------------------------------------------
const (
	REQ_ERR                 = "Creation request response error."
	NOT_FOUND_STATUS        = "404 Not Found"
	NOT_FOUND_MSG           = "404 Not Found{\"detail\":\"Not found.\"}"
	UNAUTHORIZED_STATUS     = "401 Unauthorized"
	UNAUTHORIZED_MSG        = "401 Unauthorized{\"detail\":\"Token non valide.\"}"
	DESTROY_WRONG_MSG       = "{\"detail\":\"Destroying resource wrong body message\"}"
	CHECK_REDIRECT_FAILURE  = "CheckRedirectReqFailure"
	VDC_DESTROY_FAILURE_MSG = "Destroying the VDC now"
	VM_DESTROY_FAILURE_MSG  = "Destroying the VM now"
	VM_RESOURCE_TYPE        = "vm"
	VDC_RESOURCE_TYPE       = "vdc"
	WRONG_RESOURCE_TYPE     = "a_non_supported_resource_type"
)

var (
	TEST_VDC_CREATION_MAP = map[string]interface{}{
		"name":       "Unit test vdc",
		"enterprise": "sewan-rd-cloud-beta",
		"datacenter": "dc1",
		"vdc_resources": []interface{}{
			map[string]interface{}{
				"resource": "sewan-rd-cloud-beta-mono-ram",
				"total":    "20",
			},
			map[string]interface{}{
				"resource": "sewan-rd-cloud-beta-mono-cpu",
				"total":    "1",
			},
			map[string]interface{}{
				"resource": "sewan-rd-cloud-beta-mono-storage_enterprise",
				"total":    "10",
			},
			map[string]interface{}{
				"resource": "sewan-rd-cloud-beta-mono-storage_performance",
				"total":    "10",
			},
			map[string]interface{}{
				"resource": "sewan-rd-cloud-beta-mono-storage_high_performance",
				"total":    "10",
			},
		},
	}
	TEST_VDC_READ_RESPONSE_MAP = map[string]interface{}{
		"name":       "Unit test vdc",
		"enterprise": "sewan-rd-cloud-beta",
		"datacenter": "dc1",
		"vdc_resources": []interface{}{
			map[string]interface{}{
				"resource": "sewan-rd-cloud-beta-mono-ram",
				"used":     "0",
				"total":    "20",
				"slug":     "sewan-rd-cloud-beta-dc1-vdc_te-sewan-rd-cloud-beta-mono-ram",
			},
			map[string]interface{}{
				"resource": "sewan-rd-cloud-beta-mono-cpu",
				"used":     "0",
				"total":    "1",
				"slug":     "sewan-rd-cloud-beta-dc1-vdc_te-sewan-rd-cloud-beta-mono-cpu",
			},
			map[string]interface{}{
				"resource": "sewan-rd-cloud-beta-mono-storage_enterprise",
				"used":     "0",
				"total":    "10",
				"slug":     "sewan-rd-cloud-beta-dc1-vdc_te-sewan-rd-cloud-beta-mono-storage_enterprise",
			},
			map[string]interface{}{
				"resource": "sewan-rd-cloud-beta-mono-storage_performance",
				"used":     "0",
				"total":    "10",
				"slug":     "sewan-rd-cloud-beta-dc1-vdc_te-sewan-rd-cloud-beta-mono-storage_performance",
			},
			map[string]interface{}{
				"resource": "sewan-rd-cloud-beta-mono-storage_high_performance",
				"used":     "0",
				"total":    "10",
				"slug":     "sewan-rd-cloud-beta-dc1-vdc_te-sewan-rd-cloud-beta-mono-storage_high_performance",
			},
		},
		"slug":          "sewan-rd-cloud-beta-dc1-vdc_te",
		"dynamic_field": "",
	}
	TEST_VM_MAP = map[string]interface{}{
		"name":     "Unit test resource",
		"template": "",
		"state":    "UP",
		"os":       "Debian",
		"ram":      "8",
		"cpu":      "4",
		"disks": []interface{}{
			map[string]interface{}{
				"name":   "disk 1",
				"size":   "24",
				"v_disk": "v_disk",
				"slug":   "slug",
			},
		},
		"nics": []interface{}{
			map[string]interface{}{
				"vlan":       "vlan 1 update",
				"mac_adress": "24",
				"connected":  "true",
			},
			map[string]interface{}{
				"vlan":       "vlan 2",
				"mac_adress": "24",
				"connected":  "true",
			},
		},
		"vdc":               "vdc",
		"boot":              "on disk",
		"vdc_resource_disk": "vdc_disk",
		"slug":              "42",
		"token":             "424242",
		"backup":            "backup-no_backup",
		"disk_image":        "",
		"platform_name":     "42",
		"backup_size":       "42",
		"comment":           "42",
		"outsourcing":       "42",
		"dynamic_field":     "42",
	}
)

func resource_vdc() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdc_resources": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"used": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"total": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"slug": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resource_vm() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"template": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ram": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"disks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"size": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"v_disk": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"slug": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"nics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"mac_adress": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"connected": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"vdc": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"boot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdc_resource_disk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"disk_image": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"platform_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"outsourcing": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resource(resourceType string) *schema.Resource {

	resource := &schema.Resource{}
	switch resourceType {
	case "vdc":
		resource = resource_vdc()
	case "vm":
		resource = resource_vm()
	default:
		//return a false resource
		resource = &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		}
	}
	return resource
}

type Resp_Body struct {
	Detail string `json:"detail"`
}

// Error response *ClientTooler
type ErrorResponse_HttpClienter struct{}

func (client ErrorResponse_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	return nil, errors.New(REQ_ERR)
}

// Response body error *ClientTooler
type BadBodyResponseContentType_HttpClienter struct{}

func (client BadBodyResponseContentType_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "image")
	resp.Body = ioutil.NopCloser(bytes.NewBufferString("Internal Server Error"))
	resp.StatusCode = http.StatusInternalServerError
	return &resp, nil
}

// Response body error *ClientTooler
type StatusInternalServerError_HttpClienter struct{}

func (client StatusInternalServerError_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "text/html")
	resp.Body = ioutil.NopCloser(bytes.NewBufferString("<h1>Server Error (500)</h1>"))
	resp.StatusCode = http.StatusInternalServerError
	return &resp, nil
}

// Response body error *ClientTooler
type BadBodyResponse_StatusCreated_HttpClienter struct{}

func (client BadBodyResponse_StatusCreated_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.Body = ioutil.NopCloser(bytes.NewBufferString("{\"detail\"\"Invalid json string}}.\"}"))
	resp.StatusCode = http.StatusCreated
	return &resp, nil
}

// Response body error *ClientTooler
type BadBodyResponse_StatusOK_HttpClienter struct{}

func (client BadBodyResponse_StatusOK_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.Body = ioutil.NopCloser(bytes.NewBufferString("{\"detail\"\"Invalid json string}}.\"}"))
	resp.StatusCode = http.StatusOK
	return &resp, nil
}

// 401 Reponse code error *ClientTooler
type Error401_HttpClienter struct{}

func (client Error401_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusUnauthorized
	resp.Status = UNAUTHORIZED_STATUS
	body := Resp_Body{"Token non valide."}
	js, _ := json.Marshal(body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

// 404 Reponse code error *ClientTooler
type Error404_HttpClienter struct{}

func (client Error404_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusNotFound
	resp.Status = NOT_FOUND_STATUS
	body := Resp_Body{"Not found."}
	js, _ := json.Marshal(body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

// Creation success *ClientTooler
type VDC_CreationSuccess_HttpClienter struct{}

func (client VDC_CreationSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusCreated
	js, _ := json.Marshal(TEST_VDC_READ_RESPONSE_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

// Read success *ClientTooler
type VDC_ReadSuccess_HttpClienter struct{}

func (client VDC_ReadSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	js, _ := json.Marshal(TEST_VDC_READ_RESPONSE_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

// Update success *ClientTooler
type VDC_UpdateSuccess_HttpClienter struct{}

func (client VDC_UpdateSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	js, _ := json.Marshal(TEST_VDC_CREATION_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

type VDC_DeleteSuccess_HttpClienter struct{}

func (client VDC_DeleteSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	body := Resp_Body{VDC_DESTROY_FAILURE_MSG}
	js, _ := json.Marshal(body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

// Creation success *ClientTooler
type VM_CreationSuccess_HttpClienter struct{}

func (client VM_CreationSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusCreated
	js, _ := json.Marshal(TEST_VM_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

// Read success *ClientTooler
type VM_ReadSuccess_HttpClienter struct{}

func (client VM_ReadSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	js, _ := json.Marshal(TEST_VM_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

// Update success *ClientTooler
type VM_UpdateSuccess_HttpClienter struct{}

func (client VM_UpdateSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	js, _ := json.Marshal(TEST_VM_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

type VM_DeleteSuccess_HttpClienter struct{}

func (client VM_DeleteSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	body := Resp_Body{VM_DESTROY_FAILURE_MSG}
	js, _ := json.Marshal(body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

type DeleteWRONGResponseBody_HttpClienter struct{}

func (client DeleteWRONGResponseBody_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{}
	resp.Header.Add("Content-Type", "application/json")
	resp.StatusCode = http.StatusOK
	resp.Body = ioutil.NopCloser(bytes.NewBufferString(DESTROY_WRONG_MSG))
	return &resp, nil
}

// req failure *ClientTooler
type CheckRedirectReqFailure_HttpClienter struct{}

func (client CheckRedirectReqFailure_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	return &resp, errors.New(CHECK_REDIRECT_FAILURE)
}
