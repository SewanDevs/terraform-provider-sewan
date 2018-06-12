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
		},
	}
}

func resource_vdc_create(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resource_vdc_read(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resource_vdc_update(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resource_vdc_delete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
