package repository

import (
	"context"
	"errors"
	"sync"

	"pix-generation/src/client"
	"pix-generation/src/model"
)

var (
	instanceRepositoryUser RepositoryUser
	onceRepositoryUser     sync.Once
)

type RepositoryUser interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.User, error)
	Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.User, error)
}

type repositoryUser struct{}

func GetInstanceUser() RepositoryUser {
	onceRepositoryUser.Do(func() {
		instanceRepositoryUser = &repositoryUser{}
	})
	return instanceRepositoryUser
}

func (ru *repositoryUser) Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.User, error) {

	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDB")
	}

	var content []model.User
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return content, nil
}

func (ru *repositoryUser) FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.User, error) {

	var user model.User
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	if err != nil {
		return user, errors.New("Error Repository: Error find query in mongoDb")
	}
	result.Decode(&user)

	return user, nil
}
