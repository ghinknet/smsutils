package client

import (
	"encoding/json"

	openapiV2 "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	dysmsapi20180501 "github.com/alibabacloud-go/dysmsapi-20180501/v2/client"
	"github.com/ghinknet/smsutils/v2/aliyun"
	model "github.com/ghinknet/smsutils/v2/config"
	model2 "github.com/ghinknet/smsutils/v2/model"
	"github.com/ghinknet/smsutils/v2/qcloud"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	smsv20210111 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// createAliyunClient creates a aliyun client
func createAliyunClient(config *model.Config) (*aliyun.Client, error) {
	// Struct aliyun client config
	clientConfigV2 := &openapiV2.Config{
		AccessKeyId:     &config.Aliyun.KeyID,
		AccessKeySecret: &config.Aliyun.KeySecret,
	}

	// Set aliyun endpoint
	endpoint := model2.AliyunEndpoint
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
func createQCloudClient(config *model.Config) (*qcloud.Client, error) {
	// Construct qcloud client config
	clientCredential := common.NewCredential(
		config.QCloud.SecretID,
		config.QCloud.SecretKey,
	)

	// Create client profile
	cpf := profile.NewClientProfile()

	// Set QCloud endpoint
	endpoint := model2.QCloudEndpoint
	region := model2.QCloudRegion
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

// CreateClient creates a unified client
func CreateClient(config model.Config) (*Client, error) {
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

	return smsClient, nil
}
