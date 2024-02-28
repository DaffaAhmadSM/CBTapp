package routers

import (
	"gorm.io/gorm"
	"net/http"
)

func RouterInit(db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()
	api(mux, db)
	return mux

}
