package sewan

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	sdk "gitlab.com/rd/sewan_go_sdk"
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
			VdcSuccessfullCrudOperationsAirDrumAPIer{},
			nil,
		},
		{
			2,
			VdcFailureCrudOperationsAirDrumAPIer{},
			errors.New(vdcCreationFailure),
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
					errTestResultDiffs, testCase.Id, err, testCase.Creation_Err)
			}
		case err.Error() != testCase.Creation_Err.Error():
			t.Errorf("\n\nTC %d : VDC creation error was incorrect,"+
				errTestResultDiffs,
				testCase.Id, err.Error(), testCase.Creation_Err.Error())
		}
	}
}

func TestResourceVdcRead(t *testing.T) {
	testCases := []struct {
		Id       int
		TC_apier sdk.APIer
		Read_Err error
	}{
		{
			1,
			VdcSuccessfullCrudOperationsAirDrumAPIer{},
			nil,
		},
		{
			2,
			VdcFailureCrudOperationsAirDrumAPIer{},
			nil,
		},
		{
			3,
			VdcReadFailureCrudOperationsAirDrumAPIer{},
			errors.New(vdcReadFailure),
		},
		{
			4,
			VdcNotFoundErrorOnReadOperationsAirDrumAPIer{},
			nil,
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
				t.Errorf(errorTcIdAndWrongVdcUpdateError+
					errTestResultDiffs, testCase.Id, err, testCase.Read_Err)
			}
		case err.Error() != testCase.Read_Err.Error():
			t.Errorf(errorTcIdAndWrongVdcUpdateError+
				errTestResultDiffs,
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
			VdcSuccessfullCrudOperationsAirDrumAPIer{},
			nil,
		},
		{
			2,
			VdcFailureCrudOperationsAirDrumAPIer{},
			errors.New(vdcUpdateFailure),
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
				t.Errorf(errorTcIdAndWrongVdcUpdateError+
					errTestResultDiffs, testCase.Id, err, testCase.Update_Err)
			}
		case err.Error() != testCase.Update_Err.Error():
			t.Errorf(errorTcIdAndWrongVdcUpdateError+
				errTestResultDiffs,
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
			VdcSuccessfullCrudOperationsAirDrumAPIer{},
			nil,
		},
		{
			2,
			VdcFailureCrudOperationsAirDrumAPIer{},
			errors.New(vdcDeletionFailure),
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
					errTestResultDiffs, testCase.Id, err, testCase.Delete_Err)
			}
		case err.Error() != testCase.Delete_Err.Error():
			t.Errorf("\n\nTC %d : VDC deletion error was incorrect,"+
				errTestResultDiffs,
				testCase.Id, err.Error(), testCase.Delete_Err.Error())
		}
	}
}
