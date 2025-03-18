package repository

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"pix-generation/src/client"
	"pix-generation/src/model"
)

var (
	instanceRepositoryInvoice RepositoryInvoice
	onceRepositoryInvoice     sync.Once
)

type RepositoryInvoice interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Invoice, error)
	Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Invoice, error)
}

type repositoryInvoice struct{}

func GetInstanceInvoice() RepositoryInvoice {
	onceRepositoryInvoice.Do(func() {
		instanceRepositoryInvoice = &repositoryInvoice{}
	})
	return instanceRepositoryInvoice
}

func (ru *repositoryInvoice) Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Invoice, error) {

	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDB")
	}

	var content []model.Invoice
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return content, nil
}

func (ru *repositoryInvoice) FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Invoice, error) {

	var Invoice model.Invoice
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	fmt.Println(result)
	if err != nil {
		return Invoice, errors.New("Error Repository: Error find query in mongoDb")
	}
	result.Decode(&Invoice)

	return Invoice, nil
}
