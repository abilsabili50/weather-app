package helpers

import (
	"time"

	"github.com/abilsabili50/weather-app/dto"
	"github.com/abilsabili50/weather-app/entity/weather"
)

func WeatherRequestToWeatherEntity(payload dto.WeatherRequest) weather.Weather {
	return weather.Weather{
		Wind:      payload.Wind,
		Water:     payload.Water,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
