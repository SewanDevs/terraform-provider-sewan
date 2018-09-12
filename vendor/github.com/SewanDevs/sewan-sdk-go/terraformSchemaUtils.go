package sewansdk

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"reflect"
	"strconv"
	//"strings"
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
	//UpdateVdcResourcesNames(d *schema.ResourceData) error
	ReadElement(key interface{}, value interface{}) (interface{}, error)
}

// SchemaSchemaer implements Schemaer interface
type SchemaSchemaer struct{}

// DeleteTerraformResource deletes a resource from terraform state file (.tfstate)
func (schemaer SchemaSchemaer) DeleteTerraformResource(d *schema.ResourceData) {
	d.SetId("")
}

// UpdateLocalResourceState updates resource state in .tfstate file through schema update
// - *schema.ResourceData Set error are not handled to ignore possible change of
//     data structure got from clouddc environment, and to ignore non-relevant
//     resource data.
func (schemaer SchemaSchemaer) UpdateLocalResourceState(resourceState map[string]interface{},
	d *schema.ResourceData, schemaTools *SchemaTooler) error {
	var (
		err       error
		readValue interface{}
	)
	for key, value := range resourceState {
		readValue, err = schemaTools.SchemaTools.ReadElement(key, value)
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
				err = errors.New("Format of " + key + "(" +
					reflect.TypeOf(value).Kind().String() + ") not handled.")
			}
			d.SetId(sID)
		} else {
			d.Set(key, readValue)
		}
		if err != nil {
			return err
		}
		readValue = nil
	}
	return err
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
		err       error
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
				err = schemaer.ReadElement(mapKey,
				mapValue)
			readMapValue[mapKey] = mapItem
		}
		readValue = readMapValue
	case []interface{}:
		var readListValue []interface{}
		var listItem interface{}
		for listKey, listValue := range valueType {
			listItem,
				err = schemaer.ReadElement(listKey,
				listValue)
			readListValue = append(readListValue, listItem)
		}
		readValue = readListValue
	default:
		if value == nil {
			readValue = nil
		} else {
			err = errors.New("Format " +
				reflect.TypeOf(valueType).Kind().String() + " not handled.")
		}
	}
	return readValue, err
}
