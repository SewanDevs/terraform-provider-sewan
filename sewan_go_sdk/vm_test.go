package sewan_go_sdk

import (
	"testing"
	"github.com/hashicorp/terraform/helper/schema"
)

func TestGet_vm_url(t *testing.T) {
	test_cases := []struct {
		api_url string
		vm_id   string
		vm_url  string
	}{
		{"https://next.cloud-datacenter.fr/api/clouddc/vm/", "42",
			"https://next.cloud-datacenter.fr/api/clouddc/vm/42/"},
		{"http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/vm/", "42",
			"http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/vm/42/"},
	}

	for _, test_case := range test_cases {
		s_vm_url := get_vm_url(test_case.api_url, test_case.vm_id)
		if s_vm_url != test_case.vm_url {
			t.Errorf("VM url was incorrect, got: %s, want: %s.", s_vm_url, test_case.vm_url)
		}
	}
}

func TestVmInstanceCreate(t *testing.T) {
}

func TestCreate_vm_resource(t *testing.T) {
}

func TestRead_vm_resource(t *testing.T) {
}

func TestUpdate_vm_resource(t *testing.T) {
}

func TestDelete_vm_resource(t *testing.T) {
}
