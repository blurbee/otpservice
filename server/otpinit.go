package main

import (
	"fmt"
	"os"

	"github.com/blurbee/otpserver/api"
	"github.com/blurbee/otpserver/util"
)

func InitEnvironment(cfgFile string) (cfg *util.Config, err api.StatusCode) {
	util.InitLogs()
	file := "./config/config.yaml"
	f, er := os.Open(file)
	if er != nil {
		fmt.Println("Test config file could not be opened: ", file, er.Error())
		os.Exit(-1)
	}
	defer f.Close()
	cfg = new(util.Config)
	err = util.LoadConfig(f, cfg)
	if err != api.OK {
		fmt.Println("load config failed", err)
		os.Exit(-1)
	}

	return cfg, api.OK

}
