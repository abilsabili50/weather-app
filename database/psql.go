package database

import (
	"log"

	"github.com/abilsabili50/weather-app/config"
	"github.com/abilsabili50/weather-app/entity/weather"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	dsn := config.GetPsqlConfig()
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err = db.AutoMigrate(weather.Weather{}); err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("database connected")
}

func GetDbInstance() *gorm.DB {
	return db
}
