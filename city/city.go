package city

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type City struct {
	Name []string
}

type Location struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

func FindCityLocation(c City) []Location {

	var locations []Location

	for _, name := range c.Name {
		url := "https://geocode.maps.co/search?city=" + name + "&api_key=6a6797bc4c7d4603243282kgsaf9cd9"

		response, err := http.Get(url)

		defer response.Body.Close()

		if err != nil {
			break
		}

		var result []Location

		resultJsonApi := json.NewDecoder(response.Body).Decode(&result)

		if resultJsonApi != nil {
			fmt.Println("erro decode json")
			break
		}
		locations = append(locations, result...)
	}

	return locations

}
