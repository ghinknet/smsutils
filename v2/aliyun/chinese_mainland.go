package aliyun

import (
	"fmt"

	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/ghinknet/toolbox/pointer"
)

// sendMessageToChineseMainlandRaw is the raw method to send an SMS to Chinese Mainland
func sendMessageToChineseMainlandRaw(c *Client, to string, from string, templateCode string, templateParams string) error {
	resp, err := c.CN.SendSmsWithOptions(&dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &to,
		SignName:      &from,
		TemplateCode:  &templateCode,
		TemplateParam: &templateParams,
	}, new(util.RuntimeOptions))
	if err != nil {
		return err
	}
	if pointer.SafeDeref(resp.Body.Code) != "OK" {
		return fmt.Errorf("%s", pointer.SafeDeref(resp.Body.Code))
	}

	return nil
}

// SendMessageToChineseMainland sends a message to a Chinese mainland phone
func (c *Client) SendMessageToChineseMainland(to string, from string, templateCode string, templateParams map[string]string) error {
	// Marshal params
	marshalledParam, err := c.Config.Marshal(templateParams)
	if err != nil {
		return err
	}

	// Send message
	return sendMessageToChineseMainlandRaw(c, to, from, templateCode, string(marshalledParam))
}
