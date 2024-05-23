package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
)

const apiKey = ""

type WeatherResponse struct {
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["city"]
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("q", city).
		SetQueryParam("appid", apiKey).
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

	temperature := weatherData["main"].(map[string]interface{})["temp"].(float64)
	description := weatherData["weather"].([]interface{})[0].(map[string]interface{})["description"].(string)

	response := WeatherResponse{
		Temperature: temperature,
		Description: description,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}", getWeather).Methods("GET")
	http.ListenAndServe(":8080", r)
}
