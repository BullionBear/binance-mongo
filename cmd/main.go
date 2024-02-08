package main

import (
	"flag"
	"fmt"

	"github.com/BullionBear/binance-mongo/config"
)

func main() {
	// Parse configuration file path from command-line arguments
	configPath := flag.String("config", "config.json", "path to config file")
	flag.Parse()

	// Load the configuration
	config := config.LoadConfig(*configPath)
	fmt.Printf("Loaded config: %+v", config)
}
