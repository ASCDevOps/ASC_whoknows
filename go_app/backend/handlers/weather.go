package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type WeatherHandler struct{}

var weatherTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (*WeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	_ = weatherTemplate.Execute(w, nil)
}

type WeatherAPIHandler struct{}

type WeatherResponse struct {
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	GenerationTime float64 `json:"generationtime_ms"`
	UTCOffsetSec   int     `json:"utc_offset_seconds"`
	Timezone       string  `json:"timezone"`
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
		WindDir     float64 `json:"winddirection"`
		WeatherCode int     `json:"weathercode"`
		Time        string  `json:"time"`
	} `json:"current_weather"`
}

// GET /api/weather - Location is set to Copenhagen right now.
func (*WeatherAPIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	weather, err := GetCopenhagenWeather()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch weather: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(weather); err != nil {
		http.Error(w, fmt.Sprintf("Failed encoding JSON: %v", err), http.StatusInternalServerError)
	}
}

// GetCopenhagenWeather GETS weatherdata from Open-Meteo API
func GetCopenhagenWeather() (*WeatherResponse, error) {
	url := "https://api.open-meteo.com/v1/forecast?latitude=55.6761&longitude=12.5683&current_weather=true"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, fmt.Errorf("Failed parsing JSON: %v", err)
	}

	return &weather, nil
}
