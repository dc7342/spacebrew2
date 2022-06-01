package repository

import "github.com/je09/spacebrew2/internal/entity"

type Task interface {
	GetAll() ([]entity.Task, error)
	Get(id int) (entity.Task, error)
	Close(id int) error
	Open(id int) error
}
