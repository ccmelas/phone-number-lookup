package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func PhoneNumbersHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	phoneNumberParam := request.URL.Query().Get("phoneNumber")
	countryCodeParam := request.URL.Query().Get("countryCode")

	// pre-validate
	err := ValidatePhoneNumber(phoneNumberParam)
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(writer).Encode(ValidationError{
			PhoneNumber: phoneNumberParam,
			Error:       map[string]string{"phoneNumber": "invalid phone number"},
		})
		return
	}

	// check if we need a country code
	requiresCountryCode := DetermineIfPhoneNumberRequiresCountryCode(phoneNumberParam)
	if requiresCountryCode {
		// validate country code
		if err := ValidateCountryCode(countryCodeParam); err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(writer).Encode(ValidationError{
				PhoneNumber: phoneNumberParam,
				Error:       map[string]string{"countryCode": "required value is missing"},
			})
			return
		}
	}

	phoneNumber := NewPhoneNumber(phoneNumberParam, countryCodeParam, requiresCountryCode)
	details, err := phoneNumber.ResolveDetails()
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(writer).Encode(ValidationError{
			PhoneNumber: phoneNumberParam,
			Error:       map[string]string{"phoneNumber": err.Error()},
		})
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(details)
}

func main() {
	http.HandleFunc("/v1/phone-numbers", PhoneNumbersHandler)
	println("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
