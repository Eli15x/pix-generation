package repository

import (
	"context"
	"errors"
	"sync"

	"pix-generation/src/client"
	"pix-generation/src/model"
)

var (
	instanceRepositoryEmissor RepositoryEmissor
	onceRepositoryEmissor     sync.Once
)

type RepositoryEmissor interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Emissor, error)
	Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Emissor, error)
}

type repositoryEmissor struct{}

func GetInstanceEmissor() RepositoryEmissor {
	onceRepositoryEmissor.Do(func() {
		instanceRepositoryEmissor = &repositoryEmissor{}
	})
	return instanceRepositoryEmissor
}

func (ru *repositoryEmissor) Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Emissor, error) {
	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in MongoDB")
	}

	var content []model.Emissor
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information MongoDB")
	}

	return content, nil
}

func (ru *repositoryEmissor) FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Emissor, error) {
	var emissor model.Emissor
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	if err != nil {
		return emissor, errors.New("Error Repository: Error find query in MongoDB")
	}
	result.Decode(&emissor)

	return emissor, nil
}
