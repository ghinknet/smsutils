package volc

import (
	"github.com/ghinknet/smsutils/v2/config"
	"github.com/volcengine/volc-sdk-golang/base"
)

type Client struct {
	Client     *base.Client
	SmsAccount string
	// Config
	Config *config.Config
}
