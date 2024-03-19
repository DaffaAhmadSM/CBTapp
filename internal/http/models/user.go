package models

import (
	"time"
)

type User struct {
	ID        uint       `json:"id" gorm:"primaryKey; autoIncrement:true; not null; unique"`
	Name      string     `json:"name"`
	Email     *string    `json:"email"`
	Password  string     `json:"password"`
	Image     *string    `json:"image"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserRole struct {
	ID        uint       `json:"id" gorm:"primaryKey; autoIncrement:true; not null; unique"`
	UserID    uint       `json:"user_id"`
	User      User       `gorm:"references:ID"`
	RoleID    uint       `json:"role_id"`
	Role      Role       `gorm:"references:ID"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Role struct {
	ID        uint       `json:"id" gorm:"primaryKey; autoIncrement:true; not null; unique"`
	Name      string     `json:"name"`
	CodeName  *string    `json:"code_name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
