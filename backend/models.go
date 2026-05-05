package main

import (
	"sync"
	"time"
)

type weatherData struct {
	Name     string `json:"name"`
	Timezone int    `json:"timezone"` // Offset in seconds from UTC
	Dt       int64  `json:"dt"`       // Current Unix timestamp
	Sys      struct {
		Country string `json:"country"`
	} `json:"sys"`
	Main struct {
		Kelvin   float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`

	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`

	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
}

type cacheItem struct {
	data      map[string]interface{}
	expiresAt time.Time
}

type WeatherCache struct {
	data  map[string]cacheItem
	mutex sync.RWMutex
}