package client

import (
	"encoding/json"

	openapiV2 "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	dysmsapi20180501 "github.com/alibabacloud-go/dysmsapi-20180501/v2/client"
	"github.com/ghinknet/smsutils/v2/aliyun"
	"github.com/ghinknet/smsutils/v2/config"
	"github.com/ghinknet/smsutils/v2/model"
	"github.com/ghinknet/smsutils/v2/qcloud"
	"github.com/ghinknet/smsutils/v2/volc"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	smsv20210111 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/volcengine/volc-sdk-golang/service/sms"
)

// createAliyunClient creates a aliyun client
func createAliyunClient(config *config.Config) (*aliyun.Client, error) {
	// Struct aliyun client config
	clientConfigV2 := &openapiV2.Config{
		AccessKeyId:     &config.Aliyun.KeyID,
		AccessKeySecret: &config.Aliyun.KeySecret,
	}

	// Set aliyun endpoint
	endpoint := model.AliyunEndpoint
	if config.Aliyun.Endpoint != "" {
		endpoint = config.Aliyun.Endpoint
	}
	clientConfigV2.Endpoint = &endpoint

	// Create aliyun client
	globeResult := new(dysmsapi20180501.Client)
	globeResult, err := dysmsapi20180501.NewClient(clientConfigV2)
	if err != nil {
		return new(aliyun.Client), err
	}
	cnResult := new(dysmsapi20170525.Client)
	cnResult, err = dysmsapi20170525.NewClient(clientConfigV2)
	if err != nil {
		return new(aliyun.Client), err
	}

	return &aliyun.Client{
		Globe:  globeResult,
		CN:     cnResult,
		Config: config,
	}, nil
}

// createQCloudClient creates a qcloud client
func createQCloudClient(config *config.Config) (*qcloud.Client, error) {
	// Construct qcloud client config
	clientCredential := common.NewCredential(
		config.QCloud.SecretID,
		config.QCloud.SecretKey,
	)

	// Create client profile
	cpf := profile.NewClientProfile()

	// Set QCloud endpoint
	endpoint := model.QCloudEndpoint
	region := model.QCloudRegion
	if config.QCloud.Endpoint != "" {
		endpoint = config.QCloud.Endpoint
	}
	if config.QCloud.Region != "" {
		region = config.QCloud.Region
	}
	cpf.HttpProfile.Endpoint = endpoint

	// Create QCloud client
	client, err := smsv20210111.NewClient(clientCredential, region, cpf)
	if err != nil {
		return nil, err
	}

	return &qcloud.Client{
		Client: client,
		Config: config,
	}, nil
}

// createVolcClient creates a volc client
func createVolcClient(config *config.Config) (*volc.Client, error) {
	// Create volc client
	client := sms.NewInstance().Client

	// Set access key and secret key
	sms.DefaultInstance.Client.SetAccessKey(config.Volc.AccessKey)
	sms.DefaultInstance.Client.SetSecretKey(config.Volc.SecretKey)

	return &volc.Client{
		Client:     client,
		Config:     config,
		SmsAccount: config.Volc.SmsAccount,
	}, nil
}

// CreateClient creates a unified client
func CreateClient(config config.Config) (*Client, error) {
	// Create smsutils client
	smsClient := new(Client)

	// Check methods
	if config.Marshal == nil {
		config.Marshal = json.Marshal
	}
	if config.Unmarshal == nil {
		config.Unmarshal = json.Unmarshal
	}

	// Save config
	smsClient.Config = &config

	// Create aliyun client
	if config.Aliyun != nil {
		aliyunClient, err := createAliyunClient(&config)
		if err != nil {
			return new(Client), err
		}

		smsClient.Aliyun = aliyunClient
	}

	// Create qcloud client
	if config.QCloud != nil {
		qcloudClient, err := createQCloudClient(&config)
		if err != nil {
			return new(Client), err
		}

		smsClient.QCloud = qcloudClient
	}

	// Create volc client
	if config.Volc != nil {
		volcClient, err := createVolcClient(&config)
		if err != nil {
			return new(Client), err
		}

		smsClient.Volc = volcClient
	}

	return smsClient, nil
}
