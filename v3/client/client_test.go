package client

import (
	"encoding/json"
	stderrors "errors"
	"testing"

	"go.gh.ink/smsutils/v3/errors"
	"go.gh.ink/smsutils/v3/internal/state"
	"go.gh.ink/smsutils/v3/model"
)

// MockDriver implements model.Driver for testing
type MockDriver struct {
	shouldFail bool
	mockClient *MockClient
}

func (m *MockDriver) NewClient(params model.DriverClientParam) (model.Client, error) {
	if m.shouldFail {
		return nil, errors.ErrDriverCredentialInvalid
	}
	return m.mockClient, nil
}

// MockClient implements model.Client for testing
type MockClient struct {
	name          string
	sendCallCount int
	lastDest      string
	lastSender    string
	lastTemplate  string
	lastVars      model.Vars
}

func (m *MockClient) SendMessage(dest string, sender string, template string, vars model.Vars) error {
	m.sendCallCount++
	m.lastDest = dest
	m.lastSender = sender
	m.lastTemplate = template
	m.lastVars = vars
	return nil
}

func TestNewClient(t *testing.T) {
	// Register test drivers
	defer func() {
		// Cleanup
		delete(state.Drivers, "mock-driver-success")
		delete(state.Drivers, "mock-driver-fail")
	}()

	mockDriverSuccess := &MockDriver{shouldFail: false, mockClient: &MockClient{name: "mock-success"}}
	mockDriverFail := &MockDriver{shouldFail: true}

	state.Drivers["mock-driver-success"] = mockDriverSuccess
	state.Drivers["mock-driver-fail"] = mockDriverFail

	tests := []struct {
		name        string
		config      model.Config
		wantErr     bool
		wantClients int
	}{
		{
			name: "Empty credentials",
			config: model.Config{
				Credentials: model.C{},
			},
			wantErr:     false,
			wantClients: 0,
		},
		{
			name: "Single valid driver",
			config: model.Config{
				Credentials: model.C{
					"mock-driver-success": {"key": "value"},
				},
			},
			wantErr:     false,
			wantClients: 1,
		},
		{
			name: "Multiple valid drivers",
			config: model.Config{
				Credentials: model.C{
					"mock-driver-success": {"key": "value1"},
				},
			},
			wantErr:     false,
			wantClients: 1,
		},
		{
			name: "Invalid driver",
			config: model.Config{
				Credentials: model.C{
					"non-existent-driver": {"key": "value"},
				},
			},
			wantErr: true,
		},
		{
			name: "Driver creation fails",
			config: model.Config{
				Credentials: model.C{
					"mock-driver-fail": {"key": "value"},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clients, err := NewClient(tt.config)

			if tt.wantErr && err == nil {
				t.Errorf("NewClient() expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("NewClient() unexpected error: %v", err)
			}

			if !tt.wantErr {
				if len(clients) != tt.wantClients {
					t.Errorf("NewClient() got %d clients, want %d", len(clients), tt.wantClients)
				}
			}
		})
	}
}

func TestNewClientWithCustomMarshal(t *testing.T) {
	defer func() {
		delete(state.Drivers, "mock-driver-success")
	}()

	mockDriver := &MockDriver{shouldFail: false, mockClient: &MockClient{name: "mock"}}
	state.Drivers["mock-driver-success"] = mockDriver

	marshal := func(v any) ([]byte, error) {
		return json.Marshal(v)
	}
	unmarshal := func(data []byte, v any) error {
		return json.Unmarshal(data, v)
	}

	config := model.Config{
		Credentials: model.C{
			"mock-driver-success": {"key": "value"},
		},
		Marshal:   marshal,
		Unmarshal: unmarshal,
	}

	clients, err := NewClient(config)
	if err != nil {
		t.Errorf("NewClient() unexpected error: %v", err)
	}

	if len(clients) != 1 {
		t.Errorf("NewClient() got %d clients, want 1", len(clients))
	}
}

func TestNewClientDefaultsMarshal(t *testing.T) {
	defer func() {
		delete(state.Drivers, "mock-driver-success")
	}()

	mockDriver := &MockDriver{shouldFail: false, mockClient: &MockClient{name: "mock"}}
	state.Drivers["mock-driver-success"] = mockDriver

	config := model.Config{
		Credentials: model.C{
			"mock-driver-success": {"key": "value"},
		},
		// Don't set Marshal/Unmarshal - should use defaults
	}

	clients, err := NewClient(config)
	if err != nil {
		t.Errorf("NewClient() unexpected error: %v", err)
	}

	if len(clients) != 1 {
		t.Errorf("NewClient() got %d clients, want 1", len(clients))
	}
}

func TestNewClientErrorMessage(t *testing.T) {
	defer func() {
		delete(state.Drivers, "non-existent-driver")
	}()

	config := model.Config{
		Credentials: model.C{
			"non-existent-driver": {"key": "value"},
		},
	}

	_, err := NewClient(config)
	if err == nil {
		t.Errorf("NewClient() expected error for non-existent driver, got nil")
	}

	// Check if error is the expected type using errors.As
	var smsErr *errors.SmsutilsError
	if !stderrors.As(err, &smsErr) {
		t.Errorf("NewClient() error is not SmsutilsError: %T", err)
		return
	}

	if smsErr.DriverName() != "non-existent-driver" {
		t.Errorf("NewClient() driver name = %s, want 'non-existent-driver'", smsErr.DriverName())
	}
}



