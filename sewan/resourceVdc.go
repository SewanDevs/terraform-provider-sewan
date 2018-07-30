package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVdcResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			RESOURCE_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			USED_FIELD: &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			TOTAL_FIELD: &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			SLUG_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVdc() *schema.Resource {
	return &schema.Resource{
		Create: resourceVdcCreate,
		Read:   resourceVdcRead,
		Update: resourceVdcUpdate,
		Delete: resourceVdcDelete,
		Schema: map[string]*schema.Schema{
			NAME_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			ENTERPRISE_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			DATACENTER_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			VDC_RESOURCE_FIELD: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceVdcResource(),
			},
			SLUG_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			DYNAMIC_FIELD: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVdcCreate(d *schema.ResourceData, m interface{}) error {
	var creationError error
	creationError = nil
	var apiCreationResponse map[string]interface{}
	sewan := m.(*Client).sewan
	creationError, apiCreationResponse = m.(*Client).sewanApiTooler.Api.CreateResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanTemplatesTooler,
		m.(*Client).sewanSchemaTooler,
		VDC_RESOURCE_TYPE,
		sewan)

	if creationError == nil {
		creationError = m.(*Client).sewanSchemaTooler.SchemaTools.UpdateLocalResourceState(apiCreationResponse,
			d,
			m.(*Client).sewanSchemaTooler)
	}
	return creationError
}

func resourceVdcRead(d *schema.ResourceData, m interface{}) error {
	var readError error
	readError = nil
	var resource_exists bool
	var apiCreationResponse map[string]interface{}
	sewan := m.(*Client).sewan
	sewanSchemaTooler := m.(*Client).sewanSchemaTooler
	sewanSchemaTools := sewanSchemaTooler.SchemaTools
	readError, apiCreationResponse, resource_exists = m.(*Client).sewanApiTooler.Api.ReadResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanTemplatesTooler,
		m.(*Client).sewanSchemaTooler,
		VDC_RESOURCE_TYPE,
		sewan)

	if resource_exists == false {
		sewanSchemaTools.DeleteTerraformResource(d)
	} else {
		if readError == nil {
			readError = sewanSchemaTools.UpdateLocalResourceState(apiCreationResponse,
				d, sewanSchemaTooler)
		}
	}
	if readError == nil {
		readError = sewanSchemaTools.UpdateVdcResourcesNames(d)
	}
	return readError
}

func resourceVdcUpdate(d *schema.ResourceData, m interface{}) error {
	var updateError error
	updateError = nil
	sewan := m.(*Client).sewan
	updateError = m.(*Client).sewanApiTooler.Api.UpdateResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanTemplatesTooler,
		m.(*Client).sewanSchemaTooler,
		VDC_RESOURCE_TYPE,
		sewan)
	return updateError
}

func resourceVdcDelete(d *schema.ResourceData, m interface{}) error {

	var deleteError error
	deleteError = nil
	sewan := m.(*Client).sewan
	deleteError = m.(*Client).sewanApiTooler.Api.DeleteResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanTemplatesTooler,
		m.(*Client).sewanSchemaTooler,
		VDC_RESOURCE_TYPE,
		sewan)
	if deleteError == nil {
		m.(*Client).sewanSchemaTooler.SchemaTools.DeleteTerraformResource(d)
	}
	return deleteError
}
