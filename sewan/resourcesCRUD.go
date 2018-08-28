package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
	sdk "github.com/SewanDevs/sewan-sdk-go"
)

func createResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	apiResponse, err := m.(*clientStruct).sewanAPIImplementerTooler.APIImplementer.CreateResource(d,
		m.(*clientStruct).sewanClientTooler,
		m.(*clientStruct).sewanTemplatesTooler,
		m.(*clientStruct).sewanResourceTooler,
		resourceType,
		m.(*clientStruct).sewan)
	if err != nil {
		return err
	}
	return m.(*clientStruct).sewanSchemaTooler.SchemaTools.UpdateLocalResourceState(apiResponse,
		d,
		m.(*clientStruct).sewanSchemaTooler)
}

func readResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	apiResponse, err := m.(*clientStruct).sewanAPIImplementerTooler.APIImplementer.ReadResource(d,
		m.(*clientStruct).sewanClientTooler,
		m.(*clientStruct).sewanResourceTooler,
		resourceType,
		m.(*clientStruct).sewan)
	switch {
	case err == sdk.ErrResourceNotExist:
		m.(*clientStruct).sewanSchemaTooler.SchemaTools.DeleteTerraformResource(d)
		return nil
	case err != nil:
		return err
	default:
		return m.(*clientStruct).sewanSchemaTooler.SchemaTools.UpdateLocalResourceState(apiResponse,
			d,
			m.(*clientStruct).sewanSchemaTooler)
	}
}

func updateResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	return m.(*clientStruct).sewanAPIImplementerTooler.APIImplementer.UpdateResource(d,
		m.(*clientStruct).sewanClientTooler,
		m.(*clientStruct).sewanTemplatesTooler,
		m.(*clientStruct).sewanResourceTooler,
		resourceType,
		m.(*clientStruct).sewan)
}

func deleteResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	err := m.(*clientStruct).sewanAPIImplementerTooler.APIImplementer.DeleteResource(d,
		m.(*clientStruct).sewanClientTooler,
		m.(*clientStruct).sewanResourceTooler,
		resourceType,
		m.(*clientStruct).sewan)
	if err == nil {
		m.(*clientStruct).sewanSchemaTooler.SchemaTools.DeleteTerraformResource(d)
	}
	return err
}
