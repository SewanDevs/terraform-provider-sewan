package sewan

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	sdk "gitlab.com/sewan_go_sdk"
	"testing"
)

func vdcCRUDTestInit() (*Client, *schema.ResourceData) {
	vdcResource := resourceVdc()
	d := vdcResource.TestResourceData()
	return ResourceCRUDTestInit(), d
}

func TestResourceVdcCreate(t *testing.T) {
	testCases := []struct {
		Id           int
		TC_apier     sdk.APIer
		Creation_Err error
	}{
		{
			1,
			VDC_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VDC_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VDC_CREATION_FAILURE),
		},
	}
	var (
		err        error
		metaStruct *Client
		d          *schema.ResourceData
	)
	metaStruct, d = vdcCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVdcCreate(d, metaStruct)
		switch {
		case err == nil || testCase.Creation_Err == nil:
			if !(err == nil && testCase.Creation_Err == nil) {
				t.Errorf("\n\nTC %d : VDC creation error was incorrect,"+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Creation_Err)
			}
		case err.Error() != testCase.Creation_Err.Error():
			t.Errorf("\n\nTC %d : VDC creation error was incorrect,"+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Creation_Err.Error())
		}
	}
}

func TestResourceVdcRead(t *testing.T) {
	testCases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Read_Err   error
		Res_exists bool
	}{
		{
			1,
			VDC_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
			true,
		},
		{
			2,
			VDC_failure_CRUD_operations_AirDrumAPIer{},
			nil,
			false,
		},
		{
			3,
			VDC_readfailure_CRUD_operations_AirDrumAPIer{},
			errors.New(VDC_READ_FAILURE),
			false,
		},
	}
	var (
		err        error
		metaStruct *Client
		d          *schema.ResourceData
	)
	metaStruct, d = vdcCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVdcRead(d, metaStruct)
		switch {
		case err == nil || testCase.Read_Err == nil:
			if !(err == nil && testCase.Read_Err == nil) {
				t.Errorf(ERROR_TC_ID_AND_WRONG_VDC_UPDATE_ERR+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Read_Err)
			}
		case err.Error() != testCase.Read_Err.Error():
			t.Errorf(ERROR_TC_ID_AND_WRONG_VDC_UPDATE_ERR+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Read_Err.Error())
		}
	}
}

func TestResourceVdcUpdate(t *testing.T) {
	testCases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Update_Err error
	}{
		{
			1,
			VDC_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VDC_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VDC_UPDATE_FAILURE),
		},
	}
	var (
		err        error
		metaStruct *Client
		d          *schema.ResourceData
	)
	metaStruct, d = vdcCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVdcUpdate(d, metaStruct)
		switch {
		case err == nil || testCase.Update_Err == nil:
			if !(err == nil && testCase.Update_Err == nil) {
				t.Errorf(ERROR_TC_ID_AND_WRONG_VDC_UPDATE_ERR+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Update_Err)
			}
		case err.Error() != testCase.Update_Err.Error():
			t.Errorf(ERROR_TC_ID_AND_WRONG_VDC_UPDATE_ERR+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Update_Err.Error())
		}
	}
}

func TestResourceVdcDelete(t *testing.T) {
	testCases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Delete_Err error
	}{
		{
			1,
			VDC_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VDC_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VDC_DELETION_FAILURE),
		},
	}
	var (
		err        error
		metaStruct *Client
		d          *schema.ResourceData
	)
	metaStruct, d = vdcCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVdcDelete(d, metaStruct)
		switch {
		case err == nil || testCase.Delete_Err == nil:
			if !(err == nil && testCase.Delete_Err == nil) {
				t.Errorf("\n\nTC %d : VDC deletion error was incorrect,"+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Delete_Err)
			}
		case err.Error() != testCase.Delete_Err.Error():
			t.Errorf("\n\nTC %d : VDC deletion error was incorrect,"+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Delete_Err.Error())
		}
	}
}
