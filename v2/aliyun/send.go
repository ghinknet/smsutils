package aliyun

import (
	"strconv"
	"strings"

	"github.com/ghinknet/smsutils/v2/utils"
)

// SendMessage sends a message to a phone
func (c *Client) SendMessage(to string, from string, templateCode string, templateParams map[string]string) error {
	// Try to parse number
	countryCode, nationalNumber, regionCode, err := utils.ParseNumber(to)
	if err != nil {
		return err
	}
	if regionCode == "" && len([]rune(to)) == 11 && strings.HasPrefix(to, "1") {
		to = strings.Join([]string{"86", to}, "")
		regionCode = "CN"
	} else {
		to = strings.Join([]string{
			strconv.FormatInt(countryCode, 10), strconv.FormatInt(nationalNumber, 10),
		}, "")
	}

	if regionCode == "CN" {
		return c.SendMessageToChineseMainland(to, from, templateCode, templateParams)
	}
	return c.SendMessageToGlobe(to, from, templateCode, templateParams)
}
