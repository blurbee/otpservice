package util

import (
	"fmt"
	"os"
	"testing"

	"github.com/blurbee/otpserver/api"
)

var cfg Config

func TestMain(m *testing.M) {
	InitLogs()

	code := m.Run()

	// Tear down code

	os.Exit(code)
}

func TestLoadBadConfig(t *testing.T) {
	file := "../testdata/config-bad.yaml"
	f, err := os.Open(file)
	if err != nil {
		t.Fatal("Test config file not found: ", file)
	}
	defer f.Close()
	err = LoadConfig(f, &cfg)
	if err == api.OK {
		t.Fatal("load bad config failed", err)
	}
}

func TestLoadConfig(t *testing.T) {
	file := "../testdata/test-config.yaml"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Test config file not found: ", file)
		os.Exit(-1)
	}
	defer f.Close()
	err = LoadConfig(f, &cfg)
	if err != api.OK {
		fmt.Println("load config failed", err)
		os.Exit(-1)
	}

	s, err := cfg.GetScenario("loginauth")
	if err != api.OK {
		Debug(s)
		t.Fatal("load config failed", err)
	}

	if s.KeyLeadingText != "boo-" {
		t.Log("load config failed", err)
		t.Fatal()
	}

	s, err = cfg.GetScenario("pwdreset")
	if err != api.OK {
		t.Fatal("load config failed", err)
	}

	if s.AllowWhatsapp || s.KeyLeadingText != "pwd-" {
		Debug(s)
		t.Fatal("load config failed", err)
	}

}
