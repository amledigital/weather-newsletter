package models

import "time"

type NDFDHourly struct {
	NDFDHourly struct {
		Data []WeatherData `json:"data"`
	} `json:"ndfd_hourly,omitempty"`
}

type WeatherData struct {
	Precipitation struct {
		Probability struct {
			Value float64 `json:"value"`
			Units string  `json:"units"`
		} `json:"probability"`
		Potential struct {
			Value float64 `json:"value"`
			Units string  `json:"units"`
		} `json:"potential"`
	} `json:"precipitation"`
	Temperature struct {
		Value     float64 `json:"value"`
		WindChill float64 `json:"wind_chill"`
		DewPoint  float64 `json:"dew_point"`
		Apparent  float64 `json:"apparent"`
		Units     string  `json:"units"`
	} `json:"temperature"`
	CloudCover struct {
		Value float64 `json:"value"`
		Units string  `json:"units"`
		Text  string  `json:"text"`
	} `json:"cloud_cover"`
	RelativeHumidity struct {
		Value float64 `json:"value"`
		Units string  `json:"units"`
	} `json:"relative_humidity"`
	Wind struct {
		Speed      float64 `json:"speed"`
		SpeedUnits string  `json:"speed_units"`
		Gust       float64 `json:"gust"`
		Dir        int     `json:"dir"`
		DirUnits   string  `json:"dir_units"`
	} `json:"wind"`
	WeatherCode struct {
		Value string `json:"value"`
		Text  string `json:"text"`
	} `json:"weather_code"`
	Daylight   bool      `json:"daylight"`
	ValidBegin time.Time `json:"valid_begin"`
}
