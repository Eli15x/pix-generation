package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseCenter struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CentroExpenseID string             `bson:"centroExpenseID" json:"centroExpense_id"`
	NomeCentro      string             `bson:"nomeCentro" json:"nome_centro"`
	CreatedAt       time.Time          `bson:"createdAt" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updatedAt" json:"updated_at"`
}

type ExpenseCenterReceive struct {
	NomeCentro string `json:"nome_centro" example:"Administrativo"`
}

type ExpenseCenterDeleteRequest struct {
	ID string `json:"id" example:"647a8f9c0bde123456789abc"`
}
