package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Bootstrap struct {
	Database Database
	Logger   Logger
	Server   Server
	Jwt      JWTConfig
}

type Server struct {
	Name string
	Host string
	Port string
}

type JWTConfig struct {
	AccessTokenKey  string
	AccessTokenExp  string
	RefreshTokenKey string
	RefreshTokenExp string
}

type Database struct {
	Host                  string
	Port                  string
	Name                  string
	User                  string
	Pass                  string
	Tz                    string
	IdleConnection        string
	MaxConnection         string
	MaxLifeTimeConnection string
}

type Logger struct {
	Level string
}

func Get() *Bootstrap {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load configuration: ", err.Error())
	}

	return &Bootstrap{
		Server: Server{
			Name: os.Getenv("SERVER_NAME"),
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Logger: Logger{
			Level: os.Getenv("LOG_LEVEL"),
		},
		Database: Database{
			Host:                  os.Getenv("DB_HOST"),
			Port:                  os.Getenv("DB_PORT"),
			User:                  os.Getenv("DB_USER"),
			Pass:                  os.Getenv("DB_PASS"),
			Name:                  os.Getenv("DB_NAME"),
			Tz:                    os.Getenv("DB_TZ"),
			IdleConnection:        os.Getenv("DB_POOL_IDLE"),
			MaxConnection:         os.Getenv("DB_POOL_MAX"),
			MaxLifeTimeConnection: os.Getenv("DB_POOL_LIFETIME"),
		},
		Jwt: JWTConfig{
			AccessTokenKey:  os.Getenv("JWT_ACCESS_TOKEN_KEY"),
			AccessTokenExp:  os.Getenv("JWT_ACCESS_TOKEN_EXP"),
			RefreshTokenKey: os.Getenv("JWT_REFRESH_TOKEN_KEY"),
			RefreshTokenExp: os.Getenv("JWT_REFRESH_TOKEN_EXP"),
		},
	}
}
