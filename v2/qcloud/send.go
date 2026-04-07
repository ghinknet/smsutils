package qcloud

import (
	"fmt"

	"github.com/ghinknet/smsutils/v2/utils"
	"github.com/ghinknet/toolbox/pointer"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	smsv20210111 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// SendMessage sends a message to a phone
func (c *Client) SendMessage(to string, from string, templateCode string, templateParams []string) error {
	// Try to parse number
	to, _, _, _, err := utils.ProcessNumberForChinese(to)

	// Construct a request
	request := smsv20210111.NewSendSmsRequest()

	// Set request params
	request.PhoneNumberSet = []*string{&to}
	request.SmsSdkAppId = &c.Config.QCloud.SmsSdkAppID
	request.TemplateId = &templateCode
	request.SignName = &from
	request.TemplateParamSet = common.StringPtrs(templateParams)

	// Send requests
	response, err := c.Client.SendSms(request)
	if err != nil {
		return err
	}
	for _, status := range response.Response.SendStatusSet {
		if pointer.SafeDeref(status.Code) != "Ok" {
			return fmt.Errorf("%s", pointer.SafeDeref(status.Code))
		}
	}

	return nil
}
