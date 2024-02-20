package main

import (
	"github.com/JohnKucharsky/go_echo_sqlc/db"
	"github.com/JohnKucharsky/go_echo_sqlc/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

func routes(e *echo.Echo, dbConnectionString string) {
	database := db.DatabaseConnection(dbConnectionString)
	h := handler.DatabaseController{Database: database}

	e.GET(
		"/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		},
	)
	e.POST("/users", h.UserPost)
}
