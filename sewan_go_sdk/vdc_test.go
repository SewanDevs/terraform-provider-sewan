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

// Creation success *ClientTooler
type VDC_CreationSuccess_HttpClienter struct{}

func (client VDC_CreationSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{"Content-Type": {"application/json"}}
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
	resp.Header = map[string][]string{"Content-Type": {"application/json"}}
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
	resp.Header = map[string][]string{"Content-Type": {"application/json"}}
	resp.StatusCode = http.StatusOK
	js, _ := json.Marshal(TEST_VDC_CREATION_MAP)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

type VDC_DeleteSuccess_HttpClienter struct{}

func (client VDC_DeleteSuccess_HttpClienter) Do(api *API,
	req *http.Request) (*http.Response, error) {

	resp := http.Response{}
	resp.Header = map[string][]string{"Content-Type": {"application/json"}}
	resp.StatusCode = http.StatusOK
	body := Resp_Body{DESTROY_MSG}
	js, _ := json.Marshal(body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(js))
	return &resp, nil
}

//------------------------------------------------------------------------------
//-------------Units tests------------------------------------------------------
//------------------------------------------------------------------------------
func TestCreate_vdc_resource(t *testing.T) {
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
			errors.New("Creation of \"Unit test vdc\" failed, response body json " +
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
			VDC_CreationSuccess_HttpClienter{},
			nil,
			TEST_VDC_READ_RESPONSE_MAP,
		},
		{
			5,
			CheckRedirectReqFailure_HttpClienter{},
			errors.New("Creation of \"Unit test vdc\" failed, response reception " +
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
	vdc_res := resource_vdc()
	d := vdc_res.TestResourceData()
	d.SetId("UnitTest vdc1")
	d.Set("name", "Unit test vdc")
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		fake_client_tooler.Client = test_case.TC_clienter
		err, resp_creation_map = Apier.Create_vdc_resource(d, &fake_client_tooler, sewan)
		t.Log("------")
		t.Log("TC ", test_case.Id)
		switch {
		case err == nil || test_case.Creation_Err == nil:
			if !(err == nil && test_case.Creation_Err == nil) {
				t.Errorf("TC %d : vdc creation error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Creation_Err)
			} else {
				switch {
				case !reflect.DeepEqual(test_case.Created_resource, resp_creation_map):
					t.Log("resp_creation_map :", resp_creation_map)
					t.Log("test_case.Created_resource :", test_case.Created_resource)
					t.Log("!reflect.DeepEqual(test_case.Created_resource, resp_creation_map)",
						!reflect.DeepEqual(test_case.Created_resource, resp_creation_map))
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
		case err.Error() != test_case.Creation_Err.Error():
			t.Errorf("TC %d : vdc creation error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Creation_Err.Error())
		case !reflect.DeepEqual(test_case.Created_resource, resp_creation_map):
			t.Log("resp_creation_map :", resp_creation_map)
			t.Log("test_case.Created_resource :", test_case.Created_resource)
			t.Log("!reflect.DeepEqual(test_case.Created_resource, resp_creation_map)",
				!reflect.DeepEqual(test_case.Created_resource, resp_creation_map))
			t.Errorf("TC %d : Wrong created resource map,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, resp_creation_map, test_case.Created_resource)
		}
	}
}

//------------------------------------------------------------------------------
func TestRead_vdc_resource(t *testing.T) {
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
			errors.New("Read of \"Unit test vdc\" failed, response body json " +
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
			nil,
			nil,
			false,
		},
		{
			5,
			CheckRedirectReqFailure_HttpClienter{},
			errors.New("Read of \"Unit test vdc\" state failed, response reception " +
				"error : CheckRedirectReqFailure"),
			nil,
			true,
		},
		{
			6,
			VDC_ReadSuccess_HttpClienter{},
			nil,
			TEST_VDC_READ_RESPONSE_MAP,
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
	vdc_res := resource_vdc()
	d := vdc_res.TestResourceData()
	d.SetId("UnitTest vdc1")
	d.Set("name", "Unit test vdc")
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		fake_client_tooler.Client = test_case.TC_clienter
		err, resp_creation_map, res_exists = Apier.Read_vdc_resource(d, &fake_client_tooler, sewan)
		t.Log("--")
		switch {
		case err == nil || test_case.Read_Err == nil && !(err == nil && test_case.Read_Err == nil):
			t.Log("1")
			if !(err == nil && test_case.Read_Err == nil) {
				t.Log("1.1")
				t.Errorf("TC %d : vdc read error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Read_Err)
			} else {
				switch {
				case res_exists != test_case.Resource_exists:
					t.Log("4")
					t.Errorf("TC %d : Wrong read vdc exists value"+
						"\n\rgot: \"%v\"\n\rwant: \"%v\"",
						test_case.Id, res_exists, test_case.Resource_exists)
				case !reflect.DeepEqual(test_case.Read_resource, resp_creation_map):
					t.Log("5")
					t.Errorf("TC %d : Wrong vdc read resource map,"+
						"\n\rgot: \"%s\"\n\rwant: \"%s\"",
						test_case.Id, resp_creation_map, test_case.Read_resource)
				}
			}
		case err != nil && test_case.Read_Err != nil:
			t.Log("2")
			if resp_creation_map != nil {
				t.Log("2.2")
				t.Errorf("TC %d : Wrong created resource map,"+
					" it should be nil as error is not nil,"+
					"\n\rgot map: \n\r\"%s\"\n\rwant map: \n\r\"%s\"\n\r",
					test_case.Id, resp_creation_map, test_case.Read_resource)
			}
		case err.Error() != test_case.Read_Err.Error():
			t.Log("3")
			t.Errorf("TC %d : vdc read error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Read_Err.Error())
		case res_exists != test_case.Resource_exists:
			t.Log("4")
			t.Errorf("TC %d : Wrong read vdc exists value"+
				"\n\rgot: \"%v\"\n\rwant: \"%v\"",
				test_case.Id, res_exists, test_case.Resource_exists)
		case !reflect.DeepEqual(test_case.Read_resource, resp_creation_map):
			t.Log("5")
			t.Errorf("TC %d : Wrong vdc read resource map,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, resp_creation_map, test_case.Read_resource)
		}
	}
}

//------------------------------------------------------------------------------
func TestUpdate_vdc_resource(t *testing.T) {
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
			errors.New("Update of \"Unit test vdc\" failed, response body json " +
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
			errors.New("Update of \"Unit test vdc\" state failed, response reception " +
				"error : CheckRedirectReqFailure"),
		},
		{
			6,
			VDC_UpdateSuccess_HttpClienter{},
			nil,
		},
	}
	Apier := AirDrumAPIer{}
	vdc_res := resource_vdc()
	d := vdc_res.TestResourceData()
	d.SetId("UnitTest vdc1")
	d.Set("name", "Unit test vdc")
	var (
		sewan *API
		err   error
	)
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		fake_client_tooler.Client = test_case.TC_clienter
		err = Apier.Update_vdc_resource(d, &fake_client_tooler, sewan)
		t.Log("--TC ",test_case.Id)
		t.Log("err", err)
		switch {
		case err == nil || test_case.Update_Err == nil:
			if !(err == nil && test_case.Update_Err == nil) {
				t.Errorf("TC %d : vdc update error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Update_Err)
			}
		case err.Error() != test_case.Update_Err.Error():
			t.Errorf("TC %d : vdc update error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Update_Err.Error())
		}
	}
}

//------------------------------------------------------------------------------
func TestDelete_vdc_resource(t *testing.T) {
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
			errors.New("Read of \"Unit test vdc\" failed, response body json " +
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
			errors.New("Deletion of \"Unit test vdc\" state failed, response reception " +
				"error : CheckRedirectReqFailure"),
		},
		{
			6,
			VDC_DeleteSuccess_HttpClienter{},
			nil,
		},
		{
			7,
			DeleteWRONGResponseBody_HttpClienter{},
			errors.New(DESTROY_WRONG_MSG),
		},
	}
	Apier := AirDrumAPIer{}
	vdc_res := resource_vdc()
	d := vdc_res.TestResourceData()
	d.SetId("UnitTest vdc1")
	d.Set("name", "Unit test vdc")
	var (
		sewan *API
		err   error
	)
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}
	fake_client_tooler := ClientTooler{}

	for _, test_case := range test_cases {
		fake_client_tooler.Client = test_case.TC_clienter
		err = Apier.Delete_vdc_resource(d, &fake_client_tooler, sewan)
		switch {
		case err == nil || test_case.Delete_Err == nil:
			if !(err == nil && test_case.Delete_Err == nil) {
				t.Errorf("TC %d : vdc read error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"", test_case.Id, err, test_case.Delete_Err)
			}
		case err.Error() != test_case.Delete_Err.Error():
			t.Errorf("TC %d : vdc read error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Delete_Err.Error())
		}
	}
}

//------------------------------------------------------------------------------
func TestvdcInstanceCreate(t *testing.T){
	//slice elements "disks" and "nics" not tested, ref=TD-35489-UT-35737-1
	vdc_res := resource_vdc()
	d := vdc_res.TestResourceData()
	var (
		vdcInstance VDC
		res_data_value   interface{}
		vdcInstance_value interface{}
	)

	d.SetId("UnitTest vdc1")
	d.Set("enterprise", "enterprise")
	d.Set("datacenter", "datacenter")
	d.Set("vdc_resources", nil)
	d.Set("slug", "slug")
	d.Set("dynamic_field", "42")

	vdcInstance = vdcInstanceCreate(d)
	val := reflect.ValueOf(vdcInstance)

	for i := 0; i < val.Type().NumField(); i++ {
		switch value_type := val.Field(i).Kind(); value_type {
		case reflect.String:
			res_data_value = d.Get(val.Type().Field(i).Tag.Get("json")).(string)
			vdcInstance_value = val.Field(i).Interface().(string)
			if res_data_value != vdcInstance_value {
				t.Errorf("vdc instance was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"",
					vdcInstance_value, res_data_value)
			}
		case reflect.Slice:
			//slice elements "disks" and "nics" not tested, ref=TD-35489-UT-35737-1
		}
	}
}
