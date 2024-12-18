package models

type GeoCode struct {
	GeoCode struct {
		Data []GeoCodeLocation `json:"data,omitempty"`
	} `json:"geocode,omitempty"`
}

type GeoCodeLocation struct {
	City                   string `json:"city,omitempty"`
	GeoCodeLocationCountry `json:"country"`
	Coordinates            []float32 `json:"coordinates,omitempty"`
	GeoCodeLocationRegion  `json:"region"`
	GeoCodeLocationCounty  `json:"county"`
}

type GeoCodeLocationCountry struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type GeoCodeLocationRegion struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type GeoCodeLocationCounty struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
