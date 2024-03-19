package routers

import (
	"github.com/DaffaAhmadSM/CBTapp/internal/http/controllers"
	"github.com/labstack/echo/v4"
)

func api(app *echo.Echo, ctr *controllers.MainController) {

	api := app.Group("/api")
	api.POST("/get-token", ctr.GetTokenSiswa)
}
