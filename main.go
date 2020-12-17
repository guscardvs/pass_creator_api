package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		passPhrase := GetPassPhrase()
		return c.JSON(http.StatusOK, passPhrase)
	})
	e.Logger.Fatal(e.Start(":3001"))
}
