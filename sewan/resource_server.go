package sewan

import (
  "github.com/hashicorp/terraform/helper/schema"
  //"terraform-poc/terraform-provider-sewan/httpReqBuilder"
)

// Create an HTTP client instance

func resourceServer() *schema.Resource {
  return &schema.Resource{
    Create: resourceServerCreate,
    Read:   resourceServerRead,
    Update: resourceServerUpdate,
    Delete: resourceServerDelete,

    Schema: map[string]*schema.Schema{
        "address": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
        },
    },
  }
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
  //Here will be implemented the POST REST call
  address := d.Get("address").(string)
  d.SetId(address)
  return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
  //Here will be implemented the GET REST call

  //Example :
  //client := m.(*MyClient)

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

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
  //Here will be implemented the PUT REST call
  return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
  //Here will be implemented the DELETE REST call
  d.SetId("")
  return nil
}
