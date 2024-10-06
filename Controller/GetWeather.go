package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	model "weather/Model"
)

const weatherApiUrl = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline"

type WeatherResponse struct {
	Latitude        float64     `json:"latitude"`
	Longitude       float64     `json:"longitude"`
	ResolvedAddress string      `json:"resolvedAddress"`
	Days            []model.Day `json:"days"`
}

func GetWeather(Location string) (*WeatherResponse, error) {
	//Get api key
	apiKey := os.Getenv("WEATHER_API_KEY")
	//get apiUrl
	apiUrl := fmt.Sprintf("%s/%s?key=%s", weatherApiUrl, Location, apiKey)

	log.Printf("Fetching weather on :%v", Location)

	// Build the URL with query parameters
	//-> Seperate url into components
	apiUrlBuilder, err := url.Parse(apiUrl)

	if err != nil {
		return nil, fmt.Errorf("Some thing wrong when parse baseURL : %v", err)
	}
	// Add query parameters
	//Values is a map which store every query parameters and form values
	//Values{} "key":"value "
	params := url.Values{}
	//Add into the Values
	params.Add("key", apiKey)
	params.Add("contentType", "json")
	//Using metric system like {wind speed: km/h,temp:*c}..
	params.Add("unitGroup", "metric")

	// Make the GET request
	resp, err := http.Get(apiUrlBuilder.String())
	if err != nil {
		return nil, fmt.Errorf("Error making Get request : %v", err)
	}
	defer resp.Body.Close()
	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code :%v", err)
	}
	// Decode the JSON response into weather struct
	var apiResponse WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("Cant decode into weather struct")
	}
	log.Printf("API response : %+v", apiResponse)
	return &apiResponse, nil
}
