package main

import (
	"errors"
	country "github.com/mikekonan/go-countries"
	"github.com/nyaruka/phonenumbers"
	"regexp"
	"strings"
	"unicode"
)

type ValidationError struct {
	PhoneNumber string            `json:"phoneNumber"`
	Error       map[string]string `json:"error"`
}

func ValidatePhoneNumber(phoneNumber string) error {
	// E.164 numbers can have a max of 15 digits. The country code is between 1-3
	// we can only have 2 spaces - after the country code. This can be at position 1,2 or 3
	// The second space can come anywhere between the area code and local number
	// Define the regular expression pattern
	pattern := `^(?:\+?[1-9]{1,3}\s?)?\d{1,14}\s?\d{1,14}$`
	// Create a regular expression object
	regex := regexp.MustCompile(pattern)
	// Check if the phone number matches the pattern
	matches := regex.MatchString(phoneNumber)

	if !matches {
		return errors.New("invalid phone number")
	}

	// The final check is if we have a valid number of digits (<= 15)
	digits := 0
	for i := 0; i < len(phoneNumber); i++ {
		if unicode.IsDigit(rune(phoneNumber[i])) {
			digits++
		}

		if digits > 15 {
			return errors.New("invalid phone number")
		}
	}

	return nil
}

func ValidateCountryCode(countryCode string) error {
	_, ok := country.ByAlpha2Code(country.Alpha2Code(countryCode))
	if ok {
		return nil
	}
	return errors.New("invalid country code")
}

func DetermineIfPhoneNumberRequiresCountryCode(phoneNumber string) bool {
	if strings.HasPrefix(phoneNumber, "+") {
		return false
	}
	pn, _ := phonenumbers.Parse("+"+phoneNumber, "")
	return !phonenumbers.IsValidNumber(pn)
}
