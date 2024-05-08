package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"mongo/internal/models"
)

type Repository struct {
	Library
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Library: NewBookRepository(db),
	}
}

type Library interface {
	GetAll() ([]models.Book, error)
	GetById(id string) (models.Book, error)
	Create(book models.Book) (string, error)
	Update(updatedBook models.Book, id string) error
	Delete(id string) error
}
