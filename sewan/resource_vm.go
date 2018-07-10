package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resource_vm_disk() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			NAME_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			SIZE_FIELD: &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			STORAGE_CLASS_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			SLUG_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			DELETION_FIELD: &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resource_vm_nic() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			VLAN_NAME_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			MAC_ADRESS_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			CONNECTED_FIELD: &schema.Schema{
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
			NAME_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			ENTERPRISE_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			TEMPLATE_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			STATE_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			OS_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			RAM_FIELD: &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			CPU_FIELD: &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			DISKS_FIELD: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resource_vm_disk(),
			},
			NICS_FIELD: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resource_vm_nic(),
			},
			VDC_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			BOOT_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			STORAGE_CLASS_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			SLUG_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			TOKEN_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			BACKUP_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			DISK_IMAGE_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			PLATFORM_NAME_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			BACKUP_SIZE_FIELD: &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			COMMENT_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			DYNAMIC_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
		m.(*Client).sewan_templatesTooler,
		m.(*Client).sewan_schemaTooler,
		VM_RESOURCE_TYPE,
		sewan)

	if creationError == nil {
		creationError = m.(*Client).sewan_schemaTooler.SchemaTools.Update_local_resource_state(apiCreationResponse,
			d,
			m.(*Client).sewan_schemaTooler)
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
		m.(*Client).sewan_templatesTooler,
		m.(*Client).sewan_schemaTooler,
		VM_RESOURCE_TYPE,
		sewan)

	if resource_exists == false {
		m.(*Client).sewan_schemaTooler.SchemaTools.Delete_terraform_resource(d)
	} else {
		if readError == nil {
			readError = m.(*Client).sewan_schemaTooler.SchemaTools.Update_local_resource_state(apiReadResponse,
				d,
				m.(*Client).sewan_schemaTooler)
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
		m.(*Client).sewan_templatesTooler,
		m.(*Client).sewan_schemaTooler,
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
		m.(*Client).sewan_templatesTooler,
		m.(*Client).sewan_schemaTooler,
		VM_RESOURCE_TYPE,
		sewan)
	if deleteError == nil {
		m.(*Client).sewan_schemaTooler.SchemaTools.Delete_terraform_resource(d)
	}
	return deleteError
}
