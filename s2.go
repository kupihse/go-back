package main

import (
	"github.com/labstack/echo"
)

func Asd() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return
	})
}
