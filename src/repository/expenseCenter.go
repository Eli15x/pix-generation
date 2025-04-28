package repository

import (
	"context"
	"errors"
	"sync"

	"pix-generation/src/client"
	"pix-generation/src/model"
)

var (
	instanceRepositoryExpenseCenter RepositoryExpenseCenter
	onceRepositoryExpenseCenter     sync.Once
)

type RepositoryExpenseCenter interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.ExpenseCenter, error)
	Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.ExpenseCenter, error)
}

type repositoryExpenseCenter struct{}

func GetInstanceExpenseCenter() RepositoryExpenseCenter {
	onceRepositoryExpenseCenter.Do(func() {
		instanceRepositoryExpenseCenter = &repositoryExpenseCenter{}
	})
	return instanceRepositoryExpenseCenter
}

func (ru *repositoryExpenseCenter) Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.ExpenseCenter, error) {
	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDB")
	}

	var content []model.ExpenseCenter
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return content, nil
}

func (ru *repositoryExpenseCenter) FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.ExpenseCenter, error) {
	var expense model.ExpenseCenter
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	if err != nil {
		return expense, errors.New("Error Repository: Error find query in mongoDB")
	}
	result.Decode(&expense)

	return expense, nil
}
