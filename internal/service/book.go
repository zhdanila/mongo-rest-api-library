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
	return s.repo.GetAll()
}

func (s *BookService) GetById(id string) (models.Book, error) {
	return s.repo.GetById(id)
}

func (s *BookService) Create(book models.Book) (string, error) {
	return s.repo.Create(book)
}

func (s *BookService) Update(updatedBook models.Book, id string) error {
	return s.repo.Update(updatedBook, id)
}

func (s *BookService) Delete(id string) error {
	return s.repo.Delete(id)
}
