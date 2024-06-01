package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		http.Error(w, "API key not found", http.StatusInternalServerError)
		return
	}

	city := mux.Vars(r)["city"]

	client := resty.New()
	resp, err := client.R().
		SetQueryParam("q", city).
		SetQueryParam("appid", apiKey).
		SetQueryParam("units", "metric").
		Get("https://api.openweathermap.org/data/2.5/weather")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var weatherData map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &weatherData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	main, ok := weatherData["main"].(map[string]interface{})
	if !ok {
		http.Error(w, "Invalid response from weather API", http.StatusInternalServerError)
		return
	}
	weather, ok := weatherData["weather"].([]interface{})
	if !ok || len(weather) == 0 {
		http.Error(w, "Invalid response from weather API", http.StatusInternalServerError)
		return
	}
	weatherDetails, ok := weather[0].(map[string]interface{})
	if !ok {
		http.Error(w, "Invalid response from weather API", http.StatusInternalServerError)
		return
	}

	temperature, ok := main["temp"].(float64)
	if !ok {
		http.Error(w, "Invalid temperature data", http.StatusInternalServerError)
		return
	}
	description, ok := weatherDetails["description"].(string)
	if !ok {
		http.Error(w, "Invalid description data", http.StatusInternalServerError)
		return
	}

	response := WeatherResponse{
		Temperature: temperature,
		Description: description,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := 8080
	if val, exists := os.LookupEnv("PORT"); exists {
		if p, err := strconv.Atoi(val); err == nil {
			port = p
		}
	}

	r := mux.NewRouter()

	// Serve static files from the ./static directory
	staticDir := "./static"
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))

	r.HandleFunc("/weather/{city}", getWeather).Methods("GET")

	log.Printf("Server listening on port %d", port)
	http.ListenAndServe(":"+strconv.Itoa(port), r)
}
