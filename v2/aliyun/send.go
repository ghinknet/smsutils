package aliyun

import (
	"github.com/ghinknet/smsutils/v2/utils"
)

// SendMessage sends a message to a phone
func (c *Client) SendMessage(to string, from string, templateCode string, templateParams map[string]string) error {
	// Try to parse number
	to, _, _, regionCode, err := utils.ProcessNumberForChinese(to)
	if err != nil {
		return err
	}

	if regionCode == "CN" {
		return c.SendMessageToChineseMainland(to, from, templateCode, templateParams)
	}
	return c.SendMessageToGlobe(to, from, templateCode, templateParams)
}
