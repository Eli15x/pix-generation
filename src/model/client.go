package model

import (
	"time"
)

type Client struct {
	ClientID  string    `bson:"ClientID,omitempty" json:"client_id,omitempty"`
	UserID    string    `bson:"UserID" json:"user_id"`
	Nome      string    `bson:"nome" json:"nome"`
	CPF       string    `bson:"cpf" json:"cpf"`
	Email     string    `bson:"email" json:"email"`
	Celular   string    `bson:"celular" json:"celular"`
	CreatedAt time.Time `bson:"createdAt" json:"created_at"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updated_at"`
}

type ClientReceive struct {
	UserID  string `bson:"UserID" json:"user_id"`
	Nome    string `json:"nome" binding:"required"`
	CPF     string `json:"cpf" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Celular string `json:"celular" binding:"required"`
}

type ClientRequest struct {
	ID string `json:"id" binding:"required"`
}

type ClientUserRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type ClientCpfRequest struct {
	CPF string `json:"cpf" binding:"required"`
}
