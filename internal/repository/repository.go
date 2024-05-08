package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	Library
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{

	}
}

type Library interface {
	GetAll()
}
