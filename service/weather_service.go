package service

import (
	"math/rand"
	"time"

	"github.com/abilsabili50/weather-app/dto"
	"github.com/abilsabili50/weather-app/entity/weather"
	"github.com/abilsabili50/weather-app/helpers"
	"github.com/abilsabili50/weather-app/helpers/errs"
)

type weatherService struct {
	weatherRepo weather.IWeatherRepo
}

func NewWeatherService(weatherRepo weather.IWeatherRepo) IWeatherService {
	return &weatherService{weatherRepo: weatherRepo}
}

func (w *weatherService) CreateWeather(payload dto.WeatherRequest) (dto.IResponse, errs.IErrResponse) {
	rowCount, err := w.weatherRepo.GetRecordCount()
	if err != nil {
		return nil, err
	}

	if *rowCount > 0 {
		if err := w.weatherRepo.DeleteAllWeather(); err != nil {
			return nil, err
		}
	}

	weather := helpers.WeatherRequestToWeatherEntity(payload)
	weather.GenerateStatus()

	createdWeather, err := w.weatherRepo.CreateWeather(weather)
	if err != nil {
		return nil, err
	}

	weatherResponse := dto.WeatherResponse{
		ID:          createdWeather.ID,
		Water:       createdWeather.Water,
		Wind:        createdWeather.Wind,
		WaterStatus: createdWeather.WaterStatus,
		WindStatus:  createdWeather.WindStatus,
		CreatedAt:   &createdWeather.CreatedAt,
	}

	response := dto.NewCreatedResponse(false, "weather created successfully", weatherResponse)

	return response, nil
}

func (w *weatherService) ShowWeather() (dto.IResponse, errs.IErrResponse) {
	getWeather, err := w.weatherRepo.GetLastWeather()
	if err != nil {
		return nil, err
	}

	weatherResponse := dto.WeatherResponse{
		ID:          getWeather.ID,
		Water:       getWeather.Water,
		Wind:        getWeather.Wind,
		WaterStatus: getWeather.WaterStatus,
		WindStatus:  getWeather.WindStatus,
		CreatedAt:   &getWeather.CreatedAt,
		UpdatedAt:   &getWeather.UpdatedAt,
		Time:        time.Now().Format("15:04"),
		Degree:      rand.Intn(35),
	}

	response := dto.NewOKResponse(false, "weather found", weatherResponse)

	return response, nil
}

func (w *weatherService) UpdateLastWeather() (dto.IResponse, errs.IErrResponse) {

	water := rand.Intn(100)
	wind := rand.Intn(100)

	lastWeather, err := w.weatherRepo.GetLastWeather()
	if err != nil {
		return nil, err
	}

	newWeather := weather.Weather{
		ID:          lastWeather.ID,
		Water:       lastWeather.Water,
		Wind:        lastWeather.Wind,
		WaterStatus: lastWeather.WaterStatus,
		WindStatus:  lastWeather.WindStatus,
		CreatedAt:   lastWeather.CreatedAt,
		UpdatedAt:   lastWeather.UpdatedAt,
	}

	newWeather.UpdateWeather(water, wind)
	newWeather.GenerateStatus()

	updatedWeather, err := w.weatherRepo.UpdateWeather(newWeather)
	if err != nil {
		return nil, err
	}

	weatherResponse := dto.WeatherResponse{
		ID:          updatedWeather.ID,
		Water:       updatedWeather.Water,
		Wind:        updatedWeather.Wind,
		WaterStatus: updatedWeather.WaterStatus,
		WindStatus:  updatedWeather.WindStatus,
		CreatedAt:   &updatedWeather.CreatedAt,
		UpdatedAt:   &updatedWeather.UpdatedAt,
	}

	response := dto.NewOKResponse(false, "weather updated successfully", weatherResponse)

	return response, nil
}
