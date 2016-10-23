package weather

import (
	"errors"
	"fmt"

	"gopkg.in/h2non/gentleman.v1"
)

type weatherDescription struct {
	Main        string
	Description string
}
type cityPayload struct {
	Name string
	Sys  struct {
		Country string
	}
	Weather []weatherDescription
	Wind    map[string]float64
	Main    struct {
		Temp     float64
		Humidity float64
	}
}
type cityGroupPayload struct {
	List []cityPayload
}

var client = gentleman.New()

func getCityGroupWeather() (*cityGroupPayload, error) {
	req := client.Request().URL("api.openweathermap.org")
	req.Path("/data/2.5/group")

	var nantes, paloAlto, prague = 2990969, 5380748, 3067696
	req.SetQuery("id", fmt.Sprintf("%v,%v,%v", paloAlto, nantes, prague))
	req.SetQuery("units", "metric")
	req.SetQuery("appid", apiKey)

	res, err := req.Send()
	if err != nil {
		return nil, errors.New("Request error: " + err.Error())
	}
	if !res.Ok {
		return nil, errors.New(fmt.Sprintf("Invalid server response: %d\n", res.StatusCode))
	}

	weather := &cityGroupPayload{}
	err = res.JSON(weather)
	if err != nil {
		return nil, err
	}

	return weather, nil
}

// TODO: find a way to remove duplicated code
func getCityWeather(city string) (*cityPayload, error) {
	req := client.Request().URL("api.openweathermap.org")
	req.Path("/data/2.5/weather")

	req.SetQuery("q", city)
	req.SetQuery("units", "metric")
	req.SetQuery("appid", apiKey)

	res, err := req.Send()
	if err != nil {
		return nil, errors.New("Request error: " + err.Error())
	}
	if !res.Ok {
		return nil, errors.New(fmt.Sprintf("Invalid server response: %d\n", res.StatusCode))
	}

	weather := &cityPayload{}
	err = res.JSON(weather)
	if err != nil {
		return nil, err
	}

	return weather, nil
}
