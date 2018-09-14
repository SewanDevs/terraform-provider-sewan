package sewan

import (
	sdk "github.com/SewanDevs/sewan-sdk-go"
)

const (
	vdcResourceField    = sdk.VdcResourceField
	vdcField            = sdk.VdcField
	nameField           = sdk.NameField
	enterpriseField     = sdk.EnterpriseField
	dataCenterField     = sdk.DataCenterField
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
	APIToken   string
	APIURL     string
	Enterprise string
}

type clientToolerStruct struct {
	SewanAPITooler       *sdk.APITooler
	SewanClientTooler    *sdk.ClientTooler
	SewanTemplatesTooler *sdk.TemplatesTooler
	SewanResourceTooler  *sdk.ResourceTooler
	SewanSchemaTooler    *sdk.SchemaTooler
}

type clientStruct struct {
	Sewan        *sdk.API
	ToolerStruct clientToolerStruct
}

func (c *configStruct) clientStruct(apiTooler *sdk.APITooler) (*clientStruct,
	error) {
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
	api := apiTooler.Initialyser.New(
		c.APIToken,
		c.APIURL,
		c.Enterprise,
	)
	err1 := apiTooler.Initialyser.CheckCloudDcStatus(api,
		&clientTooler,
		&resourceTooler)
	if err1 != nil {
		return nil, err1
	}
	clientStructAPIMeta, err2 := apiTooler.Initialyser.GetClouddcEnvMeta(api,
		&clientTooler)
	if err2 != nil {
		return nil, err2
	}
	api.Meta = *clientStructAPIMeta
	return &clientStruct{api,
			clientToolerStruct{
				apiTooler,
				&clientTooler,
				&templatesTooler,
				&resourceTooler,
				&schemaTooler},
		},
		nil
}
