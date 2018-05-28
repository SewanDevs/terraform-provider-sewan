package sewan_go_sdk

import (
	"errors"
	"testing"
)

//------------------------------------------------------------------------------
func TestNew(t *testing.T) {
	New("42", "tata")

}

//------------------------------------------------------------------------------
func TestValidate_status() (t *testing.T) {

}

//------------------------------------------------------------------------------
func TestGet_vm_creation_url(t *testing.T) {
	// Old test arch backup
	//a := AirdrumURL_Builder{}
	//test_cases := []struct {
	//	api_url         URL
	//	vm_creation_url string
	//}{
	//	{URL{"https://next.cloud-datacenter.fr/api/clouddc/", nil},
	//		"https://next.cloud-datacenter.fr/api/clouddc/vm/"},
	//	{URL{"http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/", nil},
	//		"http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/vm/"},
	//}
	//for _, test_case := range test_cases {
	//	s_vm_creation_url := a.Get_vm_creation_url(test_case.api_url)
	//	if s_vm_creation_url != test_case.vm_creation_url {
	//		t.Errorf("VM url was incorrect, got: %s, want: %s.", s_vm_creation_url, test_case.vm_creation_url)
	//	}
	//}
}

//------------------------------------------------------------------------------
func TestGet_vm_url(t *testing.T) {
	// Old test arch backup
	//a := AirdrumURL_Builder{}
	//test_cases := []struct {
	//	api_url URL
	//	vm_id   string
	//	vm_url  string
	//}{
	//	{URL{"https://next.cloud-datacenter.fr/api/clouddc/", nil}, "42",
	//		"https://next.cloud-datacenter.fr/api/clouddc/vm/42/"},
	//	{URL{"http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/", nil}, "42",
	//		"http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/vm/42/"},
	//}
	//for _, test_case := range test_cases {
	//	s_vm_url := a.Get_vm_url(test_case.api_url, test_case.vm_id)
	//	if s_vm_url != test_case.vm_url {
	//		t.Errorf("VM url was incorrect, got: %s, want: %s.", s_vm_url, test_case.vm_url)
	//	}
	//}
}
