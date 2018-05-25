package sewan_go_sdk

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {

}

func TestCreate_and_Validate_API_URL(t *testing.T) {
	a := AirdrumURL_Builder{}
	err := errors.New("Sewan's Airdrum API's URL is wrong, accepted urls : https://next.cloud-datacenter.fr/api/clouddc/, ")
	pass_test_cases := []struct {
		api_url    string
		return_url URL
	}{
		{
			"https://next.cloud-datacenter.fr/api/clouddc/",
			URL{"https://next.cloud-datacenter.fr/api/clouddc/", nil},
		},
	}
	fail_test_cases := []struct {
		api_url    string
		return_url URL
	}{
		{
			"http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/",
			URL{"", err},
		},
		{
			"https://next.cloud-datacenter.fr/api/clouddc/youkoulélé",
			URL{"", err},
		},
		{
			"https://next.cloud-datacenter.fr:8081/api/clouddc/",
			URL{"", err},
		},
	}
	for _, test_case := range pass_test_cases {
		test_api_url := a.Create_and_Validate_API_URL(test_case.api_url)
		if test_api_url.S_url != test_case.return_url.S_url {
			t.Errorf("VM url was incorrect, got: %s, want: %s.", test_api_url.S_url, test_case.return_url.S_url)
		}
		if test_api_url.Err != test_case.return_url.Err {
			t.Errorf("Err should be nil as url is correct, got: %s, want: %s.", test_api_url.Err, test_case.return_url.Err)
		}
	}
	for _, test_case := range fail_test_cases {
		test_api_url := a.Create_and_Validate_API_URL(test_case.api_url)
		if test_api_url.S_url != test_case.return_url.S_url {
			t.Errorf("Returned url should be empty, got: %s, want: %s.", test_api_url.S_url, test_case.return_url.S_url)
		}
		if test_api_url.Err.Error() != test_case.return_url.Err.Error() {
			t.Errorf("Error msg was incorrect, got: %s, want: %s.", test_api_url.Err.Error(), test_case.return_url.Err.Error())
		}
	}
}

func TestGet_vm_creation_url(t *testing.T) {
	a := AirdrumURL_Builder{}
	test_cases := []struct {
		api_url         URL
		vm_creation_url string
	}{
		{URL{"https://next.cloud-datacenter.fr/api/clouddc/", nil},
			"https://next.cloud-datacenter.fr/api/clouddc/vm/"},
		{URL{"http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/", nil},
			"http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/vm/"},
	}
	for _, test_case := range test_cases {
		s_vm_creation_url := a.Get_vm_creation_url(test_case.api_url)
		if s_vm_creation_url != test_case.vm_creation_url {
			t.Errorf("VM url was incorrect, got: %s, want: %s.", s_vm_creation_url, test_case.vm_creation_url)
		}
	}
}

func TestGet_vm_url(t *testing.T) {
	a := AirdrumURL_Builder{}
	test_cases := []struct {
		api_url URL
		vm_id   string
		vm_url  string
	}{
		{URL{"https://next.cloud-datacenter.fr/api/clouddc/", nil}, "42",
			"https://next.cloud-datacenter.fr/api/clouddc/vm/42/"},
		{URL{"http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/", nil}, "42",
			"http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/vm/42/"},
	}
	for _, test_case := range test_cases {
		s_vm_url := a.Get_vm_url(test_case.api_url, test_case.vm_id)
		if s_vm_url != test_case.vm_url {
			t.Errorf("VM url was incorrect, got: %s, want: %s.", s_vm_url, test_case.vm_url)
		}
	}
}
