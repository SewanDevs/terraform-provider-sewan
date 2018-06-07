package sewan_go_sdk

import (
	"testing"
	"github.com/hashicorp/terraform/helper/schema"
)

//------------------------------------------------------------------------------
//--Structures init, interface implementation fakes, various test items etc.----
//------------------------------------------------------------------------------
func resource_vdc() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//------------------------------------------------------------------------------
//-------------Units tests------------------------------------------------------
//------------------------------------------------------------------------------
func TestCreate_vdc_resource(t *testing.T) {
	t.Errorf("RED TDD step, cycle1")
}

//------------------------------------------------------------------------------
func TestRead_vdc_resource(t *testing.T) {
	t.Errorf("RED TDD step, cycle1")
}

//------------------------------------------------------------------------------
func TestUpdate_vdc_resource(t *testing.T) {
	t.Errorf("RED TDD step, cycle1")
}

//------------------------------------------------------------------------------
func TestDelete_vdc_resource(t *testing.T) {
	t.Errorf("RED TDD step, cycle1")
}
