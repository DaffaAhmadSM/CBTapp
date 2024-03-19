package routers

import (
	"github.com/DaffaAhmadSM/CBTapp/internal/http/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouterInit(db *gorm.DB) *echo.Echo {
	mux := echo.New()
	ctr := &controllers.MainController{Db: db}

	web(mux, ctr)
	api(mux, ctr)

	return mux
}
