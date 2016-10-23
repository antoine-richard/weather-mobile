package weather

import (
	"encoding/json"
	"errors"
	"fmt"

	"strings"
)

type CityWeather struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Temp string `json:"temp"`
}

func FetchDefaultCities() ([]byte, error) {
	apiPayload, err := getCityGroupWeather()
	if err != nil {
		return nil, err
	}
	citiesWeather := cityGroupToApp(apiPayload)
	data, err := json.Marshal(citiesWeather)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FetchCustomCity(city string) ([]byte, error) {
	city = strings.TrimSpace(city)
	if city == "" {
		return nil, errors.New("Please provide a city name")
	}
	apiPayload, err := getCityWeather(city)
	if err != nil {
		return nil, err
	}
	cityWeather := citytoApp(apiPayload)
	data, err := json.Marshal(cityWeather)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func cityGroupToApp(apiPayload *cityGroupPayload) (res []CityWeather) {
	for _, v := range apiPayload.List {
		res = append(res, *citytoApp(&v))
	}
	return
}

func citytoApp(apiPayload *cityPayload) (res *CityWeather) {
	return &CityWeather{
		Name: getFormatedName(apiPayload),
		Desc: getFormatedDescription(apiPayload),
		Temp: getFormatedTemperature(apiPayload),
	}
}

func getFormatedName(cityWeather *cityPayload) string {
	return fmt.Sprintf("%v, %v", cityWeather.Name, cityWeather.Sys.Country)
}

func getFormatedDescription(cityWeather *cityPayload) (fullDescription string) {
	descriptions := cityWeather.Weather
	if len(descriptions) == 0 {
		return "No description"
	}
	for i, desc := range descriptions {
		if i == 0 {
			fullDescription += strings.ToUpper(string(desc.Description[0])) + strings.ToLower(desc.Description[1:])
		} else {
			fullDescription += ", " + strings.ToLower(desc.Description)
		}
	}
	return fullDescription
}

func getFormatedTemperature(cityWeather *cityPayload) string {
	return fmt.Sprintf("%.1f°C", cityWeather.Main.Temp)
}
