package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *BookRepository) GetById(id string) (models.Book, error) {
	return models.Book{}, nil
}

func (r *BookRepository) Create(book models.Book) (string, error) {
	collection := r.db.Database("rest").Collection("library")

	insertResult, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		fmt.Println("Error inserting document:", err)
		return "", err
	}

	insertedID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("inserted ID is not an ObjectID")
	}

	return insertedID.Hex(), nil
}

func (r *BookRepository) Update(updatedBook models.Book, id string) error {
	return nil
}

func (r *BookRepository) Delete(id string) error {
	return nil
}
