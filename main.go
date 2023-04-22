package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("./statics"))

	http.Handle("/statics/", http.StripPrefix("/statics/", fs))
	db := http.FileServer(http.Dir("./java"))

	http.Handle("/java/", http.StripPrefix("/java/", db))
	http.HandleFunc("/", HomeHandler)
	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Weather(w, r)
}

type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Dt int64 `json:"dt"`
}

func Weather(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	location, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		panic(err)
	}

	url := "http://api.openweathermap.org/data/2.5/weather?q=Paris&lang=fr&appid=24183512b38965eb7657303937bfbd9a&units=metric"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var weatherResponse WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)

	timestamp := weatherResponse.Dt
	temperature := weatherResponse.Main.Temp
	temperatureStr := fmt.Sprintf("%.0f", temperature)
	humidity := weatherResponse.Main.Humidity
	description := weatherResponse.Weather[0].Description
	currentTime := time.Unix(timestamp, 0)
	currentTimeInFrance := currentTime.In(location)
	data := struct {
		Temperature string
		Humidity    float64
		Description string
	}{
		Temperature: temperatureStr,
		Humidity:    humidity,
		Description: description,
	}

	fmt.Printf("Le temps actuel à Paris est %s avec une température de %s°C et une humidité de %.0f%% à %s\n", description, temperatureStr, humidity, currentTimeInFrance.Format(time.RFC1123))
	tmpl.Execute(w, data)
}
