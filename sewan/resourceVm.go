package sewan

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVmDisk() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			nameField: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			sizeField: &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			storageClassField: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			slugField: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			vDiskField: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func resourceVmNic() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			vlanNameField: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			macAdressField: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			connectedField: &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}
func resourceVm() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmCreate,
		Read:   resourceVmRead,
		Update: resourceVmUpdate,
		Delete: resourceVmDelete,
		Schema: map[string]*schema.Schema{
			nameField: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			instanceNumberField: &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			enterpriseField: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			templateField: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			stateField: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			osField: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			ramField: &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			cpuField: &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			disksField: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceVmDisk(),
			},
			nicsField: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceVmNic(),
			},
			vdcField: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			bootField: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			storageClassField: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			slugField: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			tokenField: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			backupField: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			diskImageField: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			platformNameField: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			backupSizeField: &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			commentField: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			dynamicField: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			outsourcingField: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func resourceVmCreate(d *schema.ResourceData, m interface{}) error {
	return createResource(d, m, vmResourceType)
}
func resourceVmRead(d *schema.ResourceData, m interface{}) error {
	return readResource(d, m, vmResourceType)
}
func resourceVmUpdate(d *schema.ResourceData, m interface{}) error {
	return updateResource(d, m, vmResourceType)
}
func resourceVmDelete(d *schema.ResourceData, m interface{}) error {
	return deleteResource(d, m, vmResourceType)
}
