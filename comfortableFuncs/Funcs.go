package comfortablefuncs

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func BadRequest(c echo.Context, Message string) error {
	return c.JSON(http.StatusBadRequest, Responce{
		Status:  "Error",
		Message: Message,
	})
}

func OK(c echo.Context, Message string) error {
	return c.JSON(http.StatusOK, Responce{
		Status:  "Success",
		Message: Message,
	})
}