package models

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DefaultModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func InitDefaultModel() DefaultModel {
	now := time.Now()
	return DefaultModel{
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type Model struct {
	*gorm.DB
}

func NewModel() (*Model, error) {
	db, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Model{db}, nil
}
