package sewan

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	sdk "gitlab.com/sewan_go_sdk"
	"testing"
)

func vmCRUDTestInit() (*Client, *schema.ResourceData) {
	vmResource := resourceVm()
	d := vmResource.TestResourceData()
	return ResourceCRUDTestInit(), d
}

func TestResourceVmCreate(t *testing.T) {
	testCases := []struct {
		Id           int
		TC_apier     sdk.APIer
		Creation_Err error
	}{
		{
			1,
			VMSuccessfullCrudOperationsAirDrumAPIer{},
			nil,
		},
		{
			2,
			VMFailureCrudOperationsAirDrumAPIer{},
			errors.New(vdcCreationFailure),
		},
	}
	var (
		err        error
		metaStruct *Client
		d          *schema.ResourceData
	)
	metaStruct, d = vmCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVmCreate(d, metaStruct)
		switch {
		case err == nil || testCase.Creation_Err == nil:
			if !(err == nil && testCase.Creation_Err == nil) {
				t.Errorf("\n\nTC %d : VM creation error was incorrect,"+
					errTestResultDiffs, testCase.Id, err, testCase.Creation_Err)
			}
		case err.Error() != testCase.Creation_Err.Error():
			t.Errorf("\n\nTC %d : VM creation error was incorrect,"+
				errTestResultDiffs,
				testCase.Id, err.Error(), testCase.Creation_Err.Error())
		}
	}
}

func TestResourceVmRead(t *testing.T) {
	testCases := []struct {
		Id       int
		TC_apier sdk.APIer
		Read_Err error
	}{
		{
			1,
			VMSuccessfullCrudOperationsAirDrumAPIer{},
			nil,
		},
		{
			2,
			VMFailureCrudOperationsAirDrumAPIer{},
			nil,
		},
		{
			3,
			VMReadFailureCrudOperationsAirDrumAPIer{},
			errors.New(vdcReadFailure),
		},
		{
			4,
			VMNotFoundErrorOnReadOperationsAirDrumAPIer{},
			nil,
		},
	}
	var (
		err        error
		metaStruct *Client
		d          *schema.ResourceData
	)
	metaStruct, d = vmCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVmRead(d, metaStruct)
		switch {
		case err == nil || testCase.Read_Err == nil:
			if !(err == nil && testCase.Read_Err == nil) {
				t.Errorf(errorTcIdAndWrongVmUpdateError+
					errTestResultDiffs, testCase.Id, err, testCase.Read_Err)
			}
		case err.Error() != testCase.Read_Err.Error():
			t.Errorf(errorTcIdAndWrongVmUpdateError+
				errTestResultDiffs,
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
			VMSuccessfullCrudOperationsAirDrumAPIer{},
			nil,
		},
		{
			2,
			VMFailureCrudOperationsAirDrumAPIer{},
			errors.New(vdcUpdateFailure),
		},
	}
	var (
		err        error
		metaStruct *Client
		d          *schema.ResourceData
	)
	metaStruct, d = vmCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVmUpdate(d, metaStruct)
		switch {
		case err == nil || testCase.Update_Err == nil:
			if !(err == nil && testCase.Update_Err == nil) {
				t.Errorf(errorTcIdAndWrongVmUpdateError+
					errTestResultDiffs, testCase.Id, err, testCase.Update_Err)
			}
		case err.Error() != testCase.Update_Err.Error():
			t.Errorf(errorTcIdAndWrongVmUpdateError+
				errTestResultDiffs,
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
			VMSuccessfullCrudOperationsAirDrumAPIer{},
			nil,
		},
		{
			2,
			VMFailureCrudOperationsAirDrumAPIer{},
			errors.New(vmDeletionFailure),
		},
	}
	var (
		err        error
		metaStruct *Client
		d          *schema.ResourceData
	)
	metaStruct, d = vmCRUDTestInit()
	for _, testCase := range testCases {
		metaStruct.sewanApiTooler.Api = testCase.TC_apier
		err = resourceVmDelete(d, metaStruct)
		switch {
		case err == nil || testCase.Delete_Err == nil:
			if !(err == nil && testCase.Delete_Err == nil) {
				t.Errorf("\n\nTC %d : VM deletion error was incorrect,"+
					errTestResultDiffs, testCase.Id, err, testCase.Delete_Err)
			}
		case err.Error() != testCase.Delete_Err.Error():
			t.Errorf("\n\nTC %d : VM deletion error was incorrect,"+
				errTestResultDiffs,
				testCase.Id, err.Error(), testCase.Delete_Err.Error())
		}
	}
}
