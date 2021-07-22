package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Note model
type Note struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
}
