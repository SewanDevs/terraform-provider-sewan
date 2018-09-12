package sewan

import (
	sdk "github.com/SewanDevs/sewan-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func createResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	apiResponse, err := m.(*clientStruct).ToolerStruct.SewanAPITooler.Implementer.CreateResource(d,
		m.(*clientStruct).ToolerStruct.SewanClientTooler,
		m.(*clientStruct).ToolerStruct.SewanTemplatesTooler,
		m.(*clientStruct).ToolerStruct.SewanResourceTooler,
		resourceType,
		m.(*clientStruct).Sewan)
	if err != nil {
		return err
	}
	return m.(*clientStruct).ToolerStruct.SewanSchemaTooler.SchemaTools.UpdateLocalResourceState(apiResponse,
		d,
		m.(*clientStruct).ToolerStruct.SewanSchemaTooler)
}

func readResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	apiResponse, err := m.(*clientStruct).ToolerStruct.SewanAPITooler.Implementer.ReadResource(d,
		m.(*clientStruct).ToolerStruct.SewanClientTooler,
		m.(*clientStruct).ToolerStruct.SewanResourceTooler,
		resourceType,
		m.(*clientStruct).Sewan)
	switch {
	case err == sdk.ErrResourceNotExist:
		m.(*clientStruct).ToolerStruct.SewanSchemaTooler.SchemaTools.DeleteTerraformResource(d)
		return nil
	case err != nil:
		return err
	default:
		return m.(*clientStruct).ToolerStruct.SewanSchemaTooler.SchemaTools.UpdateLocalResourceState(apiResponse,
			d,
			m.(*clientStruct).ToolerStruct.SewanSchemaTooler)
	}
}

func updateResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	return m.(*clientStruct).ToolerStruct.SewanAPITooler.Implementer.UpdateResource(d,
		m.(*clientStruct).ToolerStruct.SewanClientTooler,
		m.(*clientStruct).ToolerStruct.SewanTemplatesTooler,
		m.(*clientStruct).ToolerStruct.SewanResourceTooler,
		resourceType,
		m.(*clientStruct).Sewan)
}

func deleteResource(d *schema.ResourceData,
	m interface{},
	resourceType string) error {
	err := m.(*clientStruct).ToolerStruct.SewanAPITooler.Implementer.DeleteResource(d,
		m.(*clientStruct).ToolerStruct.SewanClientTooler,
		m.(*clientStruct).ToolerStruct.SewanResourceTooler,
		resourceType,
		m.(*clientStruct).Sewan)
	if err == nil {
		m.(*clientStruct).ToolerStruct.SewanSchemaTooler.SchemaTools.DeleteTerraformResource(d)
	}
	return err
}
