package store

import (
	"fmt"
	"os"
	"testing"

	"caluxor.com/api"
	"caluxor.com/util"
)

func TestInitMongo(t *testing.T) {

	// t.Setenv("MONGO_URL", "mongodb://admin:blueh5ashMy%24test@localhost:27017/projects")

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

	err = InitMongo(&cfg)
	if err != api.OK {
		fmt.Println("Connection failed")
		os.Exit(-1)
	}
}

func TestGetAttribute(t *testing.T) {
	// t.Setenv("MONGO_URL", "mongodb://admin:blueh5ashMy%24test@localhost:27017/projects")

	util.InitLogs()
	file := "../testdata/test-config.yaml"
	f, err := os.Open(file)
	if err != nil {
		t.Fatal("Test config file not found: ", file)
	}
	defer f.Close()
	err = util.LoadConfig(f, &cfg)
	if err != api.OK {
		t.Fatal("load config failed", err)
	}

	err = InitMongo(&cfg)
	if err != api.OK {
		t.Fatal("Connection failed")
	}

	m, err := GetDB("018d0df5-a50d-7017-883b-8431e708d4b2")
	if err != api.OK {
		t.Fatal("DB instance not found")
	}

	email, err := m.GetEmail("3F190455-6D4B-446E-AAAE-DAEF761B318B")
	if err != api.OK || email != "test@testemail.com" {
		t.Fatal("Email not fetched.")
	}

}
