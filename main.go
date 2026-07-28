package main

import (
	"go-weather-homework/city"
)

func main() {

	cities := city.City{
		Name: []string{"Tashkent", "Navoiy"},
	}
	city.FindCityLocation(cities)
}
