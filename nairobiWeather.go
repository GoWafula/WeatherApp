package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type WeatherData struct {
	Main struct {
		temp float64 `json:"temp"`
	} `json:"main"`

	Weather []struct {
		Description string `json: "description`
	} `json:"weather"`
}

func getWeatherData(city, apiKey string) (*WeatherData, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}
	return &weatherData, nil
}

func main() {
	apiKey := "01351b39fd9c1ffaa313ee6f021208d4"
	city := "Nairobi"

	weatherData, err := getWeatherData(city, apiKey)
	if err != nil {
		fmt.Print("Failed to fetch data ", err)
		os.Exit(1)
	}

	fmt.Printf("Weather in %s:\n", city)
	fmt.Printf("Temperature: %.1f°C\n", weatherData.Main.temp)
	fmt.Printf("Description: %s\n", weatherData.Weather[0].Description)
}
