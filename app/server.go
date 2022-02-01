package app

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func NewServer(router *echo.Echo) *http.Server {
	return &http.Server{Addr: os.Getenv("PORT"), Handler: router}
}
