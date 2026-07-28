package city

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
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

func FindCityLocation(c City) <-chan Location {

	locationChan := make(chan Location)
	var wg sync.WaitGroup
	// gorountine anonymous - yani func ichida func ishlatish
	// nima uchun aynan buyerda gorountine ishlatilyapti, chunki manashuyerda for-loopda bir nechta malumotlar aylanyapti
	// va ular navbat-navbat bolib ishlaydi, bularni bir vaqtda ishaltish uchun esa gorountine ishltiladi

	for _, name := range c.Name {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			url := "https://geocode.maps.co/search?city=" + name + "&api_key=6a6797bc4c7d4603243282kgsaf9cd9"

			response, err := http.Get(url)

			if err != nil {
				return
			}
			defer response.Body.Close()

			var locations []Location

			resultJsonApi := json.NewDecoder(response.Body).Decode(&locations)

			if resultJsonApi != nil {
				fmt.Println("erro decode json")
			}
			for _, location := range locations {
				locationChan <- location
			}

		}(name)

	}
	go func() {
		wg.Wait()
		close(locationChan)
	}()
	return locationChan
}
