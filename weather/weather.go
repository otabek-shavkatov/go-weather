package weather

import (
	"fmt"
	"go-weather-homework/city"
)

type Weather struct {
	gradus float64
	city.City
	city.Location
}

func getCityWeather(l city.Location) {
	fmt.Println(l)
}
