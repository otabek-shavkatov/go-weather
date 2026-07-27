package main

import (
	"go-weather-homework/city"
	"go-weather-homework/weather"
)

func main() {

	cities := city.City{Name: "tashkent"}
	weather.FindCityLocation(cities)
}
