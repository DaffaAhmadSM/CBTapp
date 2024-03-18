package routers

import (
	"github.com/DaffaAhmadSM/CBTapp/internal/http/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func api(app *echo.Echo, db *gorm.DB) {
	ctr := &controllers.MainController{Database: db}
	api := app.Group("/api")
	api.POST("/get-token", ctr.GetTokenSiswa)
}
