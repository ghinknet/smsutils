package driver

import (
	"testing"

	"go.gh.ink/smsutils/v3/internal/state"
	"go.gh.ink/smsutils/v3/model"
)

// TestDriver implements model.Driver for testing
type TestDriver struct {
	name string
}

func (t *TestDriver) NewClient(params model.DriverClientParam) (model.Client, error) {
	return nil, nil
}

func TestRegister(t *testing.T) {
	// Store original drivers
	originalDrivers := make(map[string]model.Driver)
	for k, v := range state.Drivers {
		originalDrivers[k] = v
	}

	defer func() {
		// Restore original drivers
		state.Drivers = originalDrivers
	}()

	tests := []struct {
		name       string
		driverName string
		driver     model.Driver
	}{
		{
			name:       "Register single driver",
			driverName: "test-driver-1",
			driver:     &TestDriver{name: "test-driver-1"},
		},
		{
			name:       "Register multiple drivers",
			driverName: "test-driver-2",
			driver:     &TestDriver{name: "test-driver-2"},
		},
		{
			name:       "Override existing driver",
			driverName: "test-driver-1",
			driver:     &TestDriver{name: "test-driver-1-updated"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.driverName, tt.driver)

			// Check if driver is registered
			registeredDriver, exists := state.Drivers[tt.driverName]
			if !exists {
				t.Errorf("Register() driver not found in state.Drivers")
			}

			if registeredDriver != tt.driver {
				t.Errorf("Register() registered driver is different from the one passed")
			}
		})
	}
}

func TestRegisterMultipleDrivers(t *testing.T) {
	// Store original drivers
	originalDrivers := make(map[string]model.Driver)
	for k, v := range state.Drivers {
		originalDrivers[k] = v
	}

	defer func() {
		// Restore original drivers
		state.Drivers = originalDrivers
	}()

	// Register multiple drivers
	driver1 := &TestDriver{name: "driver-1"}
	driver2 := &TestDriver{name: "driver-2"}
	driver3 := &TestDriver{name: "driver-3"}

	Register("driver-1", driver1)
	Register("driver-2", driver2)
	Register("driver-3", driver3)

	// Verify all drivers are registered
	if d1, ok := state.Drivers["driver-1"]; !ok || d1 != driver1 {
		t.Errorf("driver-1 not properly registered")
	}
	if d2, ok := state.Drivers["driver-2"]; !ok || d2 != driver2 {
		t.Errorf("driver-2 not properly registered")
	}
	if d3, ok := state.Drivers["driver-3"]; !ok || d3 != driver3 {
		t.Errorf("driver-3 not properly registered")
	}
}

func TestRegisterReplaceDriver(t *testing.T) {
	// Store original drivers
	originalDrivers := make(map[string]model.Driver)
	for k, v := range state.Drivers {
		originalDrivers[k] = v
	}

	defer func() {
		// Restore original drivers
		state.Drivers = originalDrivers
	}()

	driverName := "replaceable-driver"

	// First registration
	driver1 := &TestDriver{name: "driver-v1"}
	Register(driverName, driver1)

	if d, ok := state.Drivers[driverName]; !ok || d != driver1 {
		t.Errorf("First registration failed")
	}

	// Replace with new driver
	driver2 := &TestDriver{name: "driver-v2"}
	Register(driverName, driver2)

	if d, ok := state.Drivers[driverName]; !ok || d != driver2 {
		t.Errorf("Driver replacement failed")
	}

	if d := state.Drivers[driverName]; d == driver1 {
		t.Errorf("Old driver was not replaced")
	}
}

func TestRegisterEmptyName(t *testing.T) {
	// Store original drivers
	originalDrivers := make(map[string]model.Driver)
	for k, v := range state.Drivers {
		originalDrivers[k] = v
	}

	defer func() {
		// Restore original drivers
		state.Drivers = originalDrivers
	}()

	driver := &TestDriver{name: "test"}
	Register("", driver)

	// Empty name is allowed, it will just be registered with empty string as key
	if d, ok := state.Drivers[""]; !ok || d != driver {
		t.Errorf("Empty name registration failed")
	}
}

