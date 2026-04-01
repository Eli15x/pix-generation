package repository

import (
	"context"
	"errors"
	"sync"

	"pix-generation/src/client"
	"pix-generation/src/model"
)

var (
	instanceRepositoryAdminUser RepositoryAdminUser
	onceRepositoryAdminUser     sync.Once
)

type RepositoryAdminUser interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.AdminUser, error)
	Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.AdminUser, error)
}

type repositoryAdminUser struct{}

func GetInstanceAdminUser() RepositoryAdminUser {
	onceRepositoryAdminUser.Do(func() {
		instanceRepositoryAdminUser = &repositoryAdminUser{}
	})
	return instanceRepositoryAdminUser
}

func (ru *repositoryAdminUser) Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.AdminUser, error) {

	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDB")
	}

	var content []model.AdminUser
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return content, nil
}

func (ru *repositoryAdminUser) FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.AdminUser, error) {

	var user model.AdminUser
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	if err != nil {
		return user, errors.New("Error Repository: Error find query in mongoDb")
	}
	result.Decode(&user)

	return user, nil
}
