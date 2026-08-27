package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string `gorm:"not null" json:"title"`
	Status string `gorm:"default:'pending'" json:"status"`
}
