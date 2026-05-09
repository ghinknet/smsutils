package utils

import (
	"strconv"
	"strings"
)

// ProcessNumberForChinese provides a method to process number to standard format for Chinese applications
// Special handling: 11-digit numbers starting with 1 and with second digit 3-9 are treated as Chinese numbers
// (e.g., 17601205205 -> +8617601205205, where second digit 7 is in range [3-9])
func ProcessNumberForChinese(to string) (
	toProcessed string,
	countryCode int64,
	nationalNumber int64,
	regionCode string,
	err error,
) {
	originalTo := to

	// Special case: detect Chinese mobile number format (11 digits, starts with 1, second digit 3-9)
	// Chinese mobile: starts with 13x-19x (1 followed by 3-9 as second digit, then 9 more digits)
	// US number: +1 followed by 10 digits (area code 2-9, exchange 2-9, then 8 more digits)
	// This helps distinguish: 17601205205 (China) vs 12025551234 (US where 2 is area code first digit)
	isChinaFormat := func(s string) bool {
		if len([]rune(s)) != 11 {
			return false
		}
		if !strings.HasPrefix(s, "1") || strings.HasPrefix(s, "+") {
			return false
		}
		// Check if second character is 3-9 (Chinese mobile operator ID)
		if len(s) < 2 {
			return false
		}
		secondChar := s[1]
		return secondChar >= '3' && secondChar <= '9'
	}

	// Standard parsing
	countryCode, nationalNumber, regionCode, err = ParseNumber(to)
	if err != nil {
		return "", 0, 0, "", err
	}

	// If ParseNumber thinks it's US but matches China format, override it
	if countryCode == 1 && regionCode == "US" && isChinaFormat(originalTo) {
		chinaNumber := strings.Join([]string{"+86", originalTo}, "")
		countryCode, nationalNumber, regionCode, _ = ParseNumber(chinaNumber)
		return chinaNumber, countryCode, nationalNumber, regionCode, nil
	}

	// Normal formatting for successfully parsed numbers
	if regionCode != "" {
		to = strings.Join([]string{
			"+", strconv.FormatInt(countryCode, 10), strconv.FormatInt(nationalNumber, 10),
		}, "")
	} else {
		// Fallback: unknown region code
		to = strings.Join([]string{"+86", originalTo}, "")
		regionCode = "CN"
	}

	return to, countryCode, nationalNumber, regionCode, nil
}
