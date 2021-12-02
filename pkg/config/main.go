package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tahsinature/future-proof-gin/pkg/utilities"
)

var (
	DB    dbConfig
	App   appConfig
	JWT   jwtConfig
	Redis redisConfig
)

func Validate() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	DB = dbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBname:   os.Getenv("DB_NAME"),
	}

	App = appConfig{
		Port:         os.Getenv("PORT"),
		Environment:  os.Getenv("ENV"),
		SSL:          os.Getenv("SSL") == "TRUE",
		IsProduction: os.Getenv("ENV") == "PRODUCTION",
		APIVersion:   os.Getenv("API_VERSION"),
	}

	JWT = jwtConfig{
		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}

	Redis = redisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       os.Getenv("REDIS_DB"),
	}

	utilities.ValidateMultipleStruct(DB, App, JWT, Redis)
}
