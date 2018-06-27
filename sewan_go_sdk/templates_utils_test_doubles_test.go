package sewan_go_sdk

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type TemplaterDummy struct{}

func (templaterFake TemplaterDummy) FetchTemplateFromList(template_name string,
  templateList []interface{}) (map[string]interface{}, error) {

  return nil,nil
}
func (templaterFake TemplaterDummy) UpdateSchema(d *schema.ResourceData,
  template map[string]interface{},
  templatesTooler *TemplatesTooler) error {

  return nil
}
func (templaterFake TemplaterDummy) UpdateSchemaDisks(d *schema.ResourceData,
  disks []interface{}) error {

  return nil
}
func (templaterFake TemplaterDummy) UpdateSchemaNics(d *schema.ResourceData) error {

  return nil
}

//------------------------------------------------------------------------------
type EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP_TemplaterFake struct{}

func (templaterFake EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP_TemplaterFake) FetchTemplateFromList(template_name string,
  templateList []interface{}) (map[string]interface{}, error) {

  return map[string]interface{}{
    "id":         82,
    "name":       "centos7-rd-DC1",
    "slug":       "centos7-rd-dc1",
    "ram":        1,
    "cpu":        1,
    "os":         "CentOS",
    "enterprise": "sewan-rd-cloud-beta",
    "disks": []interface{}{
      map[string]interface{}{"name": "disk-centos7-rd-DC1-1",
        "size":          20,
        "storage_class": "storage_enterprise",
        "slug":          "disk-centos7-rd-dc1-1",
      },
    },
    "datacenter": "dc1",
    "nics": []interface{}{
      map[string]interface{}{"vlan": "sewanrd-mgt-th3",
        "mac_address": "00:50:56:21:7c:ab",
        "connected":   true,
      },
      map[string]interface{}{"vlan": "sewanrd-priv-th3",
        "mac_address": "00:50:56:21:7c:ac",
        "connected":   true,
      },
    },
    "login":         "",
    "password":      "",
    "dynamic_field": "",
  },nil
}
func (templaterFake EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP_TemplaterFake) UpdateSchema(d *schema.ResourceData,
  template map[string]interface{},
  templatesTooler *TemplatesTooler) error {

  d.Set("name","Unit test template no disc add on vm resource")
  d.Set("enterprise","sewan-rd-cloud-beta")
  d.Set("template","centos7-rd-DC1")
  d.Set("ram",1)
  d.Set("cpu",1)
  d.Set("disks",
    []interface{}{
    map[string]interface{}{"name": "disk-centos7-rd-DC1-1",
      "size":          20,
      "storage_class": "storage_enterprise",
      "slug":          "disk-centos7-rd-dc1-1",
    },
  },
  )
  return nil
}
func (templaterFake EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP_TemplaterFake) UpdateSchemaDisks(d *schema.ResourceData,
  disks []interface{}) error {

  return nil
}
func (templaterFake EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP_TemplaterFake) UpdateSchemaNics(d *schema.ResourceData) error {

  return nil
}
