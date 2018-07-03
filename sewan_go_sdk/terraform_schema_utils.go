package sewan_go_sdk

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"reflect"
	"strconv"
)

type SchemaTooler struct {
	SchemaTools Schemaer
}
type Schemaer interface {
	Delete_terraform_resource(d *schema.ResourceData)
	Update_local_resource_state(resource_state map[string]interface{},
		d *schema.ResourceData, schemaTools *SchemaTooler) error
	Read_element(key interface{}, value interface{},
		logger *log.Logger) (interface{}, error)
}
type Schema_Schemaer struct{}

func (schemaer Schema_Schemaer) Delete_terraform_resource(d *schema.ResourceData) {
	d.SetId("")
}

func (schemaer Schema_Schemaer) Update_local_resource_state(resource_state map[string]interface{},
	d *schema.ResourceData, schemaTools *SchemaTooler) error {

	var (
		updateError error = nil
		read_value  interface{}
	)
	logger := LoggerCreate("update_local_resource_state_" +
		d.Get("name").(string) + ".log")
	for key, value := range resource_state {
		read_value,
			updateError = schemaTools.SchemaTools.Read_element(key,
			value,
			logger)
		logger.Println("Set \"", key, "\" to \"", read_value, "\"")
		if key == "id" {
			var s_id string = ""
			switch {
			case reflect.TypeOf(value).Kind() == reflect.Float64:
				s_id = strconv.FormatFloat(value.(float64), 'f', -1, 64)
			case reflect.TypeOf(value).Kind() == reflect.Int:
				s_id = strconv.Itoa(value.(int))
			case reflect.TypeOf(value).Kind() == reflect.String:
				s_id = value.(string)
			default:
				updateError = errors.New("Format of " + key + "(" +
					reflect.TypeOf(value).Kind().String() + ") not handled.")
			}
			d.SetId(s_id)
		} else {
			updateError = d.Set(key, read_value)
		}
		read_value = nil
	}
	return updateError
}

func (schemaer Schema_Schemaer) Read_element(key interface{}, value interface{},
	logger *log.Logger) (interface{}, error) {

	var (
		readError  error = nil
		read_value interface{}
	)
	switch value_type := value.(type) {
	case string:
		read_value = value.(string)
	case bool:
		read_value = value.(bool)
	case float64:
		read_value = int(value.(float64))
	case int:
		read_value = value.(int)
	case map[string]interface{}:
		var read_map_value map[string]interface{}
		read_map_value = make(map[string]interface{})
		var map_item interface{}
		for map_key, map_value := range value_type {
			map_item,
				readError = schemaer.Read_element(map_key,
				map_value,
				logger)
			read_map_value[map_key] = map_item
		}
		read_value = read_map_value
	case []interface{}:
		var read_list_value []interface{}
		var list_item interface{}
		for list_key, list_value := range value_type {
			list_item,
				readError = schemaer.Read_element(list_key,
				list_value,
				logger)
			read_list_value = append(read_list_value, list_item)
		}
		read_value = read_list_value
	default:
		if value == nil {
			read_value = nil
		} else {
			readError = errors.New("Format " +
				reflect.TypeOf(value_type).Kind().String() + " not handled.")
		}
	}
	return read_value, readError
}
