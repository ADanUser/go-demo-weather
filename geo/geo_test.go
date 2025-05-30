package geo_test

import (
	"demo/weather/geo"
	"errors"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	//Arrange - подготовка, expected результат, данные для функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}
	// Act - действие, вызов функции
	got, err := geo.GetMyLocation(city)
	// Assert - проверка, что полученный результат совпадает с ожидаемым
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, Получение %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "London123"
	_, err := geo.GetMyLocation(city)
	if !errors.Is(err, geo.ErrorNoCity) {
		t.Errorf("Ожидалось %v, Получение %v", geo.ErrorNoCity, err)
	}
}
