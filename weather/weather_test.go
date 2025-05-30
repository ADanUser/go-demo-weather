package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "London"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3
	result, _ := weather.GetWeather(geoData, format)
	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, Получение %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 155},
	{name: "Zero format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Moscow"
			geoData := geo.GeoData{
				City: expected,
			}
			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrWrongFormat {
				t.Errorf("Ожидалось %v, Получение %v", weather.ErrWrongFormat, err)
			}
		})
	}
}
