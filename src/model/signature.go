package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Signature struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SignatureID    string             `bson:"signatureID" json:"signature_id"`
	ClienteID      string             `bson:"clienteID" json:"cliente_id"`
	DiaLancamento  int                `bson:"dia_lancamento" json:"dia_lancamento"`
	DiaVencimento  int                `bson:"dia_vencimento" json:"dia_vencimento"`
	QtdParcelas    int                `bson:"qdta_parcelas" json:"qtd_parcelas"`
	CentroCustoID  primitive.ObjectID `bson:"centroCusto" json:"centro_custo"`
	ValorOperacao  float64            `bson:"valorOperacao" json:"valor_operacao"`
	EmitidoEsteMes bool               `bson:"emitidoEsteMes" json:"emitido_este_mes"`
	VencidoEsteMes bool               `bson:"vencidoEsteMes" json:"vencido_este_mes"`
	CreatedAt      time.Time          `bson:"createdAt" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updated_at"`
}

type SignatureReceive struct {
	ClienteID     string  `json:"cliente_id" binding:"required"`
	DiaLancamento int     `json:"dia_lancamento" binding:"required"`
	DiaVencimento int     `json:"dia_vencimento" binding:"required"`
	QtdParcelas   int     `json:"qtd_parcelas" binding:"required"`
	CentroCustoID string  `json:"centro_custo" binding:"required"`
	ValorOperacao float64 `json:"valor_operacao" binding:"required"`
}

type SignatureDeleteRequest struct {
	ID string `json:"id" binding:"required"`
}

type SignatureClienteRequest struct {
	ClienteID string `json:"cliente_id" binding:"required"`
}
