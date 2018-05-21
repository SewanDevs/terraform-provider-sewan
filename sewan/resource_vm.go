package sewan

import (
  "github.com/hashicorp/terraform/helper/schema"
  "terraform-provider-sewan/sewan_go_sdk"
)

func resource_vm() *schema.Resource {
  return &schema.Resource{
    Create: resource_vm_create,
    Read:   resource_vm_read,
    Update: resource_vm_update,
    Delete: resource_vm_delete,
    Schema: map[string]*schema.Schema{
      "name": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "state": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "os": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "ram": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "cpu": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "disks": &schema.Schema{
        Type:     schema.TypeList,
        Required: true,
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
            "name": &schema.Schema{
              Type:     schema.TypeString,
              Required: true,
            },
            "size": &schema.Schema{
              Type:     schema.TypeString,
              Required: true,
            },
            "v_disk": &schema.Schema{
              Type:     schema.TypeString,
              Required: true,
            },
            "slug": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
          },
        },
      },
      "nics": &schema.Schema{
        Type:     schema.TypeList,
        Optional: true,
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
            "vlan": &schema.Schema{
              Type:     schema.TypeString,
              Required: true,
            },
            "mac_adress": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
            "connected": &schema.Schema{
              Type:     schema.TypeBool,
              Optional: true,
            },
          },
        },
      },
      "vdc": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "boot": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "vdc_resource_disk": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      //"template": &schema.Schema{
      //  Type:     schema.TypeString,
      //  Optional: true,
      //},
      "slug": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "token": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "backup": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "disk_image": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "platform_name": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "backup_size": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "comment": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "outsourcing": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "dynamic_field": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
    },
  }
}

func resource_vm_create(d *schema.ResourceData, m interface{}) error {
  var creationError error
  creationError = nil
  var apiCreationResponse map[string]interface{}
  client := sewan_go_sdk.ClientCreate()
  creationError,apiCreationResponse = sewan_go_sdk.Create_vm_resource(d,client)
  if creationError==nil {
    creationError = Update_local_resource_state(apiCreationResponse, d)
  }
  return creationError
}

func resource_vm_read(d *schema.ResourceData, m interface{}) error {
  var readError error
  readError = nil
  var resource_exists bool
  var apiCreationResponse map[string]interface{}
  client := sewan_go_sdk.ClientCreate()
  readError,apiCreationResponse,resource_exists = sewan_go_sdk.Read_vm_resource(d,client)
  if readError==nil {
    if resource_exists==true {
      readError = Update_local_resource_state(apiCreationResponse, d)
    } else{
      Delete_resource(d)
    }
  }
  return readError
}

func resource_vm_update(d *schema.ResourceData, m interface{}) error {
  var updateError error
  updateError = nil
  client := sewan_go_sdk.ClientCreate()
  updateError = sewan_go_sdk.Update_vm_resource(d,client)
  return updateError
}

func resource_vm_delete(d *schema.ResourceData, m interface{}) error {
  var deleteError error
  deleteError = nil
  client := sewan_go_sdk.ClientCreate()
  deleteError= sewan_go_sdk.Delete_vm_resource(d,client)
  if deleteError==nil {
    Delete_resource(d)
  }
  return deleteError
}
