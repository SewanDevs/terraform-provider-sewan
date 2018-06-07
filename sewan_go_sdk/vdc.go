package sewan_go_sdk

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type VDC struct {
	Name string `json:"name"`
}

func vdcInstanceCreate(d *schema.ResourceData) VDC {
	return VDC{
		Name: d.Get("name").(string),
	}
}

//------------------------------------------------------------------------------
func (apier AirDrumAPIer) Create_vdc_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	sewan *API) (error, map[string]interface{}) {

	return nil,nil
}

//------------------------------------------------------------------------------
func (apier AirDrumAPIer) Read_vdc_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	sewan *API) (error, map[string]interface{}, bool) {

  return nil,nil,false
}

//------------------------------------------------------------------------------
func (apier AirDrumAPIer) Update_vdc_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	sewan *API) (error) {

  return nil
}

//------------------------------------------------------------------------------
func (apier AirDrumAPIer) Delete_vdc_resource(d *schema.ResourceData,
	clientTooler *ClientTooler,
	sewan *API) (error) {

  return nil
}
