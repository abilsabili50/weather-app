package http_handler

import (
	"fmt"
	"log"
	"time"

	"github.com/abilsabili50/weather-app/dto"
	"github.com/abilsabili50/weather-app/helpers/errs"
	"github.com/abilsabili50/weather-app/service"
	"github.com/gin-gonic/gin"
)

type IWeatherHandler interface {
	CreateWeather(ctx *gin.Context)
	ShowWeather(ctx *gin.Context)
	UpdateWeather()
}

type weatherHandler struct {
	weatherService service.IWeatherService
}

func NewWeatherHandler(weatherService service.IWeatherService) IWeatherHandler {
	return &weatherHandler{weatherService: weatherService}
}

func (w *weatherHandler) CreateWeather(ctx *gin.Context) {
	var requestBody dto.WeatherRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErrBind := errs.NewBadRequestError("invalid body request")
		ctx.AbortWithStatusJSON(newErrBind.GetCode(), newErrBind)
		return
	}

	response, err := w.weatherService.CreateWeather(requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(err.GetCode(), err)
		return
	}

	ctx.JSON(response.GetCode(), response)
}

func (w *weatherHandler) ShowWeather(ctx *gin.Context) {
	response, err := w.weatherService.ShowWeather()
	if err != nil {
		ctx.AbortWithStatusJSON(err.GetCode(), err)
		return
	}

	ctx.HTML(response.GetCode(), "weather.html", response.GetData())
}

func (w *weatherHandler) UpdateWeather() {
	for {
		response, err := w.weatherService.UpdateLastWeather()
		if err != nil {
			log.Fatalln(err.Error())
		}

		weather := response.GetData()

		fmt.Printf("{water: %d, wind: %d}\n", weather.Water, weather.Wind)
		fmt.Printf("status water : %s\n", weather.WaterStatus)
		fmt.Printf("wtatus wind : %s\n\n", weather.WindStatus)

		time.Sleep(time.Second * 5)
	}
}
