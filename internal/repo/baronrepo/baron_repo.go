package baronrepo

import (
	"crypto/hmac"
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/amledigital/weather-newsletter/internal/config"
	"github.com/amledigital/weather-newsletter/internal/models"
	"github.com/go-resty/resty/v2"
)

type BaronRepo struct {
	cfg    *config.AppConfig
	conn   *sql.DB
	client *resty.Client
}

func NewBaronServicer(a *config.AppConfig, db *sql.DB, restClient *resty.Client) *BaronRepo {
	return &BaronRepo{
		cfg:    a,
		conn:   db,
		client: restClient,
	}
}

// Sign accepts a string to sign and a secret and produces a SHA1 hash and returns a string
// in the context of baron request, you are passing in key:ts as the string value to sign, and the baron provided secret
func (br *BaronRepo) Sign(stringToSign, secret string) string {

	if stringToSign == "" || secret == "" {
		return ""
	}
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(stringToSign))
	hashBytes := mac.Sum(nil)
	hashString := base64.StdEncoding.EncodeToString(hashBytes)

	return hashString

}

// SignRequest signs the api request to baron based on their documentation
// https://developer.baronweather.com/overview/authentication-authorization/
func (br *BaronRepo) SignRequest(urlStr, key, secret string) (string, error) {
	ts := time.Now().UTC().Unix()

	stringToSign := fmt.Sprintf("%s:%d", key, ts)
	signature := br.Sign(stringToSign, secret)

	// URL encode the signature
	encodedSig := url.QueryEscape(signature)

	// Determine correct separator (? or &)
	separator := "?"
	if strings.Contains(urlStr, "?") {
		separator = "&"
	}

	// Construct final URL
	signedURL := fmt.Sprintf("%s%ssig=%s&ts=%d",
		urlStr,
		separator,
		encodedSig,
		ts,
	)

	return signedURL, nil

}

func (br *BaronRepo) GetStation(station string) (any, error) {

	fmtURL := fmt.Sprintf("%s/%s/reports/metar/station/%s.json",
		br.cfg.ApiURL,
		br.cfg.AccessKey,
		station)

	signedRequest, err := br.SignRequest(fmtURL, br.cfg.AccessKey, br.cfg.SecretKey)

	resp, err := br.client.R().
		EnableTrace().
		Get(signedRequest)

	if err != nil {
		return nil, err
	}
	return resp.String(), nil
}

func (br *BaronRepo) FetchGeo(zipcode string) (*models.GeoCode, error) {
	fmtURL := fmt.Sprintf("%s/%s/reports/geocode/zip.json?zip=%s&us=1", br.cfg.ApiURL, br.cfg.AccessKey, zipcode)

	signedRequest, err := br.SignRequest(fmtURL, br.cfg.AccessKey, br.cfg.SecretKey)

	if err != nil {
		return nil, err
	}

	resp, err := br.client.R().EnableTrace().Get(signedRequest)

	if err != nil {
		return nil, err
	}

	if resp.Body() != nil {
		geoCode := models.GeoCode{}

		err = json.Unmarshal(resp.Body(), &geoCode)

		if err != nil {
			return nil, err
		}

		return &geoCode, nil
	}

	return nil, nil

}

func (br *BaronRepo) FetchHourlyPointForecast(lat, lon float32, hours int) (*models.NDFDHourly, error) {

	dateOnly := time.Now().Format(time.DateOnly)
	timeOnly := time.Now().Format(time.TimeOnly)

	fmtURL := fmt.Sprintf("%s/%s/reports/ndfd/hourly.json?lat=%.2f&lon=%.2f&hours=%d&utc=%s",
		br.cfg.ApiURL,
		br.cfg.AccessKey,
		lon,
		lat,
		hours,
		dateOnly+"T"+timeOnly,
	)

	fmt.Println(fmtURL)

	signedRequest, err := br.SignRequest(fmtURL, br.cfg.AccessKey, br.cfg.SecretKey)

	if err != nil {
		return nil, err
	}

	resp, err := br.client.R().EnableTrace().Get(signedRequest)

	if err != nil {
		return nil, err
	}

	if resp.Body() != nil {
		ndfdData := models.NDFDHourly{}

		err = json.Unmarshal(resp.Body(), &ndfdData)

		if err != nil {
			return nil, err
		}
		return &ndfdData, nil
	}

	return nil, nil
}
