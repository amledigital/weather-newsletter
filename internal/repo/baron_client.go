package repo

import "github.com/amledigital/weather-newsletter/internal/models"

type BaronServicer interface {
	Sign(stringToSign, secret string) string
	SignRequest(urlStr, key, secret string) (string, error)
	GetStation(station string) (any, error)
	FetchGeo(zipcode string) (*models.GeoCode, error)
	FetchHourlyPointForecast(lat, lon float32, hours int) (*models.NDFDHourly, error)
}
