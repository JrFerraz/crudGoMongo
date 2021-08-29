package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//el omitempty quiere decir que lo omita si el campo viene vacio
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id",omitempty`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}

type Users []*User
