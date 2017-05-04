package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloHandler(c echo.Context) error {
	response := HelloResponse{
		Message: "Hello",
	}
	return c.JSON(http.StatusOK, response)
}
