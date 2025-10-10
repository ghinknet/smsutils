package aliyun

import (
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20180501/client"
	"strings"
)

// sendMessageToGlobeRaw is the raw method to send an SMS to Globe
func sendMessageToGlobeRaw(c *Client, to string, message string, from string) error {
	req := &dysmsapi.SendMessageToGlobeRequest{
		To:      &to,
		Message: &message,
		From:    &from,
	}
	_, err := c.Globe.SendMessageToGlobe(req)
	if err != nil {
		return err
	}

	return nil
}

// SendMessageToGlobe sends a message to a globe phone
func (c *Client) SendMessageToGlobe(to string, from string, templateContent string, templateParam map[string]string) error {
	// Render template
	for k, v := range templateParam {
		templateContent = strings.Replace(templateContent, "${"+k+"}", v, -1)
	}

	// Send message
	return sendMessageToGlobeRaw(c, to, templateContent, from)
}
