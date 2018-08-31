package sewansdk

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"reflect"
	"strconv"
	"strings"
)

// SchemaTooler contains implementation of Schemaer interface
type SchemaTooler struct {
	SchemaTools Schemaer
}

// Schemaer interface is responsible of operations on :
// * github.com/hashicorp/terraform/helper/schema
// * terraform state file (.tfstate)
type Schemaer interface {
	DeleteTerraformResource(d *schema.ResourceData)
	UpdateLocalResourceState(resourceState map[string]interface{},
		d *schema.ResourceData, schemaTools *SchemaTooler) error
	UpdateVdcResourcesNames(d *schema.ResourceData) error
	ReadElement(key interface{}, value interface{}) (interface{}, error)
}

// SchemaSchemaer implements Schemaer interface
type SchemaSchemaer struct{}

// DeleteTerraformResource deletes a resource from terraform state file (.tfstate)
func (schemaer SchemaSchemaer) DeleteTerraformResource(d *schema.ResourceData) {
	d.SetId("")
}

// UpdateLocalResourceState updates resource state in .tfstate file through schema update
func (schemaer SchemaSchemaer) UpdateLocalResourceState(resourceState map[string]interface{},
	d *schema.ResourceData, schemaTools *SchemaTooler) error {
	var (
		updateError error
		readValue   interface{}
	)
	for key, value := range resourceState {
		readValue, updateError = schemaTools.SchemaTools.ReadElement(key, value)
		if key == IDField {
			var sID string
			switch {
			case reflect.TypeOf(value).Kind() == reflect.Float64:
				sID = strconv.FormatFloat(value.(float64), 'f', -1, 64)
			case reflect.TypeOf(value).Kind() == reflect.Int:
				sID = strconv.Itoa(value.(int))
			case reflect.TypeOf(value).Kind() == reflect.String:
				if value == nil {
					sID = ""
				} else {
					sID = value.(string)
				}
			default:
				updateError = errors.New("Format of " + key + "(" +
					reflect.TypeOf(value).Kind().String() + ") not handled.")
			}
			d.SetId(sID)
		} else {
			updateError = d.Set(key, readValue)
		}
		readValue = nil
	}
	return updateError
}

// UpdateVdcResourcesNames trims meaningless part of vdc resource name to store
// a shorter name locally, example :
// * "<enterprise name>-mono-ram" -> "ram"
func (schemaer SchemaSchemaer) UpdateVdcResourcesNames(d *schema.ResourceData) error {
	var (
		vdcResourcesList       = d.Get(VdcResourceField).([]interface{})
		vdcResourcesListUpdate = []interface{}{}
		enterpriseName         = d.Get(EnterpriseField).(string)
		resourceName           string
	)
	for _, resource := range vdcResourcesList {
		resourceName = resource.(map[string]interface{})[ResourceField].(string)
		resourceName = strings.Replace(resourceName,
			enterpriseName, "", 1)
		resourceName = strings.Replace(resourceName,
			monoField, "", 1)
		resource.(map[string]interface{})[ResourceField] = resourceName
		vdcResourcesListUpdate = append(vdcResourcesListUpdate, resource)
	}
	return d.Set(VdcResourceField, vdcResourcesListUpdate)
}

// ReadElement formats Element(key,value) value type to a type accepted by terraform :
//
// * value type -> terraform accepted type
//
// * string -> string
//
// * bool -> bool
//
// * float64 -> int (rounded to nearest int)
//
// * int -> int
//
// * map -> map (recursive call of function for each map element)
//
// * slice -> slice (recursive call of function for each slice element)
//
// * other types -> return error
func (schemaer SchemaSchemaer) ReadElement(key interface{},
	value interface{}) (interface{}, error) {
	var (
		readError error
		readValue interface{}
	)
	switch valueType := value.(type) {
	case string:
		readValue = value.(string)
	case bool:
		readValue = value.(bool)
	case float64:
		readValue = int(value.(float64))
	case int:
		readValue = value.(int)
	case map[string]interface{}:
		var readMapValue map[string]interface{}
		readMapValue = make(map[string]interface{})
		var mapItem interface{}
		for mapKey, mapValue := range valueType {
			mapItem,
				readError = schemaer.ReadElement(mapKey,
				mapValue)
			readMapValue[mapKey] = mapItem
		}
		readValue = readMapValue
	case []interface{}:
		var readListValue []interface{}
		var listItem interface{}
		for listKey, listValue := range valueType {
			listItem,
				readError = schemaer.ReadElement(listKey,
				listValue)
			readListValue = append(readListValue, listItem)
		}
		readValue = readListValue
	default:
		if value == nil {
			readValue = nil
		} else {
			readError = errors.New("Format " +
				reflect.TypeOf(valueType).Kind().String() + " not handled.")
		}
	}
	return readValue, readError
}
