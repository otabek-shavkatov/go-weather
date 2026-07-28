package main

import (
	"encoding/json"
	"fmt"
	"go-weather-homework/city"
	"go-weather-homework/weather"
)

func main() {

	cities := city.City{
		Name: []string{"Tashkent", "Navoiy"},
	}
	resultLocationCity := city.FindCityLocation(cities)

	var locations []city.Location
	var weatherCities []weather.Weather

	for location := range resultLocationCity {
		locations = append(locations, location)
	}

	fmt.Println(locations)

	for _, cityLocation := range locations {
		resultWeather := weather.GetCityWeather(cityLocation)
		var weatherCity weather.Weather
		fmt.Println("weatherCity ===== ", weatherCity)
		resultJsonWeather := json.NewDecoder(resultWeather.Body).Decode(&weatherCity)

		if resultJsonWeather != nil {
			fmt.Println("decode error")
			continue
		}

		weatherCities = append(weatherCities, weatherCity)

	}

	// fmt.Println(weatherCities)
}
