package sewan_go_sdk

import (
	"github.com/hashicorp/terraform/helper/schema"
  "encoding/json"
  "bytes"
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

func JsonStub() map[string]interface{}{

  var jsonStub interface{}
  simple_json,_ := json.Marshal(Resp_Body{Detail: "a simple json"})
  jsonBytes := ioutil.NopCloser(bytes.NewBuffer(simple_json))
  readBytes,_ := ioutil.ReadAll(jsonBytes)
  _ = json.Unmarshal(readBytes, &jsonStub)

	return jsonStub.(map[string]interface{})
}
