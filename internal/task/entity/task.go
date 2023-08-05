package entity

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
