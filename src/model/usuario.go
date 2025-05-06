package model

type Usuario struct {
	UsuarioID string `bson:"UsuarioID" json:"usuario_id"`
	UserID    string `bson:"UserID" json:"user_id"`
	Nome      string `bson:"nome" json:"nome"`
	Email     string `bson:"email" json:"email"`
	Senha     string `bson:"senha" json:"senha"`
	Nivel     int    `bson:"nivel" json:"nivel"`
	Setor     string `bson:"setor" json:"setor"`
	Celular   string `bson:"celular" json:"celular"`
	Loja      string `bson:"loja" json:"loja"`
}

type UsuarioReceive struct {
	Nome    string `json:"nome" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Senha   string `json:"senha" binding:"required"`
	Nivel   int    `json:"nivel" binding:"required"`
	Setor   string `json:"setor" binding:"required"`
	Celular string `json:"celular" binding:"required"`
	Loja    string `json:"loja" binding:"required"`
}

type UsuarioDeleteRequest struct {
	UsuarioID string `json:"usuario_id" binding:"required"`
}

type UsuarioEmailRequest struct {
	Email string `json:"email" binding:"required"`
}
