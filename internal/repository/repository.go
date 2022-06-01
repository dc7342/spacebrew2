package repository

import "gorm.io/gorm"

type Repository struct {
	Task
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Task: NewTaskGORM(db),
	}
}
