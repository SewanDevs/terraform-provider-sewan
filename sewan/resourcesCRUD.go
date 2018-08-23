package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
	sdk "github.com/SewanDevs/sewan_go_sdk"
)

func createResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	apiResponse, err := m.(*Client).sewanApiTooler.Api.CreateResource(d,
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
	apiResponse, err := m.(*Client).sewanApiTooler.Api.ReadResource(d,
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
	return m.(*Client).sewanApiTooler.Api.UpdateResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanTemplatesTooler,
		m.(*Client).sewanResourceTooler,
		resourceType,
		m.(*Client).sewan)
}

func deleteResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	err := m.(*Client).sewanApiTooler.Api.DeleteResource(d,
		m.(*Client).sewanClientTooler,
		m.(*Client).sewanResourceTooler,
		resourceType,
		m.(*Client).sewan)
	if err == nil {
		m.(*Client).sewanSchemaTooler.SchemaTools.DeleteTerraformResource(d)
	}
	return err
}
