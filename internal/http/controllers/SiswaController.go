package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctr *MainController) SiswaHomeView(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "home_siswa", nil)
}
