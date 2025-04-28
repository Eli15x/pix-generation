package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Emissor struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmissorID      string             `bson:"emissorID,omitempty" json:"emissor_id,omitempty"`
	Nome           string             `bson:"nome" json:"nome"`
	Email          string             `bson:"email" json:"email"`
	Senha          string             `bson:"senha" json:"senha"`
	Documento      string             `bson:"documento" json:"documento"`
	TokenConfraPix string             `bson:"tokenConfraPix" json:"token_confra_pix"`
	CreatedAt      time.Time          `bson:"createdAt" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updated_at"`
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
