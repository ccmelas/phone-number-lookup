package main

import (
	"testing"
)

func TestValidatePhoneNumber(t *testing.T) {
	if err := ValidatePhoneNumber("+12125690123"); err != nil {
		t.Error("Should not return an error for a valid number: +12125690123")
	}
	if err := ValidatePhoneNumber("+52 631 3118150"); err != nil {
		t.Error("Should not return an error for a valid number: +52 631 3118150")
	}
	if err := ValidatePhoneNumber("34 915 872200"); err != nil {
		t.Error("Should not return an error for a valid number: 34 915 872200")
	}
	if err := ValidatePhoneNumber("351 21 094 2000"); err == nil {
		t.Error("Should return an error for an invalid number: 351 21 094 2000")
	}
}

func TestValidateCountryCode(t *testing.T) {
	if err := ValidateCountryCode("US"); err != nil {
		t.Error("Should not return an error for a valid country code: US")
	}
	if err := ValidateCountryCode("ZZ"); err == nil {
		t.Error("Should return an error for an invalid country code: ZZ")
	}
	if err := ValidateCountryCode(""); err == nil {
		t.Error("Should return an error for an empty country code")
	}
}

func TestDetermineIfPhoneNumberRequiresCountryCode(t *testing.T) {
	if requiresCountryCode := DetermineIfPhoneNumberRequiresCountryCode("+12125690123"); requiresCountryCode != false {
		t.Error("Should return false if a number does not need a country code: +12125690123")
	}
	if requiresCountryCode := DetermineIfPhoneNumberRequiresCountryCode("34 915 872200"); requiresCountryCode != false {
		t.Error("Should return false if a number does not need a country code: 34 915 872200")
	}
	if requiresCountryCode := DetermineIfPhoneNumberRequiresCountryCode("438 765 3673"); requiresCountryCode != true {
		t.Error("Should return true if a number needs a country code")
	}
}
