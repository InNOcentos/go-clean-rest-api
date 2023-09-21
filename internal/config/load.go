package config

import "github.com/spf13/viper"

func Load() error {
	viper.AddConfigPath("../internal/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}
