package handler

import (
	"github.com/abilsabili50/weather-app/database"
	"github.com/abilsabili50/weather-app/handler/http_handler"
	"github.com/abilsabili50/weather-app/repository/pg_repo"
	"github.com/abilsabili50/weather-app/service"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	db := database.GetDbInstance()
	weatherRepo := pg_repo.NewWeatherRepo(db)
	weatherService := service.NewWeatherService(weatherRepo)
	weatherHandler := http_handler.NewWeatherHandler(weatherService)

	route := gin.Default()

	route.LoadHTMLGlob("templates/*")

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"project": "weather app with autoreload data",
		})
	})

	v1route := route.Group("/api/v1")

	weatherRoute := v1route.Group("/weather")
	{
		go weatherHandler.UpdateWeather()
		weatherRoute.POST("/", weatherHandler.CreateWeather)
		weatherRoute.GET("/", weatherHandler.ShowWeather)
	}

	route.Run(":8080")
}
