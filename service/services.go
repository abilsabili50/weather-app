package service

import (
	"github.com/abilsabili50/weather-app/dto"
	"github.com/abilsabili50/weather-app/helpers/errs"
)

type IWeatherService interface {
	CreateWeather(dto.WeatherRequest) (dto.IResponse, errs.IErrResponse)
	ShowWeather() (dto.IResponse, errs.IErrResponse)
	UpdateLastWeather() (dto.IResponse, errs.IErrResponse)
}
