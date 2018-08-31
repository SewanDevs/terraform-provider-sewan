package sewansdk

import (
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

// TemplatesTooler contains implementation of Templater interface
type TemplatesTooler struct {
	TemplatesTools Templater
}

// Templater interface is responsible of operations on :
// * Sewan's clouddc managed templates (see user doc : https://github.com/SewanDevs/terraform-provider-sewan/blob/github_release/website/docs/r/vm.html.md)
type Templater interface {
	FetchTemplateFromList(templateName string,
		templateList []interface{}) (map[string]interface{}, error)
	validateTemplate(template map[string]interface{}) error
	updateSchemaFromTemplateOnResourceCreation(d *schema.ResourceData,
		template map[string]interface{}) error
	createVMTemplateOverrideConfig(d *schema.ResourceData,
		template map[string]interface{}) (string, error)
}

// TemplateTemplater implements Templater interface
type TemplateTemplater struct{}

type diskModifiableFields struct {
	Name         string `json:"name"`
	Size         int    `json:"size"`
	StorageClass string `json:"storage_class"`
}

type nicModifiableFields struct {
	Vlan      string `json:"vlan"`
	Connected bool   `json:"connected"`
}

type templateCreatedVMOverride struct {
	Name      string        `json:"name"`
	OS        string        `json:"os"`
	RAM       int           `json:"ram"`
	CPU       int           `json:"cpu"`
	Disks     []interface{} `json:"disks,omitempty"`
	Nics      []interface{} `json:"nics,omitempty"`
	Vdc       string        `json:"vdc"`
	Boot      string        `json:"boot"`
	Backup    string        `json:"backup"`
	DiskImage string        `json:"disk_image"`
}

// FetchTemplateFromList extracts a template from the received list
// Known implementation limitation :
//  * Redmine ticket #35489/#36874
func (templater TemplateTemplater) FetchTemplateFromList(templateName string,
	templateList []interface{}) (map[string]interface{}, error) {
	var (
		template          map[string]interface{}
		templateListError error
	)
	for i := 0; i < len(templateList); i++ {
		switch reflect.TypeOf(templateList[i]).Kind() {
		case reflect.Map:
			var (
				listTemplateName = templateList[i].(map[string]interface{})[NameField].(string)
			)
			if listTemplateName == templateName {
				template = templateList[i].(map[string]interface{})
				break
			}
		default:
			templateListError = errors.New("One of the fetch template " +
				"has a wrong format." +
				"\ngot : " + reflect.TypeOf(templateList[i]).Kind().String() +
				"\nwant : " + reflect.Map.String())
			break
		}
	}
	if template == nil && templateListError == nil {
		templateListError = errors.New("template \"" + templateName +
			"\" does not exists, please validate it's name")
	}
	return template, templateListError
}

// Validate a template is correctly formated  and has the required fields
// correctly set :
// "name", "os", "ram", "cpu", "enterprise", "disks"
// It too alidate that "nics" fields is a slice if exists
func (templater TemplateTemplater) validateTemplate(template map[string]interface{}) error {
	var (
		templateError              error
		templateRequiredFieldSlice = []string{NameField, OsField, RAMField,
			CPUField, EnterpriseField, DisksField}
		missingFieldsList strings.Builder
	)
	for _, elem := range templateRequiredFieldSlice {
		if _, ok := template[elem]; !ok {
			missingFieldsList.WriteString("\"")
			missingFieldsList.WriteString(elem)
			missingFieldsList.WriteString("\" ")
		}
	}
	if missingFieldsList.String() != "" {
		templateError = errors.New("Template missing fields : " +
			missingFieldsList.String())
	} else {
		_, ok := template[NicsField]
		if ok && (reflect.TypeOf(template[NicsField]).Kind() != reflect.Slice) {
			templateError = errors.New("Template " + NicsField +
				" is not a list as required but a " +
				reflect.TypeOf(template[NicsField]).Kind().String())
		}
	}
	return templateError
}

func (templater TemplateTemplater) updateSchemaFromTemplateOnResourceCreation(d *schema.ResourceData,
	template map[string]interface{}) error {
	if d.Id() != "" {
		return errors.New("Template field should not be set on " +
			"an existing resource, please review the configuration field." +
			"\n : The resource schema has not been updated.")
	}
	for key, value := range template {
		if reflect.ValueOf(key).IsValid() && reflect.ValueOf(value).IsValid() {
			updateSchemaFieldOnResourceCreation(d, key, value)
		}
	}
	return nil
}

// Creation of a json override configuration file with additional vm resource
// fields fetch from template. An override file is created because it is
// not possible not wanted to modify initial configuration file.
// Warning : The override file must be manually deleted after a deletion of all
// resource created from the template.
func (templater TemplateTemplater) createVMTemplateOverrideConfig(d *schema.ResourceData,
	template map[string]interface{}) (string, error) {
	vm := templateCreatedVMOverride{
		RAM:       d.Get(RAMField).(int),
		CPU:       d.Get(CPUField).(int),
		Vdc:       d.Get(VdcField).(string),
		Boot:      d.Get(BootField).(string),
		Backup:    d.Get(BackupField).(string),
		DiskImage: d.Get(DiskImageField).(string),
	}
	var (
		schemaer     SchemaSchemaer
		err          error
		listItem     interface{}
		overrideFile strings.Builder
		vmName       strings.Builder
	)
	switch {
	case d.Get(TemplateField) == "":
		return "", errors.New("Schema \"Template\" field is empty, " +
			"can not create a template override configuration.")
	default:
		overrideFile.WriteString(d.Get(TemplateField).(string))
		overrideFile.WriteString("_Template_override.tf.json")
		vmName.WriteString(d.Get(NameField).(string))
		_, isSet := d.GetOk(InstanceNumberField)
		if isSet {
			vmName.WriteString(resourceNameCountSeparator)
			vmName.WriteString(resourceDynamicInstanceNumber)
		}
		vm.OS = template[OsField].(string)
		vm.Name = vmName.String()
		if _, err := os.Stat(overrideFile.String()); os.IsNotExist(err) {
			readListValue := []interface{}{}
			for listKey, listValue := range template[DisksField].([]interface{}) {
				listItem, _ = schemaer.ReadElement(listKey, listValue)
				disk := diskModifiableFields{
					Name:         listItem.(map[string]interface{})[NameField].(string),
					Size:         listItem.(map[string]interface{})[SizeField].(int),
					StorageClass: listItem.(map[string]interface{})[StorageClassField].(string),
				}
				readListValue = append(readListValue, disk)
			}
			vm.Disks = readListValue
			readListValue = []interface{}{}
			for listKey, listValue := range d.Get(NicsField).([]interface{}) {
				listItem, _ = schemaer.ReadElement(listKey, listValue)
				nic := nicModifiableFields{
					Vlan:      listItem.(map[string]interface{})[VlanNameField].(string),
					Connected: listItem.(map[string]interface{})[ConnectedField].(bool),
				}
				readListValue = append(readListValue, nic)
			}
			vm.Nics = readListValue
			vmFieldsMap := map[string]interface{}{d.Get(NameField).(string): vm}
			vmMap := map[string]interface{}{"sewan_clouddc_vm": vmFieldsMap}
			resourcesMap := map[string]interface{}{"resource": vmMap}
			vmJSON, _ := json.Marshal(resourcesMap)
			err = ioutil.WriteFile(overrideFile.String(),
				vmJSON, 0644)
		}
		return overrideFile.String(), err
	}
}

func conformizeNicsSliceOnResourceCreation(d *schema.ResourceData,
	templateParamName string,
	value []interface{}) []interface{} {
	var (
		nicMap          map[string]interface{}
		schemaNicsSlice []interface{}
	)
	for _, nic := range value {
		nicMap = map[string]interface{}{}
		for nicParamName, nicParamValue := range nic.(map[string]interface{}) {
			switch nicParamName {
			case VlanNameField:
				nicMap[nicParamName] = nicParamValue
			case ConnectedField:
				nicMap[nicParamName] = nicParamValue
			default:
			}
		}
		schemaNicsSlice = append(schemaNicsSlice, nicMap)
	}
	for _, nic := range d.Get(templateParamName).([]interface{}) {
		schemaNicsSlice = append(schemaNicsSlice,
			nic.(map[string]interface{}))
	}
	return schemaNicsSlice
}

func updateSchemaFieldOnResourceCreation(d *schema.ResourceData,
	key string,
	value interface{}) {
	var (
		templateParamName                 = reflect.ValueOf(key).String()
		interfaceTemplateName interface{} = reflect.ValueOf(value).Interface()
		templateParamValue                = reflect.ValueOf(value).String()
	)
	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		switch {
		case templateParamName == IDField:
		case templateParamName == OsField:
		case templateParamName == NameField:
		case templateParamName == DatacenterField:
		case d.Get(templateParamName) == "":
			d.Set(templateParamName, templateParamValue)
		default:
		}
	case reflect.Int:
		switch {
		case templateParamName == IDField:
		case d.Get(templateParamName).(int) == 0:
			d.Set(templateParamName, int(interfaceTemplateName.(int)))
		default:
		}
	case reflect.Slice:
		switch {
		case key == DisksField:
		case key == NicsField:
			schemaNicsSlice := conformizeNicsSliceOnResourceCreation(d,
				templateParamName,
				value.([]interface{}))
			d.Set(templateParamName, schemaNicsSlice)
		default:
		}
	default:
	}
}
