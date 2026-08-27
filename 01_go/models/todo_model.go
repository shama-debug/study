package models

import "gorm.io/gorm"

// Todoテーブル
type Todo struct {
	gorm.Model
	UserID uint   `gorm:"index" json:"-"`
	Title  string `gorm:"not null" json:"title"`
	Status string `gorm:"default:'pending'" json:"status"`
	Memo   string `gorm:"default:'pending'" json:"memo"`
}
