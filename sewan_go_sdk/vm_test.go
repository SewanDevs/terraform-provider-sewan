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
//--Structures init, interface implementation fakes, various test items etc.----
//------------------------------------------------------------------------------
const (
	REQ_ERR = "Creation request response error."
	NOT_FOUND_STATUS = "404 Not Found"
	NOT_FOUND_MSG = "404 Not Found{\"detail\":\"Not found.\"}"
	UNAUTHORIZED_STATUS = "401 Unauthorized"
	UNAUTHORIZED_MSG = "401 Unauthorized{\"detail\":\"Token non valide.\"}"
	DESTROY_WRONG_MSG = "{\"detail\":\"Destroying VM wrong body message\"}"
	DESTROY_MSG = "Destroying the VM now"
	CHECK_REDIRECT_FAILURE = "CheckRedirectReqFailure"
)

var (
	TEST_VM_MAP = map[string]interface{}{"name": "Unit test vm",
		"state": "UP",
		"os":    "Debian",
		"ram":   "8",
		"cpu":   "4",
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
		"vdc_resource_disk": "vdc_disk", //"template":"template name",
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

func resource_vm() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ram": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"disks": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
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
				Required: true,
			},
			//"template": &schema.Schema{
			//  Type:     schema.TypeString,
			//  Optional: true,
			//},
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
type BadBodyResponse_StatusCreated_HttpClienter struct{}

func (client BadBodyResponse_StatusCreated_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{"Content-Type": {"application/json"}}
	resp.Body = ioutil.NopCloser(bytes.NewBufferString("{\"detail\"\"Invalid json string}}.\"}"))
	resp.StatusCode = http.StatusCreated
	return &resp, nil
}

// Response body error *ClientTooler
type BadBodyResponse_StatusOK_HttpClienter struct{}

func (client BadBodyResponse_StatusOK_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{"Content-Type": {"application/json"}}
	resp.Body = ioutil.NopCloser(bytes.NewBufferString("{\"detail\"\"Invalid json string}}.\"}"))
	resp.StatusCode = http.StatusOK
	return &resp, nil
}

// 401 Reponse code error *ClientTooler
type Error401_HttpClienter struct{}

func (client Error401_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
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
	resp.StatusCode = http.StatusNotFound
	resp.Status = NOT_FOUND_STATUS
	body := Resp_Body{"Not found."}
	js, _ := json.Marshal(body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

// Creation success *ClientTooler
type CreationSuccess_HttpClienter struct{}

func (client CreationSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{"Content-Type": {"application/json"}}
	resp.StatusCode = http.StatusCreated
	resp.Body = req.Body
	return &resp, nil
}

// Read success *ClientTooler
type ReadSuccess_HttpClienter struct{}

func (client ReadSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{"Content-Type": {"application/json"}}
	resp.StatusCode = http.StatusOK
	js, _ := json.Marshal(TEST_VM_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

// Update success *ClientTooler
type UpdateSuccess_HttpClienter struct{}

func (client UpdateSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{"Content-Type": {"application/json"}}
	resp.StatusCode = http.StatusOK
	js, _ := json.Marshal(TEST_VM_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

type DeleteSuccess_HttpClienter struct{}

func (client DeleteSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{"Content-Type": {"application/json"}}
	resp.StatusCode = http.StatusOK
	body := Resp_Body{DESTROY_MSG}
	js, _ := json.Marshal(body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

type DeleteWRONGResponseBody_HttpClienter struct{}

func (client DeleteWRONGResponseBody_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{"Content-Type": {"application/json"}}
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

//------------------------------------------------------------------------------
//-------------Units tests------------------------------------------------------
//------------------------------------------------------------------------------
func TestCreate_vm_resource(t *testing.T) {
	test_cases := []struct {
		Id               int
		TC_clienter      Clienter
		Creation_Err     error
		Created_resource map[string]interface{}
	}{
		{
			1,
			ErrorResponse_HttpClienter{},
			errors.New(REQ_ERR),
			nil,
		},
		{
			2,
			BadBodyResponse_StatusCreated_HttpClienter{},
			errors.New("Creation of \"Unit test vm\" failed, response body json " +
				"error :\n\r\"invalid character '\"' after object key\""),
			nil,
		},
		{
			3,
			Error401_HttpClienter{},
			errors.New(UNAUTHORIZED_MSG),
			nil,
		},
		{
			4,
			CreationSuccess_HttpClienter{},
			nil,
			TEST_VM_MAP,
		},
		{
			5,
			CheckRedirectReqFailure_HttpClienter{},
			errors.New("Creation of \"Unit test vm\" failed, response reception " +
				"error : CheckRedirectReqFailure"),
			nil,
		},
	}
	var (
		sewan             *API
		err               error
		resp_creation_map map[string]interface{}
	)
	Apier := AirDrumAPIer{}
	vm_res := resource_vm()
	d := vm_res.TestResourceData()
	d.SetId("UnitTest vm1")
	d.Set("name", "Unit test vm")
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		fake_client_tooler.Client = test_case.TC_clienter
		err, resp_creation_map = Apier.Create_vm_resource(d, &fake_client_tooler, sewan)
		t.Log("resp_creation_map, test_case.Created_resource",
			resp_creation_map, test_case.Created_resource)
		switch {
		case err == nil || test_case.Creation_Err == nil:
			if !(err == nil && test_case.Creation_Err == nil) {
				t.Errorf("TC %d : VM creation error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Creation_Err)
			}
		case err.Error() != test_case.Creation_Err.Error():
			t.Errorf("TC %d : VM creation error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Creation_Err.Error())
		case !reflect.DeepEqual(test_case.Created_resource, resp_creation_map):
			t.Errorf("TC %d : Wrong created resource map,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, resp_creation_map, test_case.Created_resource)
		}
	}
}

//------------------------------------------------------------------------------
func TestRead_vm_resource(t *testing.T) {
	test_cases := []struct {
		Id              int
		TC_clienter     Clienter
		Read_Err        error
		Read_resource   map[string]interface{}
		Resource_exists bool
	}{
		{
			1,
			ErrorResponse_HttpClienter{},
			errors.New(REQ_ERR),
			nil,
			true,
		},
		{
			2,
			BadBodyResponse_StatusOK_HttpClienter{},
			errors.New("Read of \"Unit test vm\" failed, response body json " +
				"error :\n\r\"invalid character '\"' after object key\""),
			nil,
			true,
		},
		{
			3,
			Error401_HttpClienter{},
			errors.New(UNAUTHORIZED_MSG),
			nil,
			true,
		},
		{
			4,
			Error404_HttpClienter{},
			errors.New(NOT_FOUND_MSG),
			nil,
			false,
		},
		{
			5,
			CheckRedirectReqFailure_HttpClienter{},
			errors.New("Read of \"Unit test vm\" state failed, response reception " +
				"error : CheckRedirectReqFailure"),
			nil,
			true,
		},
		{
			6,
			ReadSuccess_HttpClienter{},
			nil,
			TEST_VM_MAP,
			true,
		},
	}
	var (
		sewan             *API
		err               error
		resp_creation_map map[string]interface{}
		res_exists        bool
	)
	Apier := AirDrumAPIer{}
	vm_res := resource_vm()
	d := vm_res.TestResourceData()
	d.SetId("UnitTest vm1")
	d.Set("name", "Unit test vm")
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		fake_client_tooler.Client = test_case.TC_clienter
		err, resp_creation_map, res_exists = Apier.Read_vm_resource(d, &fake_client_tooler, sewan)
		switch {
		case err == nil || test_case.Read_Err == nil:
			if !(err == nil && test_case.Read_Err == nil) {
				t.Errorf("TC %d : VM read error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Read_Err)
			}
		case err.Error() != test_case.Read_Err.Error():
			t.Errorf("TC %d : VM read error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Read_Err.Error())
		case res_exists != test_case.Resource_exists:
			t.Errorf("TC %d : Wrong read vm exists value"+
				"\n\rgot: \"%v\"\n\rwant: \"%v\"",
				test_case.Id, res_exists, test_case.Resource_exists)
		case !reflect.DeepEqual(test_case.Read_resource, resp_creation_map):
			t.Errorf("TC %d : Wrong vm read resource map,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, resp_creation_map, test_case.Read_resource)
		}
	}
}

//------------------------------------------------------------------------------
func TestUpdate_vm_resource(t *testing.T) {
	test_cases := []struct {
		Id          int
		TC_clienter Clienter
		Update_Err  error
	}{
		{
			1,
			ErrorResponse_HttpClienter{},
			errors.New(REQ_ERR),
		},
		{
			2,
			BadBodyResponse_StatusOK_HttpClienter{},
			errors.New("Read of \"Unit test vm\" failed, response body json " +
				"error :\n\r\"invalid character '\"' after object key"),
		},
		{
			3,
			Error401_HttpClienter{},
			errors.New(UNAUTHORIZED_MSG),
		},
		{
			4,
			Error404_HttpClienter{},
			errors.New(NOT_FOUND_MSG),
		},
		{
			5,
			CheckRedirectReqFailure_HttpClienter{},
			errors.New("Update of \"Unit test vm\" state failed, response reception " +
				"error : CheckRedirectReqFailure"),
		},
		{
			6,
			UpdateSuccess_HttpClienter{},
			nil,
		},
	}
	Apier := AirDrumAPIer{}
	vm_res := resource_vm()
	d := vm_res.TestResourceData()
	d.SetId("UnitTest vm1")
	d.Set("name", "Unit test vm")
	var (
		sewan *API
		err   error
	)
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		fake_client_tooler.Client = test_case.TC_clienter
		err = Apier.Update_vm_resource(d, &fake_client_tooler, sewan)
		switch {
		case err == nil || test_case.Update_Err == nil:
			if !(err == nil && test_case.Update_Err == nil) {
				t.Errorf("TC %d : VM read error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Update_Err)
			}
		case err.Error() != test_case.Update_Err.Error():
			t.Errorf("TC %d : VM read error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Update_Err.Error())
		}
	}
}

//------------------------------------------------------------------------------
func TestDelete_vm_resource(t *testing.T) {
	test_cases := []struct {
		Id          int
		TC_clienter Clienter
		Delete_Err  error
	}{
		{
			1,
			ErrorResponse_HttpClienter{},
			errors.New(REQ_ERR),
		},
		{
			2,
			BadBodyResponse_StatusOK_HttpClienter{},
			errors.New("Read of \"Unit test vm\" failed, response body json " +
				"error :\n\r\"invalid character '\"' after object key"),
		},
		{
			3,
			Error401_HttpClienter{},
			errors.New(UNAUTHORIZED_MSG),
		},
		{
			4,
			Error404_HttpClienter{},
			errors.New(NOT_FOUND_MSG),
		},
		{
			5,
			CheckRedirectReqFailure_HttpClienter{},
			errors.New("Deletion of \"Unit test vm\" state failed, response reception " +
				"error : CheckRedirectReqFailure"),
		},
		{
			6,
			DeleteSuccess_HttpClienter{},
			nil,
		},
		{
			7,
			DeleteWRONGResponseBody_HttpClienter{},
			errors.New(DESTROY_WRONG_MSG),
		},
	}
	Apier := AirDrumAPIer{}
	vm_res := resource_vm()
	d := vm_res.TestResourceData()
	d.SetId("UnitTest vm1")
	d.Set("name", "Unit test vm")
	var (
		sewan *API
		err   error
	)
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		fake_client_tooler.Client = test_case.TC_clienter
		err = Apier.Delete_vm_resource(d, &fake_client_tooler, sewan)
		switch {
		case err == nil || test_case.Delete_Err == nil:
			if !(err == nil && test_case.Delete_Err == nil) {
				t.Errorf("TC %d : VM read error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Delete_Err)
			}
		case err.Error() != test_case.Delete_Err.Error():
			t.Errorf("TC %d : VM read error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Delete_Err.Error())
		}
	}
}

//------------------------------------------------------------------------------
func TestVmInstanceCreate(t *testing.T) {
	//slice elements "disks" and "nics" not tested, ref=TD-35489-UT-35737-1
	vm_res := resource_vm()
	d := vm_res.TestResourceData()
	var (
		vmInstance       VM
		res_data_value   interface{}
		vmInstance_value interface{}
	)

	d.SetId("UnitTest vm1")
	d.Set("name", "Unit test vm")
	d.Set("state", "UP")
	d.Set("os", "Debian")
	d.Set("ram", "4")
	d.Set("cpu", "2")
	d.Set("disks", nil)
	d.Set("nics", nil)
	d.Set("vdc", "vdc1")
	d.Set("boot", "on disk")
	d.Set("vdc_resource_disk", "vdc_resource_disk")
	//d.Get("template","")
	d.Set("slug", "slug")
	d.Set("token", "424242")
	d.Set("backup", "backup_no_backup")
	d.Set("disk_image", "disk img")
	d.Set("platform_name", "plateforme name")
	d.Set("backup_size", "42")
	d.Set("comment", "42")
	d.Set("outsourcing", "false")
	d.Set("dynamic_field", "42")

	vmInstance = vmInstanceCreate(d)
	val := reflect.ValueOf(vmInstance)

	for i := 0; i < val.Type().NumField(); i++ {
		switch value_type := val.Field(i).Kind(); value_type {
		case reflect.String:
			res_data_value = d.Get(val.Type().Field(i).Tag.Get("json")).(string)
			vmInstance_value = val.Field(i).Interface().(string)
			if res_data_value != vmInstance_value {
				t.Errorf("VM instance was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"",
					vmInstance_value, res_data_value)
			}
		case reflect.Slice:
			//slice elements "disks" and "nics" not tested, ref=TD-35489-UT-35737-1
		}
	}
}
