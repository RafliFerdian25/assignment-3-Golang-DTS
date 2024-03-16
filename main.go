package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/robfig/cron/v3"
)

type Weather struct {
	Status Status
}

type Status struct {
	Water uint
	Wind  uint
}

func main() {
	weather := Weather{
		Status{
			Water: 0,
			Wind:  0,
		},
	}

	c := cron.New()
	c.AddFunc("@every 5s", func() {
		getWeather(&weather)

		JsonData, err := json.Marshal(weather)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		fmt.Println("JSON Data:")
		fmt.Println(string(JsonData))

		fmt.Println("Cuaca Status:")
		printWaterStatus(weather.Status.Water)
		printWindStatus(weather.Status.Wind)
		fmt.Println()
	})

	c.Start()

	select {}
}

func getWeather(weather *Weather) *Weather {
	weather.Status.Water = uint(rand.Intn(100) + 1)
	weather.Status.Wind = uint(rand.Intn(100) + 1)

	return weather
}

func printWaterStatus(water uint) {
	switch {
	case water <= 5:
		fmt.Println("Status air : Aman")
	case water <= 8:
		fmt.Println("Status air : Siaga")
	default:
		fmt.Println("Status air : Bahaya")
	}
}

func printWindStatus(wind uint) {
	switch {
	case wind <= 6:
		fmt.Println("Status angin : Aman")
	case wind <= 15:
		fmt.Println("Status angin : Siaga")
	default:
		fmt.Println("Status angin : Bahaya")
	}
}
