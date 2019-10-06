package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (a *Api) home(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, "Hello world! "+time.Now().Format("2006-01-02 15:04:05"))
}
