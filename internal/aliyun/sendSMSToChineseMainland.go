package aliyun

import (
	"encoding/json"
	"fmt"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
)

// sendMessageToChineseMainlandRaw is the raw method to send an SMS to Chinese Mainland
func sendMessageToChineseMainlandRaw(c *Client, to string, from string, templateCode string, templateParam string) error {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &to,
		SignName:      &from,
		TemplateCode:  &templateCode,
		TemplateParam: &templateParam,
	}
	runtime := &util.RuntimeOptions{}

	resp, err := c.CN.SendSmsWithOptions(sendSmsRequest, runtime)
	if err != nil {
		return err
	}
	if *resp.Body.Code != "OK" {
		return fmt.Errorf("%s", *resp.Body.Code)
	}

	return nil
}

// SendMessageToChineseMainland sends a message to a chinese mainland phone
func (c *Client) SendMessageToChineseMainland(to string, from string, templateCode string, templateParam map[string]string) error {
	// Marshal params
	marshalledParam, err := json.Marshal(templateParam)
	if err != nil {
		return err
	}
	payload := string(marshalledParam)

	// Send message
	return sendMessageToChineseMainlandRaw(c, to, from, templateCode, payload)
}
