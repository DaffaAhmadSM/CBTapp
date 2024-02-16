package controllers

import (
	"fmt"
	"net/http"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (db *MainController) GetData(w http.ResponseWriter, r *http.Request) {
	var user User
	err := db.Database.Raw("SELECT * FROM users WHERE id = ?", 2).Scan(&user).Error
	if err != nil {

	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<div style=\"color: red; font-size: 20px;\">Hello %s</div>", user.Name)
}
