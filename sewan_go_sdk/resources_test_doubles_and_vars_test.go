package sewan_go_sdk

import (
	"bytes"
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
)

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
	ENTERPRISE_SLUG         = "unit test enterprise"
)

var (
	VDC_CREATION_MAP = map[string]interface{}{
		"name":       "Unit test vdc resource",
		"enterprise": "unit test enterprise",
		"datacenter": "dc1",
		"vdc_resources": []interface{}{
			map[string]interface{}{
				"resource": "ram",
				"total":    20,
			},
			map[string]interface{}{
				"resource": "cpu",
				"total":    1,
			},
			map[string]interface{}{
				"resource": "storage_enterprise",
				"total":    10,
			},
			map[string]interface{}{
				"resource": "storage_performance",
				"total":    10,
			},
			map[string]interface{}{
				"resource": "storage_high_performance",
				"total":    10,
			},
		},
	}
	VDC_READ_RESPONSE_MAP = map[string]interface{}{
		"name":       "Unit test vdc",
		"enterprise": "unit test enterprise",
		"datacenter": "dc1",
		"vdc_resources": []interface{}{
			map[string]interface{}{
				"resource": "ram",
				"used":     "0",
				"total":    "20",
				"slug":     "unit test enterprise-dc1-vdc_te-ram",
			},
			map[string]interface{}{
				"resource": "cpu",
				"used":     "0",
				"total":    "1",
				"slug":     "unit test enterprise-dc1-vdc_te-cpu",
			},
			map[string]interface{}{
				"resource": "storage_enterprise",
				"used":     "0",
				"total":    "10",
				"slug":     "unit test enterprise-dc1-vdc_te-storage_enterprise",
			},
			map[string]interface{}{
				"resource": "storage_performance",
				"used":     "0",
				"total":    "10",
				"slug":     "unit test enterprise-dc1-vdc_te-storage_performance",
			},
			map[string]interface{}{
				"resource": "storage_high_performance",
				"used":     "0",
				"total":    "10",
				"slug":     "unit test enterprise-dc1-vdc_te-storage_high_performance",
			},
		},
		"slug":          "unit test enterprise-dc1-vdc_te",
		"dynamic_field": "",
	}
	NO_TEMPLATE_VM_MAP = map[string]interface{}{
		"name":  "Unit test no template vm resource",
		"state": "UP",
		"os":    "Debian",
		"ram":   8,
		"cpu":   4,
		"disks": []interface{}{
			map[string]interface{}{
				"name":          "disk 1",
				"size":          24,
				"storage_class": "storage_class",
			},
		},
		"nics": []interface{}{
			map[string]interface{}{
				"vlan":        "vlan 1 update",
				"mac_address": "00:21:21:21:21:21",
				"connected":   true,
			},
			map[string]interface{}{
				"vlan":        "vlan 2",
				"mac_address": "00:21:21:21:21:21",
				"connected":   true,
			},
		},
		"vdc":           "vdc",
		"boot":          "on disk",
		"storage_class": "storage_enterprise",
		"slug":          "42",
		"token":         "424242",
		"backup":        "backup_no_backup",
		"disk_image":    "",
		"platform_name": "42",
		"backup_size":   42,
		"comment":       "42",
	}
	EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP = map[string]interface{}{
		"name":          "Unit test template no disc add on vm resource",
		"enterprise":    "unit test enterprise",
		"template":      "template1",
		"state":         "UP",
		"vdc":           "vdc",
		"boot":          "on disk",
		"storage_class": "storage_enterprise",
		"slug":          "42",
		"token":         "424242",
		"backup":        "backup_no_backup",
	}
	EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP = map[string]interface{}{
		"name":       "EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP",
		"enterprise": "unit test enterprise",
		"template":   "template1",
		"state":      "UP",
		"os":         "Debian",
		"ram":        8,
		"cpu":        4,
		"disks": []interface{}{
			map[string]interface{}{
				"name":          "disk 1",
				"size":          24,
				"storage_class": "storage_class",
			},
		},
		"nics": []interface{}{
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
		"vdc":           "vdc",
		"boot":          "on disk",
		"storage_class": "storage_enterprise",
		"slug":          "42",
		"token":         "424242",
		"backup":        "backup_no_backup",
		"disk_image":    "",
		"platform_name": "42",
		"backup_size":   42,
		"comment":       "42",
	}
	EXISTING_TEMPLATE_WITH_MODIFIED_NIC_AND_DISK_VM_MAP = map[string]interface{}{
		"name":       "EXISTING_TEMPLATE_WITH_MODIFIED_NIC_AND_DISK_VM_MAP",
		"enterprise": "unit test enterprise",
		"template":   "template1",
		"state":      "UP",
		"os":         "Debian",
		"ram":        8,
		"cpu":        4,
		"disks": []interface{}{
			map[string]interface{}{
				"name":          "template1 disk1",
				"size":          24,
				"storage_class": "storage_class",
				"slug":          "template1 disk1 slug",
			},
		},
		"nics": []interface{}{
			map[string]interface{}{
				"mac_address": "00:21:21:21:21:22",
				"connected":   true,
				"vlan":        "unit test vlan1",
			},
			map[string]interface{}{
				"vlan":        "unit test vlan2",
				"mac_address": "00:21:21:21:21:21",
				"connected":   true,
			},
		},
		"vdc":           "vdc",
		"boot":          "on disk",
		"storage_class": "storage_enterprise",
		"slug":          "42",
		"token":         "424242",
		"backup":        "backup_no_backup",
		"disk_image":    "",
		"platform_name": "42",
		"backup_size":   42,
		"comment":       "42",
	}
	EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP = map[string]interface{}{
		"id":         "EXISTING_TEMPLATE_AND_VM_INSTANCE_WITH_DELETED_DISK_VM_MAP",
		"name":       "EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP",
		"enterprise": "unit test enterprise",
		"template":   "template1",
		"state":      "UP",
		"os":         "Debian",
		"ram":        8,
		"cpu":        4,
		"disks": []interface{}{
			map[string]interface{}{
				"name":          "template1 disk1",
				"size":          24,
				"storage_class": "storage_class",
				"deletion":      true,
			},
			map[string]interface{}{
				"name":          "disk 1",
				"size":          24,
				"storage_class": "storage_class",
			},
		},
		"nics":          []interface{}{},
		"vdc":           "vdc",
		"boot":          "on disk",
		"storage_class": "storage_enterprise",
		"slug":          "42",
		"token":         "424242",
		"backup":        "backup_no_backup",
		"disk_image":    "",
		"platform_name": "42",
		"backup_size":   42,
		"comment":       "42",
		"dynamic_field": "{\"terraform_provisioned\":true,\"creation_template\":\"template1\",\"disks_created_from_template\":null}",
	}
	NON_EXISTING_TEMPLATE_VM_MAP = map[string]interface{}{
		"name":          "windows95 vm",
		"enterprise":    "unit test enterprise",
		"template":      "windows95",
		"state":         "UP",
		"ram":           8,
		"cpu":           4,
		"vdc":           "vdc",
		"boot":          "on disk",
		"storage_class": "storage_enterprise",
		"slug":          "42",
		"token":         "424242",
		"backup":        "backup_no_backup",
		"disk_image":    "",
	}
	TEMPLATES_LIST = []interface{}{
		map[string]interface{}{
			"id":         40,
			"name":       "template2",
			"slug":       "unit test disk goulouglougoulouglou",
			"ram":        1,
			"cpu":        1,
			"os":         "CentOS",
			"enterprise": "unit test enterprise",
			"disks": []interface{}{
				map[string]interface{}{
					"name":          "unit test disk goulouglouglou",
					"size":          20,
					"storage_class": "storage_enterprise",
					"slug":          "unit test disk goulouglou slug",
				},
			},
			"datacenter":    "dc2",
			"nics":          []interface{}{},
			"login":         "",
			"password":      "",
			"dynamic_field": "",
		},
		map[string]interface{}{
			"id":         82,
			"name":       "template1",
			"slug":       "template1 slug",
			"ram":        1,
			"cpu":        1,
			"os":         "CentOS",
			"enterprise": "unit test enterprise",
			"disks": []interface{}{
				map[string]interface{}{
					"name":          "unit test disk template1",
					"size":          20,
					"storage_class": "storage_enterprise",
					"slug":          "unit test disk slug",
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
		},
		map[string]interface{}{
			"id":         41,
			"name":       "template3",
			"slug":       "unit test template3 slug",
			"ram":        1,
			"cpu":        1,
			"os":         "Debian",
			"enterprise": "unit test enterprise",
			"disks": []interface{}{
				map[string]interface{}{
					"name":          "unit test disk2",
					"size":          20,
					"storage_class": "storage_enterprise",
					"slug":          "unit test disk slug 2",
				},
			},
			"datacenter":    "dc2",
			"nics":          []interface{}{},
			"login":         "",
			"password":      "",
			"dynamic_field": "",
		},
		map[string]interface{}{
			"id":         43,
			"name":       "template4",
			"slug":       "tpl-centos7-rd",
			"ram":        1,
			"cpu":        1,
			"os":         "CentOS",
			"enterprise": "unit test enterprise",
			"disks": []interface{}{
				map[string]interface{}{
					"name":          "unit test disk 1",
					"size":          20,
					"storage_class": "storage_enterprise",
					"slug":          "unit test disk slug",
				},
			},
			"datacenter": "dc2",
			"nics": []interface{}{
				map[string]interface{}{
					"vlan":        "unit test vlan1",
					"mac_address": "00:50:56:00:00:23",
					"connected":   true,
				},
				map[string]interface{}{
					"vlan":        "unit test vlan2",
					"mac_address": "00:50:56:00:00:24",
					"connected":   true,
				},
			},
			"login":         nil,
			"password":      nil,
			"dynamic_field": nil,
		},
		map[string]interface{}{
			"id":         58,
			"name":       "template windaube7",
			"slug":       "slug windows7",
			"ram":        1,
			"cpu":        1,
			"os":         "Windows Serveur 64bits",
			"enterprise": "unit test enterprise",
			"disks": []interface{}{
				map[string]interface{}{
					"name":          "disk-Template-Windows",
					"size":          60,
					"storage_class": "storage_enterprise",
					"slug":          "disk-template-windows7",
				},
			},
			"datacenter":    "dc2",
			"nics":          []interface{}{},
			"login":         nil,
			"password":      nil,
			"dynamic_field": nil,
		},
		map[string]interface{}{
			"id":         69,
			"name":       "template5",
			"slug":       "template5-slug",
			"ram":        1,
			"cpu":        1,
			"os":         "Debian",
			"enterprise": "unit test enterprise",
			"disks": []interface{}{
				map[string]interface{}{
					"name":          "disk-debian9-rd-1",
					"size":          10,
					"storage_class": "storage_enterprise",
					"slug":          "disk-debian9-rd-1",
				},
			},
			"datacenter": "dc2",
			"nics": []interface{}{
				map[string]interface{}{
					"vlan":        "unit test vlan1",
					"mac_address": "00:50:56:00:01:de",
					"connected":   true,
				},
				map[string]interface{}{
					"vlan":        "unit test vlan2",
					"mac_address": "00:50:56:00:01:df",
					"connected":   true,
				},
			},
			"login":         nil,
			"password":      nil,
			"dynamic_field": nil,
		},
	}
	TEST_UPDATE_VM_MAP = map[string]interface{}{
		"id":    "unit test vm",
		"name":  "Unit test vm",
		"state": "DOWN",
		"os":    "CentOS",
		"ram":   16,
		"cpu":   8,
		"disks": []interface{}{
			map[string]interface{}{
				"name":          "disk 1 update",
				"size":          42,
				"storage_class": "storage_class update",
				"slug":          "slug update",
				"v_disk":        "",
				"deletion":      false,
			},
			map[string]interface{}{
				"name":          "disk 2 update",
				"size":          42,
				"storage_class": "storage_class update",
				"slug":          "slug update",
				"v_disk":        "",
				"deletion":      false,
			},
		},
		"nics": []interface{}{
			map[string]interface{}{
				"vlan":        "vlan 1 update",
				"mac_address": "42",
				"connected":   false,
			},
		},
		"vdc":           "vdc update",
		"boot":          "on disk update",
		"storage_class": "storage_enterprise update",
		"slug":          "42 update",
		"token":         "424242 update",
		"backup":        "backup_no_backup update",
		"disk_image":    " update",
		"platform_name": "",
		"backup_size":   42,
		"comment":       "",
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
				"name":          "disk 1 update",
				"size":          42,
				"storage_class": "storage_class update",
				"slug":          "slug update",
				"v_disk":        "",
				"deletion":      false,
			},
			map[string]interface{}{
				"name":          "disk 2 update",
				"size":          42,
				"storage_class": "storage_class update",
				"slug":          "slug update",
				"v_disk":        "",
				"deletion":      false,
			},
		},
		"nics": []interface{}{
			map[string]interface{}{
				"vlan":        "vlan 1 update",
				"mac_address": "42",
				"connected":   false,
			},
		},
		"vdc":           "vdc update",
		"boot":          "on disk update",
		"storage_class": "storage_enterprise update",
		"slug":          "42 update",
		"token":         "424242 update",
		"backup":        "backup_no_backup update",
		"disk_image":    " update",
		"platform_name": "",
		"backup_size":   43,
		"comment":       "",
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
				"name":          "disk 1 update",
				"size":          42,
				"storage_class": "storage_class update",
				"slug":          "slug update",
				"v_disk":        "",
				"deletion":      false,
			},
			map[string]interface{}{
				"name":          "disk 2 update",
				"size":          42,
				"storage_class": "storage_class update",
				"slug":          "slug update",
				"v_disk":        "",
				"deletion":      false,
			},
		},
		"nics": []interface{}{
			map[string]interface{}{
				"vlan":        "vlan 1 update",
				"mac_address": "42",
				"connected":   false,
			},
		},
		"vdc":           "vdc update",
		"boot":          "on disk update",
		"storage_class": "storage_enterprise update",
		"slug":          "42 update",
		"token":         "424242 update",
		"backup":        "backup_no_backup update",
		"disk_image":    " update",
		"platform_name": "",
		"backup_size":   42,
		"comment":       "",
	}
)

func resource_vdc_resource() *schema.Resource {
	return &schema.Resource{
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
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

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
				Elem:     resource_vdc_resource(),
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

func resource_vm_disk() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"storage_class": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"v_disk": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"deletion": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resource_vm_nic() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connected": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
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
			"enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     resource_vm_disk(),
			},
			"nics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resource_vm_nic(),
			},
			"vdc": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"boot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_class": &schema.Schema{
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
				Type:     schema.TypeInt,
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
				Computed: true,
			},
		},
	}
}

func Fake_vdcInstance_VDC_CREATION_MAP() VDC {
	return VDC{
		Name:       "Unit test vdc resource",
		Enterprise: "unit test enterprise",
		Datacenter: "dc1",
		Vdc_resources: []interface{}{
			map[string]interface{}{
				"resource": "ram",
				"used":     0,
				"total":    20,
				"slug":     "",
			},
			map[string]interface{}{
				"resource": "cpu",
				"used":     0,
				"total":    1,
				"slug":     "",
			},
			map[string]interface{}{
				"resource": "storage_enterprise",
				"used":     0,
				"total":    10,
				"slug":     "",
			},
			map[string]interface{}{
				"resource": "storage_performance",
				"used":     0,
				"total":    10,
				"slug":     "",
			},
			map[string]interface{}{
				"resource": "storage_high_performance",
				"used":     0,
				"total":    10,
				"slug":     "",
			},
		},
		Slug:          "",
		Dynamic_field: "",
	}
}

func vdcInstanceFake() VDC {
	return VDC{
		Name:       "Unit test vdc resource",
		Enterprise: "Unit Test value",
		Datacenter: "Unit Test value",
		Vdc_resources: []interface{}{
			map[string]interface{}{
				"resource": "Resource1",
				"used":     1,
				"total":    2,
				"slug":     "Unit Test value1",
			},
			map[string]interface{}{
				"Resource": "Resource2",
				"used":     1,
				"total":    2,
				"slug":     "Unit Test value2",
			},
			map[string]interface{}{
				"resource": "Resource3",
				"used":     1,
				"total":    2,
				"slug":     "Unit Test value3",
			},
		},
		Slug:          "Unit Test value",
		Dynamic_field: "Unit Test value",
	}
}

func vmInstanceNO_TEMPLATE_VM_MAP() VM {
	return VM{
		Name:  "Unit test no template vm resource",
		State: "UP",
		OS:    "Debian",
		RAM:   8,
		CPU:   4,
		Disks: []interface{}{
			map[string]interface{}{
				"name":          "disk 1",
				"size":          24,
				"storage_class": "storage_class",
				"slug":          "",
				"v_disk":        "",
				"deletion":      false,
			},
		},
		Nics: []interface{}{
			map[string]interface{}{
				"vlan":        "vlan 1 update",
				"mac_address": "00:21:21:21:21:21",
				"connected":   true,
			},
			map[string]interface{}{
				"vlan":        "vlan 2",
				"mac_address": "00:21:21:21:21:21",
				"connected":   true,
			},
		},
		Vdc:           "vdc",
		Boot:          "on disk",
		Storage_class: "storage_enterprise",
		Slug:          "42",
		Token:         "424242",
		Backup:        "backup_no_backup",
		Disk_image:    "",
		Platform_name: "42",
		Backup_size:   42,
		Comment:       "",
		Dynamic_field: "{\"terraform_provisioned\":true,\"creation_template\":\"\",\"disks_created_from_template\":null}",
	}
}

func Fake_vmInstance_EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP() VM {
	return VM{
		Name:       "Unit test template no disc add on vm resource",
		Enterprise: "unit test enterprise",
		Template:   "template1",
		State:      "UP",
		RAM:        1,
		CPU:        1,
		Disks: []interface{}{
			map[string]interface{}{
				"name":          "template1 disk1",
				"size":          20,
				"storage_class": "storage_enterprise",
				"slug":          "template1 disk1 slug",
				"v_disk":        "",
				"deletion":      false,
			},
		},
		Nics:          []interface{}{},
		Vdc:           "vdc",
		Boot:          "on disk",
		Storage_class: "storage_enterprise",
		Slug:          "42",
		Token:         "424242",
		Backup:        "backup_no_backup",
		Dynamic_field: "{\"terraform_provisioned\":true,\"creation_template\":\"template1\",\"disks_created_from_template\":[{\"name\":\"template1 disk1\",\"size\":20,\"slug\":\"template1 disk1 slug\",\"storage_class\":\"storage_enterprise\"}]}",
	}
}

func Fake_vmInstance_EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP() VM {
	return VM{
		Name:       "EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP",
		Enterprise: "unit test enterprise",
		Template:   "template1",
		State:      "UP",
		OS:         "Debian",
		RAM:        8,
		CPU:        4,
		Disks: []interface{}{
			map[string]interface{}{
				"name":          "disk 1",
				"size":          24,
				"storage_class": "storage_enterprise",
				"slug":          "",
				"v_disk":        "",
				"deletion":      false,
			},
			map[string]interface{}{
				"name":          "template1 disk1",
				"size":          25,
				"storage_class": "storage_enterprise",
				"slug":          "template1 disk1 slug",
				"v_disk":        "",
				"deletion":      false,
			},
		},
		Nics: []interface{}{
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
		Vdc:           "vdc",
		Boot:          "on disk",
		Storage_class: "storage_enterprise",
		Slug:          "42",
		Token:         "424242",
		Backup:        "backup_no_backup",
		Disk_image:    "",
		Platform_name: "42",
		Backup_size:   42,
		Comment:       "",
		Dynamic_field: "{\"terraform_provisioned\":true,\"creation_template\":\"template1\",\"disks_created_from_template\":[{\"name\":\"template1 disk1\",\"size\":20,\"slug\":\"template1 disk1 slug\",\"storage_class\":\"storage_enterprise\"}]}",
	}
}

func Fake_vmInstance_EXISTING_TEMPLATE_WITH_MODIFIED_NIC_AND_DISK_VM_MAP() VM {
	return VM{
		Name:       "EXISTING_TEMPLATE_WITH_MODIFIED_NIC_AND_DISK_VM_MAP",
		Enterprise: "unit test enterprise",
		Template:   "template1",
		State:      "UP",
		OS:         "Debian",
		RAM:        8,
		CPU:        4,
		Disks: []interface{}{
			map[string]interface{}{
				"name":          "template1 disk1",
				"size":          24,
				"storage_class": "storage_class",
				"slug":          "template1 disk1 slug",
				"v_disk":        "",
				"deletion":      false,
			},
		},
		Nics: []interface{}{
			map[string]interface{}{
				"vlan":        "unit test vlan1",
				"mac_address": "00:21:21:21:21:22",
				"connected":   true,
			},
			map[string]interface{}{
				"vlan":        "unit test vlan2",
				"mac_address": "00:21:21:21:21:21",
				"connected":   true,
			},
		},
		Vdc:           "vdc",
		Boot:          "on disk",
		Storage_class: "storage_enterprise",
		Slug:          "42",
		Token:         "424242",
		Backup:        "backup_no_backup",
		Disk_image:    "",
		Platform_name: "42",
		Backup_size:   42,
		Comment:       "",
		Dynamic_field: "{\"terraform_provisioned\":true,\"creation_template\":\"template1\",\"disks_created_from_template\":null}",
	}
}

func Fake_vmInstance_EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP() VM {
	return VM{
		Name:       "EXISTING_TEMPLATE_WITH_DELETED_DISK_VM_MAP",
		Enterprise: "unit test enterprise",
		Template:   "",
		State:      "UP",
		OS:         "Debian",
		RAM:        8,
		CPU:        4,
		Disks: []interface{}{
			map[string]interface{}{
				"name":          "disk 1",
				"size":          24,
				"storage_class": "storage_enterprise",
				"slug":          "",
				"v_disk":        "",
				"deletion":      false,
			},
		},
		Nics:          []interface{}{},
		Vdc:           "vdc",
		Boot:          "on disk",
		Storage_class: "storage_enterprise",
		Slug:          "42",
		Token:         "424242",
		Backup:        "backup_no_backup",
		Disk_image:    "",
		Platform_name: "42",
		Backup_size:   42,
		Comment:       "",
		Dynamic_field: "{\"terraform_provisioned\":true,\"creation_template\":\"template1\",\"disks_created_from_template\":null}",
	}
}

func vmInstanceNoTemplateFake() VM {
	return VM{
		Name:       "Unit test vm resource",
		Enterprise: "Unit Test value",
		Template:   "",
		State:      "Unit Test value",
		OS:         "Unit Test value",
		RAM:        1,
		CPU:        1,
		Disks: []interface{}{
			map[string]interface{}{
				"name":          "name1",
				"size":          10,
				"storage_class": "Unit Test value",
			},
			map[string]interface{}{
				"name":          "name2",
				"size":          10,
				"storage_class": "Unit Test value",
			},
		},
		Nics: []interface{}{
			map[string]interface{}{
				"vlan":        "vlan1",
				"mac_address": "00:21:21:21:21:21",
				"connected":   true,
			},
			map[string]interface{}{
				"vlan":        "vlan1",
				"mac_address": "00:21:21:21:21:21",
				"connected":   false,
			},
		},
		Vdc:           "Unit Test value",
		Boot:          "Unit Test value",
		Storage_class: "Unit Test value",
		Slug:          "Unit Test value",
		Token:         "Unit Test value",
		Backup:        "Unit Test value",
		Disk_image:    "Unit Test value",
		Platform_name: "Unit Test value",
		Backup_size:   42,
		Comment:       "",
	}
}

func vdc_schema_init(vdc map[string]interface{}) *schema.ResourceData {
	d := resource_vdc().TestResourceData()

	schemaTooler := SchemaTooler{
		SchemaTools: Schema_Schemaer{},
	}
	schemaTooler.SchemaTools.Update_local_resource_state(vdc, d, &schemaTooler)

	return d
}

func vm_schema_init(vm map[string]interface{}) *schema.ResourceData {
	d := resource_vm().TestResourceData()

	schemaTooler := SchemaTooler{
		SchemaTools: Schema_Schemaer{},
	}
	schemaTooler.SchemaTools.Update_local_resource_state(vm, d, &schemaTooler)

	return d
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

func JsonStub() map[string]interface{} {

	var jsonStub interface{}
	simple_json, _ := json.Marshal(Resp_Body{Detail: "a simple json"})
	jsonBytes := ioutil.NopCloser(bytes.NewBuffer(simple_json))
	readBytes, _ := ioutil.ReadAll(jsonBytes)
	_ = json.Unmarshal(readBytes, &jsonStub)

	return jsonStub.(map[string]interface{})
}

func JsonTemplateListFake() []interface{} {
	var jsonFake interface{}
	fake_json, _ := json.Marshal(TEMPLATES_LIST)
	jsonBytes := ioutil.NopCloser(bytes.NewBuffer(fake_json))
	readBytes, _ := ioutil.ReadAll(jsonBytes)
	_ = json.Unmarshal(readBytes, &jsonFake)

	return jsonFake.([]interface{})
}
