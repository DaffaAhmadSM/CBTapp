package models

import "time"

type Student struct {
	ID        uint       `json:"id" gorm:"primaryKey; autoIncrement:true; not null; unique"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Group struct {
	ID        uint       `json:"id" gorm:"primaryKey; autoIncrement:true; not null; unique"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type StudentGroup struct {
	ID        uint       `json:"id" gorm:"primaryKey; autoIncrement:true; not null; unique"`
	StudentID uint       `json:"student_id"`
	Student   Student    `gorm:"references:ID"`
	GroupID   uint       `json:"group_id"`
	Group     Group      `gorm:"references:ID"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
type StudentAnswer struct {
	ID         uint       `json:"id" gorm:"primaryKey; autoIncrement:true; not null; unique"`
	StudentID  uint       `json:"student_id"`
	Student    Student    `gorm:"references:ID"`
	QuestionID uint       `json:"question_id"`
	Question   Question   `gorm:"references:ID"`
	Answer     string     `json:"answer"`
	IsCorrect  bool       `json:"is_correct"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
