package service

import (
	"mongo/internal/models"
	"mongo/internal/repository"
)

type Service struct {
	Library
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Library: NewBookService(repos.Library),
	}
}

type Library interface {
	GetAll() ([]models.Book, error)
	GetById(id int) (models.Book, error)
	Create(book models.Book) (int, error)
	Update(updatedBook models.Book, id int) error
	Delete(id int) error
}
