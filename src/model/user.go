package model

import (
	"time"
)

type User struct {
	UserID         string    `bson:"UserID,omitempty" json:"user_id,omitempty"`
	Document       string    `bson:"Document" json:"document"`
	Valid          bool      `bson:"Valid" json:"valid"`
	Name           string    `bson:"Name" json:"name"`
	Email          string    `bson:"Email" json:"email"`
	Password       string    `bson:"Password" json:"password"`
	TokenConfraPix string    `bson:"TokenConfraPix,omitempty" json:"token_confra_pix,omitempty"`
	CreatedAt      time.Time `bson:"CreatedAt" json:"createdAt"`
	UpdatedAt      time.Time `bson:"UpdatedAt" json:"UpdatedAt"`

	//TaxaTotal    	string             	`bson:"TaxaTotal" json:"taxaTotal"`
	//TaxaFaltante 	string             	`bson:"TaxaFaltante" json:"taxaFaltante"`
	//ClientID     	string            	`bson:"ClientID" json:"client_id"`
}

type ResponseUser struct {
	UserID string `bson:"user_id,omitempty" json:"user_id,omitempty"`
	JWT    string `bson:"JWT,omitempty" json:"JWT,omitempty"`
}

type UserLoginRequest struct {
	Email    string `json:"email" example:"usuario@email.com"`
	Password string `json:"password" example:"minhasenha123"`
}

type UserIDRequest struct {
	UserID string `json:"user_id" example:"647a8f9c0bde123456789abc"`
}

type UserDeleteRequest struct {
	Document string `json:"document" example:"12345678900"`
}
