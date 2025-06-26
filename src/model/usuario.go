package model

type Usuario struct {
	UsuarioID string `bson:"UsuarioID" json:"usuario_id"`
	UserID    string `bson:"UserID" json:"user_id"`
	Nome      string `bson:"Nome" json:"nome"`
	Email     string `bson:"Email" json:"email"`
	Senha     string `bson:"Senha" json:"senha"`
	Nivel     int    `bson:"Nivel" json:"nivel"`
	Setor     string `bson:"Setor" json:"setor"`
	Celular   string `bson:"Celular" json:"celular"`
	Loja      string `bson:"Loja" json:"loja"`
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
