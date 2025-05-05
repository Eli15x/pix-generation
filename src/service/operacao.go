package service

import (
	"context"
	"errors"
	"pix-generation/src/client"
	"pix-generation/src/model"
	"pix-generation/src/repository"
	"pix-generation/src/utils"
	"sync"

	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceServiceOperacao ServiceOperacao
	onceServiceOperacao     sync.Once
)

type ServiceOperacao interface {
	CreateOperacao(ctx context.Context, operacao model.OperacaoReceive) error
	GetOperacaoByID(ctx context.Context, id string) (model.Operacao, error)
	GetAllOperacao(ctx context.Context) ([]model.Operacao, error)
	UpdateOperacao(ctx context.Context, id string, operacao model.OperacaoReceive) error
	DeleteOperacao(ctx context.Context, id string) error
}

type Operacao struct{}

func GetInstanceOperacao() ServiceOperacao {
	onceServiceOperacao.Do(func() {
		instanceServiceOperacao = &Operacao{}
	})
	return instanceServiceOperacao
}

func (o *Operacao) CreateOperacao(ctx context.Context, operacaoReceive model.OperacaoReceive) error {
	operacao := model.Operacao{
		OperacaoID: utils.CreateCodeId(),
		Nome:       operacaoReceive.Nome,
	}

	operacaoMap := structs.Map(operacao)
	_, err := client.GetInstance().Insert(ctx, "Operacao", operacaoMap)
	if err != nil {
		return errors.New("Create Operacao: problem to insert into MongoDB")
	}
	return nil
}

func (o *Operacao) GetOperacaoByID(ctx context.Context, id string) (model.Operacao, error) {
	filter := map[string]interface{}{"operacaoID": id}
	return repository.GetInstanceOperacao().FindOne(ctx, "Operacao", filter)
}

func (o *Operacao) GetAllOperacao(ctx context.Context) ([]model.Operacao, error) {
	filter := map[string]interface{}{}
	return repository.GetInstanceOperacao().Find(ctx, "Operacao", filter)
}

func (o *Operacao) UpdateOperacao(ctx context.Context, id string, operacao model.OperacaoReceive) error {
	filter := bson.M{"OperacaoID": id}
	updateData := bson.M{
		"$set": bson.M{
			"nome": operacao.Nome,
		},
	}

	_, err := client.GetInstance().UpdateOne(ctx, "Operacao", filter, updateData)
	if err != nil {
		return errors.New("Update Operacao: problem to update into MongoDB")
	}
	return nil
}

func (o *Operacao) DeleteOperacao(ctx context.Context, id string) error {
	filter := map[string]interface{}{"OperacaoID": id}

	err := client.GetInstance().Remove(ctx, "Operacao", filter)
	if err != nil {
		return errors.New("Delete Operacao: problem to delete from MongoDB")
	}
	return nil
}
