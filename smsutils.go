package smsutils

import (
	"git.ghink.net/ghink/smsutils/internal/method"
	"git.ghink.net/ghink/smsutils/internal/model"
)

type AliyunConfig = model.AliyunConfig
type Config = model.Config

type Client = model.Client

var CreateClient = method.CreateClient
