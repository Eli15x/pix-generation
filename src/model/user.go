package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
	ClientID  string             `bson:"ClientID" json:"client_id"`
	Document  string             `bson:"Document" json:"document"`
	AcessID   string             `bson:"AcessID" json:"acess_id"`
	Name      string             `bson:"Name" json:"name"`
	Email     string             `bson:"Email" json:"email"`
	Password  string             `bson:"Password" json:"password"`
	CreatedAt time.Time          `bson:"CreatedAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"UpdatedAt" json:"UpdatedAt"`
}

type ResponseUser struct {
	UserID string `bson:"user_id,omitempty" json:"user_id,omitempty"`
	JWT    string `bson:"JWT,omitempty" json:"JWT,omitempty"`
}
