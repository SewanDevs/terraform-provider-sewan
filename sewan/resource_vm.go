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
      "address": &schema.Schema{
      Type:     schema.TypeString,
      Required: true,
      },
      "cpus": &schema.Schema{
      Type:     schema.TypeInt,
      Required: true,
      },
      "ram": &schema.Schema{
      Type:     schema.TypeInt,
      Required: true,
      },
    },
  }
}

type Machine struct {
	Name string
  Address string
  CPUs int
	RAM  int
}

type AirDrumClient struct {
	ApiToken   string
	Endpoint   string
	Timeout    int
	MaxRetries int
}

func Configure(d *schema.ResourceData) (interface{}, error) {
	Client := AirdrumClient{
		ApiToken:   d.Get("api_key").(string),
		Endpoint:   d.Get("endpoint").(string),
		Timeout:    d.Get("timeout").(int),
		MaxRetries: d.Get("max_retries").(int),
	}
  // Validate here the api token has not expired
	return &Client, nil
}

func resourceVMCreate(d *schema.ResourceData, m interface{}) error {

  //sewanClient := m.(*MyClient)

  machine := Machine{
		Name: d.Get("name").(string),
    Address: d.Get("address").(string),
    CPUs: d.Get("cpus").(int),
		RAM:  d.Get("ram").(int),
	}

  //err := client.CreateMachine(&machine)
  //if err != nil {
  //  return err
  //}

  d.SetId(machine.Name)

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
