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

	for _, name := range c.Name {
		wg.Add(1)

		// gorountine anonymous - yani func ichida func ishlatish
		// nima uchun aynan buyerda gorountine ishlatilyapti, chunki manashuyerda for-loopda bir nechta malumotlar aylanyapti
		// va ular navbat-navbat bolib ishlaydi, bularni bir vaqtda ishaltish uchun esa gorountine ishlatiladi
		// wg ga 1 qoshamiz yani bitta wg bitta ishni kutadi  degani
		go func(name string) {
			// buyerda go ichiga yani gorountine ichiga kirib bolganligi uchun uni done qilamiz yani bitta gorountine ishini boldi deganday ni -1 qilamiz
			// deferni nima uchun boshiga qoydik oxiriga qoysak ham ishalydi lekin orada resultJsonApi nil bolib qolsa wg abadiy kutadi shuning uchun wg Done bolishni boshdia reja qilib oqyadi orada nima bolsa ham oxirida Done qilib yubroadi

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

	// buyerda ochilgan kanalni yopamiz kanal qachon yopilishini bilamdyi shuning uchun qolda yopishimiz kerak boaldi
	// yani wg 0 bolgancha kutadi va close qialdi boldi hech qnadya wg kelmaydi chanell kelmaydi boldi yopaver deydi
	go func() {
		wg.Wait()
		close(locationChan)
	}()
	return locationChan
}
