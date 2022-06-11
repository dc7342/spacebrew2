package repository

import "gorm.io/gorm"

type TaskModel struct {
	gorm.Model
	ID          int64  `gorm:"primaryKey"`
	Open        bool   `gorm:"column:open"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	AuthorID    int64  `gorm:"column:author"`
}
