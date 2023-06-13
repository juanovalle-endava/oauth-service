package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	Environment          string `mapstructure:"environment"`
	PrivateKeyPath       string `mapstructure:"private-key-path"`
	TokenDefaultDuration string `mapstructure:"token-default-duration"`
}

func NewEnv() Config {

	// AddConfigPath called multiple times for testing purposes (viper look for config file from the path we call NewEnv)
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	config := Config{}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("☠️ cannot read configuration", err)
		} else {
			fmt.Println("☠️ config file was found but another error was produced", err)
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("☠️ environment can't be loaded: ", err)
	}

	ForceMapping(&config)

	return config
}

func ForceMapping(env *Config) {

	if env.Environment == "" {
		env.Environment = viper.GetString("environment")
	}

	if env.PrivateKeyPath == "" {
		env.PrivateKeyPath = viper.GetString("private-key-path")
	}

	if env.TokenDefaultDuration == "" {
		env.TokenDefaultDuration = viper.GetString("token-default-duration")
	}
}

var Module = fx.Provide(NewEnv)
