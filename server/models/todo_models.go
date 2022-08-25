package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    Todo     string             `json:"todo,omitempty" validate:"required"`
}