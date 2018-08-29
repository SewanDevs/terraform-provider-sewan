package sewan

import (
	"testing"
)

func TestClientStruct(t *testing.T) {
	testConfigStruct := configStruct{"unit test token", "unit test url"}
	TestclientStruct, err := testConfigStruct.clientStruct()
	switch {
	case TestclientStruct.sewan == nil:
		t.Errorf("clientStruct API is nil, it should be initialized.")
	case TestclientStruct.sewanAPITooler == nil:
		t.Errorf("clientStruct APITooler is nil, it should be initialized.")
	case TestclientStruct.sewanClientTooler == nil:
		t.Errorf("clientStruct ClientTooler is nil, it should be initialized.")
	case TestclientStruct.sewanTemplatesTooler == nil:
		t.Errorf("clientStruct TemplatesTooler is nil, it should be initialized.")
	case TestclientStruct.sewanResourceTooler == nil:
		t.Errorf("clientStruct ResourceTooler is nil, it should be initialized.")
	case TestclientStruct.sewanSchemaTooler == nil:
		t.Errorf("clientStruct SchemaTooler is nil, it should be initialized.")
	case err == nil:
		t.Errorf("err should not be nil as testConfigStruct{} token and url are " +
			"wrongly formatted")
	}
}
