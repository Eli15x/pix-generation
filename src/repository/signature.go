package repository

import (
	"context"
	"errors"
	"sync"

	"pix-generation/src/client"
	"pix-generation/src/model"
)

var (
	instanceRepositorySignature RepositorySignature
	onceRepositorySignature     sync.Once
)

type RepositorySignature interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Signature, error)
	Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Signature, error)
}

type repositorySignature struct{}

func GetInstanceSignature() RepositorySignature {
	onceRepositorySignature.Do(func() {
		instanceRepositorySignature = &repositorySignature{}
	})
	return instanceRepositorySignature
}

func (ru *repositorySignature) Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Signature, error) {
	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in MongoDB")
	}

	var content []model.Signature
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information MongoDB")
	}

	return content, nil
}

func (ru *repositorySignature) FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Signature, error) {
	var signature model.Signature
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	if err != nil {
		return signature, errors.New("Error Repository: Error find query in MongoDB")
	}
	result.Decode(&signature)

	return signature, nil
}
