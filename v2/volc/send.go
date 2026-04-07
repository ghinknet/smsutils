package volc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ghinknet/smsutils/v2/utils"
	"github.com/volcengine/volc-sdk-golang/service/sms"
)

// SendMessage sends a message to a phone
func (c *Client) SendMessage(to string, from string, templateCode string, templateParams map[string]string) error {
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

	// Marshal params
	marshalledParam, err := c.Config.Marshal(templateParams)
	if err != nil {
		return err
	}

	// Construct request
	req := &sms.SmsRequest{
		SmsAccount:    c.SmsAccount,
		Sign:          from,
		TemplateID:    templateCode,
		TemplateParam: string(marshalledParam),
		PhoneNumbers:  to,
	}

	// Send request
	result, statusCode, err := sms.DefaultInstance.Send(req)
	if err != nil {
		return err
	}
	if statusCode != 200 {
		if result.ResponseMetadata.Error != nil {
			return fmt.Errorf("%s", result.ResponseMetadata.Error.Message)
		}
		return errors.New("unknown error happened while sending message")
	}

	return nil
}
