package service

import "github.com/je09/spacebrew2/internal/entity"

type Post interface {
	New(task entity.Task) error
	Edit(task entity.Task) error
	EditTitle(id int64, text string) error
	EditDescription(id int64, text string) error
	Close(id int64) error
	Text(task entity.Task) string
	Pages(perPage int) (int64, error)
}

type Pin interface {
	Update(channel string, page int, perPage int) (string, error)
}
