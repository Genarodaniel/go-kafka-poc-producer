package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	GinMode    string `mapstructure:"GIN_MODE"`
	KafkaHost  string `mapstructure:"KAFKA_HOST"`
	KafkaPort  string `mapstructure:"KAFKA_PORT"`
}

var Config Env

func Load() error {
	viper.AddConfigPath("../")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	logger := log.Default()
	logger.Print(&Config)

	return nil
}
