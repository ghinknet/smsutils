package model

type AliyunConfig struct {
	KeyID     string
	KeySecret string
}

type Config struct {
	Aliyun AliyunConfig
	Debug  bool
}
