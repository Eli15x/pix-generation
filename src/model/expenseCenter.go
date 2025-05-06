package model

import (
	"time"
)

type ExpenseCenter struct {
	CentroExpenseID string    `bson:"CentroExpenseID" json:"centroExpense_id"`
	NomeCentro      string    `bson:"NomeCentro" json:"nome_centro"`
	CreatedAt       time.Time `bson:"CreatedAt" json:"created_at"`
	UpdatedAt       time.Time `bson:"UpdatedAt" json:"updated_at"`
}

type ExpenseCenterReceive struct {
	NomeCentro string `json:"nome_centro" example:"Administrativo"`
}

type ExpenseCenterDeleteRequest struct {
	CentroExpenseID string `json:"id" example:"647a8f9c0bde123456789abc"`
}
