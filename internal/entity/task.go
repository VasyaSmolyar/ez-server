package entity

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID     string `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	UserID int
	User   *User `json:"user"`
}

type User struct {
	gorm.Model
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
