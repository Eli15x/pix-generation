package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	InvoiceID   string             `bson:"invoiceID,omitempty" json:"invoiceID,omitempty"`
	CnpjCliente string             `bson:"cnpjCliente,omitempty" json:"cnpjCliente,omitempty"`
	IdTransacao int                `bson:"idTransacao" json:"idTransacao"`
	Uuid        string             `bson:"uuid" json:"uuid"`
	Amount      float32            `bson:"amount" json:"amount"`
	Emitido     time.Time          `bson:"emitido" json:"emitido"`
	Expira      string             `bson:"expira" json:"expira"`
	Pago        string             `bson:"pago" json:"pago"`
	TxId        string             `bson:"txId" json:"txId"`
	TaxaPaga    bool               `bson:"taxaPaga" json:"taxaPaga"`
}

type InvoiceReceive struct {
	CnpjCliente string  `bson:"cnpjCliente,omitempty" json:"cnpjCliente,omitempty"`
	IdTransacao int     `bson:"idTransacao" json:"idTransacao"`
	Uuid        string  `bson:"uuid" json:"uuid"`
	Amount      float32 `bson:"amount" json:"amount"`
	Emitido     string  `bson:"emitido" json:"emitido"`
	Expira      string  `bson:"expira" json:"expira"`
	Pago        string  `bson:"pago" json:"pago"`
	TxId        string  `bson:"txId" json:"txId"`
	TaxaPaga    bool    `bson:"taxaPaga" json:"taxaPaga"`
}

type InvoiceIDRequest struct {
	InvoiceID string `json:"invoice_id" example:"INV-123456"`
}

type InvoiceCNPJRequest struct {
	CnpjCliente string `json:"cnpj_cliente" example:"12345678000199"`
}

type InvoiceDeleteRequest struct {
	InvoiceID string `bson:"invoiceID,omitempty" json:"invoiceID,omitempty"`
}
