package comms

import (
	"fmt"
	"os"
	"testing"

	"github.com/blurbee/otpserver/api"
	"github.com/blurbee/otpserver/util"
)

var cfg util.Config

func TestTwilioInit(t *testing.T) {
	util.InitLogs()
	file := "../testdata/test-config.yaml"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Test config file not found: ", file)
		os.Exit(-1)
	}
	defer f.Close()
	err = util.LoadConfig(f, &cfg)
	if err != api.OK {
		fmt.Println("load config failed", err)
		os.Exit(-1)
	}

	tsender, err := TwilioInit(&cfg)
	if err != api.OK {
		t.Fatal("Failed to create twilio.", err)
	}

	if tsender == nil {
		t.Fatal("Creating Twilio text sender failed.")
	}
}

func TestTwilioMessage(t *testing.T) {
	util.InitLogs()
	file := "../testdata/test-config.yaml"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Test config file not found: ", file)
		os.Exit(-1)
	}
	defer f.Close()
	err = util.LoadConfig(f, &cfg)
	if err != api.OK {
		fmt.Println("load config failed", err)
		os.Exit(-1)
	}

	tsender, err := TwilioInit(&cfg)
	if err != api.OK {
		t.Fatal("Failed to create twilio.", err)
	}

	if tsender == nil {
		t.Fatal("Creating Twilio text sender failed.")
	}
	err = tsender.SendText("+14083739954", "Yo! Lakshmi, How are ya?")

	if err != api.OK {
		t.Fatal("Unable to send text message:", err)
	}
}
