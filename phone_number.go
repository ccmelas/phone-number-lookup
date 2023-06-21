package main

import (
	"errors"
	"github.com/nyaruka/phonenumbers"
	"strings"
	"unicode"
)

type PhoneNumber struct {
	phoneNumber         string
	countryCode         string
	requiresCountryCode bool
}

type PhoneNumberDetails struct {
	PhoneNumber      string `json:"phoneNumber"`
	CountryCode      string `json:"countryCode"`
	AreaCode         string `json:"areaCode"`
	LocalPhoneNumber string `json:"localPhoneNumber"`
}

func (phoneNumber PhoneNumber) StartsWithPlus() bool {
	return strings.HasPrefix(phoneNumber.phoneNumber, "+")
}

func (phoneNumber PhoneNumber) GetPrefixedPhoneNumber() string {
	if phoneNumber.StartsWithPlus() {
		return phoneNumber.phoneNumber
	}
	return "+" + phoneNumber.phoneNumber
}

func (phoneNumber PhoneNumber) GetFormattedPhoneNumber() string {
	return strings.ReplaceAll(phoneNumber.GetPrefixedPhoneNumber(), " ", "")
}

func (phoneNumber PhoneNumber) ResolveDetails() (*PhoneNumberDetails, error) {
	phone := phoneNumber.phoneNumber
	if !phoneNumber.requiresCountryCode {
		phone = phoneNumber.GetPrefixedPhoneNumber()
	}

	pN, _ := phonenumbers.Parse(phone, phoneNumber.countryCode)
	if !phonenumbers.IsValidNumber(pN) {
		return nil, errors.New("invalid phone number")
	}

	countryCode := phonenumbers.GetRegionCodeForNumber(pN)

	// hard-coded US intentionally
	formattedNumber := phonenumbers.FormatOutOfCountryCallingNumber(pN, "US")
	parts := make([]string, 0)
	current := ""
	for i := 0; i < len(formattedNumber); i++ {
		if unicode.IsDigit(rune(formattedNumber[i])) {
			current += string(formattedNumber[i])
			continue
		}

		if len(current) > 0 {
			parts = append(parts, current)
			current = ""
		}
	}

	if len(current) > 0 {
		parts = append(parts, current)
	}

	areaCode := parts[1]
	localPhoneNumber := ""
	for i := 2; i < len(parts); i++ {
		localPhoneNumber += parts[i]
	}

	return &PhoneNumberDetails{
		PhoneNumber:      "+" + strings.ReplaceAll(formattedNumber, " ", ""),
		LocalPhoneNumber: localPhoneNumber,
		AreaCode:         areaCode,
		CountryCode:      countryCode,
	}, nil
}

func NewPhoneNumber(phoneNumber string, countryCode string, requiresCountryCode bool) *PhoneNumber {
	return &PhoneNumber{
		phoneNumber:         phoneNumber,
		countryCode:         countryCode,
		requiresCountryCode: requiresCountryCode,
	}
}
