package qcloud

import (
	"github.com/ghinknet/smsutils/v2/config"
	smsv20210111 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Client struct {
	Client *smsv20210111.Client
	// Config
	Config *config.Config
}
