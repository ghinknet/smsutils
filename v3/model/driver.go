package model

type Driver interface {
	NewClient(params DriverClientParam) Client
}

type DriverClientParam struct {
	// Platform credential
	Credential map[string]string
	// JSON
	Unmarshal func(data []byte, v any) error
	Marshal   func(v any) ([]byte, error)
}
