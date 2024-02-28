package controllers

import (
	"github.com/DaffaAhmadSM/CBTapp/internal/http/models"
	"gorm.io/gorm"
	"net/http"
)

func (c *MainController) GetData(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	err := c.Database.First(&user, "id = ?", 1).Error
	if err == gorm.ErrRecordNotFound {
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = templ.Render(w, "index", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
