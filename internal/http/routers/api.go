package routers

import (
	"github.com/DaffaAhmadSM/CBTapp/internal/http/controllers"
	"gorm.io/gorm"
	"net/http"
)

func api(mux *http.ServeMux, db *gorm.DB) {
	ctr := controllers.MainController{Database: db}
	mux.HandleFunc("GET /", ctr.GetData)
}
