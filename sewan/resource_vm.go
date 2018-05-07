package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVM() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMCreate,
		Read:   resourceVMRead,
		Update: resourceVMUpdate,
		Delete: resourceVMDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdc": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ram": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"disk_image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"boot": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdc_resource_disk": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"backup": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type VM struct {
	Name              string
	Vdc               string
	RAM               int
	CPU               int
	Disk_image        string
	Boot              string
	Template          string
	Vdc_resource_disk string
	Backup            string
}

func resourceVMCreate(d *schema.ResourceData, m interface{}) error {

	//sewanClient := m.(*sewanClient)

	VMInstance := VM{
		Name:              d.Get("name").(string),
		Vdc:               d.Get("vdc").(string),
		CPU:               d.Get("cpu").(int),
		RAM:               d.Get("ram").(int),
		Disk_image:        d.Get("disk_image").(string),
		Boot:              d.Get("boot").(string),
		Template:          d.Get("template").(string),
		Vdc_resource_disk: d.Get("vdc_resource_disk").(string),
		Backup:            d.Get("backup").(string),
	}

	//err := SewanClient.CreateMachine(&VMInstance)
	//if err != nil {
	//  return err
	//}

	d.SetId(VMInstance.Name)

	return nil
}

func resourceVMRead(d *schema.ResourceData, m interface{}) error {
	//Here will be implemented the GET REST call
	//sewanClient := m.(*MyClient)

	// Attempt to read from an upstream API
	//obj, ok := client.Get(d.Id())

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	//if !ok {
	//  d.SetId("")
	//  return nil
	//}

	//d.Set("address", obj.Address)
	return nil
}

func resourceVMUpdate(d *schema.ResourceData, m interface{}) error {
	//Here will be implemented the PUT REST call
	return nil
}

func resourceVMDelete(d *schema.ResourceData, m interface{}) error {
	//Here will be implemented the DELETE REST call
	d.SetId("")
	return nil
}
