package main

import (
	"donation-feed-backend/config"
	"github.com/Viva-con-Agua/vcago"
	"github.com/labstack/echo/v4"
	"strconv"
)

func main() {
	cfg := config.LoadFromEnv()
	e := echo.New()
	e.HTTPErrorHandler = vcago.HTTPErrorHandler
	e.Validator = vcago.JSONValidator
	e.Use(vcago.CORS.Init())
	e.Use(vcago.Logger.Init("donation-feed-backend"))

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(cfg.AppPort)))
}
