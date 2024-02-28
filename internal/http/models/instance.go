package models

import (
	"gorm.io/gorm"
	"time"
)

type Instance struct {
	gorm.Model
	ID        uint       `json:"id" gorm:"primaryKey; autoIncrement:true; not null; unique"`
	Name      string     `json:"nama"`
	Address   *string    `json:"alamat"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
