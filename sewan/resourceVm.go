package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVmDisk() *schema.Resource {
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
			V_DISK_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVmNic() *schema.Resource {
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

func resourceVm() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmCreate,
		Read:   resourceVmRead,
		Update: resourceVmUpdate,
		Delete: resourceVmDelete,
		Schema: map[string]*schema.Schema{
			NAME_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			INSTANCE_NUMBER_FIELD: &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
				Elem:     resourceVmDisk(),
			},
			NICS_FIELD: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceVmNic(),
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
			OUTSOURCING_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVmCreate(d *schema.ResourceData, m interface{}) error {
	var creationError error
	creationError = nil
	var apiCreationResponse map[string]interface{}
	sewan := m.(*Client).sewan
	creationError,
		apiCreationResponse = m.(*Client).sewanApiTooler.Api.CreateResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanTemplatesTooler,
		m.(*Client).sewanSchemaTooler,
		VM_RESOURCE_TYPE,
		sewan)

	if creationError == nil {
		creationError = m.(*Client).sewanSchemaTooler.SchemaTools.UpdateLocalResourceState(apiCreationResponse,
			d,
			m.(*Client).sewanSchemaTooler)
	}
	return creationError
}

func resourceVmRead(d *schema.ResourceData, m interface{}) error {
	var readError error
	readError = nil
	var resource_exists bool
	var apiReadResponse map[string]interface{}
	sewan := m.(*Client).sewan
	readError,
		apiReadResponse,
		resource_exists = m.(*Client).sewanApiTooler.Api.ReadResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanTemplatesTooler,
		m.(*Client).sewanSchemaTooler,
		VM_RESOURCE_TYPE,
		sewan)

	if !resource_exists {
		m.(*Client).sewanSchemaTooler.SchemaTools.DeleteTerraformResource(d)
	} else {
		if readError == nil {
			readError = m.(*Client).sewanSchemaTooler.SchemaTools.UpdateLocalResourceState(apiReadResponse,
				d,
				m.(*Client).sewanSchemaTooler)
		}
	}
	return readError
}

func resourceVmUpdate(d *schema.ResourceData, m interface{}) error {
	var updateError error
	updateError = nil
	sewan := m.(*Client).sewan
	updateError = m.(*Client).sewanApiTooler.Api.UpdateResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanTemplatesTooler,
		m.(*Client).sewanSchemaTooler,
		VM_RESOURCE_TYPE,
		sewan)
	return updateError
}

func resourceVmDelete(d *schema.ResourceData, m interface{}) error {
	var deleteError error
	deleteError = nil
	sewan := m.(*Client).sewan
	deleteError = m.(*Client).sewanApiTooler.Api.DeleteResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanTemplatesTooler,
		m.(*Client).sewanSchemaTooler,
		VM_RESOURCE_TYPE,
		sewan)
	if deleteError == nil {
		m.(*Client).sewanSchemaTooler.SchemaTools.DeleteTerraformResource(d)
	}
	return deleteError
}
