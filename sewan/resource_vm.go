package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
	sdk "terraform-provider-sewan/sewan_go_sdk"
)

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
			"v_disk": &schema.Schema{
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
		Create: resource_vm_create,
		Read:   resource_vm_read,
		Update: resource_vm_update,
		Delete: resource_vm_delete,
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

func resource_vm_create(d *schema.ResourceData, m interface{}) error {
	var creationError error
	creationError = nil
	var apiCreationResponse map[string]interface{}
	sewan := m.(*Client).sewan
	creationError,
		apiCreationResponse = m.(*Client).sewan_apiTooler.Api.Create_resource(d,
		m.(*Client).sewan_clientTooler,
		m.(*Client).sewan_TemplatesTooler,
		VM_RESOURCE_TYPE,
		sewan)

	if creationError == nil {
		creationError = sdk.Update_local_resource_state(apiCreationResponse, d)
	}
	return creationError
}

func resource_vm_read(d *schema.ResourceData, m interface{}) error {
	var readError error
	readError = nil
	var resource_exists bool
	var apiReadResponse map[string]interface{}
	sewan := m.(*Client).sewan
	readError,
		apiReadResponse,
		resource_exists = m.(*Client).sewan_apiTooler.Api.Read_resource(d,
		m.(*Client).sewan_clientTooler,
		m.(*Client).sewan_TemplatesTooler,
		VM_RESOURCE_TYPE,
		sewan)

	if resource_exists == false {
		sdk.Delete_terraform_resource(d)
	} else {
		if readError == nil {
			readError = sdk.Update_local_resource_state(apiReadResponse, d)
		}
	}
	return readError
}

func resource_vm_update(d *schema.ResourceData, m interface{}) error {
	var updateError error
	updateError = nil
	sewan := m.(*Client).sewan
	updateError = m.(*Client).sewan_apiTooler.Api.Update_resource(d,
		m.(*Client).sewan_clientTooler,
		m.(*Client).sewan_TemplatesTooler,
		VM_RESOURCE_TYPE,
		sewan)
	return updateError
}

func resource_vm_delete(d *schema.ResourceData, m interface{}) error {
	var deleteError error
	deleteError = nil
	sewan := m.(*Client).sewan
	deleteError = m.(*Client).sewan_apiTooler.Api.Delete_resource(d,
		m.(*Client).sewan_clientTooler,
		m.(*Client).sewan_TemplatesTooler,
		VM_RESOURCE_TYPE,
		sewan)
	if deleteError == nil {
		sdk.Delete_terraform_resource(d)
	}
	return deleteError
}
