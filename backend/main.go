package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const allowedOrigin = "http://localhost:4200" // allow frontend requests
var cache = NewWeatherCache()

func main() {
    godotenv.Load()
    port := ":8080"

     fmt.Println("Server is running on port" + port)

    http.HandleFunc("/weather", weatherHandler)
    http.HandleFunc("/weather/", weatherHandler)

	http.ListenAndServe(port, nil)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %s %s | Query: %s\n", r.Method, r.URL.Path, r.URL.RawQuery)
 
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")

    cityName := strings.ToLower(getCityName(r))

    if cityName == "" {
        http.Error(w, "City required", http.StatusBadRequest)
        return
    }

    // Cache check
    if data, ok := cache.Get(cityName); ok {
        fmt.Printf("Cache FROM_CACHE: %s\n", cityName)
        w.Header().Set("X-Cache", "FROM_CACHE") // Optional: custom header to inform the client/frontend that data was served from local memory
        json.NewEncoder(w).Encode(data)
        return
    }

    // Fetch with openweathermap API
    wdata, err := fetchWeather(cityName)
    if err != nil || wdata.Name == "" {
		fmt.Printf("cityName akouza: n"  )
        http.Error(w, "Weather data not found", http.StatusNotFound)
        return
    }

    response := formatWeatherResponse(wdata)

    // Cache set
    cache.Set(cityName, response, 1*time.Minute)

    w.Header().Set("X-Cache", "FROM_EXT_API") // Optional: custom header to inform the client that a real API call was required
    json.NewEncoder(w).Encode(response)
}