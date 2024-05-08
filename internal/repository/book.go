package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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
	var books []models.Book

	collection := r.db.Database("rest").Collection("library")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) GetById(id string) (models.Book, error) {
	var book models.Book

	collection := r.db.Database("rest").Collection("library")
	filter := bson.M{"name": id}

	err := collection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
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

func (r *BookRepository) Update(updatedBook models.Book, name string) error {
	collection := r.db.Database("rest").Collection("library")
	filter := bson.M{"name": name}

	update := bson.M{
		"$set": bson.M{
			"name":           updatedBook.Name,
			"description":    updatedBook.Description,
			"production_year": updatedBook.ProductionYear,
		},
	}
	fmt.Println(update)
	fmt.Println(filter)
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) Delete(id string) error {
	collection := r.db.Database("rest").Collection("library")
	filter := bson.M{"name": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
