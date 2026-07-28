package weather

import (
	"fmt"
	"go-weather-homework/city"
	"net/http"
)

type Weather struct {
	Name string `json:"name"`

	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
}

func GetCityWeather(l city.Location) *http.Response {
	fmt.Println(l)

	weatherUrl := "https://api.openweathermap.org/data/2.5/weather" +
		"?lat=" + l.Lat +
		"&lon=" + l.Lon +
		"&units=metric" +
		"&appid=8857272b48c13776ef7c9b3cb0f5e3bf"
	fmt.Println(weatherUrl)
	response, err := http.Get(weatherUrl)

	fmt.Println("response ====== ", response)

	if err != nil {
		return nil
	}
	return response
}
