package service

import (
	"fmt"
	"github.com/je09/spacebrew2/internal/entity"
	"github.com/je09/spacebrew2/internal/repository"
)

const (
	// postFormat describes post structure with fields: Title, Description, isOpen
	postFormat = "%s\n\n%s\n#%s"
)

type PostService struct {
	repo repository.Task
}

func NewPostService(repos repository.Task) *PostService {
	return &PostService{repo: repos}
}

func (p *PostService) New(task entity.Task) error {
	return p.repo.Add(task)
}

func (p *PostService) Text(task entity.Task) string {
	status := "open"
	if !task.Open {
		status = "closed"
	}

	return fmt.Sprintf(postFormat, task.Title, task.Description, status)
}

func (p *PostService) Edit(task entity.Task) error {
	return p.repo.Update(task)
}

func (p *PostService) EditDescription(id int64, text string) error {
	task, err := p.repo.Get(id)
	if err != nil {
		return err
	}

	task.Description = text
	return p.Edit(task)
}

func (p *PostService) EditTitle(id int64, text string) error {
	task, err := p.repo.Get(id)
	if err != nil {
		return err
	}

	task.Title = text
	return p.Edit(task)
}

func (p *PostService) Close(id int64) error {
	return p.repo.Close(id)
}

func (p *PostService) Pages(perPage int) (int64, error) {
	c, err := p.repo.Count()
	if err != nil {
		return 0, err
	}

	return c / int64(perPage), nil
}
