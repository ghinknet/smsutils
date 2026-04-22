package utils

import (
	"strings"

	"github.com/nyaruka/phonenumbers"
	"go.gh.ink/toolbox/pointer"
)

// ParseNumber splits a phone number into country code, national number and region code
func ParseNumber(number string) (countryCode int64, nationalNumber int64, regionCode string, err error) {
	// Pre-process number
	if !strings.HasPrefix(number, "+") {
		number = strings.Join([]string{"+", number}, "")
	}

	// Try to split
	split, err := phonenumbers.Parse(number, "")
	if err != nil {
		return 0, 0, "", err
	}

	return int64(pointer.SafeDeref(split.CountryCode)),
		int64(pointer.SafeDeref(split.NationalNumber)),
		phonenumbers.GetRegionCodeForNumber(split), nil
}
