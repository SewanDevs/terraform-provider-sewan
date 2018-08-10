package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVdcResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			resourceField: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			usedField: &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			totalField: &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			slugField: &schema.Schema{
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
			nameField: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			enterpriseField: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			datacenterField: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			vdcResourceField: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceVdcResource(),
			},
			slugField: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			dynamicField: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceVdcCreate(d *schema.ResourceData, m interface{}) error {
	return createResource(d, m, vdcResourceType)
}
func resourceVdcRead(d *schema.ResourceData, m interface{}) error {
	return readResource(d, m, vdcResourceType)
}
func resourceVdcUpdate(d *schema.ResourceData, m interface{}) error {
	return updateResource(d, m, vdcResourceType)
}
func resourceVdcDelete(d *schema.ResourceData, m interface{}) error {
	return deleteResource(d, m, vdcResourceType)
}
