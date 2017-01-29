package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"github.com/taitai42/simple-cms-api/config"
	"github.com/taitai42/simple-cms-api/context"
	"github.com/taitai42/simple-cms-api/db"
	"github.com/taitai42/simple-cms-api/api"
)

func main() {
	config.Config()

	context.Init()
	defer db.GetDB().Close()
	e := echo.New()

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: func(key string, c echo.Context) bool {
			return config.HasToken(key)
		},
	}))


	e.GET("/home", api.GetHome)
	e.POST("/home", api.GetHome)
	e.PUT("/home", api.GetHome)

	log.Fatal(e.Start("0.0.0.0:" + viper.GetString("port")))
}
