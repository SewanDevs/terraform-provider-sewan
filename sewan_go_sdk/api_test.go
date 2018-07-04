package sewan_go_sdk

import (
	"errors"
	"net/http"
	"testing"
)

const (
	RIGHT_API_URL             = "https://api_url/api/bonjour/"
	RIGHT_VM_CREATION_API_URL = "https://api_url/api/bonjour/vm/"
	RIGHT_VM_URL_PATATE       = "https://api_url/api/bonjour/vm/PATATE/"
	RIGHT_VM_URL_42           = "https://api_url/api/bonjour/vm/42/"
	WRONG_API_URL             = "a wrong url"
	WRONG_API_URL_ERROR       = "Wrong api url msg"
	NO_RESP_API_URL           = "https://NO_RESP_API_URL.fr"
	NO_RESP_BODY_API_URL      = "https://NO_BODY_API_URL.org"
	NOT_JSON_RESP_API_URL     = "https://tata.fr"
	RIGHT_API_TOKEN           = "42424242424242424242424242424242"
	WRONG_API_TOKEN           = "a wrong token"
	WRONG_TOKEN_ERROR         = "Wrong api token msg"
)

//------------------------------------------------------------------------------
func TestNew(t *testing.T) {

	test_cases := []struct {
		Id          int
		Input_token string
		Input_url   string
		Output_api  API
	}{
		{1,
			WRONG_API_TOKEN,
			RIGHT_API_URL,
			API{WRONG_API_TOKEN, RIGHT_API_URL, nil},
		},
		{2,
			RIGHT_API_TOKEN,
			WRONG_API_URL,
			API{RIGHT_API_TOKEN, WRONG_API_URL, nil},
		},
		{3,
			WRONG_API_TOKEN,
			WRONG_API_URL,
			API{WRONG_API_TOKEN, WRONG_API_URL, nil},
		},
		{4,
			RIGHT_API_TOKEN,
			RIGHT_API_URL,
			API{RIGHT_API_TOKEN, RIGHT_API_URL, nil},
		},
	}

	fake_api_tools := APITooler{
		Api: FakeAirDrumResource_APIer{},
	}

	for _, test_case := range test_cases {
		api := fake_api_tools.New(
			test_case.Input_token,
			test_case.Input_url,
		)

		switch {
		case api.Token != test_case.Output_api.Token:
			t.Errorf("\n\nTC %d : API token error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, api.Token, test_case.Output_api.Token)
		case api.URL != test_case.Output_api.URL:
			t.Errorf("\n\nTC %d : API token error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, api.URL, test_case.Output_api.URL)
		}
	}
}

//------------------------------------------------------------------------------
func TestCheckStatus(t *testing.T) {

	test_cases := []struct {
		Id        int
		Input_api *API
		Err       error
	}{
		{1,
			&API{
				WRONG_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			errors.New(WRONG_TOKEN_ERROR),
		},
		{2,
			&API{
				RIGHT_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New(WRONG_API_URL_ERROR),
		},
		{3,
			&API{
				WRONG_API_TOKEN,
				WRONG_API_URL,
				&http.Client{},
			},
			errors.New(WRONG_API_URL_ERROR),
		},
		{4,
			&API{
				RIGHT_API_TOKEN,
				RIGHT_API_URL,
				&http.Client{},
			},
			nil,
		},
	}

	fake_api_tools := APITooler{}

	for _, test_case := range test_cases {
		fake_api_tools.Api = FakeAirDrumResource_APIer{}
		err := fake_api_tools.CheckStatus(test_case.Input_api)
		switch {
		case err == nil || test_case.Err == nil:
			if !(err == nil && test_case.Err == nil) {
				t.Errorf("\n\nTC %d : Check API error was incorrect,"+
					"\n\rgot: \"%s\"\n\rwant: \"%s\"",
					test_case.Id, err, test_case.Err)
			}
		case err.Error() != test_case.Err.Error():
			t.Errorf("\n\nTC %d : Check API error was incorrect,"+
				"\n\rgot: \"%s\"\n\rwant: \"%s\"",
				test_case.Id, err.Error(), test_case.Err.Error())
		}
	}
}
