package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"mongo/internal/models"
)

type BookRepository struct {
	db *mongo.Client
}

func NewBookRepository(db *mongo.Client) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	return nil, nil
}

func (r *BookRepository) GetById(id int) (models.Book, error) {
	return models.Book{}, nil
}

func (r *BookRepository) Create(book models.Book) (int, error) {
	return 0, nil
}

func (r *BookRepository) Update(updatedBook models.Book, id int) error {
	return nil
}

func (r *BookRepository) Delete(id int) error {
	return nil
}
