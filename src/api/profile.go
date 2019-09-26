package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/trilobit/go-chat/src/models"
)

func (a *Api) profile(ctx echo.Context) error {
	user := ctx.Get("user").(*models.User)
	return ctx.JSON(http.StatusOK, user)
}
