package main

import (
	"fmt"
	"log"

	"github.com/amledigital/weather-newsletter/internal/config"
	"github.com/amledigital/weather-newsletter/internal/repo"
	"github.com/amledigital/weather-newsletter/internal/repo/baronrepo"
	"github.com/go-resty/resty/v2"
)

type BaronHandler struct {
	client repo.BaronServicer
}

var app = config.NewAppConfig()

func main() {

	readConfig(app)

	var bh BaronHandler

	bh.client = baronrepo.NewBaronServicer(app, nil, resty.New())

	data, err := bh.client.GetStation("KTVC")

	if err != nil {
		log.Fatalln(err)
	}
	data = data
	zipData, err := bh.client.FetchGeo("49685")

	if err != nil {
		log.Fatalln(err)
	}

	hourlyData, err := bh.client.FetchHourlyPointForecast(zipData.GeoCode.Data[0].Coordinates[0], zipData.GeoCode.Data[0].Coordinates[1], 6)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(hourlyData)

}
