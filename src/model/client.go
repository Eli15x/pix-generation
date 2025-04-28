package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`            //ver qual o motivo disso aqui, ta tentando simular o valor do mongo?
	ClientID  string             `bson:"clientID,omitempty" json:"clientID,omitempty"` //ver qual o motivo disso aqui, ta tentando simular o valor do mongo?
	Nome      string             `bson:"nome" json:"nome"`
	CPF       string             `bson:"cpf" json:"cpf"`
	Email     string             `bson:"email" json:"email"`
	Celular   string             `bson:"celular" json:"celular"`
	CreatedAt time.Time          `bson:"createdAt" json:"created_at"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updated_at"`
}

type ClientReceive struct {
	Nome    string `json:"nome" binding:"required"`
	CPF     string `json:"cpf" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Celular string `json:"celular" binding:"required"`
}

type ClientDeleteRequest struct {
	ID string `json:"id" binding:"required"`
}

type ClientCpfRequest struct {
	CPF string `json:"cpf" binding:"required"`
}
