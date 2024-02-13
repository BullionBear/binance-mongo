package utils

import (
	"flag"
	"fmt"

	"github.com/BullionBear/binance-mongo/env"
)

var (
	versionFlag   bool
	commitFlag    bool
	buildTimeFlag bool
)

func MetadataParser() bool {
	// MetadataParser must put after the flag setup
	flag.BoolVar(&versionFlag, "version", false, "Print the version number")
	flag.BoolVar(&commitFlag, "commit", false, "Print the commit hash")
	flag.BoolVar(&buildTimeFlag, "build_time", false, "Print the build time")
	if versionFlag {
		fmt.Println("Version:", env.Version)
		return true
	}
	if commitFlag {
		fmt.Println("Commit:", env.CommitHash)
		return true
	}
	if buildTimeFlag {
		fmt.Println("Build Time:", env.BuildTime)
		return true
	}
	return false
}
