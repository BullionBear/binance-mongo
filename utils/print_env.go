package utils

import (
	"github.com/BullionBear/binance-mongo/env"
	"github.com/golang/glog"
)

func PrintEnv(name string) {
	// PrintEnv must put after the flag setup
	glog.Infoln("Process Name: ", name)
	glog.Infoln("Commit Hash: ", env.CommitHash)
	glog.Infoln("Version: ", env.Version)
	glog.Infoln("Build time: ", env.BuildTime)
}
