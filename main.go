package main

import (
	"encoding/json"
	"fmt"
	"go-weather-homework/city"
	"go-weather-homework/weather"
	"sync"
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

	var wg sync.WaitGroup
	weatherChannel := make(chan weather.Weather)

	for _, cityLocation := range locations {
		wg.Add(1)
		go func(cityLocation city.Location) {
			defer wg.Done()

			resultWeather := weather.GetCityWeather(cityLocation)
			var weatherCity weather.Weather

			fmt.Println("weatherCity ===== ", weatherCity)
			resultJsonWeather := json.NewDecoder(resultWeather.Body).Decode(&weatherCity)
			resultWeather.Body.Close()
			if resultJsonWeather != nil {
				fmt.Println("decode error")
				return
			}
			weatherChannel <- weatherCity

		}(cityLocation)

	}
	go func() {
		wg.Wait()
		close(weatherChannel)
	}()

	for weatherCity := range weatherChannel {
		weatherCities = append(weatherCities, weatherCity)
	}
	fmt.Println(weatherCities)
}
