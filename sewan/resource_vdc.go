package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resource_vdc() *schema.Resource {
	return &schema.Resource{
		Create: resource_vdc_create,
		Read:   resource_vdc_read,
		Update: resource_vdc_update,
		Delete: resource_vdc_delete,
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

func resource_vdc_create(d *schema.ResourceData, m interface{}) error {
	var creationError error
	creationError = nil
	var apiCreationResponse map[string]interface{}
	sewan := m.(*Client).sewan
	creationError, apiCreationResponse = m.(*Client).sewan_apiTooler.Api.Create_resource(d,
		m.(*Client).sewan_clientTooler,
		VDC_RESOURCE_TYPE,
		sewan)

	if creationError == nil {
		creationError = Update_local_resource_state(apiCreationResponse, d)
	}
	return creationError
}

func resource_vdc_read(d *schema.ResourceData, m interface{}) error {
	var readError error
	readError = nil
	var resource_exists bool
	var apiCreationResponse map[string]interface{}
	sewan := m.(*Client).sewan
	readError, apiCreationResponse, resource_exists = m.(*Client).sewan_apiTooler.Api.Read_resource(d,
		m.(*Client).sewan_clientTooler,
		VDC_RESOURCE_TYPE,
		sewan)

	if resource_exists == false {
		Delete_resource(d)
	} else {
		if readError == nil {
			readError = Update_local_resource_state(apiCreationResponse, d)
		}
	}
	return readError
}

func resource_vdc_update(d *schema.ResourceData, m interface{}) error {
	var updateError error
	updateError = nil
	sewan := m.(*Client).sewan
	updateError = m.(*Client).sewan_apiTooler.Api.Update_resource(d,
		m.(*Client).sewan_clientTooler,
		VDC_RESOURCE_TYPE,
		sewan)
	return updateError
}

func resource_vdc_delete(d *schema.ResourceData, m interface{}) error {
	// Delete vdc in cloud not implemented for security reasons :
	// nulify the risk of deleting the wrong vdc during tests
	//
	//
	//var deleteError error
	//deleteError = nil
	//sewan := m.(*Client).sewan
	//deleteError = m.(*Client).sewan_apiTooler.Api.Delete_resource(d,
	//	m.(*Client).sewan_clientTooler,
	//	VDC_RESOURCE_TYPE,
	//	sewan)
	//if deleteError == nil {
	//	Delete_resource(d)
	//}
	//return deleteError
	Delete_resource(d)
	return nil
}
