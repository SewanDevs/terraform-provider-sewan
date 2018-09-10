package sewan

import (
	sdk "github.com/SewanDevs/sewan-sdk-go"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestClientStruct(t *testing.T) {
	testCases := []struct {
		ID               int
		TCConfigStruct   configStruct
		TCAPIInitialyser sdk.APIInitialyser
		TCclientStruct   *clientStruct
		TCErr            error
	}{
		{
			1,
			configStruct{
				unitTestToken,
				unitTestAPIURL,
				unitTestEnterprise,
			},
			initSuccess{},
			resourceTestInit(),
			nil,
		},
		{
			2,
			configStruct{
				unitTestToken,
				unitTestAPIURL,
				unitTestEnterprise,
			},
			getEnvMetaFailure{},
			nil,
			errGetEnvMetaFailure,
		},
		{
			3,
			configStruct{
				unitTestToken,
				unitTestAPIURL,
				unitTestEnterprise,
			},
			checkCloudDcStatusFailure{},
			nil,
			errCheckCloudDcStatusFailure,
		},
	}
	apiTooler := sdk.APITooler{}
	for _, testCase := range testCases {
		apiTooler.Initialyser = testCase.TCAPIInitialyser
		clientStruct, err := testCase.TCConfigStruct.clientStruct(&apiTooler)
		diffs := cmp.Diff(clientStruct, testCase.TCclientStruct)
		switch {
		case err == nil || testCase.TCErr == nil:
			if !(err == nil && testCase.TCErr == nil) {
				t.Errorf("\n\nTC %d : TestClientStruct error was incorrect,"+
					errTestResultDiffs, testCase.ID, err, testCase.TCErr)
			}
			if diffs != "" {
				t.Errorf("\n\nTC %d : Wrong TestClientStruct returned structure (-got +want) \n%s",
					testCase.ID, diffs)
			}
		case err.Error() != testCase.TCErr.Error():
			t.Errorf("\n\nTC %d : TestClientStruct error was incorrect,"+
				errTestResultDiffs,
				testCase.ID, err.Error(), testCase.TCErr.Error())
		case diffs != "":
			t.Errorf("\n\nTC %d : Wrong TestClientStruct returned structure (-got +want) \n%s",
				testCase.ID, diffs)
		}
	}
}
