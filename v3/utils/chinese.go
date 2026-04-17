package utils

import (
	"strconv"
	"strings"
)

// ProcessNumberForChinese provides a method to process number to standard format for Chinese applications
func ProcessNumberForChinese(to string) (
	toProcessed string,
	countryCode int64,
	nationalNumber int64,
	regionCode string,
	err error,
) {
	countryCode, nationalNumber, regionCode, err = ParseNumber(to)
	if err != nil {
		return "", 0, 0, "", err
	}
	if regionCode == "" && len([]rune(to)) == 11 && strings.HasPrefix(to, "1") {
		to = strings.Join([]string{"+86", to}, "")
		regionCode = "CN"
	} else {
		to = strings.Join([]string{
			"+", strconv.FormatInt(countryCode, 10), strconv.FormatInt(nationalNumber, 10),
		}, "")
	}

	return to, countryCode, nationalNumber, regionCode, nil
}
