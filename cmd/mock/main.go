package main

import (
	"fmt"
	"hoge/internal"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	srvHost := internal.GetEnvDefault("SERVER_HOST", "http://localhost:9000")
	prt := strings.Split(srvHost, ":")[2]

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "msg from mock"})
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", prt)))
}
