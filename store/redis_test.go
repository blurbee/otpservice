package store

import (
	"fmt"
	"os"
	"testing"
	"time"

	"caluxor.com/api"

	"caluxor.com/util"
)

var cfg util.Config

func TestMain(m *testing.M) {
	// Setup code
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

	code := m.Run()

	// Teardown code
	os.Exit(code)
}

func TestInit(t *testing.T) {
	var k RedisStore
	err := k.Init(&cfg)
	if err != api.OK {
		t.Log("Redis connection failed")
		t.Fatal()
	}

	if k.Close() != api.OK {
		t.Log("Failed connection close.")
		t.Fatal()
	}
}

func TestSetKey(t *testing.T) {
	var k RedisStore
	err := k.Init(&cfg)
	if err != api.OK {
		t.Log("Redis connection failed")
		t.Fatal()
	}
	defer k.Close()

	TEST_VALUE := "12345"
	err = k.StoreKey("my-session", TEST_VALUE, time.Second)
	if err != api.OK {
		t.Fatal()
	}

	val, err := k.RetrieveKey("my-session")
	if err != api.OK || val != TEST_VALUE {
		t.Fatal()
	}

	// check key expiration
	time.Sleep(time.Second)
	val, err = k.RetrieveKey("my-session")
	if err == api.OK || val == TEST_VALUE {
		t.Fatal()
	}

}
