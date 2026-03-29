package model

type adminUser struct {
	UserID   string `bson:"UserID,omitempty" json:"user_id,omitempty"`
	Name     string `bson:"Name" json:"name"`
	Email    string `bson:"Email" json:"email"`
	Password string `bson:"Password" json:"password"`
}

type AdminResponseUser struct {
	UserID string `bson:"user_id,omitempty" json:"user_id,omitempty"`
	JWT    string `bson:"JWT,omitempty" json:"JWT,omitempty"`
}

type AdminUserLoginRequest struct {
	Email    string `json:"email" example:"usuario@email.com"`
	Password string `json:"password" example:"minhasenha123"`
}
