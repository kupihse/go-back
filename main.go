package main

import (
	"github.com/labstack/echo"

	"github.com/satori/go.uuid"
)

var port = ":8080"

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var productStorage = make(map[string]*Product)

func main() {
	router := echo.New()
	productGroup := router.Group("/pr")
	productGroup.POST("/new", func(c echo.Context) error {
		var p Product
		c.Bind(&p)
		p.ID = uuid.NewV4().String()
		productStorage[p.ID] = &p
		return c.NoContent(200)
	})
	productGroup.GET("/all", func(c echo.Context) error {
		return c.JSON(200, productStorage)
	})
	router.Start(port)
}
