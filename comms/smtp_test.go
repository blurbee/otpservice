package comms

import (
	"fmt"
	"os"
	"testing"

	"github.com/blurbee/otpserver/api"
	"github.com/blurbee/otpserver/util"
)

func TestSMTPTest(t *testing.T) {
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

	esender, err := SMTPInit(&cfg)
	if err != api.OK {
		t.Fatal("Failed to create smtp.", err)
	}

	if esender == nil {
		t.Fatal("Creating Email sender failed.")
	}
}
