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
	ENTERPRISE_SLUG         = "sewan-rd-cloud-beta"
)

var (
	VDC_CREATION_MAP = map[string]interface{}{
		"name":       "Unit test vdc resource",
		"enterprise": "sewan-rd-cloud-beta",
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
		"enterprise": "sewan-rd-cloud-beta",
		"datacenter": "dc1",
		"vdc_resources": []interface{}{
			map[string]interface{}{
				"resource": "ram",
				"used":     "0",
				"total":    "20",
				"slug":     "sewan-rd-cloud-beta-dc1-vdc_te-ram",
			},
			map[string]interface{}{
				"resource": "cpu",
				"used":     "0",
				"total":    "1",
				"slug":     "sewan-rd-cloud-beta-dc1-vdc_te-cpu",
			},
			map[string]interface{}{
				"resource": "storage_enterprise",
				"used":     "0",
				"total":    "10",
				"slug":     "sewan-rd-cloud-beta-dc1-vdc_te-storage_enterprise",
			},
			map[string]interface{}{
				"resource": "storage_performance",
				"used":     "0",
				"total":    "10",
				"slug":     "sewan-rd-cloud-beta-dc1-vdc_te-storage_performance",
			},
			map[string]interface{}{
				"resource": "storage_high_performance",
				"used":     "0",
				"total":    "10",
				"slug":     "sewan-rd-cloud-beta-dc1-vdc_te-storage_high_performance",
			},
		},
		"slug":          "sewan-rd-cloud-beta-dc1-vdc_te",
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
		"outsourcing":   "42",
		"dynamic_field": "42",
	}
	EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP = map[string]interface{}{
		"name":          "Unit test template no disc add on vm resource",
		"enterprise":    "sewan-rd-cloud-beta",
		"template":      "centos7-rd-DC1",
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
		"enterprise": "sewan-rd-cloud-beta",
		"template":   "centos7-rd-DC1",
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
		"outsourcing":   "42",
		"dynamic_field": "42",
	}
	EXISTING_TEMPLATE_WITH_MODIFIED_NIC_AND_DISK_VM_MAP = map[string]interface{}{
		"name":       "EXISTING_TEMPLATE_WITH_MODIFIED_NIC_AND_DISK_VM_MAP",
		"enterprise": "sewan-rd-cloud-beta",
		"template":   "centos7-rd-DC1",
		"state":      "UP",
		"os":         "Debian",
		"ram":        8,
		"cpu":        4,
		"disks": []interface{}{
			map[string]interface{}{
				"name":          "disk-centos7-rd-DC1-1",
				"size":          24,
				"storage_class": "storage_class",
				"slug":          "disk-centos7-rd-dc1-1",
			},
		},
		"nics": []interface{}{
			map[string]interface{}{
				"mac_address": "00:21:21:21:21:22",
				"connected":   true,
				"vlan":        "sewanrd-mgt-th3",
			},
			map[string]interface{}{
				"vlan":        "sewanrd-priv-th3",
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
		"outsourcing":   "42",
		"dynamic_field": "42",
	}
	NON_EXISTING_TEMPLATE_VM_MAP = map[string]interface{}{
		"name":          "windows95 vm",
		"enterprise":    "sewan-rd-cloud-beta",
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
			"name":       "CentOS 7 Classique",
			"slug":       "TPL-CentOS7-x64-20Go-1vCPU-1Go-2GoS",
			"ram":        1,
			"cpu":        1,
			"os":         "CentOS",
			"enterprise": "sewan-rd-cloud-beta",
			"disks": []interface{}{
				map[string]interface{}{
					"name":          "/",
					"size":          20,
					"storage_class": "storage_enterprise",
					"slug":          "centos7-classic-disk1",
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
		},
		map[string]interface{}{
			"id":         41,
			"name":       "Debian 8 Classique",
			"slug":       "TPL-Debian8-x64-20Go-1vCPU-1Go-2GoS",
			"ram":        1,
			"cpu":        1,
			"os":         "Debian",
			"enterprise": "sewan-rd-cloud-beta",
			"disks": []interface{}{
				map[string]interface{}{
					"name":          "/",
					"size":          20,
					"storage_class": "storage_enterprise",
					"slug":          "debian-8-classic-disk-1",
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
			"name":       "tpl-CentOS7 R&D",
			"slug":       "tpl-centos7-rd",
			"ram":        1,
			"cpu":        1,
			"os":         "CentOS",
			"enterprise": "sewan-rd-cloud-beta",
			"disks": []interface{}{
				map[string]interface{}{
					"name":          "disk-tpl-CentOS7 R&D-1",
					"size":          20,
					"storage_class": "storage_enterprise",
					"slug":          "disk-tpl-centos7-rd-1",
				},
			},
			"datacenter": "dc2",
			"nics": []interface{}{
				map[string]interface{}{
					"vlan":        "sewanrd-mgt-tc3",
					"mac_address": "00:50:56:00:00:23",
					"connected":   true,
				},
				map[string]interface{}{
					"vlan":        "sewanrd-priv-tc3",
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
			"name":       "Template-Windows7",
			"slug":       "template-windows7",
			"ram":        1,
			"cpu":        1,
			"os":         "Windows Serveur 64bits",
			"enterprise": "sewan-rd-cloud-beta",
			"disks": []interface{}{
				map[string]interface{}{
					"name":          "disk-Template-Windows7-1",
					"size":          60,
					"storage_class": "storage_enterprise",
					"slug":          "disk-template-windows7-1",
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
			"name":       "debian9-rd",
			"slug":       "debian9-rd",
			"ram":        1,
			"cpu":        1,
			"os":         "Debian",
			"enterprise": "sewan-rd-cloud-beta",
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
					"vlan":        "sewanrd-mgt-tc3",
					"mac_address": "00:50:56:00:01:de",
					"connected":   true,
				},
				map[string]interface{}{
					"vlan":        "sewanrd-priv-tc3",
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
			},
			map[string]interface{}{
				"name":          "disk 2 update",
				"size":          42,
				"storage_class": "storage_class update",
				"slug":          "slug update",
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
		"outsourcing":   "42 update",
		"dynamic_field": "42 update",
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
			},
			map[string]interface{}{
				"name":          "disk 2 update",
				"size":          42,
				"storage_class": "storage_class update",
				"slug":          "slug update",
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
		"outsourcing":   "42 update",
		"dynamic_field": "42 update",
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
			},
			map[string]interface{}{
				"name":          "disk 2 update",
				"size":          42,
				"storage_class": "storage_class update",
				"slug":          "slug update",
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
		"outsourcing":   "42 update",
		"dynamic_field": "42 update",
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
				Required: true,
			},
			"storage_class": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
				Optional: true,
			},
		},
	}
}

func Fake_vdcInstance_VDC_CREATION_MAP() VDC {
	return VDC{
		Name:       "Unit test vdc resource",
		Enterprise: "sewan-rd-cloud-beta",
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
		Comment:       "42",
		Outsourcing:   "42",
		Dynamic_field: "42",
	}
}

func Fake_vmInstance_EXISTING_TEMPLATE_NO_ADDITIONAL_DISK_VM_MAP() VM {
	return VM{
		Name:       "Unit test template no disc add on vm resource",
		Enterprise: "sewan-rd-cloud-beta",
		Template:   "centos7-rd-DC1",
		State:      "UP",
		RAM:        1,
		CPU:        1,
		Disks: []interface{}{
			map[string]interface{}{
				"name":          "disk-centos7-rd-DC1-1",
				"size":          20,
				"storage_class": "storage_enterprise",
				"slug":          "disk-centos7-rd-dc1-1",
			},
		},
		Nics:          []interface{}{},
		Vdc:           "vdc",
		Boot:          "on disk",
		Storage_class: "storage_enterprise",
		Slug:          "42",
		Token:         "424242",
		Backup:        "backup_no_backup",
	}
}

func Fake_vmInstance_EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP() VM {
	return VM{
		Name:       "EXISTING_TEMPLATE_WITH_ADDITIONAL_AND_MODIFIED_NICS_AND_DISKS_VM_MAP",
		Enterprise: "sewan-rd-cloud-beta",
		Template:   "centos7-rd-DC1",
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
			},
			map[string]interface{}{
				"name":          "disk-centos7-rd-DC1-1",
				"size":          25,
				"storage_class": "storage_enterprise",
				"slug":          "disk-centos7-rd-dc1-1",
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
		Comment:       "42",
		Outsourcing:   "42",
		Dynamic_field: "42",
	}
}

func Fake_vmInstance_EXISTING_TEMPLATE_WITH_MODIFIED_NIC_AND_DISK_VM_MAP() VM {
	return VM{
		Name:       "EXISTING_TEMPLATE_WITH_MODIFIED_NIC_AND_DISK_VM_MAP",
		Enterprise: "sewan-rd-cloud-beta",
		Template:   "centos7-rd-DC1",
		State:      "UP",
		OS:         "Debian",
		RAM:        8,
		CPU:        4,
		Disks: []interface{}{
			map[string]interface{}{
				"name":          "disk-centos7-rd-DC1-1",
				"size":          24,
				"storage_class": "storage_class",
				"slug":          "disk-centos7-rd-dc1-1",
			},
		},
		Nics: []interface{}{
			map[string]interface{}{
				"vlan":        "sewanrd-mgt-th3",
				"mac_address": "00:21:21:21:21:22",
				"connected":   true,
			},
			map[string]interface{}{
				"vlan":        "sewanrd-priv-th3",
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
		Comment:       "42",
		Outsourcing:   "42",
		Dynamic_field: "42",
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
		Comment:       "Unit Test value",
		Outsourcing:   "Unit Test value",
		Dynamic_field: "Unit Test value",
	}
}

func vdc_schema_init(vdc map[string]interface{}) *schema.ResourceData {
	d := resource_vdc().TestResourceData()

	Update_local_resource_state(vdc, d)

	return d
}

func vm_schema_init(vm map[string]interface{}) *schema.ResourceData {
	d := resource_vm().TestResourceData()

	Update_local_resource_state(vm, d)

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
