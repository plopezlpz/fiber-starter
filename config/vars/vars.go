// Package vars is used to hold global variables read from .env
package vars

import (
	"os"
	"strconv"
	"time"
)

// Global variables accessible by the application
var (
	LoggerSetting   logger
	AppSetting      app
	DatabaseSetting database
)

type logger struct {
	Name  string
	Level string
}

type app struct {
	HttpPort     int
	Name         string
	Mode         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

type database struct {
	Host string
	Port int
	Name string
	User string
	Pass string
}

func Setup() {
	dbPort, _ := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
	DatabaseSetting = database{
		Host: os.Getenv("DB_HOST"),
		Port: int(dbPort),
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
	}
	appPort, _ := strconv.ParseInt(os.Getenv("APP_PORT"), 10, 32)
	AppSetting = app{
		HttpPort:     int(appPort),
		Name:         os.Getenv("SERVICE_NAME"),
		Mode:         os.Getenv("SERVICE_MODE"),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	LoggerSetting = logger{
		Name:  os.Getenv("LOGGER_NAME"),
		Level: os.Getenv("LOG_LEVEL"),
	}
}
