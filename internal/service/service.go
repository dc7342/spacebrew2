package service

import "github.com/je09/spacebrew2/internal/repository"

type Service struct {
	Post
	Pin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Post: NewPostService(repos.Task),
		Pin:  NewPin(repos.Task),
	}
}
