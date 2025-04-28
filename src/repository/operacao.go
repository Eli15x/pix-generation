package repository

import (
	"context"
	"errors"
	"sync"

	"pix-generation/src/client"
	"pix-generation/src/model"
)

var (
	instanceRepositoryOperacao RepositoryOperacao
	onceRepositoryOperacao     sync.Once
)

type RepositoryOperacao interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Operacao, error)
	Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Operacao, error)
}

type repositoryOperacao struct{}

func GetInstanceOperacao() RepositoryOperacao {
	onceRepositoryOperacao.Do(func() {
		instanceRepositoryOperacao = &repositoryOperacao{}
	})
	return instanceRepositoryOperacao
}

func (ru *repositoryOperacao) Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Operacao, error) {
	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in MongoDB")
	}

	var content []model.Operacao
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information MongoDB")
	}

	return content, nil
}

func (ru *repositoryOperacao) FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Operacao, error) {
	var operacao model.Operacao
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	if err != nil {
		return operacao, errors.New("Error Repository: Error find query in MongoDB")
	}
	result.Decode(&operacao)

	return operacao, nil
}
