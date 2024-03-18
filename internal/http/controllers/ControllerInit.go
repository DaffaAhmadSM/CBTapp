package controllers

import (
	"gorm.io/gorm"
)

// Template Parse HTML templates

type MainController struct {
	Database *gorm.DB
}
