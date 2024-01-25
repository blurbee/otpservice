package api

import "time"

type KeyStore interface {
	//	GetPartialEmail(fullEmail string) (obfustcatedEmail string)

	//	GetPartialPhone(fullPhone string) (obfuscatedPhone string)

	StoreKey(sessionid string, otpcode string, expires time.Duration) (err StatusCode)

	RetrieveKey(sessionid string) (otpcode string, err StatusCode)
}

type UserStore interface {
	// Given user id, get the user's email address
	GetEmail(userid string) (email string, err StatusCode)

	// Given user id, get the user's email address
	GetPhone(userid string) (phone string, err StatusCode)

	// Given user id, get the whatsapp number
	GetWhatsapp(userid string) (whatsapp string, err StatusCode)

	// Given user id, get the text number
	GetText(userid string) (text string, err StatusCode)
}

type TextSender interface {
	SendText(dest string, message string) (err StatusCode)
}

type EmailSender interface {
	SendEmail(dest string, subject string, message string) (err StatusCode)
}

type OTPService interface {
	InitService(userStore *UserStore, keystore *KeyStore,
		texter *TextSender, emailer *EmailSender) (err StatusCode)
	CreateOTPSession(userid string, scenarioid string) (sessionid string, err StatusCode)
	VerifyUser(sessionid string, otpresponse string) (success bool, err StatusCode)
}
