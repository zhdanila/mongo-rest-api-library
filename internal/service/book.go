package service

import (
	"mongo/internal/models"
	"mongo/internal/repository"
)

type BookService struct {
	repo repository.Library
}

func NewBookService(repo repository.Library) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAll() ([]models.Book, error) {
	return nil, nil
}

func (s *BookService) GetById(id string) (models.Book, error) {
	return models.Book{}, nil
}

func (s *BookService) Create(book models.Book) (string, error) {
	return s.repo.Create(book)
}

func (s *BookService) Update(updatedBook models.Book, id string) error {
	return nil
}

func (s *BookService) Delete(id string) error {
	return nil
}
