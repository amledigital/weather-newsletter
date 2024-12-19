package main

import (
	"github.com/amledigital/weather-newsletter/internal/config"
	"github.com/spf13/viper"
)

func readConfig(cfg *config.AppConfig) error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return err
		} else {
			return err
		}
	}

	cfg.Version = viper.Get("version").(string)
	cfg.BaronConfig.AccessKey = viper.Get("baron_access_key").(string)
	cfg.BaronConfig.SecretKey = viper.Get("baron_secret_key").(string)
	cfg.BaronConfig.ApiURL = viper.Get("baron_api_url").(string)

	return nil

}
