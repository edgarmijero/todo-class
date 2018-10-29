package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func Healthz(c echo.Context) error {
	return c.String(http.StatusOK, "Hellow world")
}
