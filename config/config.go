package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Channel    string `mapstructure:"channel"`
	MongoDBUrl string `mapstructure:"mongodb"`
}

func LoadConfig(configPath string) Config {
	var config Config

	viper.SetConfigFile(configPath) // Set the path of the config file
	viper.AutomaticEnv()            // Override values from environment variables

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %s", err)
	}

	return config
}
