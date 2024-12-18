package config

import "github.com/amledigital/weather-newsletter/internal/models"

type AppConfig struct {
	Version string `json:"version"`
	models.BaronConfig
	models.MailChimpConfig
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		BaronConfig:     models.NewBaronConfig(),
		MailChimpConfig: models.NewMailChimpConfig(),
	}
}
