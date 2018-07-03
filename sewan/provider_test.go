package sewan

import (
	"testing"
)

func TestProvider(t *testing.T) {
	err := Provider().InternalValidate()
	if err != nil {
		t.Fatalf("Provider shema error: %s", err)
	}
}
