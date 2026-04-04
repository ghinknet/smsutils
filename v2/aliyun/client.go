package aliyun

import (
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	dysmsapi20180501 "github.com/alibabacloud-go/dysmsapi-20180501/v2/client"
	"github.com/ghinknet/smsutils/v2/config"
)

type Client struct {
	Globe *dysmsapi20180501.Client
	CN    *dysmsapi20170525.Client
	// Config
	Config *config.Config
}
