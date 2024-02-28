package main

import (
	"fmt"
	"github.com/DaffaAhmadSM/CBTapp/internal/http/database"
	"github.com/DaffaAhmadSM/CBTapp/internal/http/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db := database.PostgresInit()
	err = db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.UserRole{},
		&models.Instance{},
		&models.Question{},
		&models.QuestionType{},
		&models.BankQuestion{},
		&models.Student{},
		&models.Group{},
		&models.StudentGroup{},
		&models.StudentAnswer{})
	if err != nil {
		return
	}

	fmt.Printf("Auto migration success\n")
}
