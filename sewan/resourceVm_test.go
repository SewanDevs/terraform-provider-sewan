package sewan

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	sdk "gitlab.com/rd/sewan_go_sdk"
	"testing"
)

func vmCRUDTestInit() (*Client, *schema.ResourceData) {
	vmResource := resourceVM()
	d := vmResource.TestResourceData()
	return ResourceCRUDTestInit(), d
}

func TestResourceVMCreate(t *testing.T) {
	testCases := []struct {
		ID          int
		TC_apier    sdk.APIer
		CreationErr error
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
		metaStruct.sewanAPIImplementerTooler.APIImplementer = testCase.TC_apier
		err = resourceVMCreate(d, metaStruct)
		switch {
		case err == nil || testCase.CreationErr == nil:
			if !(err == nil && testCase.CreationErr == nil) {
				t.Errorf("\n\nTC %d : VM creation error was incorrect,"+
					errTestResultDiffs, testCase.ID, err, testCase.CreationErr)
			}
		case err.Error() != testCase.CreationErr.Error():
			t.Errorf("\n\nTC %d : VM creation error was incorrect,"+
				errTestResultDiffs,
				testCase.ID, err.Error(), testCase.CreationErr.Error())
		}
	}
}

func TestResourceVMRead(t *testing.T) {
	testCases := []struct {
		ID       int
		TC_apier sdk.APIer
		ReadErr  error
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
		metaStruct.sewanAPIImplementerTooler.APIImplementer = testCase.TC_apier
		err = resourceVMRead(d, metaStruct)
		switch {
		case err == nil || testCase.ReadErr == nil:
			if !(err == nil && testCase.ReadErr == nil) {
				t.Errorf(errorTcIDAndWrongVmUpdateError+
					errTestResultDiffs, testCase.ID, err, testCase.ReadErr)
			}
		case err.Error() != testCase.ReadErr.Error():
			t.Errorf(errorTcIDAndWrongVmUpdateError+
				errTestResultDiffs,
				testCase.ID, err.Error(), testCase.ReadErr.Error())
		}
	}
}

func TestResourceVmUpdate(t *testing.T) {
	testCases := []struct {
		ID        int
		TC_apier  sdk.APIer
		UpdateErr error
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
		metaStruct.sewanAPIImplementerTooler.APIImplementer = testCase.TC_apier
		err = resourceVMUpdate(d, metaStruct)
		switch {
		case err == nil || testCase.UpdateErr == nil:
			if !(err == nil && testCase.UpdateErr == nil) {
				t.Errorf(errorTcIDAndWrongVmUpdateError+
					errTestResultDiffs, testCase.ID, err, testCase.UpdateErr)
			}
		case err.Error() != testCase.UpdateErr.Error():
			t.Errorf(errorTcIDAndWrongVmUpdateError+
				errTestResultDiffs,
				testCase.ID, err.Error(), testCase.UpdateErr.Error())
		}
	}
}

func TestResourceVmDelete(t *testing.T) {
	testCases := []struct {
		ID        int
		TC_apier  sdk.APIer
		DeleteErr error
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
		metaStruct.sewanAPIImplementerTooler.APIImplementer = testCase.TC_apier
		err = resourceVMDelete(d, metaStruct)
		switch {
		case err == nil || testCase.DeleteErr == nil:
			if !(err == nil && testCase.DeleteErr == nil) {
				t.Errorf("\n\nTC %d : VM deletion error was incorrect,"+
					errTestResultDiffs, testCase.ID, err, testCase.DeleteErr)
			}
		case err.Error() != testCase.DeleteErr.Error():
			t.Errorf("\n\nTC %d : VM deletion error was incorrect,"+
				errTestResultDiffs,
				testCase.ID, err.Error(), testCase.DeleteErr.Error())
		}
	}
}
