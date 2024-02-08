package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Channel    string `json:"channel"`
	MongoDBUrl string `json:"mongodb"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	file, err := os.Open(path)
	if err != nil {
		return config, err // Return the zero value of Config and the error
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err // Return the zero value of Config and the error
	}

	return config, nil // Return the loaded config and no error
}
