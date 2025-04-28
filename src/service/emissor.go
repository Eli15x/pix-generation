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
	instanceServiceEmissor ServiceEmissor
	onceServiceEmissor     sync.Once
)

type ServiceEmissor interface {
	CreateEmissor(ctx context.Context, emissor model.EmissorReceive) error
	GetEmissorByID(ctx context.Context, id string) (model.Emissor, error)
	GetAllEmissor(ctx context.Context) ([]model.Emissor, error)
	UpdateEmissor(ctx context.Context, id string, update model.EmissorReceive) error
	DeleteEmissor(ctx context.Context, id string) error
	GetEmissorByDocumento(ctx context.Context, documento string) (model.Emissor, error)
}

type Emissor struct{}

func GetInstanceEmissor() ServiceEmissor {
	onceServiceEmissor.Do(func() {
		instanceServiceEmissor = &Emissor{}
	})
	return instanceServiceEmissor
}

func (e *Emissor) CreateEmissor(ctx context.Context, emissorReceive model.EmissorReceive) error {
	var emissor model.Emissor

	emissor.EmissorID = utils.CreateCodeId()
	emissor.Nome = emissorReceive.Nome
	emissor.Email = emissorReceive.Email
	emissor.Senha = emissorReceive.Senha
	emissor.Documento = emissorReceive.Documento
	emissor.TokenConfraPix = emissorReceive.TokenConfraPix

	emissorMap := structs.Map(emissor)
	_, err := client.GetInstance().Insert(ctx, "Emissor", emissorMap)
	if err != nil {
		return errors.New("Create Emissor: problem to insert into MongoDB")
	}

	return nil
}

func (e *Emissor) GetEmissorByID(ctx context.Context, id string) (model.Emissor, error) {
	var emissor model.Emissor
	filter := map[string]interface{}{"emissorID": id}

	emissor, err := repository.GetInstanceEmissor().FindOne(ctx, "Emissor", filter)
	if err != nil {
		return emissor, errors.New("Get Emissor: problem to find by EmissorID")
	}
	return emissor, nil
}

func (e *Emissor) GetAllEmissor(ctx context.Context) ([]model.Emissor, error) {
	filter := map[string]interface{}{}

	emissores, err := repository.GetInstanceEmissor().Find(ctx, "Emissor", filter)
	if err != nil {
		return nil, errors.New("Get All Emissores: problem to find in MongoDB")
	}
	return emissores, nil
}

func (e *Emissor) UpdateEmissor(ctx context.Context, id string, update model.EmissorReceive) error {
	filter := bson.M{"emissorID": id}
	updateData := structs.Map(update)
	_, err := client.GetInstance().UpdateOne(ctx, "Emissor", filter, updateData)
	if err != nil {
		return errors.New("Update Emissor: problem to update in MongoDB")
	}

	return nil
}

func (e *Emissor) DeleteEmissor(ctx context.Context, id string) error {
	filter := map[string]interface{}{"emissorID": id}

	err := client.GetInstance().Remove(ctx, "Emissor", filter)
	if err != nil {
		return errors.New("Delete Emissor: problem to delete from MongoDB")
	}

	return nil
}

func (e *Emissor) GetEmissorByDocumento(ctx context.Context, documento string) (model.Emissor, error) {
	var emissor model.Emissor
	filter := map[string]interface{}{"documento": documento}

	emissor, err := repository.GetInstanceEmissor().FindOne(ctx, "Emissor", filter)
	if err != nil {
		return emissor, errors.New("Get Emissor: problem to find by Documento")
	}
	return emissor, nil
}
