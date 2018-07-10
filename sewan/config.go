package sewan

import (
	sdk "gitlab.com/sewan_go_sdk"
	"net/http"
)

const (
	VM_RESOURCE_TYPE  = "vm"
	VDC_RESOURCE_TYPE = VDC_FIELD
	NAME_FIELD     = sdk.NAME_FIELD
	ENTERPRISE_FIELD     = sdk.ENTERPRISE_FIELD
	DATACENTER_FIELD = sdk.DATACENTER_FIELD
	VDC_RESOURCE_FIELD = sdk.VDC_RESOURCE_FIELD
	RESOURCE_FIELD = sdk.RESOURCE_FIELD
	TOTAL_FIELD = sdk.TOTAL_FIELD
	USED_FIELD = sdk.USED_FIELD
	SLUG_FIELD = sdk.SLUG_FIELD
	STATE_FIELD = sdk.STATE_FIELD
	OS_FIELD       = sdk.OS_FIELD
	RAM_FIELD = sdk.RAM_FIELD
	CPU_FIELD = sdk.CPU_FIELD
	DISKS_FIELD    = sdk.DISKS_FIELD
	SIZE_FIELD = sdk.SIZE_FIELD
	STORAGE_CLASS_FIELD = sdk.STORAGE_CLASS_FIELD
	DELETION_FIELD = sdk.DELETION_FIELD
	NICS_FIELD     = sdk.NICS_FIELD
	VLAN_NAME_FIELD = sdk.VLAN_NAME_FIELD
	MAC_ADRESS_FIELD = sdk.MAC_ADRESS_FIELD
	CONNECTED_FIELD = sdk.CONNECTED_FIELD
	VDC_FIELD = sdk.VDC_FIELD
	BOOT_FIELD = sdk.BOOT_FIELD
	TOKEN_FIELD = sdk.TOKEN_FIELD
	BACKUP_FIELD = sdk.BACKUP_FIELD
	DISK_IMAGE_FIELD = sdk.DISK_IMAGE_FIELD
	PLATFORM_NAME_FIELD = sdk.PLATFORM_NAME_FIELD
	BACKUP_SIZE_FIELD = sdk.BACKUP_SIZE_FIELD
	COMMENT_FIELD = sdk.COMMENT_FIELD
	TEMPLATE_FIELD = sdk.TEMPLATE_FIELD
	ID_FIELD       = sdk.ID_FIELD
	DYNAMIC_FIELD  = sdk.DYNAMIC_FIELD
)

type Config struct {
	Api_token string
	Api_url   string
}

type API struct {
	Token  string
	URL    string
	Client *http.Client
}

type Client struct {
	sewan                 *sdk.API
	sewan_apiTooler       *sdk.APITooler
	sewan_clientTooler    *sdk.ClientTooler
	sewan_templatesTooler *sdk.TemplatesTooler
	sewan_schemaTooler    *sdk.SchemaTooler
}

func (c *Config) Client() (*Client, error) {
	apiTooler := sdk.APITooler{
		Api: sdk.AirDrumResources_Apier{},
	}
	clientTooler := sdk.ClientTooler{
		Client: sdk.HttpClienter{},
	}
	templatesTooler := sdk.TemplatesTooler{
		TemplatesTools: sdk.Template_Templater{},
	}
	schemaTooler := sdk.SchemaTooler{
		SchemaTools: sdk.Schema_Schemaer{},
	}
	api := apiTooler.New(
		c.Api_token,
		c.Api_url,
	)
	err := apiTooler.CheckStatus(api)

	return &Client{api,
			&apiTooler,
			&clientTooler,
			&templatesTooler,
			&schemaTooler},
		err
}
