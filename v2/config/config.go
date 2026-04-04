package config

type AliyunConfig struct {
	KeyID     string
	KeySecret string
	Endpoint  string
}

type QCloudConfig struct {
	SecretID    string
	SecretKey   string
	SmsSdkAppID string
	Endpoint    string
	Region      string
}

type Config struct {
	// Platform
	Aliyun *AliyunConfig
	QCloud *QCloudConfig
	// Setting
	Debug bool
	// Method
	Unmarshal func(data []byte, v any) error
	Marshal   func(v any) ([]byte, error)
}
