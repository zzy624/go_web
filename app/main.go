package main

import (
	"mycode/go_web/app/api"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	log := logrus.New()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World !")
	})

	api.Routes(e)
	log.Info("Start http server:", "8116")

	if err := e.Start(":8116"); err != nil {
		log.Error("Start http server fail", err)
	}
}
