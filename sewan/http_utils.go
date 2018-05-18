package sewan

import (
  "errors"
  "github.com/hashicorp/terraform/helper/schema"
  "log"
  "reflect"
  "strconv"
)

func update_local_resource_state(body map[string]interface{}, logger *log.Logger, d *schema.ResourceData) error {
	var returnError error
	returnError = nil
	var read_value interface{}
	for key, value := range body {
		read_value,returnError = read_element(key, value, logger)
		logger.Println("Set \"", key, "\" to \"", read_value, "\"")
    if key=="id"{
      s_id := strconv.FormatFloat(value.(float64), 'f', -1, 64)
      d.SetId(s_id)
    }else{
      d.Set(key, read_value)
    }
		read_value = nil
	}
	return returnError
}

func read_element(key interface{}, value interface{}, logger *log.Logger) (interface{}, error) {
  var returnError error
	returnError = nil
  var read_value interface{}
	switch value_type := value.(type) {
	case string:
		read_value = value
	case float64:
		read_value = strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case map[string]interface{}:
		logger.Println(key, " is a map of type:", reflect.TypeOf(value))
		logger.Println("value :", value)
		logger.Println("value_type :", value_type)
    var read_map_value map[string]interface{}
    read_map_value = make(map[string]interface{})
    var map_item interface{}
    for map_key,map_value := range value_type{
      logger.Println(" --- map iterate : ",map_key,"=",map_value)
      map_item,returnError = read_element(map_key, map_value, logger)
      read_map_value[map_key] = map_item
      logger.Println(" --- map iterate : ",read_map_value,"(",reflect.TypeOf(read_map_value),")")
    }
    read_value = read_map_value
	case []interface{}:
		logger.Println(key, " is a is a list ([]interface{}) of type:", reflect.TypeOf(value))
		logger.Println("value :", value)
		logger.Println("value_type :", value_type)
    var read_list_value []interface{}
    var list_item interface{}
    for list_key,list_value := range value_type{
      logger.Println(" --- list iterate : ",list_key,"=",list_value)
      list_item,returnError = read_element(list_key, list_value, logger)
      read_list_value = append(read_list_value, list_item)
      logger.Println(" --- list iterate : read_list_value (type)",read_list_value,"(",reflect.TypeOf(read_list_value),")")
    }
    read_value = read_list_value
	default:
		if value == nil {
			read_value = nil
		} else {
			returnError = errors.New("Not able to fetch the value of" + key.(string) + "field.")
		}
	}
	return read_value,returnError
}
