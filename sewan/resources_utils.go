package sewan

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"reflect"
	"strconv"
	sdk "terraform-provider-sewan/sewan_go_sdk"
)

func Delete_resource(d *schema.ResourceData) {
	d.SetId("")
}

func Update_local_resource_state(body map[string]interface{},
	d *schema.ResourceData) error {

	var updateError, err error
	updateError = nil
	var read_value interface{}
	logger := sdk.LoggerCreate("update_local_resource_state_" +
		d.Get("name").(string) + ".log")
	for key, value := range body {
		read_value, err = read_element(key, value, logger)
		logger.Println("Set \"", key, "\" to \"", read_value, "\"")
		if key == "id" {
			var s_id string
			switch {
			case reflect.TypeOf(value).Kind().String() == "float64":
				s_id = strconv.FormatFloat(value.(float64), 'f', -1, 64)
			case reflect.TypeOf(value).Kind().String() == "int":
				s_id = strconv.Itoa(value.(int))
			case reflect.TypeOf(value).Kind().String() == "string":
				s_id = value.(string)
			default:
				err = errors.New("Format of " + key + "not handled.")
			}
			d.SetId(s_id)
		} else {
			d.Set(key, read_value)
		}
		if err != nil {
			updateError = err
		}
		read_value = nil
	}
	return updateError
}

func read_element(key interface{}, value interface{},
	logger *log.Logger) (interface{}, error) {

	var readError error
	readError = nil
	var read_value interface{}
	switch value_type := value.(type) {
	case string:
		read_value = value
	case bool:
		read_value = value
	case float64:
		read_value = strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case int:
		read_value = value
	case map[string]interface{}:
		logger.Println(key, " is a map of type:", reflect.TypeOf(value))
		logger.Println("value :", value)
		logger.Println("value_type :", value_type)
		var read_map_value map[string]interface{}
		read_map_value = make(map[string]interface{})
		var map_item interface{}
		for map_key, map_value := range value_type {
			logger.Println(" --- map iterate : ", map_key, "=", map_value)
			map_item, readError = read_element(map_key, map_value, logger)
			read_map_value[map_key] = map_item
			logger.Println(" --- map iterate : ", read_map_value, "(",
				reflect.TypeOf(read_map_value), ")")
		}
		read_value = read_map_value
	case []interface{}:
		var read_list_value []interface{}
		var list_item interface{}
		for list_key, list_value := range value_type {
			logger.Println(" --- list iterate : ", list_key, "=", list_value)
			list_item, readError = read_element(list_key, list_value, logger)
			read_list_value = append(read_list_value, list_item)
			logger.Println(" --- list iterate : read_list_value (type)",
				read_list_value, "(", reflect.TypeOf(read_list_value), ")")
		}
		read_value = read_list_value
	default:
		if value == nil {
			read_value = nil
		} else {
			readError = errors.New("Unhandled value type")
		}
	}
	return read_value, readError
}
