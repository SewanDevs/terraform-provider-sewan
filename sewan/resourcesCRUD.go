package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
	sdk "gitlab.com/rd/sewan_go_sdk"
)

func createResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	apiResponse, err := m.(*Client).sewanAPIImplementerTooler.APIImplementer.CreateResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanTemplatesTooler,
		m.(*Client).sewanResourceTooler,
		resourceType,
		m.(*Client).sewan)
	if err != nil {
		return err
	}
	return m.(*Client).sewanSchemaTooler.SchemaTools.UpdateLocalResourceState(apiResponse,
		d,
		m.(*Client).sewanSchemaTooler)
}

func readResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	apiResponse, err := m.(*Client).sewanAPIImplementerTooler.APIImplementer.ReadResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanResourceTooler,
		resourceType,
		m.(*Client).sewan)
	switch {
	case err == sdk.ErrResourceNotExist:
		m.(*Client).sewanSchemaTooler.SchemaTools.DeleteTerraformResource(d)
		return nil
	case err != nil:
		return err
	default:
		return m.(*Client).sewanSchemaTooler.SchemaTools.UpdateLocalResourceState(apiResponse,
			d,
			m.(*Client).sewanSchemaTooler)
	}
}

func updateResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	return m.(*Client).sewanAPIImplementerTooler.APIImplementer.UpdateResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanTemplatesTooler,
		m.(*Client).sewanResourceTooler,
		resourceType,
		m.(*Client).sewan)
}

func deleteResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	err := m.(*Client).sewanAPIImplementerTooler.APIImplementer.DeleteResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanResourceTooler,
		resourceType,
		m.(*Client).sewan)
	if err == nil {
		m.(*Client).sewanSchemaTooler.SchemaTools.DeleteTerraformResource(d)
	}
	return err
}
