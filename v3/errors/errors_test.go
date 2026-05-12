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
	// For newly created errors, raw is nil, so Is() checks direct pointer equality
	err1 := New("error 1")
	err2 := New("error 2")

	// Direct pointer comparison when raw is nil
	if !err1.Is(err1) {
		t.Errorf("Is() should return true when comparing error to itself")
	}

	if err1.Is(err2) {
		t.Errorf("Is() should return false for different error instances")
	}

	// For cloned errors, Is() checks if the passed error matches the raw sentinel
	cloned1 := err1.WithDriverName("driver-1")
	if !cloned1.Is(err1) {
		t.Errorf("Is() should return true for cloned error checking against original sentinel")
	}

	cloned2 := err2.WithDriverName("driver-2")
	if cloned2.Is(err1) {
		t.Errorf("Is() should return false for cloned error from different sentinel")
	}

	// Cloned errors are different instances
	if cloned1.Is(cloned2) {
		t.Errorf("Is() should not match different sentinels")
	}
}

func TestSmsutilsErrorUnwrap(t *testing.T) {
	err := New("test error")

	// Newly created error has nil raw, so Unwrap returns nil
	if unwrapped := err.Unwrap(); unwrapped != nil {
		t.Errorf("Unwrap() for new error should return nil, got %v", unwrapped)
	}

	// Cloned error has raw pointing to the original sentinel
	cloned := err.WithDriverName("driver")
	unwrapped := cloned.Unwrap()
	if unwrapped != err {
		t.Errorf("Unwrap() for cloned error should return the original sentinel")
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

	// Cloned error should reference the original sentinel
	if !err.Is(ErrDriverNotRegistered) {
		t.Errorf("Cloned predefined error should reference original sentinel with Is()")
	}
}

func TestErrorSentinelComparison(t *testing.T) {
	// Test errors.Is with predefined sentinel errors
	err := ErrDriverSendFailed.WithDriverCode("E001")

	if !errors.Is(err, ErrDriverSendFailed) {
		t.Errorf("errors.Is() should work with cloned sentinel errors")
	}

	// Different sentinel should not match
	if errors.Is(err, ErrDriverNotRegistered) {
		t.Errorf("errors.Is() should return false for different sentinels")
	}
}

func TestRawErrorPreservation(t *testing.T) {
	// Test that clone preserves raw error reference for sentinel errors
	original := New("original message")
	cloned1 := original.WithDriverName("driver1")
	cloned2 := cloned1.WithDriverCode("CODE")

	// Both cloned errors should reference the original
	if !cloned1.Is(original) {
		t.Errorf("First clone should reference original")
	}

	if !cloned2.Is(original) {
		t.Errorf("Second clone should reference original (not first clone)")
	}

	// cloned2 should not reference cloned1
	if cloned2.Is(cloned1) && cloned1 != original {
		t.Errorf("Cloned error should reference original sentinel, not intermediate clones")
	}
}
