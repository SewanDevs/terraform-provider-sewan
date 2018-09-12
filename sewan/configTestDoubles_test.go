package sewan

import (
	sdk "github.com/SewanDevs/sewan-sdk-go"
	"net/http"
)

type initSuccess struct{}

func (initialyser initSuccess) CheckCloudDcStatus(api *sdk.API,
	clientTooler *sdk.ClientTooler,
	resourceTooler *sdk.ResourceTooler) error {
	return nil
}

func (initialyser initSuccess) GetClouddcEnvMeta(api *sdk.API,
	clientTooler *sdk.ClientTooler) (*sdk.APIMeta, error) {
	return &sdk.APIMeta{
		NonCriticalResourceList: unitTestNonCriticalResourceList,
		CriticalResourceList:    unitTestCriticalResourceList,
		OtherResourceList:       unitTestOtherResourceList,
	}, nil
}

func (initialyser initSuccess) New(token string, url string, enterprise string) *sdk.API {
	return &sdk.API{
		Token:      unitTestToken,
		URL:        unitTestAPIURL,
		Enterprise: unitTestEnterprise,
		Client:     &http.Client{},
	}
}

type getEnvMetaFailure struct{}

func (initialyser getEnvMetaFailure) CheckCloudDcStatus(api *sdk.API,
	clientTooler *sdk.ClientTooler,
	resourceTooler *sdk.ResourceTooler) error {
	return nil
}

func (initialyser getEnvMetaFailure) GetClouddcEnvMeta(api *sdk.API,
	clientTooler *sdk.ClientTooler) (*sdk.APIMeta, error) {
	return nil, errGetEnvMetaFailure
}

func (initialyser getEnvMetaFailure) New(token string, url string, enterprise string) *sdk.API {
	return &sdk.API{
		Token:      unitTestToken,
		URL:        unitTestAPIURL,
		Enterprise: unitTestEnterprise,
		Client:     &http.Client{},
	}
}

type checkCloudDcStatusFailure struct{}

func (initialyser checkCloudDcStatusFailure) CheckCloudDcStatus(api *sdk.API,
	clientTooler *sdk.ClientTooler,
	resourceTooler *sdk.ResourceTooler) error {
	return errCheckCloudDcStatusFailure
}

func (initialyser checkCloudDcStatusFailure) GetClouddcEnvMeta(api *sdk.API,
	clientTooler *sdk.ClientTooler) (*sdk.APIMeta, error) {
	return &sdk.APIMeta{
		NonCriticalResourceList: unitTestNonCriticalResourceList,
		CriticalResourceList:    unitTestCriticalResourceList,
		OtherResourceList:       unitTestOtherResourceList,
	}, nil
}

func (initialyser checkCloudDcStatusFailure) New(token string, url string, enterprise string) *sdk.API {
	return &sdk.API{
		Token:      unitTestToken,
		URL:        unitTestAPIURL,
		Enterprise: unitTestEnterprise,
		Client:     &http.Client{},
	}
}
