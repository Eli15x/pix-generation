package service

import (
	"context"
	"errors"
	"pix-generation/src/client"
	"pix-generation/src/model"
	"pix-generation/src/repository"
	"pix-generation/src/utils"
	"sync"

	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceServiceUsuario ServiceUsuario
	onceServiceUsuario     sync.Once
)

type ServiceUsuario interface {
	CreateUsuario(ctx context.Context, usuario model.UsuarioReceive) error
	GetUsuarioByID(ctx context.Context, id string) (model.Usuario, error)
	GetAllUsuario(ctx context.Context) ([]model.Usuario, error)
	UpdateUsuario(ctx context.Context, id string, usuario model.UsuarioReceive) error
	DeleteUsuario(ctx context.Context, id string) error
	GetUsuarioByEmail(ctx context.Context, email string) (model.Usuario, error)
}

type Usuario struct{}

func GetInstanceUsuario() ServiceUsuario {
	onceServiceUsuario.Do(func() {
		instanceServiceUsuario = &Usuario{}
	})
	return instanceServiceUsuario
}

func (u *Usuario) CreateUsuario(ctx context.Context, usuarioReceive model.UsuarioReceive) error {
	usuario := model.Usuario{
		UsuarioID: utils.CreateCodeId(),
		Nome:      usuarioReceive.Nome,
		Email:     usuarioReceive.Email,
		Senha:     usuarioReceive.Senha,
		Nivel:     usuarioReceive.Nivel,
		Setor:     usuarioReceive.Setor,
		Celular:   usuarioReceive.Celular,
		Loja:      usuarioReceive.Loja,
	}

	usuarioMap := structs.Map(usuario)
	_, err := client.GetInstance().Insert(ctx, "Usuario", usuarioMap)
	if err != nil {
		return errors.New("Create Usuario: problem to insert into MongoDB")
	}
	return nil
}

func (u *Usuario) GetUsuarioByID(ctx context.Context, id string) (model.Usuario, error) {
	filter := map[string]interface{}{"UsuarioID": id}
	return repository.GetInstanceUsuario().FindOne(ctx, "Usuario", filter)
}

func (u *Usuario) GetAllUsuario(ctx context.Context) ([]model.Usuario, error) {
	filter := map[string]interface{}{}
	return repository.GetInstanceUsuario().Find(ctx, "Usuario", filter)
}

func (u *Usuario) UpdateUsuario(ctx context.Context, id string, usuario model.UsuarioReceive) error {
	filter := bson.M{"usuarioID": id}
	updateData := bson.M{
		"$set": bson.M{
			"nome":    usuario.Nome,
			"email":   usuario.Email,
			"senha":   usuario.Senha,
			"nivel":   usuario.Nivel,
			"setor":   usuario.Setor,
			"celular": usuario.Celular,
			"loja":    usuario.Loja,
		},
	}

	_, err := client.GetInstance().UpdateOne(ctx, "Usuario", filter, updateData)
	if err != nil {
		return errors.New("Update Usuario: problem to update into MongoDB")
	}
	return nil
}

func (u *Usuario) DeleteUsuario(ctx context.Context, id string) error {
	filter := map[string]interface{}{"UsuarioID": id}

	err := client.GetInstance().Remove(ctx, "Usuario", filter)
	if err != nil {
		return errors.New("Delete Usuario: problem to delete from MongoDB")
	}
	return nil
}

func (u *Usuario) GetUsuarioByEmail(ctx context.Context, email string) (model.Usuario, error) {
	filter := map[string]interface{}{"Email": email}
	return repository.GetInstanceUsuario().FindOne(ctx, "Usuario", filter)
}
