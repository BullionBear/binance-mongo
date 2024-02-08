package main

import (
	"flag"

	"github.com/BullionBear/binance-mongo/config"
	"github.com/golang/glog"
)

func main() {
	// Parse configuration file path from command-line arguments
	configPath := flag.String("config", "config.json", "path to config file")
	flag.Parse()

	// Load the configuration
	config := config.LoadConfig(*configPath)

	// Use logrus for logging
	// Adjust the logging level as needed
	glog.Infoln("Loaded config: ", config)
	defer glog.Flush()

}
