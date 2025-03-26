package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
	ClientID  string             `bson:"client_id" json:"client_id"`
	Document  string             `bson:"document" json:"document"`
	AcessID   string             `bson:"acess_id" json:"acess_id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"UpdatedAt" json:"UpdatedAt"`
}

type ResponseUser struct {
	UserID string `bson:"user_id,omitempty" json:"user_id,omitempty"`
	JWT    string `bson:"JWT,omitempty" json:"JWT,omitempty"`
}
