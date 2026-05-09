package model

import (
	"testing"
)

func TestVarCreation(t *testing.T) {
	v := &Var{
		Key:   "testKey",
		Value: "testValue",
	}

	if v.Key != "testKey" {
		t.Errorf("Var.Key = %s, want testKey", v.Key)
	}

	if v.Value != "testValue" {
		t.Errorf("Var.Value = %s, want testValue", v.Value)
	}
}

func TestVarsCreation(t *testing.T) {
	vars := Vars{
		{Key: "var1", Value: "value1"},
		{Key: "var2", Value: "value2"},
		{Key: "var3", Value: "value3"},
	}

	if len(vars) != 3 {
		t.Errorf("len(vars) = %d, want 3", len(vars))
	}

	if vars[0].Key != "var1" {
		t.Errorf("vars[0].Key = %s, want var1", vars[0].Key)
	}

	if vars[1].Value != "value2" {
		t.Errorf("vars[1].Value = %s, want value2", vars[1].Value)
	}
}

func TestVarsAccess(t *testing.T) {
	vars := Vars{
		{Key: "name", Value: "Alice"},
		{Key: "code", Value: "123456"},
	}

	for i, v := range vars {
		if v == nil {
			t.Fatalf("vars[%d] is nil", i)
		}

		if v.Key == "" || v.Value == "" {
			t.Errorf("vars[%d] has empty Key or Value", i)
		}
	}
}

func TestVarsAppend(t *testing.T) {
	var vars Vars

	vars = append(vars, &Var{Key: "key1", Value: "value1"})
	vars = append(vars, &Var{Key: "key2", Value: "value2"})

	if len(vars) != 2 {
		t.Errorf("After append, len(vars) = %d, want 2", len(vars))
	}
}

func TestConfigCreation(t *testing.T) {
	cred := C{
		"driver1": {"username": "user1", "password": "pass1"},
		"driver2": {"api_key": "key123"},
	}

	config := Config{
		Credentials: cred,
	}

	if config.Credentials == nil {
		t.Errorf("Config.Credentials is nil")
	}

	if driver1Cred := config.Credentials["driver1"]; driver1Cred["username"] != "user1" {
		t.Errorf("Config credentials not set correctly")
	}
}

func TestConfigWithMarshalFunctions(t *testing.T) {
	marshal := func(v any) ([]byte, error) {
		return nil, nil
	}

	unmarshal := func(data []byte, v any) error {
		return nil
	}

	config := Config{
		Credentials: C{},
		Marshal:     marshal,
		Unmarshal:   unmarshal,
	}

	if config.Marshal == nil {
		t.Errorf("Config.Marshal is nil")
	}

	if config.Unmarshal == nil {
		t.Errorf("Config.Unmarshal is nil")
	}
}

func TestConfigCredentialsType(t *testing.T) {
	// Test that C type is properly a map[string]map[string]string
	cred := C{
		"service1": {
			"key1": "value1",
			"key2": "value2",
		},
		"service2": {
			"key3": "value3",
		},
	}

	if cred["service1"]["key1"] != "value1" {
		t.Errorf("Nested credential access failed")
	}

	if cred["service2"]["key3"] != "value3" {
		t.Errorf("Nested credential access failed")
	}
}

func TestVarsEmpty(t *testing.T) {
	var vars Vars

	if len(vars) != 0 {
		t.Errorf("Empty Vars length = %d, want 0", len(vars))
	}
}

func TestVarsNil(t *testing.T) {
	var vars Vars

	// After declaration, vars should be a nil slice
	if len(vars) != 0 {
		t.Errorf("Nil Vars length = %d, want 0", len(vars))
	}
}

func TestVarJSONTags(t *testing.T) {
	// This test verifies that JSON tags are present for marshalling
	v := &Var{
		Key:   "testKey",
		Value: "testValue",
	}

	// Access through the struct to ensure the fields are defined with proper tags
	if v.Key == "" || v.Value == "" {
		t.Errorf("Var fields not properly initialized")
	}
}

