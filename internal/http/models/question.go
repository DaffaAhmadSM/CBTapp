package models

import (
	"gorm.io/datatypes"
	"time"
)

type BankQuestion struct {
	ID        uint       `json:"id" gorm:"primaryKey; autoIncrement:true; not null; unique"`
	Name      string     `json:"name"`
	CodeName  *string    `json:"code_name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Questions []Question `gorm:"foreignKey:BankQuestionID"`
}

type QuestionType struct {
	ID        uint       `json:"id" gorm:"primaryKey; autoIncrement:true; not null; unique"`
	Name      string     `json:"name"`
	CodeName  *string    `json:"code_name"`
	Questions []Question `gorm:"foreignKey:QuestionTypeID"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

type Question struct {
	ID             uint            `json:"id" gorm:"primaryKey; autoIncrement:true; not null; unique"`
	QuestionBody   *datatypes.JSON `json:"question_body"`
	BankQuestionID uint            `json:"bank_question_id"`
	BankQuestion   BankQuestion    `gorm:"references:ID"`
	QuestionTypeID uint            `json:"question_type_id"`
	QuestionType   QuestionType    `gorm:"references:ID"`
	CorrectAnswer  *string         `json:"answer"`
	Point          uint8           `json:"point"`
	CreatedAt      *time.Time      `json:"created_at"`
	UpdatedAt      *time.Time      `json:"updated_at"`
}
