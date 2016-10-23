package weather

import "testing"

func TestGetFormatedDescription(t *testing.T) {
	cases := []struct {
		in       cityPayload
		expected string
	}{
		{cityPayload{Weather: []weatherDescription{}}, "No description"},
		{cityPayload{Weather: []weatherDescription{weatherDescription{"-", "sun"}}}, "Sun"},
		{cityPayload{Weather: []weatherDescription{weatherDescription{"-", "raiN"}, weatherDescription{"-", "MilD"}}}, "Rain, mild"},
	}
	for _, c := range cases {
		got := getFormatedDescription(&c.in)
		if got != c.expected {
			t.Errorf("getFormatedDescription(%q) == %q, expected %q", c.in, got, c.expected)
		}
	}
}

func TestGetCityWeather(t *testing.T) {
	_, err := getCityWeather("Paris")
	if err != nil {
		t.Errorf("getCityWeather returned an error: %v", err)
	}
}
