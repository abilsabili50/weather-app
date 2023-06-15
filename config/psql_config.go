package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Dsn struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err.Error())
	}
}

func getEnviron() Dsn {
	return Dsn{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
}

func GetPsqlConfig() string {
	postgreConfig := getEnviron()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", postgreConfig.Host, postgreConfig.User, postgreConfig.Password, postgreConfig.Dbname, postgreConfig.Port)

	return dsn
}
