package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Post struct model
type Post struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title,omitempty"`
	Body  string             `json:"body,omitempty"`
}
