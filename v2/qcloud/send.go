package qcloud

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghinknet/smsutils/v2/utils"
	"github.com/ghinknet/toolbox/pointer"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	smsv20210111 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// SendMessage sends a message to a phone
func (c *Client) SendMessage(to string, from string, templateCode string, templateParams []string) error {
	// Try to parse number
	countryCode, nationalNumber, regionCode, err := utils.ParseNumber(to)
	if err != nil {
		return err
	}
	if regionCode == "" && len([]rune(to)) == 11 && strings.HasPrefix(to, "1") {
		to = strings.Join([]string{"+86", to}, "")
		regionCode = "CN"
	} else {
		to = strings.Join([]string{
			"+", strconv.FormatInt(countryCode, 10), strconv.FormatInt(nationalNumber, 10),
		}, "")
	}

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
