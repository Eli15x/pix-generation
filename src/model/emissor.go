package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Emissor struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmissorID      string             `bson:"EmissorID,omitempty" json:"emissor_id,omitempty"`
	Nome           string             `bson:"Nome" json:"nome"`
	Email          string             `bson:"Email" json:"email"`
	Senha          string             `bson:"Senha" json:"senha"`
	Documento      string             `bson:"Documento" json:"documento"`
	TokenConfraPix string             `bson:"TokenConfraPix,omitempty" json:"token_confra_pix,omitempty"`
	CreatedAt      time.Time          `bson:"CreatedAt" json:"created_at"`
	UpdatedAt      time.Time          `bson:"UpdatedAt" json:"updated_at"`
}

type EmissorReceive struct {
	Nome           string `json:"nome" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Senha          string `json:"senha" binding:"required"`
	Documento      string `json:"documento" binding:"required"`
	TokenConfraPix string `json:"token_confra_pix" binding:"required"`
}

type EmissorDeleteRequest struct {
	ID string `json:"id" binding:"required"`
}

type EmissorDocumentoRequest struct {
	Documento string `json:"documento" binding:"required"`
}
