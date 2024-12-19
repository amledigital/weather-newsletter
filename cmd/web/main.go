package main

import (
	"fmt"
	"log"
	"time"

	"github.com/amledigital/weather-newsletter/internal/config"
	"github.com/amledigital/weather-newsletter/internal/mailer"
	"github.com/amledigital/weather-newsletter/internal/repo"
	"github.com/amledigital/weather-newsletter/internal/repo/baronrepo"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type BaronHandler struct {
	client repo.BaronServicer
}

var app = config.NewAppConfig()

type DummyUser struct {
	Email   string
	Zipcode string
}

func main() {
	dummyUsers := []DummyUser{}

	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "JoeWarner@910MediaGroup.com",
		Zipcode: "90026",
	})
	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "JoeWarner@910mediagroup.com",
		Zipcode: "49601",
	})
	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "JoeWarner@910mediagroup.com",
		Zipcode: "32080",
	})
	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "michaelstevens@910MediaGroup.com",
		Zipcode: "90026",
	})
	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "michaelstevens@910mediagroup.com",
		Zipcode: "49601",
	})
	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "michaelstevens@910mediagroup.com",
		Zipcode: "32080",
	})
	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "annette@ledigital.com",
		Zipcode: "90026",
	})
	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "annette@ledigital.com",
		Zipcode: "49601",
	})
	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "annette@ledigital.com",
		Zipcode: "32080",
	})

	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "andrew@ledigital.com",
		Zipcode: "90026",
	})
	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "andrew@ledigital.com",
		Zipcode: "49601",
	})
	dummyUsers = append(dummyUsers, DummyUser{
		Email:   "andrew@ledigital.com",
		Zipcode: "32080",
	})

	readConfig(app)

	var mailerClient mailer.Mailer

	mailEnvironment := viper.Get("environment").(string)

	smtp := viper.Get(fmt.Sprintf("%s_smtp_host", mailEnvironment)).(string)
	port := viper.Get(fmt.Sprintf("%s_smtp_port", mailEnvironment)).(int)
	user := viper.Get(fmt.Sprintf("%s_smtp_user", mailEnvironment)).(string)
	password := viper.Get(fmt.Sprintf("%s_smtp_password", mailEnvironment)).(string)
	sender := viper.Get(fmt.Sprintf("%s_smtp_sender", mailEnvironment)).(string)
	mailerClient = mailer.New(smtp, port, user, password, sender)

	doneChan := make(chan bool)

	for _, v := range dummyUsers {
		var bh BaronHandler

		bh.client = baronrepo.NewBaronServicer(app, nil, resty.New())

		zipData, err := bh.client.FetchGeo(v.Zipcode)

		if err != nil {
			log.Fatalln(err)
		}

		hourlyData, err := bh.client.FetchHourlyPointForecast(zipData.GeoCode.Data[0].Coordinates[0], zipData.GeoCode.Data[0].Coordinates[1], 16)

		if err != nil {
			log.Fatalln(err)
		}
		tmplData := make(map[string]any)

		tmplData["HourlyData"] = hourlyData
		tmplData["EmailAddress"] = v.Email
		tmplData["Zipcode"] = v.Zipcode
		tmplData["City"] = zipData.GeoCode.Data[0].City

		time.Sleep(time.Second * 1)
		err = mailerClient.Send(v.Email, "demo-weather.tmpl", tmplData)

		if err != nil {
			log.Panic(err)
		}
	}
	doneChan <- true

	<-doneChan

}
