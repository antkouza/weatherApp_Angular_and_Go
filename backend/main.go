package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var weatherCache = make(map[string]cacheItem)
var cacheMutex sync.Mutex // Prevents crashes if two users search at once

func main() {
    godotenv.Load()
    port := ":8080"

    // Print to console before starting server
    fmt.Println("Server is running on port" + port)

    http.HandleFunc("/weather", weatherHandler)
    http.HandleFunc("/weather/", weatherHandler)

	http.ListenAndServe(port, nil)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("Request: %s %s | Query: %s\n",    r.Method,    r.URL.Path,    r.URL.RawQuery,    )

    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	cityName := getCityName(r)
	if cityName == "" {
		http.Error(w, "City required", http.StatusBadRequest)
		return
	}

	// 1. Check Cache
	cacheMutex.Lock()
	item, found := weatherCache[strings.ToLower(cityName)]
	cacheMutex.Unlock()

    if found {
        if time.Now().Before(item.expiresAt) {
            fmt.Printf("Cache HIT: %s\n", cityName)
            w.Header().Set("X-Cache", "HIT")
            json.NewEncoder(w).Encode(item.data)
            return
        } 
        
        // logic: Found but expired? Delete it!
        cacheMutex.Lock()
        delete(weatherCache, strings.ToLower(cityName))
        cacheMutex.Unlock()
        fmt.Printf("Cache EXPIRED: %s - Fetch new data\n", cityName)
    }

	// 2. Fetch from API
	wdata, err := query(cityName)
	if err != nil || wdata.Name == "" {
		http.Error(w, "Weather data not found", http.StatusNotFound)
		return
	}

	// 3. Format and Cache
	response := formatWeatherResponse(wdata)
    fmt.Printf("Response for %s: %+v\n", cityName, response)

	cacheMutex.Lock()
	weatherCache[strings.ToLower(cityName)] = cacheItem{
		data:      response,
		expiresAt: time.Now().Add(1 * time.Minute),
	}
	cacheMutex.Unlock()

	w.Header().Set("X-Cache", "MISS")
	json.NewEncoder(w).Encode(response)
}

func getCityName(r *http.Request) string {
	cityName := r.URL.Query().Get("city")
	if cityName != "" {
		return cityName
	}

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) > 1 { // /weather/London -> parts is ["weather", "London"]
		return parts[len(parts)-1]
	}
	return ""
}