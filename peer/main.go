package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type ResponseBody struct {
	Message string
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &ResponseBody{Message: "Hello World"})
	})
	e.Logger.Fatal(e.Start(":9000"))
}
