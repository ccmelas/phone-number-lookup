package main

import (
	"testing"
)

func TestPhoneNumber_ResolveDetails(t *testing.T) {
	pn := NewPhoneNumber("4387652676", "CA", true)
	details, err := pn.ResolveDetails()

	if err != nil {
		t.Error("Details should be returned for valid inputs")
	}

	if details.AreaCode != "438" {
		t.Error("Area should be correctly returned for valid inputs")
	}

	if details.LocalPhoneNumber != "7652676" {
		t.Error("LocalPhoneNumber should be correctly returned for valid inputs")
	}
}
