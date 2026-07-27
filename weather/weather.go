package weather

import (
	"fmt"
	"go-weather-homework/city"
)

type Weather struct {
}

func FindCityLocation(c city.City) {
	api := "https://geocode.maps.co/search?city=" + c.Name + "api_key=6a6797bc4c7d4603243282kgsaf9cd9"

	fmt.Println("api", api)
}
