package util

import (
	"bytes"
	"math/rand"
	"strings"
	"time"

	"github.com/blurbee/otpserver/api"
	"github.com/google/uuid"
)

type OTPDOMAIN int

const (
	OTP_ONLY_NUM OTPDOMAIN = iota
	OTP_ONLY_ALPHA
	OTP_ALPHA_NUM
)

const NUM_DOMAIN = "0123456789"
const ALPHANUM_DOMAIN = "0123456790ABCDEFGHJKLMNPQRSTUVWXYZ"
const ALPHA_DOMAIN = "ABCDEFGHJKLMNPQRSTUVWXFZ"
const MAX_OTP_LEN = 25

func GenerateRandom(l int, outType OTPDOMAIN) (otpstr string) {
	if l < 0 || l > MAX_OTP_LEN {
		return ""
	}

	vs := rand.NewSource(time.Now().UnixNano())
	v := rand.New(vs)
	var s string
	switch outType {
	case OTP_ONLY_NUM:
		s = NUM_DOMAIN
	case OTP_ONLY_ALPHA:
		s = ALPHA_DOMAIN
	case OTP_ALPHA_NUM:
		s = ALPHANUM_DOMAIN
	}

	var str strings.Builder

	for i := 0; i < l; i++ {
		str.WriteByte(s[v.Intn(len(s))])
	}

	return str.String()
}

func ObfuscateEmail(email string) (obfuscatedEmail string, err api.StatusCode) {
	atIndex := strings.Index(email, "@")
	if atIndex != -1 {
		var obfuscated bytes.Buffer

		obfuscated.WriteByte(email[0])
		obfuscated.WriteString("***")
		obfuscated.WriteByte(email[atIndex-1])
		obfuscated.WriteString(email[atIndex:])
		obfuscatedEmail = obfuscated.String()
	} else {
		return email, api.INVALID_INPUT // Invalid email, return as is
	}
	return obfuscatedEmail, api.OK
}

func GetUUID() (uid string, status api.StatusCode) {
	val, er := uuid.NewV7()
	if er != nil {
		Error("Unable to generate uuids:", er)
		return "", api.UNKNOWN_ERROR
	}
	uid = val.String()
	status = api.OK
	return
}
