package sewan_go_sdk

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"reflect"
)

const (
	NICS_PARAM     = "nics"
	DISKS_PARAM    = "disks"
	TEMPLATE_FIELD = "template"
	NAME_FIELD     = "name"
)

type TemplatesTooler struct {
	TemplatesTools Templater
}
type Templater interface {
	FetchTemplateFromList(template_name string,
		templateList []interface{}) (map[string]interface{}, error)
	UpdateSchema(d *schema.ResourceData,
		template map[string]interface{},
		templatesTooler *TemplatesTooler) error
	UpdateSchemaDisks(d *schema.ResourceData,
		disks []interface{}) error
	UpdateSchemaNics(d *schema.ResourceData) error
}
type Template_Templater struct{}

func (templater Template_Templater) FetchTemplateFromList(template_name string,
	templateList []interface{}) (map[string]interface{}, error) {

	var (
		template            map[string]interface{} = nil
		template_list_valid error                  = nil
	)
	for i := 0; i < len(templateList); i++ {
		switch reflect.TypeOf(templateList[i]).Kind() {
		case reflect.Map:
			var list_template_name string = templateList[i].(map[string]interface{})[NAME_FIELD].(string)
			if list_template_name == template_name {
				template = templateList[i].(map[string]interface{})
				break
			}
		default:
			template_list_valid = errors.New("Wrong template list format.\n" +
				"got :" + reflect.TypeOf(templateList[i]).Kind().String() +
				"want :" + reflect.Map.String())
		}
	}
	if template == nil {
		template_list_valid = errors.New("Template \"" + template_name +
			"\" does not exists, please validate it's name.")
	}
	return template, template_list_valid
}

func (templater Template_Templater) UpdateSchema(d *schema.ResourceData,
	template map[string]interface{},
	templatesTooler *TemplatesTooler) error {

	var template_handle_err error = nil
	logger := LoggerCreate("UpdateSchema" + d.Get("name").(string) + ".log")
	logger.Println("d.Get(\"disks\").([]interface{}) = ", d.Get("disks").([]interface{}))
	logger.Println("d.Get(\"nics\").([]interface{}) = ", d.Get("nics").([]interface{}))

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
						if d.Get(reflect.ValueOf(template_param_name).String()) == "" {
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
						if d.Get(reflect.ValueOf(template_param_name).String()) == "" {
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
						logger.Println("2, d.Id() = ", d.Id())
					default:
						if d.Get(reflect.ValueOf(template_param_name).String()).(int) == 0 {
							logger.Println("3, val to set = ",
								int(reflect.ValueOf(template_param_value).Interface().(float64)))
							d.Set(reflect.ValueOf(template_param_name).String(),
								int(reflect.ValueOf(template_param_value).Interface().(float64)))
						}
					}
				} else {
					switch {
					case reflect.ValueOf(template_param_name).String() == "id":
						logger.Println("2, d.Id() = ", d.Id())
					default:
						if d.Get(reflect.ValueOf(template_param_name).String()) == 0 {
							if d.Get(reflect.ValueOf(template_param_name).String()).(int) == 0 {
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
						if d.Get(reflect.ValueOf(template_param_name).String()).(int) == 0 {
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
						if d.Get(reflect.ValueOf(template_param_name).String()).(int) == 0 {
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
				switch {
				case template_param_name == NICS_PARAM:
					templatesTooler.TemplatesTools.UpdateSchemaNics(d)
				case template_param_name == DISKS_PARAM:
					templatesTooler.TemplatesTools.UpdateSchemaDisks(d,
						template_param_value.([]interface{}))
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
			default:
			}
		}
	}
	logger.Println("d.Get(\"disks\").([]interface{}) = ", d.Get("disks").([]interface{}))
	logger.Println("d.Get(\"nics\").([]interface{}) = ", d.Get("nics").([]interface{}))
	return template_handle_err
}

func (templater Template_Templater) UpdateSchemaDisks(d *schema.ResourceData,
	disks []interface{}) error {

	var (
		template_name = d.Get(TEMPLATE_FIELD).(string)
		schema_slice  []interface{}
	)
	logger := LoggerCreate("UpdateSchemaDisks" + d.Get("name").(string) + ".log")
	logger.Println("case disks")
	if d.Id() != "" {
		if len(d.Get(DISKS_PARAM).([]interface{})) == 0 {
			for _, template_slice_element := range disks {
				logger.Println("template_slice_element = ", template_slice_element)
				schema_slice = append(schema_slice,
					template_slice_element.(map[string]interface{}))
			}
			err := d.Set(reflect.ValueOf(template_name).String(), schema_slice)
			logger.Println("set err = ", err)
			logger.Println("set val = ", d.Get(DISKS_PARAM))

		} else {
			schema_slice = d.Get(DISKS_PARAM).([]interface{})
			logger.Println("schema_slice init=", schema_slice)
			for _, template_slice_element := range disks {
				var (
					elem_already_in_list = false
				)
				for schema_slice_index, schema_slice_element := range schema_slice {
					if template_slice_element.(map[string]interface{})[NAME_FIELD] == schema_slice_element.(map[string]interface{})[NAME_FIELD] {
						for map_key, map_value := range schema_slice_element.(map[string]interface{}) {
							logger.Println("map_key, map_value =", map_key, map_value)
							map_item, _ := read_element(map_key, map_value, logger)
							logger.Println("schema_slice[", schema_slice_index,
								"].(map[string]interface{})[", map_key, "] = ", map_item)
							schema_slice[schema_slice_index].(map[string]interface{})[map_key] = map_item
						}
						elem_already_in_list = true
					}
				}
				if elem_already_in_list == false {
					schema_slice = append(schema_slice,
						template_slice_element.(map[string]interface{}))
				}
			}
		}
	} else {
	}
	logger.Println("schema_slice =", schema_slice)
	d.Set(DISKS_PARAM, schema_slice)
	return nil
}

func (templater Template_Templater) UpdateSchemaNics(d *schema.ResourceData) error {

	return nil
}
