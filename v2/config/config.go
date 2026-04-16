package config

type AliyunConfig struct {
	KeyID     string
	KeySecret string
}

type QCloudConfig struct {
	SecretID    string
	SecretKey   string
	SmsSdkAppID string
	Endpoint    string
	Region      string
}

type VolcConfig struct {
	AccessKey  string
	SecretKey  string
	SmsAccount string
}

type Config struct {
	// Platform
	Aliyun *AliyunConfig
	QCloud *QCloudConfig
	Volc   *VolcConfig
	// Setting
	Debug bool
	// Method
	Unmarshal func(data []byte, v any) error
	Marshal   func(v any) ([]byte, error)
}
