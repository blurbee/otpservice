# One-Time Password Service

## Summary
The one-time password (aka otp) is a common service used by many systems as a solution for several use cases:
- multi-factor authentication
- email address or phone verification (communication control)
- password reset gateway
- transaction authorization check

The project is an open-source otp service built in Golang and uses redis to manage tokens in the backend. The user information such as email, phone, text is fetched from a user store. Currently mongo and postgres are supported as user stores.

# Getting started
1. Get the source git https://github.com/blurbee/otpservice.git
1. Use one of the configuration templates under `config` directory.
2. Substitute all properties in `<>` with appropriate values.
3. Create secrets.yaml with keys. MONGO_URL, MAIL_PASSWORD, are examples. The `secrets-template.yaml` has the list of keys that could be configured.
4. Run make img


## Configuration
### Scenario configuration
Given the variety of use cases in which OTP is used, the servic provides configuartion capability that can be used for specfic scenario. The configuration of each scenario can specify one of the following:
- key characteristics (length/character domain)
- key TTL (in seconds)
- supported dispatch channels (email/text/whatsapp)
- number of attempts to validate the OTP before it expires
- pre-key static tex

The configuration is designed to support extensibility to support multi-tenancy. The scenarios may be defined for any number of tenants. The plan is to implement configuration APIs and moving configuration into a database that will allow dynamic addition of new tenants.

Each scenario also identifies the user stores for properties such as phone, email, text or whatsapp. Given that each proprety can be defined separately, the service allows situations where these properties may be stores in different databases or different reconds in the same database. The `phonestoreconfig`, `emailstorecfg`, `whatsappstorecfg` each define which database connection should be used to retrieve the value for a given scenario.

### Database connections
Databases are configured in one of two lists: `mongostores` and `postgresstores`. These two are lists of connection parameters.

### Secrets
Properties such as user name and password for connecting to databases, redis, etc are stored in secrets file. The secrets file location is specified by `secretsfile`.

### Communication configuration
All outbound communications properties such as email, text and whatsapp are defined in `emailserverconfig`, `twilioconfig`, `phoneconfig` and `whatsappmsg`.


# Extensibility
The service is designed with key components as interfaces. Access of user data, transient storage of OTP keys, dispatching OTP can be swapped.

# How it works
Each scenario is identified by the scenarioid as defined in the config file. The following sequence diagram shows the process.

`
title This is a title

User->Application:User enters a flow that requries OTP
Application-->OTPServer:The App calls createOTPSession \n scenarioid(login/email-verification)
OTPServer->Application: return sessionid, channels
Application->User: Provide choice of how to receive OTP (channels)
User->Application: sendOTP(selected channel)
Application->OTPServer: sendOTP(ch, optional: channelValue)
`

UI componnt 
OTP sesion is required, the createOTPSession is called with the appropriate scenarioid.

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

 


