package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort      string `mapstructure:"SERVER_PORT"`
	GinMode         string `mapstructure:"GIN_MODE"`
	KafkaHost       string `mapstructure:"KAFKA_HOST"`
	KafkaPort       string `mapstructure:"KAFKA_PORT"`
	KafkaTopicOrder string `mapstructure:"KAFKA_TOPIC_ORDER"`
	KafkaTopics     []string
	KafkaSeeds      []string
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

	Config.KafkaSeeds = []string{fmt.Sprintf("%s:%s", Config.KafkaHost, Config.KafkaPort)}
	Config.KafkaTopics = []string{Config.KafkaTopicOrder}

	return nil
}
