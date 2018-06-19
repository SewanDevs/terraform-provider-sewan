package sewan_go_sdk

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"reflect"
)

func Handle_templates_list(d *schema.ResourceData,
	templateList []interface{}) (bool, error) {

	var (
		template_exists     bool   = false
		template_list_valid error  = nil
		template_handle_err error  = nil
		template            string = d.Get("template").(string)
	)
	for i := 0; i < len(templateList); i++ {
		switch reflect.TypeOf(templateList[i]).Kind() {
		case reflect.Map:
			var templateName string = templateList[i].(map[string]interface{})["name"].(string)
			if templateName == template {
				template_exists = true
				template_handle_err = Handle_template_and_set_schema(d, templateList[i].(map[string]interface{}))
				break
			}
		default:
			template_list_valid = errors.New("Wrong template list format.\n" +
				"got :" + reflect.TypeOf(templateList[i]).Kind().String() +
				"want :" + reflect.Map.String())
		}
	}
	if template_handle_err != nil {
		template_list_valid = template_handle_err
	}
	return template_exists, template_list_valid
}

func Handle_template_and_set_schema(d *schema.ResourceData,
	template map[string]interface{}) error {

	var template_handle_err error = nil
	logger := LoggerCreate("Handle_template_and_set_schema" + d.Get("name").(string) + ".log")
	logger.Println("d.Get(\"disks\").([]interface{}) = ",d.Get("disks").([]interface{}))
	logger.Println("d.Get(\"nics\").([]interface{}) = ",d.Get("nics").([]interface{}))

	for template_param_name, template_param_value := range template {
		if reflect.ValueOf(template_param_name).IsValid() && reflect.ValueOf(template_param_value).IsValid() {
			logger.Println("--")
			switch reflect.TypeOf(template_param_value).Kind() {
			case reflect.String:
				logger.Println("case String : ", template_param_name)

				if d.Id() == "" {
					switch {
					case reflect.ValueOf(template_param_name).String() == "os":
						logger.Println("1a")
					case reflect.ValueOf(template_param_name).String() == "name":
						logger.Println("2a")
					default:
						if d.Get(reflect.ValueOf(template_param_name).String()) == ""{
							logger.Println("3a : ", reflect.ValueOf(template_param_value).String())
							d.Set(reflect.ValueOf(template_param_name).String(),
								reflect.ValueOf(template_param_value).String())
						}
					}
				} else {
					switch {
					case reflect.ValueOf(template_param_name).String() == "name":
						logger.Println("2b")
					default:
						if d.Get(reflect.ValueOf(template_param_name).String()) == ""{
							logger.Println("3a : ", reflect.ValueOf(template_param_value).String())
							d.Set(reflect.ValueOf(template_param_name).String(),
								reflect.ValueOf(template_param_value).String())
						}
					}
				}

			case reflect.Float64:
				logger.Println("case float 64 : ", template_param_name, " = ",
					d.Get(reflect.ValueOf(template_param_name).String()))
				if d.Id() == "" {
					switch {
					case reflect.ValueOf(template_param_name).String() == "id":
						logger.Println("2, d.Id() = ",d.Id())
					default:
						if d.Get(reflect.ValueOf(template_param_name).String()).(int) == 0{
							logger.Println("3, val to set = ",
								int(reflect.ValueOf(template_param_value).Interface().(float64)))
							d.Set(reflect.ValueOf(template_param_name).String(),
								int(reflect.ValueOf(template_param_value).Interface().(float64)))
						}
					}
				} else {
					switch {
					case reflect.ValueOf(template_param_name).String() == "id":
						logger.Println("2, d.Id() = ",d.Id())
					default:
						if d.Get(reflect.ValueOf(template_param_name).String()) == 0 {
							if d.Get(reflect.ValueOf(template_param_name).String()).(int) == 0{
								logger.Println("3, val to set = ",
									int(reflect.ValueOf(template_param_value).Interface().(float64)))
								d.Set(reflect.ValueOf(template_param_name).String(),
									int(reflect.ValueOf(template_param_value).Interface().(float64)))
							}
						}
					}
				}
			case reflect.Int:
				logger.Println("case Int : ", template_param_name, " = ",
					d.Get(reflect.ValueOf(template_param_name).String()))
				if d.Id() == "" {
					switch {
					case reflect.ValueOf(template_param_name).String() == "id":
						logger.Println("2")
					default:
						if d.Get(reflect.ValueOf(template_param_name).String()).(int) == 0{
							logger.Println("3, val to set = ",
								int(reflect.ValueOf(template_param_value).Interface().(int)))
							d.Set(reflect.ValueOf(template_param_name).String(),
								int(reflect.ValueOf(template_param_value).Interface().(int)))
						}
					}
				} else {
					switch {
					case reflect.ValueOf(template_param_name).String() == "id":
						logger.Println("2")
					default:
						if d.Get(reflect.ValueOf(template_param_name).String()).(int) == 0{
							logger.Println("3, val to set = ",
								int(reflect.ValueOf(template_param_value).Interface().(int)))
							d.Set(reflect.ValueOf(template_param_name).String(),
								int(reflect.ValueOf(template_param_value).Interface().(int)))
						}
					}
				}
			case reflect.Slice:
				logger.Println("case Slice : ", template_param_name, " = ",
					d.Get(reflect.ValueOf(template_param_name).String()))
				var (
					schema_slice []interface{}
					map_item     interface{}
				)
				if d.Id() != "" {
					if len(d.Get(reflect.ValueOf(template_param_name).String()).([]interface{})) == 0 {
						logger.Println("6, slice elem number = 0")
						//val, err := read_element(template_param_name, template_param_value, logger)
						//logger.Println("val, err = ", val, err)
						//err = d.Set(reflect.ValueOf(template_param_name).String(), val)

					if template_param_name == "disks"{
					for _, template_slice_element := range template_param_value.([]interface{}) {
						logger.Println("template_slice_element = ",template_slice_element)

						//template_slice_element.(map[string]interface{})["mac_address"] = ""

						schema_slice = append(schema_slice,
						template_slice_element.(map[string]interface{}))
					}
				}
					err := d.Set(reflect.ValueOf(template_param_name).String(), schema_slice)
					logger.Println("set err = ", err)
					logger.Println("set val = ", d.Get(reflect.ValueOf(template_param_name).String()))

					} else {
						var (
							elem_id      string = ""
						)
						schema_slice = d.Get(reflect.ValueOf(template_param_name).String()).([]interface{})
						logger.Println("schema_slice init=", schema_slice)
						for _, template_slice_element := range template_param_value.([]interface{}) {
							var (
								elem_already_in_list = false
							)
							if template_param_name == "nics" {
								elem_id = "vlan"
							} else if template_param_name == "disks" {
								elem_id = "name"
							}
							for schema_slice_index, schema_slice_element := range schema_slice {
								if template_slice_element.(map[string]interface{})[elem_id] == schema_slice_element.(map[string]interface{})[elem_id] {
									for map_key, map_value := range schema_slice_element.(map[string]interface{}) {
										if map_key != "mac_address"{
											logger.Println("map_key, map_value =",map_key, map_value)
											map_item, _ = read_element(map_key, map_value, logger)
											logger.Println("schema_slice[",schema_slice_index,
												"].(map[string]interface{})[",map_key,"] = ",map_item)
											schema_slice[schema_slice_index].(map[string]interface{})[map_key] = map_item
										}
									}
									elem_already_in_list = true
								}
							}
							if elem_already_in_list == false {
								schema_slice = append(schema_slice,
									template_slice_element.(map[string]interface{}))
							}
						}
						logger.Println("schema_slice =", schema_slice)
						d.Set(reflect.ValueOf(template_param_name).String(), schema_slice)
					}
				} else {
					if template_param_name == "nics" {
						if len(d.Get(reflect.ValueOf(template_param_name).String()).([]interface{})) == 0 {
							logger.Println("6, slice elem number = 0")
							val, err := read_element(template_param_name, template_param_value, logger)
							logger.Println("val, err = ", val, err)
							err = d.Set(reflect.ValueOf(template_param_name).String(), val)
							logger.Println("set err = ", err)
							logger.Println("set val = ", d.Get(reflect.ValueOf(template_param_name).String()))
						} else {
							var (
								map_item     interface{}
								schema_slice []interface{} = d.Get(reflect.ValueOf(template_param_name).String()).([]interface{})
							)
							logger.Println("schema_slice init=", schema_slice)
							for _, template_slice_element := range template_param_value.([]interface{}) {
								var (
									elem_already_in_list = false
								)
								for schema_slice_index, schema_slice_element := range schema_slice {
									if template_slice_element.(map[string]interface{})[template_param_name] == schema_slice_element.(map[string]interface{})[template_param_name] {
										for map_key, map_value := range schema_slice_element.(map[string]interface{}) {
											if map_key != "mac_address" {
												map_item, _ = read_element(map_key, map_value, logger)
												schema_slice[schema_slice_index].(map[string]interface{})[map_key] = map_item
											}
										}
										elem_already_in_list = true
									}
								}
								if elem_already_in_list == false {
									schema_slice = append(schema_slice,
										template_slice_element.(map[string]interface{}))
								}
							}
							logger.Println("schema_slice =", schema_slice)
							d.Set(reflect.ValueOf(template_param_name).String(), schema_slice)
						}
					}
				}
			default:
				template_handle_err = errors.New("Handle_template_and_set_schema : Format of " + template_param_name + "(" +
					reflect.TypeOf(template_param_value).Kind().String() + ") not handled.")
			}
			if template_handle_err != nil {
				logger.Println(template_param_name, "=",
					d.Get(reflect.ValueOf(template_param_name).String()),
					"error :", template_handle_err)
				break
			}
		}
	}
	logger.Println("d.Get(\"disks\").([]interface{}) = ",d.Get("disks").([]interface{}))
	logger.Println("d.Get(\"nics\").([]interface{}) = ",d.Get("nics").([]interface{}))
	return template_handle_err
}
