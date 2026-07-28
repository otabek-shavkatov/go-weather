package main

import (
	"encoding/json"
	"fmt"
	"go-weather-homework/city"
	"go-weather-homework/weather"
	"sync"
)

func main() {

	var cityCount int

	// terminalda sorashlik uchun
	fmt.Print("nechta shaxar qidirmoqchisiz ?)")
	// scanni vazifasi kiritlganm sonni cityCount ga yozib beradi
	fmt.Scan(&cityCount)

	var cityNames []string
	// for loop da kiritlagn shaxar sonichalik olamiz
	for i := 0; i < cityCount; i++ {
		var name string

		//  sanoq 0 dan boshalandi shuning uchun i boyciha emas +1 qilib yuramiz terminalda korinishi joyiga
		fmt.Printf("%d-shahar: ", i+1)
		// kiritilgan name ni name ozgaruvchisiga yozibb beradi
		fmt.Scan(&name)

		cityNames = append(cityNames, name)
	}

	// City Struct string array qabul qiladi shuning uchun ciytName ni beruib yubrioamiz
	cities := city.City{
		Name: cityNames,
	}

	// buyerda kanaldan kelyapti
	resultLocationCity := city.FindCityLocation(cities)

	var locations []city.Location
	var weatherCities []weather.Weather
	// kanaldagilarni bittalab arrayga yigamiz
	for location := range resultLocationCity {
		locations = append(locations, location)
	}

	fmt.Println(locations)

	var wg sync.WaitGroup
	weatherChannel := make(chan weather.Weather)
	// weather api ni ham paralel ishaltish uchun gorountine ishlatamiz buyerda ham chunki bizda shaxar nomlari paralel kelgani bilan weatherga bittalab bir birini kutib boryapti, shuning uchun ularga ham kanal ochamiz
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
