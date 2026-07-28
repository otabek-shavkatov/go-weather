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

	var weatherCities []weather.Weather
	for _, cityLocation := range resultLocationCity {
		resultWeather := weather.GetCityWeather(cityLocation)
		var weatherCity weather.Weather
		fmt.Println("weatherCity ===== ", weatherCity)
		resultJsonWeather := json.NewDecoder(resultWeather.Body).Decode(&weatherCity)

		defer resultWeather.Body.Close()

		if resultJsonWeather != nil {
			fmt.Println("decode error")
			continue
		}

		weatherCities = append(weatherCities, weatherCity)

	}

	fmt.Println(weatherCities)
}
