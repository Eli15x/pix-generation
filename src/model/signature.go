package model

import (
	"time"
)

type Signature struct {
	SignatureID    string    `bson:"SignatureID" json:"signature_id"`
	ClienteID      string    `bson:"ClienteID" json:"cliente_id"`
	DiaLancamento  int       `bson:"DiaLancamento" json:"dia_lancamento"`
	DiaVencimento  int       `bson:"DiaVencimento" json:"dia_vencimento"`
	QtdParcelas    int       `bson:"QtdParcelas" json:"qtd_parcelas"`
	CentroCustoID  string    `bson:"CentroCustoID" json:"centro_custo_id"`
	ValorOperacao  float64   `bson:"ValorOperacao" json:"valor_operacao"`
	EmitidoEsteMes bool      `bson:"EmitidoEsteMes" json:"emitido_este_mes"`
	VencidoEsteMes bool      `bson:"VencidoEsteMes" json:"vencido_este_mes"`
	CreatedAt      time.Time `bson:"CreatedAt" json:"created_at"`
	UpdatedAt      time.Time `bson:"UpdatedAt" json:"updated_at"`
}

type SignatureReceive struct {
	ClienteID     string  `json:"cliente_id" binding:"required"`
	DiaLancamento int     `json:"dia_lancamento" binding:"required"`
	DiaVencimento int     `json:"dia_vencimento" binding:"required"`
	QtdParcelas   int     `json:"qtd_parcelas" binding:"required"`
	CentroCustoID string  `json:"centro_custo_id" binding:"required"`
	ValorOperacao float64 `json:"valor_operacao" binding:"required"`
}

type SignatureDeleteRequest struct {
	SignatureID string `json:"id" binding:"required"`
}

type SignatureClienteRequest struct {
	ClienteID string `json:"cliente_id" binding:"required"`
}
