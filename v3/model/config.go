package model

type Config struct {
	// Platforms' credentials
	Credentials C
	// JSON
	Unmarshal func(data []byte, v any) error
	Marshal   func(v any) ([]byte, error)
}
