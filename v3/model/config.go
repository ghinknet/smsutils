package model

type Config struct {
	// Platforms' credentials
	Credentials map[string]map[string]string
	// JSON
	Unmarshal func(data []byte, v any) error
	Marshal   func(v any) ([]byte, error)
}
