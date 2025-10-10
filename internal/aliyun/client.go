package aliyun

import (
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20180501/client"
)

type Client struct {
	Globe *dysmsapi.Client
	CN    *dysmsapi20170525.Client
}
