package user

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetName(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Mr Zhang")
}
