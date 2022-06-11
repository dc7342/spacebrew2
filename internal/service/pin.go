package service

import (
	"fmt"
	"github.com/je09/spacebrew2/internal/repository"
)

const (
	// pinFormat describes pin with fields: ID, Channel, Title
	pinFormat = "[%d](https://t.me/%s/%d) %s\n"
)

type PinService struct {
	repo repository.Task
}

func NewPin(repos repository.Task) *PinService {
	return &PinService{repo: repos}
}

func (p *PinService) Update(channel string, page int, perPage int) (string, error) {
	tasks, err := p.repo.GetByPage(page, perPage)
	if err != nil {
		return "", err
	}

	text := ""
	for _, t := range tasks {
		text += fmt.Sprintf(pinFormat, t.ID, channel, t.ID, t.Title)
	}

	return text, nil
}
