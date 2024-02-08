package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	configFile := flag.String("config", "config.json", "path to config file")
	flag.Parse()

	// Load and parse the configuration file
	config, err := config.loadConfig(*configFile)
	if err := config.loadConfig(*configFile, &config); err != nil {
		fmt.Println("Failed to load config:", err)
		os.Exit(1)
	}

	// Use the config...
	fmt.Printf("Channel: %+v\n", config)
}
