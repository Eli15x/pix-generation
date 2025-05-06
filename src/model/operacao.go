package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operacao struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OperacaoID string             `bson:"OperacaoID" json:"operacao_id"`
	Nome       string             `bson:"nome" json:"nome"`
}

type OperacaoReceive struct {
	Nome string `json:"nome" binding:"required"`
}

type OperacaoDeleteRequest struct {
	OperacaoID string `json:"operacao_id" binding:"required"`
}
