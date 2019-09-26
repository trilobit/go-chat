package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (a *Api) home(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, "Hello world! "+time.Now().String())
}
