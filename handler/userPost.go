package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (apiConfig *DatabaseController) UserPost(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
