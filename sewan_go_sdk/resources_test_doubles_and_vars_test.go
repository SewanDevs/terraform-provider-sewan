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
		"enterprise": "sewan-rd-cloud-beta",
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
	TEMPLATES_LIST = []interface{}{
		map[string]interface{}{
			"id":         "40",
			"name":       "CentOS 7 Classique",
			"slug":       "TPL-CentOS7-x64-20Go-1vCPU-1Go-2GoS",
			"ram":        "1",
			"cpu":        "1",
			"os":         "CentOS",
			"enterprise": "sewan-rd-cloud-beta",
			"disks": []interface{}{
				map[string]interface{}{
					"name":   "/",
					"size":   "20",
					"v_disk": "sewan-rd-cloud-beta-mono-storage_enterprise",
					"slug":   "centos7-classic-disk1",
				},
			},
			"datacenter":    "dc2",
			"nics":          []interface{}{},
			"login":         "",
			"password":      "",
			"dynamic_field": "nil",
		},
		map[string]interface{}{
			"id":         "41",
			"name":       "Debian 8 Classique",
			"slug":       "TPL-Debian8-x64-20Go-1vCPU-1Go-2GoS",
			"ram":        "1",
			"cpu":        "1",
			"os":         "Debian",
			"enterprise": "sewan-rd-cloud-beta",
			"disks": []interface{}{
				map[string]interface{}{
					"name":   "/",
					"size":   "20",
					"v_disk": "sewan-rd-cloud-beta-mono-storage_enterprise",
					"slug":   "debian-8-classic-disk-1",
				},
			},
			"datacenter":    "dc2",
			"nics":          []interface{}{},
			"login":         "",
			"password":      "",
			"dynamic_field": "nil",
		},
		map[string]interface{}{
			"id":         "43",
			"name":       "tpl-CentOS7 R&D",
			"slug":       "tpl-centos7-rd",
			"ram":        "1",
			"cpu":        "1",
			"os":         "CentOS",
			"enterprise": "sewan-rd-cloud-beta",
			"disks": []interface{}{
				map[string]interface{}{
					"name":   "disk-tpl-CentOS7 R&D-1",
					"size":   "20",
					"v_disk": "sewan-rd-cloud-beta-mono-storage_enterprise",
					"slug":   "disk-tpl-centos7-rd-1",
				},
			},
			"datacenter": "dc2",
			"nics": []interface{}{
				map[string]interface{}{
					"vlan":        "sewanrd-mgt-tc3",
					"mac_address": "00:50:56:00:00:23",
					"connected":   "true",
				},
				map[string]interface{}{
					"vlan":        "sewanrd-priv-tc3",
					"mac_address": "00:50:56:00:00:24",
					"connected":   "true",
				},
			},
			"login":         "nil",
			"password":      "nil",
			"dynamic_field": "nil",
		},
		map[string]interface{}{
			"id":         "58",
			"name":       "Template-Windows7",
			"slug":       "template-windows7",
			"ram":        "1",
			"cpu":        "1",
			"os":         "Windows Serveur 64bits",
			"enterprise": "sewan-rd-cloud-beta",
			"disks": []interface{}{
				map[string]interface{}{
					"name":   "disk-Template-Windows7-1",
					"size":   "60",
					"v_disk": "sewan-rd-cloud-beta-mono-storage_enterprise",
					"slug":   "disk-template-windows7-1",
				},
			},
			"datacenter":    "dc2",
			"nics":          []interface{}{},
			"login":         "nil",
			"password":      "nil",
			"dynamic_field": "nil",
		},
		map[string]interface{}{
			"id":         "69",
			"name":       "debian9-rd",
			"slug":       "debian9-rd",
			"ram":        "1",
			"cpu":        "1",
			"os":         "Debian",
			"enterprise": "sewan-rd-cloud-beta",
			"disks": []interface{}{
				map[string]interface{}{
					"name":   "disk-debian9-rd-1",
					"size":   "10",
					"v_disk": "sewan-rd-cloud-beta-mono-storage_enterprise",
					"slug":   "disk-debian9-rd-1",
				},
			},
			"datacenter": "dc2",
			"nics": []interface{}{
				map[string]interface{}{
					"vlan":        "sewanrd-mgt-tc3",
					"mac_address": "00:50:56:00:01:de",
					"connected":   "true",
				},
				map[string]interface{}{
					"vlan":        "sewanrd-priv-tc3",
					"mac_address": "00:50:56:00:01:df",
					"connected":   "true",
				},
			},
			"login":         "nil",
			"password":      "nil",
			"dynamic_field": "nil",
		},
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
			"enterprise": &schema.Schema{
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

func JsonStub() map[string]interface{} {

	var jsonStub interface{}
	simple_json, _ := json.Marshal(Resp_Body{Detail: "a simple json"})
	jsonBytes := ioutil.NopCloser(bytes.NewBuffer(simple_json))
	readBytes, _ := ioutil.ReadAll(jsonBytes)
	_ = json.Unmarshal(readBytes, &jsonStub)

	return jsonStub.(map[string]interface{})
}

func JsonTemplateListFake() []interface{} {

	return TEMPLATES_LIST
}
