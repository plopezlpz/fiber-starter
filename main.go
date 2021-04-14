package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/plopezlpz/fiber-starter/config/vars"
	"github.com/plopezlpz/fiber-starter/rest"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("loading .env")
	}
	vars.Setup()
}

func main() {
	app := rest.Setup()
	if err := app.Listen(fmt.Sprintf(":%v", vars.AppSetting.HttpPort)); err != nil {
		log.Fatalf("failed to start app %v", err)
	}
}
