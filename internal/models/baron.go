package models

type BaronConfig struct {
	AccessKey string `json:"-"`
	SecretKey string `json:"-"`
	ApiURL    string `json:"api_url"`
}

func NewBaronConfig() BaronConfig {
	return BaronConfig{}
}
