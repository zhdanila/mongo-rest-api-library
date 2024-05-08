package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	ProductionYear int                `json:"production_year"`
}
