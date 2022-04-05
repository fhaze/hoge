package main

import (
	"fmt"
	"hoge/internal"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	prt := internal.GetEnvDefault("HOST_PORT", "8000")
	srv := internal.NewService()

	e.GET("/", func(c echo.Context) error {
		if msg, err := srv.Get(); err == nil {
			c.String(http.StatusOK, msg)
		}
		return c.String(http.StatusInternalServerError, "internal server error")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", prt)))
}
