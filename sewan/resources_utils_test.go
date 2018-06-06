package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
	"reflect"
	"testing"
)

//------------------------------------------------------------------------------
//--Structures init, interface implementation fakes, various test items etc.----
//------------------------------------------------------------------------------
var (
	TEST_UPDATE_VM_MAP = map[string]interface{}{
		"id":    "unit test vm",
		"name":  "Unit test vm",
		"state": "DOWN",
		"os":    "CentOS",
		"ram":   16,
		"cpu":   8,
		"disks": []interface{}{
			map[string]interface{}{
				"name":   "disk 1 update",
				"size":   42,
				"v_disk": "v_disk update",
				"slug":   "slug update",
			},
			map[string]interface{}{
				"name":   "disk 2 update",
				"size":   42,
				"v_disk": "v_disk update",
				"slug":   "slug update",
			},
		},
		"nics": []interface{}{
			map[string]interface{}{
				"vlan":       "vlan 1 update",
				"mac_adress": "42",
				"connected":  false,
			},
		},
		"vdc":               "vdc update",
		"boot":              "on disk update",
		"vdc_resource_disk": "vdc_disk update", //"template":"template name",
		"slug":              "42 update",
		"token":             "424242 update",
		"backup":            "backup-no_backup update",
		"disk_image":        " update",
		"platform_name":     "",
		"backup_size":       "42 update",
		"comment":           "",
		"outsourcing":       "42 update",
		"dynamic_field":     "42 update",
	}
	TEST_UPDATE_VM_MAP_INTID = map[string]interface{}{
		"id":    1212,
		"name":  "Unit test vm",
		"state": "DOWN",
		"os":    "CentOS",
		"ram":   16,
		"cpu":   8,
		"disks": []interface{}{
			map[string]interface{}{
				"name":   "disk 1 update",
				"size":   42,
				"v_disk": "v_disk update",
				"slug":   "slug update",
			},
			map[string]interface{}{
				"name":   "disk 2 update",
				"size":   42,
				"v_disk": "v_disk update",
				"slug":   "slug update",
			},
		},
		"nics": []interface{}{
			map[string]interface{}{
				"vlan":       "vlan 1 update",
				"mac_adress": "42",
				"connected":  false,
			},
		},
		"vdc":               "vdc update",
		"boot":              "on disk update",
		"vdc_resource_disk": "vdc_disk update", //"template":"template name",
		"slug":              "42 update",
		"token":             "424242 update",
		"backup":            "backup-no_backup update",
		"disk_image":        " update",
		"platform_name":     "",
		"backup_size":       "42 update",
		"comment":           "",
		"outsourcing":       "42 update",
		"dynamic_field":     "42 update",
	}
	TEST_UPDATE_VM_MAP_FLOATID = map[string]interface{}{
		"id":    121212.12,
		"name":  "Unit test vm",
		"state": "DOWN",
		"os":    "CentOS",
		"ram":   16,
		"cpu":   8,
		"disks": []interface{}{
			map[string]interface{}{
				"name":   "disk 1 update",
				"size":   42,
				"v_disk": "v_disk update",
				"slug":   "slug update",
			},
			map[string]interface{}{
				"name":   "disk 2 update",
				"size":   42,
				"v_disk": "v_disk update",
				"slug":   "slug update",
			},
		},
		"nics": []interface{}{
			map[string]interface{}{
				"vlan":       "vlan 1 update",
				"mac_adress": "42",
				"connected":  false,
			},
		},
		"vdc":               "vdc update",
		"boot":              "on disk update",
		"vdc_resource_disk": "vdc_disk update", //"template":"template name",
		"slug":              "42 update",
		"token":             "424242 update",
		"backup":            "backup-no_backup update",
		"disk_image":        " update",
		"platform_name":     "",
		"backup_size":       "42 update",
		"comment":           "",
		"outsourcing":       "42 update",
		"dynamic_field":     "42 update",
	}
)

func Create_test_resource_schema(id interface{}) *schema.ResourceData {
	vm_res := resource_vm()
	d := vm_res.TestResourceData()

	d.SetId(id.(string))
	d.Set("name", "Unit test vm")
	d.Set("state", "UP")
	d.Set("os", "Debian")
	d.Set("ram", 4)
	d.Set("cpu", 2)
	d.Set("disks", []interface{}{
		map[string]interface{}{
			"name":   "disk 1 update",
			"size":   24,
			"v_disk": "v_disk update",
			"slug":   "slug update",
		},
	})
	d.Set("nics", []interface{}{
		map[string]interface{}{
			"vlan":       "vlan 1",
			"mac_adress": "24",
			"connected":  true,
		},
		map[string]interface{}{
			"vlan":       "vlan 2 update",
			"mac_adress": "24",
			"connected":  true,
		},
	})
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

	return d
}

//------------------------------------------------------------------------------
//-------------Units tests------------------------------------------------------
//------------------------------------------------------------------------------
func TestDelete_resource(t *testing.T) {
	d := Create_test_resource_schema("resource to delete")
	Delete_resource(d)
	if d.Id() != "" {
		t.Errorf("Deletion of unit test resource failed.")
	}
}

func TestUpdate_local_resource_state_AND_read_element(t *testing.T) {
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

	for _, test_case := range test_cases {
		d = Create_test_resource_schema(test_case.Vm_Id_string)
		_ = Update_local_resource_state(test_case.Vm_map, d)
		for key, value := range test_case.Vm_map {
			if key != "id" {
				if !reflect.DeepEqual(d.Get(key), value) {
					t.Errorf("TC %d : Update of field failed :\n\rGot :%s\n\rWant :%s",
						test_case.Id, d.Get(key), value)
				}
			} else {
				if d.Id() != test_case.Vm_Id_string {
					t.Errorf("TC %d : Update of Id reserved field failed :\n\rGot :%s\n\rWant :%s",
						test_case.Id, d.Id(), test_case.Vm_Id_string)
				}
			}
		}
	}
}
