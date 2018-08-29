package sewan

import (
	sdk "gitlab.com/rd/sewan-sdk-go"
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
	ramField            = sdk.RAMField
	cpuField            = sdk.CPUField
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
	idField             = sdk.IDField
	dynamicField        = sdk.DynamicField
	outsourcingField    = sdk.OutsourcingField
	instanceNumberField = sdk.InstanceNumberField
	vmResourceType      = sdk.VMResourceType
	vdcResourceType     = sdk.VdcResourceType
)

type configStruct struct {
	APIToken string
	APIURL   string
}

type clientStruct struct {
	sewan                *sdk.API
	sewanAPITooler       *sdk.APITooler
	sewanClientTooler    *sdk.ClientTooler
	sewanTemplatesTooler *sdk.TemplatesTooler
	sewanResourceTooler  *sdk.ResourceTooler
	sewanSchemaTooler    *sdk.SchemaTooler
}

func (c *configStruct) clientStruct() (*clientStruct, error) {
	apiTooler := sdk.APITooler{
		APIImplementer: sdk.AirDrumResourcesAPI{},
	}
	clientTooler := sdk.ClientTooler{
		Client: sdk.HTTPClienter{},
	}
	templatesTooler := sdk.TemplatesTooler{
		TemplatesTools: sdk.TemplateTemplater{},
	}
	schemaTooler := sdk.SchemaTooler{
		SchemaTools: sdk.SchemaSchemaer{},
	}
	resourceTooler := sdk.ResourceTooler{
		Resource: sdk.ResourceResourceer{},
	}
	api := apiTooler.New(
		c.APIToken,
		c.APIURL,
	)
	err := apiTooler.CheckCloudDcStatus(api, &clientTooler, &resourceTooler)
	return &clientStruct{api,
			&apiTooler,
			&clientTooler,
			&templatesTooler,
			&resourceTooler,
			&schemaTooler},
		err
}
