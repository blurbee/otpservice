# One-Time Password Service

## Summary
The one-time password (aka otp) is a common service used by many systems as a solution for several use cases:
- multi-factor authentication
- email address or phone verification (communication control)
- password reset gateway
- transaction authorization check

The project is an open-source otp service built in Golang and uses redis to manage tokens in the backend. The user information such as email, phone, text is fetched from a user store. Currently mongo and postgres are supported as user stores.

Configuration
Scenario configuration
Given the variety of use cases in which OTP is used, the servic provides configuartion capability that can be used for specfic scenario. The configuration of each scenario can specify one of the following:
- key characteristics (length/character domain)
- key TTL (in seconds)
- supported dispatch channels (email/text/whatsapp)
- number of attempts to validate the OTP before it expires
- pre-key static tex

Each scenario is identified by the scenarioid and is presented at the time of session creation.

User store configuration


This service has APIs to:
- create an OTP session based on pre-configured "scenario".
- if enabled, allow user to select which channel to recieve the otp on
- trigger sending on OTP on the selected channel (with or without validation of channel identity like full email or full phone number)
- validate the OTP
 
## APIs
The API specification has been described in api/openapi.yaml.

## Configuration
 There are two stores. Keystore and Userstore. Userstore is read-only and is used to retrieve the user's phone/email/text/whatsapp information. The supported userstore options are Mongo and Postgres.

 Keystore is used to store the key and expire them appropriately. Supported keystore is redis.

 


