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
		ID          int
		TC_apier    sdk.APIer
		CreationErr error
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
		metaStruct.sewanAPIImplementerTooler.APIImplementer = testCase.TC_apier
		err = resourceVdcCreate(d, metaStruct)
		switch {
		case err == nil || testCase.CreationErr == nil:
			if !(err == nil && testCase.CreationErr == nil) {
				t.Errorf("\n\nTC %d : VDC creation error was incorrect,"+
					errTestResultDiffs, testCase.ID, err, testCase.CreationErr)
			}
		case err.Error() != testCase.CreationErr.Error():
			t.Errorf("\n\nTC %d : VDC creation error was incorrect,"+
				errTestResultDiffs,
				testCase.ID, err.Error(), testCase.CreationErr.Error())
		}
	}
}

func TestResourceVdcRead(t *testing.T) {
	testCases := []struct {
		ID       int
		TC_apier sdk.APIer
		ReadErr  error
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
		metaStruct.sewanAPIImplementerTooler.APIImplementer = testCase.TC_apier
		err = resourceVdcRead(d, metaStruct)
		switch {
		case err == nil || testCase.ReadErr == nil:
			if !(err == nil && testCase.ReadErr == nil) {
				t.Errorf(errorTcIDAndWrongVdcUpdateError+
					errTestResultDiffs, testCase.ID, err, testCase.ReadErr)
			}
		case err.Error() != testCase.ReadErr.Error():
			t.Errorf(errorTcIDAndWrongVdcUpdateError+
				errTestResultDiffs,
				testCase.ID, err.Error(), testCase.ReadErr.Error())
		}
	}
}

func TestResourceVdcUpdate(t *testing.T) {
	testCases := []struct {
		ID        int
		TC_apier  sdk.APIer
		UpdateErr error
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
		metaStruct.sewanAPIImplementerTooler.APIImplementer = testCase.TC_apier
		err = resourceVdcUpdate(d, metaStruct)
		switch {
		case err == nil || testCase.UpdateErr == nil:
			if !(err == nil && testCase.UpdateErr == nil) {
				t.Errorf(errorTcIDAndWrongVdcUpdateError+
					errTestResultDiffs, testCase.ID, err, testCase.UpdateErr)
			}
		case err.Error() != testCase.UpdateErr.Error():
			t.Errorf(errorTcIDAndWrongVdcUpdateError+
				errTestResultDiffs,
				testCase.ID, err.Error(), testCase.UpdateErr.Error())
		}
	}
}

func TestResourceVdcDelete(t *testing.T) {
	testCases := []struct {
		ID        int
		TC_apier  sdk.APIer
		DeleteErr error
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
		metaStruct.sewanAPIImplementerTooler.APIImplementer = testCase.TC_apier
		err = resourceVdcDelete(d, metaStruct)
		switch {
		case err == nil || testCase.DeleteErr == nil:
			if !(err == nil && testCase.DeleteErr == nil) {
				t.Errorf("\n\nTC %d : VDC deletion error was incorrect,"+
					errTestResultDiffs, testCase.ID, err, testCase.DeleteErr)
			}
		case err.Error() != testCase.DeleteErr.Error():
			t.Errorf("\n\nTC %d : VDC deletion error was incorrect,"+
				errTestResultDiffs,
				testCase.ID, err.Error(), testCase.DeleteErr.Error())
		}
	}
}
