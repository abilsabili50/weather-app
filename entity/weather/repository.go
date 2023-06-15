package weather

import "github.com/abilsabili50/weather-app/helpers/errs"

type IWeatherRepo interface {
	CreateWeather(payload Weather) (*Weather, errs.IErrResponse)
	GetRecordCount() (*int, errs.IErrResponse)
	DeleteAllWeather() errs.IErrResponse
	GetLastWeather() (*Weather, errs.IErrResponse)
	UpdateWeather(payload Weather) (*Weather, errs.IErrResponse)
}
