package aliyun

import (
	"fmt"
	"strings"

	dysmsapi20180501 "github.com/alibabacloud-go/dysmsapi-20180501/v2/client"
	"github.com/ghinknet/toolbox/pointer"
)

// sendMessageToGlobeRaw is the raw method to send an SMS to Globe
func sendMessageToGlobeRaw(c *Client, to string, message string, from string) error {
	resp, err := c.Globe.SendMessageToGlobe(&dysmsapi20180501.SendMessageToGlobeRequest{
		To:      &to,
		Message: &message,
		From:    &from,
	})
	if err != nil {
		return err
	}
	if pointer.SafeDeref(resp.Body.ResponseCode) != "OK" {
		return fmt.Errorf("%s", pointer.SafeDeref(resp.Body.ResponseCode))
	}

	return nil
}

// SendMessageToGlobe sends a message to a globe phone
func (c *Client) SendMessageToGlobe(to string, from string, templateContent string, templateParams map[string]string) error {
	// Render template
	for k, v := range templateParams {
		templateContent = strings.Replace(templateContent, "${"+k+"}", v, -1)
	}

	// Send message
	return sendMessageToGlobeRaw(c, to, templateContent, from)
}
