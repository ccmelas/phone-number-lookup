# Phone Number Lookup

An application for retrieving phone number information

## INSTALLATION & EXECUTION

- Clone the application using provided URL
- Ensure you have docker installed
- In your terminal, navigate to application folder and run the following command ```docker-compose up -d```
- The application is now available on `localhost:8030`. If port 8030 is in use, open the docker-compose.yml file, edit LN accordingly and repeat the previous step.
- You can send requests to /v1/phone-numbers to get phone number information

## CHOICE OF PROGRAMMING LANGUAGE, FRAMEWORK, LIBRARY.

### phonenumber [github.com/nyaruka/phonenumbers]
After some research, this library had all I needed to decode information about phone numbers. It is wrapper for Google's libphonenumber library, listed on the github page
### go-countries[github.com/mikekonan/go-countries]
I use this to validate Alpha2 country codes
### golang 
I could pretty much do this in a number of languages due to the availability of different ports of the google phone numbers library. However, I chose Go as it is very efficient. It is built to natively support concurrent and parallel code execution. This problem seems like one that can benefit from these features in the long run. Also, it is one of Oxio's primary languages for work.

## DEPLOYMENT.
I would deploy this as is using GCP or AWS

## ASSUMPTIONS.
- We always only query for one phone number at a time
- No need for handling a lot of concurrent requests
- Availability and Scalability aren't issues we need to deal with at this time

## IMPROVEMENTS
- I would love to better modularize my code. Split the key parts into different modules.
- I would love to introduce the use of interfaces for resolving phone details. This way I do not have tight coupling between my handler and the library I am current using
- I would love to introduce some go routines for resolving phone number information. This way we can handle more requests.
- Add more tests to improve code coverage
- Properly modularize tests so I can properly test different scenarios