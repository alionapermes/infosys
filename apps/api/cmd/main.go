package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"infosys-api/internal/handler"
	"infosys-api/internal/handler/conn"
	"infosys-api/internal/handler/get_tracks"
)

func main() {
	defer func() {
		if handler.PgConn != nil {
			handler.PgConn.Close(context.Background())
		}
	}()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderAccept,
			echo.HeaderContentType,
			echo.HeaderOrigin,
		},
	}))

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	e.POST("/conn", conn.Handler)
	e.GET("/tracks", get_tracks.Handler)

	e.Logger.Fatal(e.Start(":8081"))
}
