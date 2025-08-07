package model

import (
	"time"
)

type Client struct {
	ClientID  string    `bson:"ClientID,omitempty" json:"client_id,omitempty"`
	UserID    string    `bson:"UserID" json:"user_id"`
	Nome      string    `bson:"Nome" json:"nome"`
	CPF       string    `bson:"CPF" json:"cpf"`
	Email     string    `bson:"Email" json:"email"`
	Celular   string    `bson:"Celular" json:"celular"`
	Rua       string    `bson:"Rua" json:"rua"`
	Cidade    string    `bson:"Cidade" json:"cidade"`
	CEP       string    `bson:"CEP" json:"cep"`
	UF        string    `bson:"UF" json:"uf"`
	CreatedAt time.Time `bson:"CreatedAt" json:"created_at"`
	UpdatedAt time.Time `bson:"UpdatedAt" json:"updated_at"`
}

type ClientReceive struct {
	UserID  string `bson:"UserID" json:"user_id"`
	Nome    string `json:"nome" binding:"required"`
	CPF     string `json:"cpf" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Celular string `json:"celular" binding:"required"`
	Rua     string `json:"rua"`
	Cidade  string `json:"cidade"`
	CEP     string `json:"cep"`
	UF      string `json:"uf"`
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

type ClientCidadeRequest struct {
	Cidade string `json:"cidade" binding:"required"`
}

type ClientUFRequest struct {
	UF string `json:"uf" binding:"required"`
}
