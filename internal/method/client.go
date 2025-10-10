package method

import (
	"git.ghink.net/ghink/smsutils/internal/aliyun"
	"git.ghink.net/ghink/smsutils/internal/model"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiV2 "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20180501/client"
	"github.com/alibabacloud-go/tea/tea"
)

// createAliyunClient creates a aliyun client
func createAliyunClient(config model.AliyunConfig) (*aliyun.Client, error) {
	// Struct aliyun client config
	clientConfigV1 := &openapi.Config{
		AccessKeyId:     tea.String(config.KeyID),
		AccessKeySecret: tea.String(config.KeySecret),
	}
	clientConfigV2 := &openapiV2.Config{
		AccessKeyId:     tea.String(config.KeyID),
		AccessKeySecret: tea.String(config.KeySecret),
	}

	// Set aliyun endpoint
	clientConfigV1.Endpoint = tea.String("dysmsapi.ap-southeast-1.aliyuncs.com")
	clientConfigV2.Endpoint = tea.String("dysmsapi.aliyuncs.com")

	// Config aliyun client
	globeResult := &dysmsapi.Client{}
	globeResult, err := dysmsapi.NewClient(clientConfigV1)
	if err != nil {
		return &aliyun.Client{}, err
	}
	cnResult := &dysmsapi20170525.Client{}
	cnResult, err = dysmsapi20170525.NewClient(clientConfigV2)
	if err != nil {
		return &aliyun.Client{}, err
	}

	return &aliyun.Client{
		Globe: globeResult,
		CN:    cnResult,
	}, nil
}

// CreateClient creates a unified client
func CreateClient(config model.Config) (*model.Client, error) {
	// Create smsutils client
	smsClient := &model.Client{}

	aliyunClient, err := createAliyunClient(config.Aliyun)
	if err != nil {
		return &model.Client{}, err
	}

	smsClient.Aliyun = aliyunClient

	return smsClient, nil
}
