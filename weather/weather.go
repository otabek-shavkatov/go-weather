package weather

import (
	"encoding/json"
	"fmt"
	"go-weather-homework/city"
	"net/http"
)

type Weather struct {
}

type Location struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

func FindCityLocation(c city.City) {
	url := "https://geocode.maps.co/search?city=" + c.Name + "&api_key=6a6797bc4c7d4603243282kgsaf9cd9"

	response, err := http.Get(url)

	defer response.Body.Close()

	if err != nil {

		fmt.Println(response)
		return
	}

	var locations []Location

	resultJsonApi := json.NewDecoder(response.Body).Decode(&locations)

	if resultJsonApi != nil {
		fmt.Println("erro decode json")
		return
	}
	fmt.Println(locations)
}
