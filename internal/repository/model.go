package repository

import "gorm.io/gorm"

type TaskModel struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	Open        bool   `gorm:"column:open"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Author      int    `gorm:"column:author"`
}
