package repository

import (
	"context"
	"errors"
	"sync"

	"pix-generation/src/client"
	"pix-generation/src/model"
)

var (
	instanceRepositoryUsuario RepositoryUsuario
	onceRepositoryUsuario     sync.Once
)

type RepositoryUsuario interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Usuario, error)
	Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Usuario, error)
}

type repositoryUsuario struct{}

func GetInstanceUsuario() RepositoryUsuario {
	onceRepositoryUsuario.Do(func() {
		instanceRepositoryUsuario = &repositoryUsuario{}
	})
	return instanceRepositoryUsuario
}

func (ru *repositoryUsuario) Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Usuario, error) {
	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in MongoDB")
	}

	var content []model.Usuario
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information MongoDB")
	}

	return content, nil
}

func (ru *repositoryUsuario) FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Usuario, error) {
	var usuario model.Usuario
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	if err != nil {
		return usuario, errors.New("Error Repository: Error find query in MongoDB")
	}
	result.Decode(&usuario)

	return usuario, nil
}
