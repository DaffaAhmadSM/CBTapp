package routers

import (
	"github.com/DaffaAhmadSM/CBTapp/internal/http/controllers"
	"github.com/DaffaAhmadSM/CBTapp/internal/http/helpers"
	"github.com/labstack/echo/v4"
)

func web(app *echo.Echo, ctr *controllers.MainController) {
	app.GET("/", ctr.LoginView)
	auth := app.Group("/home")
	auth.Use(helpers.WebIssueToken)
	auth.GET("/siswa", ctr.SiswaHomeView)
}
