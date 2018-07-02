package sewan_go_sdk

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
	"reflect"
	"testing"
)

//------------------------------------------------------------------------------
func TestResourceInstanceCreate(t *testing.T) {
	test_cases := []struct {
		Id            int
		D             *schema.ResourceData
		TC_Clienter   Clienter
		TC_Templater  Templater
		Resource_type string
		Error         error
		VmInstance    interface{}
	}{
		{
			1,
			vm_schema_init(NO_TEMPLATE_VM_MAP),
			GetTemplatesList_Success_HttpClienterFake{},
			TemplaterDummy{},
			VM_RESOURCE_TYPE,
			nil,
			vmInstanceNO_TEMPLATE_VM_MAP(),
		},
		{
			2,
			vm_schema_init(EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP),
			GetTemplatesList_Success_HttpClienterFake{},
			EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP_TemplaterFake{},
			VM_RESOURCE_TYPE,
			nil,
			Fake_vmInstance_EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP(),
		},
		{
			3,
			vm_schema_init(EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP),
			GetTemplatesList_Success_HttpClienterFake{},
			EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP_TemplaterFake{},
			VM_RESOURCE_TYPE,
			nil,
			Fake_vmInstance_EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP(),
		},
		{
			4,
			vm_schema_init(EXISTING_TEMPLATE_WITH_MODIFIED_NIC_AND_DISK_VM_MAP),
			GetTemplatesList_Success_HttpClienterFake{},
			TemplaterDummy{},
			VM_RESOURCE_TYPE,
			nil,
			Fake_vmInstance_EXISTING_TEMPLATE_WITH_MODIFIED_NIC_AND_DISK_VM_MAP(),
		},
		{
			5,
			vm_schema_init(NON_EXISTING_TEMPLATE_VM_MAP),
			GetTemplatesList_Success_HttpClienterFake{},
			Unexisting_template_TemplaterFake{},
			VM_RESOURCE_TYPE,
			errors.New("Unavailable template : windows95"),
			VM{},
		},
		{
			6,
			vdc_schema_init(VDC_CREATION_MAP),
			nil,
			TemplaterDummy{},
			VDC_RESOURCE_TYPE,
			nil,
			Fake_vdcInstance_VDC_CREATION_MAP(),
		},
		{
			7,
			vdc_schema_init(VDC_CREATION_MAP),
			GetTemplatesList_Success_HttpClienterFake{},
			TemplaterDummy{},
			WRONG_RESOURCE_TYPE,
			errors.New("Resource of type \"a_non_supported_resource_type\" not supported," +
				"list of accepted resource types :\n\r" +
				"- \"vdc\"\n\r" +
				"- \"vm\""),
			nil,
		},
		{
			8,
			vm_schema_init(NON_EXISTING_TEMPLATE_VM_MAP),
			GetTemplatesList_Failure_HttpClienterFake{},
			Unexisting_template_TemplaterFake{},
			VM_RESOURCE_TYPE,
			errors.New("GetTemplatesList() error"),
			VM{},
		},
		{
			9,
			vm_schema_init(EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP),
			GetTemplatesList_Success_HttpClienterFake{},
			EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP_TemplaterFake{},
			VM_RESOURCE_TYPE,
			nil,
			Fake_vmInstance_EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP(),
		},
	}

	var (
		sewan    *API
		err      error = nil
		instance interface{}
	)
	api_tools := APITooler{
		Api: AirDrumResources_Apier{},
	}
	fake_client_tooler := ClientTooler{}
	fake_templates_tooler := TemplatesTooler{}
	fake_schema_tooler := SchemaTooler{
		SchemaTools: Schema_Schemaer{},
	}
	sewan = &API{Token: "42", URL: "42", Client: &http.Client{}}

	for _, test_case := range test_cases {
		fake_client_tooler.Client = test_case.TC_Clienter
		fake_templates_tooler.TemplatesTools = test_case.TC_Templater
		err, instance = api_tools.Api.ResourceInstanceCreate(test_case.D,
			&fake_client_tooler,
			&fake_templates_tooler,
			&fake_schema_tooler,
			test_case.Resource_type,
			sewan)
		switch {
		case err == nil || test_case.Error == nil:
			if !(err == nil && test_case.Error == nil) {
				t.Errorf("\n\nTC %d : ResourceInstanceCreate() error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, err, test_case.Error)
			} else {
				switch {
				case !reflect.DeepEqual(test_case.VmInstance, instance):
					t.Errorf("\n\nTC %d : Wrong ResourceInstanceCreate() created instance,"+
						"\n\rgot: \"%s\"\n\rwant: \"%s\"",
						test_case.Id, instance, test_case.VmInstance)
					v := reflect.ValueOf(instance)
					v2 := reflect.ValueOf(test_case.VmInstance)
					for i := 0; i < v.NumField(); i++ {
						if !reflect.DeepEqual(v.Field(i).Interface(), v2.Field(i).Interface()) {
							t.Log("Got field difference(s) :",
								"\ngot :", v.Field(i).Interface(),
								"(", reflect.TypeOf(v.Field(i).Interface()), ")",
								"\nwant :", v2.Field(i).Interface(),
								"(", reflect.TypeOf(v2.Field(i).Interface()), ")")
						}
					}
				}
			}
		case err != nil && test_case.Error != nil:
			switch {
			case !reflect.DeepEqual(instance, test_case.VmInstance):
				t.Errorf("\n\nTC %d : Wrong ResourceInstanceCreate() created instance,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, instance, test_case.VmInstance)
			case err.Error() != test_case.Error.Error():
				t.Errorf("\n\nTC %d : resource creation error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, err.Error(), test_case.Error.Error())
			}
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
				t.Errorf("\n\nTC %d : Validate_status() error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, apiClientErr, test_case.Err)
			}
		case apiClientErr.Error() != test_case.Err.Error():
			t.Errorf("\n\nTC %d : Validate_status() error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, apiClientErr.Error(), test_case.Err.Error())
		}
	}
}

func Create_test_resource_schema(id interface{}) *schema.ResourceData {
	vm_res := resource_vm()
	d := vm_res.TestResourceData()
	d.SetId(id.(string))
	return d
}

func TestDelete_terraform_resource(t *testing.T) {
	d := Create_test_resource_schema("resource to delete")
	schemaTooler := SchemaTooler{
		SchemaTools: Schema_Schemaer{},
	}
	schemaTooler.SchemaTools.Delete_terraform_resource(d)
	if d.Id() != "" {
		t.Errorf("Deletion of unit test resource failed.")
	}
}

func TestUpdate_local_resource_state_AND_Read_element(t *testing.T) {
	test_cases := []struct {
		Id           int
		Vm_map       map[string]interface{}
		Vm_Id_string string
	}{
		{
			1,
			TEST_UPDATE_VM_MAP,
			"unit test vm",
		},
		{
			2,
			TEST_UPDATE_VM_MAP_FLOATID,
			"121212.12",
		},
		{
			3,
			TEST_UPDATE_VM_MAP_INTID,
			"1212",
		},
	}
	var d *schema.ResourceData
	schemaTooler := SchemaTooler{
		SchemaTools: Schema_Schemaer{},
	}
	for _, test_case := range test_cases {
		d = Create_test_resource_schema(test_case.Vm_Id_string)
		schemaTooler.SchemaTools.Update_local_resource_state(test_case.Vm_map,
			d,
			&schemaTooler)
		for key, value := range test_case.Vm_map {
			if key != "id" {
				if !reflect.DeepEqual(d.Get(key), value) {
					t.Errorf("\n\nTC %d : Update of %s field failed :\n\rGot :%s\n\rWant :%s",
						test_case.Id, key, d.Get(key), value)
				}
			} else {
				if d.Id() != test_case.Vm_Id_string {
					t.Errorf("\n\nTC %d : Update of Id reserved field failed :\n\rGot :%s\n\rWant :%s",
						test_case.Id, d.Id(), test_case.Vm_Id_string)
				}
			}
		}
	}
}
