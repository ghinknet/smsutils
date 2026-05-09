package errors

import (
	"errors"
	"testing"
)

func TestSmsutilsErrorNew(t *testing.T) {
	msg := "test error message"
	err := New(msg)

	if err == nil {
		t.Fatalf("New() returned nil")
	}

	if err.Error() != msg {
		t.Errorf("Error() = %s, want %s", err.Error(), msg)
	}
}

func TestSmsutilsErrorWithOptions(t *testing.T) {
	err := New("test message",
		WithDriverName("test-driver"),
		WithDriverCode("ERR001"),
		WithDriverMessage("Driver error message"),
		WithDriverRequestID("req-123"),
		WithDriverResponse(map[string]string{"key": "value"}),
	)

	if err.DriverName() != "test-driver" {
		t.Errorf("DriverName() = %s, want test-driver", err.DriverName())
	}

	if err.DriverCode() != "ERR001" {
		t.Errorf("DriverCode() = %s, want ERR001", err.DriverCode())
	}

	if err.DriverMessage() != "Driver error message" {
		t.Errorf("DriverMessage() = %s, want Driver error message", err.DriverMessage())
	}

	if err.DriverRequestID() != "req-123" {
		t.Errorf("DriverRequestID() = %s, want req-123", err.DriverRequestID())
	}

	if response, ok := err.DriverResponse().(map[string]string); !ok || response["key"] != "value" {
		t.Errorf("DriverResponse() unexpected value")
	}
}

func TestSmsutilsErrorWithDriverName(t *testing.T) {
	err := New("test error")
	err = err.WithDriverName("my-driver")

	if err.DriverName() != "my-driver" {
		t.Errorf("WithDriverName() = %s, want my-driver", err.DriverName())
	}
}

func TestSmsutilsErrorWithDriverCode(t *testing.T) {
	err := New("test error")
	err = err.WithDriverCode("CODE123")

	if err.DriverCode() != "CODE123" {
		t.Errorf("WithDriverCode() = %s, want CODE123", err.DriverCode())
	}
}

func TestSmsutilsErrorWithDriverMessage(t *testing.T) {
	err := New("test error")
	err = err.WithDriverMessage("Driver specific message")

	if err.DriverMessage() != "Driver specific message" {
		t.Errorf("WithDriverMessage() = %s, want Driver specific message", err.DriverMessage())
	}
}

func TestSmsutilsErrorWithDriverRequestID(t *testing.T) {
	err := New("test error")
	err = err.WithDriverRequestID("REQ-12345")

	if err.DriverRequestID() != "REQ-12345" {
		t.Errorf("WithDriverRequestID() = %s, want REQ-12345", err.DriverRequestID())
	}
}

func TestSmsutilsErrorWithDriverResponse(t *testing.T) {
	err := New("test error")
	response := map[string]interface{}{"status": "failed", "code": 500}
	err = err.WithDriverResponse(response)

	if respMap, ok := err.DriverResponse().(map[string]interface{}); !ok {
		t.Errorf("WithDriverResponse() type assertion failed")
	} else {
		if respMap["status"] != "failed" {
			t.Errorf("DriverResponse() status = %v, want failed", respMap["status"])
		}
		if respMap["code"] != 500 {
			t.Errorf("DriverResponse() code = %v, want 500", respMap["code"])
		}
	}
}

func TestSmsutilsErrorError(t *testing.T) {
	msg := "specific error message"
	err := New(msg)

	if err.Error() != msg {
		t.Errorf("Error() = %s, want %s", err.Error(), msg)
	}
}

func TestSmsutilsErrorIs(t *testing.T) {
	baseErr := New("base error")
	err1 := New("error 1")
	err2 := New("error 2")

	if !err1.Is(err1) {
		t.Errorf("Is() should return true for same error")
	}

	if err1.Is(err2) {
		t.Errorf("Is() should return false for different errors")
	}

	if err1.Is(baseErr) {
		t.Errorf("Is() should return false for unrelated errors")
	}
}

func TestSmsutilsErrorUnwrap(t *testing.T) {
	err := New("test error")

	if unwrapped := err.Unwrap(); !errors.Is(unwrapped, err) {
		t.Errorf("Unwrap() should return the same error")
	}
}

func TestPredefinedErrors(t *testing.T) {
	if ErrDriverNotRegistered == nil {
		t.Errorf("ErrDriverNotRegistered should not be nil")
	}

	if ErrDriverCredentialInvalid == nil {
		t.Errorf("ErrDriverCredentialInvalid should not be nil")
	}

	if ErrDriverSendFailed == nil {
		t.Errorf("ErrDriverSendFailed should not be nil")
	}
}

func TestErrorChaining(t *testing.T) {
	baseErr := New("send failed")
	err := baseErr.WithDriverName("sms-driver").
		WithDriverCode("SEND_ERROR").
		WithDriverMessage("Failed to send message").
		WithDriverRequestID("req-789")

	if err.DriverName() != "sms-driver" {
		t.Errorf("DriverName() = %s, want sms-driver", err.DriverName())
	}

	if err.DriverCode() != "SEND_ERROR" {
		t.Errorf("DriverCode() = %s, want SEND_ERROR", err.DriverCode())
	}

	if err.DriverMessage() != "Failed to send message" {
		t.Errorf("DriverMessage() = %s, want Failed to send message", err.DriverMessage())
	}

	if err.DriverRequestID() != "req-789" {
		t.Errorf("DriverRequestID() = %s, want req-789", err.DriverRequestID())
	}
}

func TestErrorDoesNotModifyOriginal(t *testing.T) {
	originalErr := New("test")
	originalDriverName := originalErr.DriverName()

	clonedErr := originalErr.WithDriverName("new-driver")

	if originalErr.DriverName() != originalDriverName {
		t.Errorf("Original error was modified")
	}

	if clonedErr.DriverName() != "new-driver" {
		t.Errorf("Cloned error not updated correctly")
	}

	// Check if they are different instances (pointer comparison)
	// This is safe because these are direct pointers to SmsutilsError, not wrapped errors
	// The Copy operation ensures they are different instances
	if clonedErr == originalErr {
		t.Errorf("WithDriverName() should return a copy, not the same instance")
	}
}

func TestErrorPredefined(t *testing.T) {
	err := ErrDriverNotRegistered.WithDriverName("my-driver")

	if err.DriverName() != "my-driver" {
		t.Errorf("Predefined error modification failed")
	}

	// Original should not be modified
	if ErrDriverNotRegistered.DriverName() != "" {
		t.Errorf("Predefined error was modified")
	}
}



