package service

import (
	"context"
	"errors"
	"pix-generation/src/client"
	"pix-generation/src/model"
	"pix-generation/src/repository"
	"pix-generation/src/utils"
	"sync"
	"time"

	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceServiceExpenseCenter ServiceExpenseCenter
	onceServiceExpenseCenter     sync.Once
)

type ServiceExpenseCenter interface {
	CreateExpenseCenter(ctx context.Context, expense model.ExpenseCenterReceive) error
	GetExpenseCenterByID(ctx context.Context, id string) (model.ExpenseCenter, error)
	UpdateExpenseCenter(ctx context.Context, id string, update model.ExpenseCenterReceive) error
	DeleteExpenseCenter(ctx context.Context, id string) error
}

type ExpenseCenter struct{}

func GetInstanceExpenseCenter() ServiceExpenseCenter {
	onceServiceExpenseCenter.Do(func() {
		instanceServiceExpenseCenter = &ExpenseCenter{}
	})
	return instanceServiceExpenseCenter
}

func (e *ExpenseCenter) CreateExpenseCenter(ctx context.Context, ec model.ExpenseCenterReceive) error {
	expense := model.ExpenseCenter{
		CentroExpenseID: utils.CreateCodeId(),
		NomeCentro:      ec.NomeCentro,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	expenseMap := structs.Map(expense)
	_, err := client.GetInstance().Insert(ctx, "ExpenseCenter", expenseMap)
	if err != nil {
		return errors.New("Create ExpenseCenter: problem to insert into MongoDB")
	}
	return nil
}

func (e *ExpenseCenter) GetExpenseCenterByID(ctx context.Context, id string) (model.ExpenseCenter, error) {
	var expense model.ExpenseCenter
	filter := map[string]interface{}{"CentroExpenseID": id}

	expense, err := repository.GetInstanceExpenseCenter().FindOne(ctx, "ExpenseCenter", filter)
	if err != nil {
		return expense, errors.New("Get ExpenseCenter: problem to find by CentroExpenseId")
	}
	return expense, nil
}

func (e *ExpenseCenter) UpdateExpenseCenter(ctx context.Context, id string, update model.ExpenseCenterReceive) error {
	filter := bson.M{"centroExpenseID": id}
	updateData := bson.M{
		"$set": bson.M{
			"nomeCentro": update.NomeCentro,
			"updatedAt":  time.Now(),
		},
	}

	_, err := client.GetInstance().UpdateOne(ctx, "ExpenseCenter", filter, updateData)
	if err != nil {
		return errors.New("Update ExpenseCenter: problem to update in MongoDB")
	}
	return nil
}

func (e *ExpenseCenter) DeleteExpenseCenter(ctx context.Context, id string) error {
	filter := map[string]interface{}{"centroExpenseID": id}
	err := client.GetInstance().Remove(ctx, "ExpenseCenter", filter)
	if err != nil {
		return errors.New("Delete ExpenseCenter: problem to delete from MongoDB")
	}
	return nil
}
