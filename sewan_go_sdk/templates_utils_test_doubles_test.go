package sewan_go_sdk

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
)

//------------------------------------------------------------------------------
type TemplaterDummy struct{}

func (templaterFake TemplaterDummy) FetchTemplateFromList(template_name string,
	templateList []interface{}) (map[string]interface{}, error) {

	return nil, nil
}
func (templaterFake TemplaterDummy) UpdateSchemaFromTemplate(d *schema.ResourceData,
	template map[string]interface{},
	templatesTooler *TemplatesTooler,
	schemaTooler *SchemaTooler) error {

	return nil
}
func (templaterFake TemplaterDummy) UpdateSchemaDisksFromTemplateDisks(d *schema.ResourceData,
	disks []interface{},
	schemaTooler *SchemaTooler) error {

	return nil
}
func (templaterFake TemplaterDummy) UpdateSchemaDisksFromTemplateNics(d *schema.ResourceData) error {

	return nil
}

//------------------------------------------------------------------------------
type Unexisting_template_TemplaterFake struct{}

func (templaterFake Unexisting_template_TemplaterFake) FetchTemplateFromList(template_name string,
	templateList []interface{}) (map[string]interface{}, error) {

	return nil, errors.New("Unavailable template : windows95")
}
func (templaterFake Unexisting_template_TemplaterFake) UpdateSchemaFromTemplate(d *schema.ResourceData,
	template map[string]interface{},
	templatesTooler *TemplatesTooler,
	schemaTooler *SchemaTooler) error {

	return nil
}
func (templaterFake Unexisting_template_TemplaterFake) UpdateSchemaDisksFromTemplateDisks(d *schema.ResourceData,
	disks []interface{},
	schemaTooler *SchemaTooler) error {

	return nil
}
func (templaterFake Unexisting_template_TemplaterFake) UpdateSchemaDisksFromTemplateNics(d *schema.ResourceData) error {

	return nil
}

//------------------------------------------------------------------------------
type EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP_TemplaterFake struct{}

func (templaterFake EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP_TemplaterFake) FetchTemplateFromList(template_name string,
	templateList []interface{}) (map[string]interface{}, error) {

	return map[string]interface{}{
		"id":         82,
		"name":       "template1",
		"slug":       "centos7-rd-dc1",
		"ram":        1,
		"cpu":        1,
		"os":         "CentOS",
		"enterprise": "unit test enterprise",
		"disks": []interface{}{
			map[string]interface{}{"name": "template1 disk1",
				"size":          20,
				"storage_class": "storage_enterprise",
				"slug":          "template1 disk1 slug",
			},
		},
		"datacenter": "dc1",
		"nics": []interface{}{
			map[string]interface{}{"vlan": "unit test vlan1",
				"mac_address": "00:50:56:21:7c:ab",
				"connected":   true,
			},
			map[string]interface{}{"vlan": "unit test vlan2",
				"mac_address": "00:50:56:21:7c:ac",
				"connected":   true,
			},
		},
		"login":         "",
		"password":      "",
		"dynamic_field": "",
	}, nil
}
func (templaterFake EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP_TemplaterFake) UpdateSchemaFromTemplate(d *schema.ResourceData,
	template map[string]interface{},
	templatesTooler *TemplatesTooler,
	schemaTooler *SchemaTooler) error {

	d.Set("name", "Unit test template no disc add on vm resource")
	d.Set("enterprise", "unit test enterprise")
	d.Set("template", "template1")
	d.Set("ram", 1)
	d.Set("cpu", 1)
	d.Set("disks",
		[]interface{}{
			map[string]interface{}{"name": "template1 disk1",
				"size":          20,
				"storage_class": "storage_enterprise",
				"slug":          "template1 disk1 slug",
			},
		},
	)
	d.Set("nics", []interface{}{})
	return nil
}
func (templaterFake EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP_TemplaterFake) UpdateSchemaDisksFromTemplateDisks(d *schema.ResourceData,
	disks []interface{},
	schemaTooler *SchemaTooler) error {

	return nil
}
func (templaterFake EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP_TemplaterFake) UpdateSchemaDisksFromTemplateNics(d *schema.ResourceData) error {

	return nil
}

//------------------------------------------------------------------------------
type EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP_TemplaterFake struct{}

func (templaterFake EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP_TemplaterFake) FetchTemplateFromList(template_name string,
	templateList []interface{}) (map[string]interface{}, error) {

	return map[string]interface{}{
		"id":         82,
		"name":       "template1",
		"slug":       "centos7-rd-dc1",
		"ram":        1,
		"cpu":        1,
		"os":         "CentOS",
		"enterprise": "unit test enterprise",
		"disks": []interface{}{
			map[string]interface{}{"name": "template1 disk1",
				"size":          20,
				"storage_class": "storage_enterprise",
				"slug":          "template1 disk1 slug",
			},
		},
		"datacenter": "dc1",
		"nics": []interface{}{
			map[string]interface{}{"vlan": "unit test vlan1",
				"mac_address": "00:50:56:21:7c:ab",
				"connected":   true,
			},
			map[string]interface{}{"vlan": "unit test vlan2",
				"mac_address": "00:50:56:21:7c:ac",
				"connected":   true,
			},
		},
		"login":         "",
		"password":      "",
		"dynamic_field": "",
	}, nil
}
func (templaterFake EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP_TemplaterFake) UpdateSchemaFromTemplate(d *schema.ResourceData,
	template map[string]interface{},
	templatesTooler *TemplatesTooler,
	schemaTooler *SchemaTooler) error {

	d.Set("name", "EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP")
	d.Set("enterprise", "unit test enterprise")
	d.Set("template", "template1")
	d.Set("ram", 8)
	d.Set("cpu", 4)
	d.Set("disks",
		[]interface{}{
			map[string]interface{}{
				"name":          "disk 1",
				"size":          24,
				"storage_class": "storage_enterprise",
				"slug":          "",
			},
			map[string]interface{}{
				"name":          "template1 disk1",
				"size":          25,
				"storage_class": "storage_enterprise",
				"slug":          "template1 disk1 slug",
			},
		},
	)
	d.Set("nics", []interface{}{
		map[string]interface{}{
			"vlan":        "non template vlan 1",
			"mac_address": "00:21:21:21:21:21",
			"connected":   true,
		},
		map[string]interface{}{
			"vlan":        "non template vlan 2",
			"mac_address": "00:21:21:21:21:22",
			"connected":   true,
		},
	},
	)
	d.Set("vdc:          ", "vdc")
	d.Set("boot:         ", "on disk")
	d.Set("storage_class:", "storage_enterprise")
	d.Set("slug:         ", "42")
	d.Set("token:        ", "424242")
	d.Set("backup:       ", "backup_no_backup")
	d.Set("disk_image:   ", "")
	d.Set("platform_name:", "42")
	d.Set("backup_size:  ", 42)
	d.Set("comment:      ", "42")
	d.Set("outsourcing:  ", "42")
	d.Set("dynamic_field:", "42")
	return nil
}
func (templaterFake EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP_TemplaterFake) UpdateSchemaDisksFromTemplateDisks(d *schema.ResourceData,
	disks []interface{},
	schemaTooler *SchemaTooler) error {

	return nil
}
func (templaterFake EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP_TemplaterFake) UpdateSchemaDisksFromTemplateNics(d *schema.ResourceData) error {

	return nil
}

//------------------------------------------------------------------------------
type EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP_TemplaterFake struct{}

func (templaterFake EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP_TemplaterFake) FetchTemplateFromList(template_name string,
	templateList []interface{}) (map[string]interface{}, error) {

	return map[string]interface{}{
		"id":         82,
		"name":       "template1",
		"slug":       "template 1 slug",
		"ram":        1,
		"cpu":        1,
		"os":         "CentOS",
		"enterprise": "unit test enterprise",
		"disks": []interface{}{
			map[string]interface{}{"name": "template1 disk1",
				"size":          20,
				"storage_class": "storage_enterprise",
				"slug":          "template1 disk1 slug",
			},
		},
		"datacenter": "dc1",
		"nics": []interface{}{
			map[string]interface{}{"vlan": "unit test vlan1",
				"mac_address": "00:50:56:21:7c:ab",
				"connected":   true,
			},
			map[string]interface{}{"vlan": "unit test vlan2",
				"mac_address": "00:50:56:21:7c:ac",
				"connected":   true,
			},
		},
		"login":         "",
		"password":      "",
		"dynamic_field": "",
	}, nil
}
func (templaterFake EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP_TemplaterFake) UpdateSchemaFromTemplate(d *schema.ResourceData,
	template map[string]interface{},
	templatesTooler *TemplatesTooler,
	schemaTooler *SchemaTooler) error {

	d.Set("name", "EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP")
	d.Set("enterprise", "unit test enterprise")
	d.Set("template", "template1")
	d.Set("ram", 8)
	d.Set("cpu", 4)
	d.Set("disks",
		[]interface{}{
			map[string]interface{}{
				"name":          "disk 1",
				"size":          24,
				"storage_class": "storage_enterprise",
				"slug":          "",
			},
			map[string]interface{}{
				"name":          "template1 disk1",
				"size":          24,
				"storage_class": "storage_class",
				"deletion":      true,
			},
		},
	)
	d.Set("nics", []interface{}{})
	d.Set("vdc:          ", "vdc")
	d.Set("boot:         ", "on disk")
	d.Set("storage_class:", "storage_enterprise")
	d.Set("slug:         ", "42")
	d.Set("token:        ", "424242")
	d.Set("backup:       ", "backup_no_backup")
	d.Set("disk_image:   ", "")
	d.Set("platform_name:", "42")
	d.Set("backup_size:  ", 42)
	d.Set("dynamic_field:  ",
		"{\"terraform_provisioned\":true,\"creation_template\":\"template1\",\"disks_created_from_template\":null}")
	return nil
}
func (templaterFake EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP_TemplaterFake) UpdateSchemaDisksFromTemplateDisks(d *schema.ResourceData,
	disks []interface{},
	schemaTooler *SchemaTooler) error {

	return nil
}
func (templaterFake EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP_TemplaterFake) UpdateSchemaDisksFromTemplateNics(d *schema.ResourceData) error {

	return nil
}
