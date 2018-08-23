package sewan

import (
	sdk "github.com/SewanDevs/sewan_go_sdk"
	"net/http"
)

const (
	vdcResourceField    = sdk.VdcResourceField
	vdcField            = sdk.VdcField
	nameField           = sdk.NameField
	enterpriseField     = sdk.EnterpriseField
	datacenterField     = sdk.DatacenterField
	resourceField       = sdk.ResourceField
	totalField          = sdk.TotalField
	usedField           = sdk.UsedField
	slugField           = sdk.SlugField
	stateField          = sdk.StateField
	osField             = sdk.OsField
	ramField            = sdk.RamField
	cpuField            = sdk.CpuField
	disksField          = sdk.DisksField
	vDiskField          = sdk.VDiskField
	sizeField           = sdk.SizeField
	storageClassField   = sdk.StorageClassField
	nicsField           = sdk.NicsField
	vlanNameField       = sdk.VlanNameField
	macAdressField      = sdk.MacAdressField
	connectedField      = sdk.ConnectedField
	bootField           = sdk.BootField
	tokenField          = sdk.TokenField
	backupField         = sdk.BackupField
	diskImageField      = sdk.DiskImageField
	platformNameField   = sdk.PlatformNameField
	backupSizeField     = sdk.BackupSizeField
	commentField        = sdk.CommentField
	templateField       = sdk.TemplateField
	idField             = sdk.IdField
	dynamicField        = sdk.DynamicField
	outsourcingField    = sdk.OutsourcingField
	instanceNumberField = sdk.InstanceNumberField
	vmResourceType      = sdk.VmResourceType
	vdcResourceType     = sdk.VdcResourceType
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
	sewan                *sdk.API
	sewanApiTooler       *sdk.APITooler
	sewanClientTooler    *sdk.ClientTooler
	sewanTemplatesTooler *sdk.TemplatesTooler
	sewanResourceTooler  *sdk.ResourceTooler
	sewanSchemaTooler    *sdk.SchemaTooler
}

func (c *Config) Client() (*Client, error) {
	apiTooler := sdk.APITooler{
		Api: sdk.AirDrumResourcesApier{},
	}
	clientTooler := sdk.ClientTooler{
		Client: sdk.HttpClienter{},
	}
	templatesTooler := sdk.TemplatesTooler{
		TemplatesTools: sdk.Template_Templater{},
	}
	schemaTooler := sdk.SchemaTooler{
		SchemaTools: sdk.SchemaSchemaer{},
	}
	resourceTooler := sdk.ResourceTooler{
		Resource: sdk.ResourceResourceer{},
	}
	api := apiTooler.New(
		c.Api_token,
		c.Api_url,
	)
	err := apiTooler.CheckCloudDcApiStatus(api, &clientTooler, &resourceTooler)
	return &Client{api,
			&apiTooler,
			&clientTooler,
			&templatesTooler,
			&resourceTooler,
			&schemaTooler},
		err
}
