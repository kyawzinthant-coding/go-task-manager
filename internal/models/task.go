package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	ID uint
	Title string
	Description string
	Deadline string
	Status string
}

func ConnectDB() (*gorm.DB, error) {
	db , err := gorm.Open(sqlite.Open("tasks.db"),&gorm.Config{})
	if err != nil {
		return nil,err
	}

	db.AutoMigrate(&Task{})
	return db,nil
}