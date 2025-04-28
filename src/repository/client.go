package repository

import (
	"context"
	"errors"
	"sync"

	"pix-generation/src/client"
	"pix-generation/src/model"
)

var (
	instanceRepositoryClient RepositoryClient
	onceRepositoryClient     sync.Once
)

type RepositoryClient interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Client, error)
	Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Client, error)
}

type repositoryClient struct{}

func GetInstanceClient() RepositoryClient {
	onceRepositoryClient.Do(func() {
		instanceRepositoryClient = &repositoryClient{}
	})
	return instanceRepositoryClient
}

func (ru *repositoryClient) Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Client, error) {
	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in MongoDB")
	}

	var content []model.Client
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information MongoDB")
	}

	return content, nil
}

func (ru *repositoryClient) FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Client, error) {
	var clientStruct model.Client
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	if err != nil {
		return clientStruct, errors.New("Error Repository: Error find query in MongoDB")
	}
	result.Decode(&clientStruct)

	return clientStruct, nil
}
