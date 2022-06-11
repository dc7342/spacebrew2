package repository

import "github.com/je09/spacebrew2/internal/entity"

type Task interface {
	GetAll() ([]entity.Task, error)
	GetByPage(page int, perPage int) ([]entity.Task, error)
	Count() (int64, error)
	Get(id int64) (entity.Task, error)
	Add(task entity.Task) error
	Update(task entity.Task) error
	Close(id int64) error
	Open(id int64) error
}
