package client

import (
	"github.com/ghinknet/smsutils/v2/aliyun"
	"github.com/ghinknet/smsutils/v2/config"
	"github.com/ghinknet/smsutils/v2/qcloud"
)

type Client struct {
	// Platform
	Aliyun *aliyun.Client
	QCloud *qcloud.Client
	// Config
	Config *config.Config
}
