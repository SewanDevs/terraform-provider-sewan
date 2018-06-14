package sewan_go_sdk

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

//------------------------------------------------------------------------------
func TestvdcInstanceCreate(t *testing.T) {
	//slice elements "disks" and "nics" not tested, ref=TD-35489-UT-35737-1
	vdc_res := resource_vdc()
	d := vdc_res.TestResourceData()
	var (
		vdcInstance       VDC
		res_data_value    interface{}
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

//------------------------------------------------------------------------------
func TestvmInstanceCreate(t *testing.T) {
	//slice elements "disks" and "nics" not tested, ref=TD-35489-UT-35737-1
	resource_res := resource(VM_RESOURCE_TYPE)
	d := resource_res.TestResourceData()
	var (
		resourceInstance       VM
		res_data_value         interface{}
		resourceInstance_value interface{}
	)

	d.SetId("UnitTest resource1")
	d.Set("name", "Unit test resource")
	d.Set("template", "")
	d.Set("state", "UP")
	d.Set("os", "Debian")
	d.Set("ram", "4")
	d.Set("cpu", "2")
	d.Set("disks", nil)
	d.Set("nics", nil)
	d.Set("vdc", "vdc1")
	d.Set("boot", "on disk")
	d.Set("vdc_resource_disk", "vdc_resource_disk")
	d.Set("slug", "slug")
	d.Set("token", "424242")
	d.Set("backup", "backup_no_backup")
	d.Set("disk_image", "disk img")
	d.Set("platform_name", "plateforme name")
	d.Set("backup_size", "42")
	d.Set("comment", "42")
	d.Set("outsourcing", "false")
	d.Set("dynamic_field", "42")

	resourceInstance = vmInstanceCreate(d)
	val := reflect.ValueOf(resourceInstance)

	for i := 0; i < val.Type().NumField(); i++ {
		switch value_type := val.Field(i).Kind(); value_type {
		case reflect.String:
			res_data_value = d.Get(val.Type().Field(i).Tag.Get("json")).(string)
			resourceInstance_value = val.Field(i).Interface().(string)
			if res_data_value != resourceInstance_value {
				t.Errorf("resource instance was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"",
					resourceInstance_value, res_data_value)
			}
		case reflect.Slice:
			//slice elements "disks" and "nics" not tested, ref=TD-35489-UT-35737-1
		}
	}
}

//------------------------------------------------------------------------------
func TestGet_resource_url(t *testing.T) {
	test_cases := []struct {
		Id     int
		api    API
		vm_id  string
		vm_url string
	}{
		{1,
			API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			"42",
			RIGHT_VM_URL_42,
		},
		{2,
			API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			"PATATE",
			RIGHT_VM_URL_PATATE,
		},
	}

	api_tools := APITooler{
		Api: AirDrumResources_Apier{},
	}

	for _, test_case := range test_cases {
		s_vm_url := api_tools.Api.Get_resource_url(&test_case.api,
			VM_RESOURCE_TYPE,
			test_case.vm_id)

		switch {
		case s_vm_url != test_case.vm_url:
			t.Errorf("VM url was incorrect,\n\rgot: \"%s\"\n\rwant: \"%s\"",
				s_vm_url, test_case.vm_url)
		}
	}
}

//------------------------------------------------------------------------------
func TestGet_resource_creation_url(t *testing.T) {
	test_cases := []struct {
		Id                    int
		api                   API
		resource_creation_url string
	}{
		{1,
			API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			RIGHT_VM_CREATION_API_URL,
		},
	}

	api_tools := APITooler{
		Api: AirDrumResources_Apier{},
	}

	for _, test_case := range test_cases {
		s_resource_creation_url := api_tools.Api.Get_resource_creation_url(&test_case.api,
			VM_RESOURCE_TYPE)

		switch {
		case s_resource_creation_url != test_case.resource_creation_url:
			t.Errorf("resource api creation url was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				s_resource_creation_url, test_case.resource_creation_url)
		}
	}
}

//------------------------------------------------------------------------------
func TestValidate_status(t *testing.T) {
	test_cases := []struct {
		Id           int
		Api          API
		Err          error
		ResourceType string
	}{
		{1,
			API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			nil,
			VM_RESOURCE_TYPE,
		},
		{2,
			API{
				WRONG_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			errors.New("401 Unauthorized{\"detail\":\"Invalid token.\"}"),
			VM_RESOURCE_TYPE,
		},
		{3,
			API{
				RIGHT_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a proper json response from \"" +
				WRONG_API_URL + "\", the api is down or this url is wrong."),
			VM_RESOURCE_TYPE,
		},
		{4,
			API{
				WRONG_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a proper json response from \"" +
				WRONG_API_URL + "\", the api is down or this url is wrong."),
			VM_RESOURCE_TYPE,
		},
		{5,
			API{
				RIGHT_API_TOKEN,
				NO_RESP_BODY_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a response body from \"" +
				NO_RESP_BODY_API_URL + "\", the api is down or this url is wrong."),
			VM_RESOURCE_TYPE,
		},
		{6,
			API{
				RIGHT_API_TOKEN,
				NO_RESP_API_URL,
				&http.Client{},
			},
			errors.New("Could not get a response from \"" +
				NO_RESP_API_URL + "\", the api is down or this url is wrong."),
			VM_RESOURCE_TYPE,
		},
	}

	apiTooler := APITooler{
		Api: AirDrumResources_Apier{},
	}
	clientTooler := ClientTooler{
		Client: FakeHttpClienter{},
	}
	var apiClientErr error

	for _, test_case := range test_cases {
		apiClientErr = apiTooler.Api.Validate_status(&test_case.Api,
			test_case.ResourceType,
			clientTooler)

		switch {
		case apiClientErr == nil || test_case.Err == nil:
			if !(apiClientErr == nil && test_case.Err == nil) {
				t.Errorf("TC %d : Validate_status() error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, apiClientErr, test_case.Err)
			}
		case apiClientErr.Error() != test_case.Err.Error():
			t.Errorf("TC %d : Validate_status() error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, apiClientErr.Error(), test_case.Err.Error())
		}
	}
}
