package util

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/blurbee/otpserver/api"
)

func TestObfuscateEmail(t *testing.T) {
	input := "xyz@gmail.com"
	output, err := ObfuscateEmail(input)
	if err != api.OK || output != "x***z@gmail.com" {
		t.Log("obfuscation failed")
		t.Fail()
	}
}

func TestObfuscateEmailWithBadInput(t *testing.T) {
	input := "aotan"
	_, err := ObfuscateEmail(input)
	if err == api.OK {
		t.Log("failure case failed")
		t.Fail()
	}
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func isAlpha(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func TestGenerateRandom(t *testing.T) {
	s := GenerateRandom(15, OTP_ONLY_NUM)
	if !isInt(s) || len(s) != 15 {
		t.Fatal("Error generating only random numbers. [", s, "]")
	}

	s = GenerateRandom(10, OTP_ONLY_ALPHA)
	if !isAlpha(s) || len(s) != 10 {
		fmt.Println(s)
		t.Fatal("Error generating only random alpha. [", s, "]")
	}

	s = GenerateRandom(15, OTP_ALPHA_NUM)
	if len(s) != 15 {
		fmt.Println(s)
		t.Fatal("Error generating alpha-numeric. [", s, "]")
	}
}
