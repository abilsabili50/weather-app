package pg_repo

import (
	"github.com/abilsabili50/weather-app/entity/weather"
	"github.com/abilsabili50/weather-app/helpers/errs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type weatherRepo struct {
	db *gorm.DB
}

func NewWeatherRepo(db *gorm.DB) weather.IWeatherRepo {
	return &weatherRepo{db: db}
}

func (w *weatherRepo) CreateWeather(payload weather.Weather) (*weather.Weather, errs.IErrResponse) {
	if err := w.db.Clauses(clause.Returning{}).Create(&payload).Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &payload, nil
}

func (w *weatherRepo) GetRecordCount() (*int, errs.IErrResponse) {
	var count int

	if err := w.db.Raw("SELECT COUNT(*) FROM weathers").Scan(&count).Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &count, nil
}

func (w *weatherRepo) DeleteAllWeather() errs.IErrResponse {
	if err := w.db.Exec("delete from weathers").Error; err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (w *weatherRepo) GetLastWeather() (*weather.Weather, errs.IErrResponse) {
	var weather weather.Weather

	if err := w.db.Last(&weather).Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &weather, nil
}

func (w *weatherRepo) UpdateWeather(payload weather.Weather) (*weather.Weather, errs.IErrResponse) {

	var updatedWeather weather.Weather

	err := w.db.Model(&updatedWeather).Where("id = ?", payload.ID).Updates(weather.Weather{Wind: payload.Wind, WindStatus: payload.WindStatus, Water: payload.Water, WaterStatus: payload.WaterStatus, UpdatedAt: payload.UpdatedAt}).Error

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &updatedWeather, nil

}
