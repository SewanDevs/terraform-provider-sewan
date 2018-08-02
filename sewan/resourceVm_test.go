package sewan

import (
	"errors"
	sdk "gitlab.com/sewan_go_sdk"
	"github.com/hashicorp/terraform/helper/schema"
	"testing"
)

func vmCRUDTestInit() (*Client,*schema.ResourceData) {
	vmResource := resourceVm()
	d := vmResource.TestResourceData()
	return ResourceCRUDTestInit(),d
}

func TestResourceVmCreate(t *testing.T) {
	testCases := []struct {
		Id           int
		TC_apier     sdk.APIer
		Creation_Err error
	}{
		{
			1,
			VM_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VM_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VM_CREATION_FAILURE),
		},
	}
	var (
		err error
		metaStruct *Client
		d *schema.ResourceData
	)
	metaStruct,d = vmCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVmCreate(d, metaStruct)
		switch {
		case err == nil || testCase.Creation_Err == nil:
			if !(err == nil && testCase.Creation_Err == nil) {
				t.Errorf("\n\nTC %d : VM creation error was incorrect,"+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Creation_Err)
			}
		case err.Error() != testCase.Creation_Err.Error():
			t.Errorf("\n\nTC %d : VM creation error was incorrect,"+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Creation_Err.Error())
		}
	}
}

func TestResourceVmRead(t *testing.T) {
	testCases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Read_Err   error
		Res_exists bool
	}{
		{
			1,
			VM_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
			true,
		},
		{
			2,
			VM_failure_CRUD_operations_AirDrumAPIer{},
			nil,
			false,
		},
		{
			3,
			VM_readfailure_CRUD_operations_AirDrumAPIer{},
			errors.New(VM_READ_FAILURE),
			false,
		},
	}
	var (
		err error
		metaStruct *Client
		d *schema.ResourceData
	)
	metaStruct,d = vmCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVmRead(d, metaStruct)
		switch {
		case err == nil || testCase.Read_Err == nil:
			if !(err == nil && testCase.Read_Err == nil) {
				t.Errorf(ERROR_TC_ID_AND_WRONG_VM_UPDATE_ERR+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Read_Err)
			}
		case err.Error() != testCase.Read_Err.Error():
			t.Errorf(ERROR_TC_ID_AND_WRONG_VM_UPDATE_ERR+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Read_Err.Error())
		}
	}
}

func TestResourceVmUpdate(t *testing.T) {
	testCases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Update_Err error
	}{
		{
			1,
			VM_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VM_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VM_UPDATE_FAILURE),
		},
	}
	var (
		err error
		metaStruct *Client
		d *schema.ResourceData
	)
	metaStruct,d = vmCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVmUpdate(d, metaStruct)
		switch {
		case err == nil || testCase.Update_Err == nil:
			if !(err == nil && testCase.Update_Err == nil) {
				t.Errorf(ERROR_TC_ID_AND_WRONG_VM_UPDATE_ERR+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Update_Err)
			}
		case err.Error() != testCase.Update_Err.Error():
			t.Errorf(ERROR_TC_ID_AND_WRONG_VM_UPDATE_ERR+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Update_Err.Error())
		}
	}
}

func TestResourceVmDelete(t *testing.T) {
	testCases := []struct {
		Id         int
		TC_apier   sdk.APIer
		Delete_Err error
	}{
		{
			1,
			VM_successfull_CRUD_operations_AirDrumAPIer{},
			nil,
		},
		{
			2,
			VM_failure_CRUD_operations_AirDrumAPIer{},
			errors.New(VM_DELETION_FAILURE),
		},
	}
	var (
		err error
		metaStruct *Client
		d *schema.ResourceData
	)
	metaStruct,d = vmCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVmDelete(d, metaStruct)
		switch {
		case err == nil || testCase.Delete_Err == nil:
			if !(err == nil && testCase.Delete_Err == nil) {
				t.Errorf("\n\nTC %d : VM deletion error was incorrect,"+
					ERROR_TEST_RESULT_DIFFS, testCase.Id, err, testCase.Delete_Err)
			}
		case err.Error() != testCase.Delete_Err.Error():
			t.Errorf("\n\nTC %d : VM deletion error was incorrect,"+
				ERROR_TEST_RESULT_DIFFS,
				testCase.Id, err.Error(), testCase.Delete_Err.Error())
		}
	}
}
