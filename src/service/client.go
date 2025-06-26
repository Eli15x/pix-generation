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
	instanceServiceClient ServiceClient
	onceServiceClient     sync.Once
)

type ServiceClient interface {
	CreateClient(ctx context.Context, client model.ClientReceive) error
	GetClientByID(ctx context.Context, id string) (model.Client, error)
	GetClientByUserID(ctx context.Context, id string) (model.Client, error)
	GetAllClient(ctx context.Context) ([]model.Client, error)
	UpdateClient(ctx context.Context, id string, update model.ClientReceive) error
	DeleteClient(ctx context.Context, id string) error
	GetClientByCpf(ctx context.Context, cpf string) (model.Client, error)
}

type Client struct{}

func GetInstanceClient() ServiceClient {
	onceServiceClient.Do(func() {
		instanceServiceClient = &Client{}
	})
	return instanceServiceClient
}

func (c *Client) CreateClient(ctx context.Context, clientReceive model.ClientReceive) error {
	var clientStruct model.Client

	clientStruct.ClientID = utils.CreateCodeId()
	clientStruct.UserID = clientReceive.UserID
	clientStruct.Nome = clientReceive.Nome
	clientStruct.CPF = clientReceive.CPF
	clientStruct.Email = clientReceive.Email
	clientStruct.Celular = clientReceive.Celular
	clientStruct.CreatedAt = time.Now()
	clientStruct.UpdatedAt = time.Now()

	clientMap := structs.Map(clientStruct)
	_, err := client.GetInstance().Insert(ctx, "Client", clientMap)
	if err != nil {
		return errors.New("Create Client: problem to insert into MongoDB")
	}

	return nil
}

func (c *Client) GetClientByID(ctx context.Context, id string) (model.Client, error) {
	var client model.Client
	filter := map[string]interface{}{"ClientID": id}

	client, err := repository.GetInstanceClient().FindOne(ctx, "Client", filter)
	if err != nil {
		return client, errors.New("Get Client: problem to find by ClientID")
	}

	if client == (model.Client{}) {
		return model.Client{}, errors.New("Get Client: not exists client with this id")
	}

	return client, nil
}

func (c *Client) GetClientByUserID(ctx context.Context, id string) (model.Client, error) {
	var client model.Client
	filter := map[string]interface{}{"UserID": id}

	client, err := repository.GetInstanceClient().FindOne(ctx, "Client", filter)
	if err != nil {
		return client, errors.New("Get Client: problem to find by ClientID")
	}

	return client, nil
}

func (c *Client) GetAllClient(ctx context.Context) ([]model.Client, error) {
	filter := map[string]interface{}{}

	clients, err := repository.GetInstanceClient().Find(ctx, "Client", filter)
	if err != nil {
		return nil, errors.New("Get All Clients: problem to find in MongoDB")
	}
	return clients, nil
}

func (c *Client) UpdateClient(ctx context.Context, id string, update model.ClientReceive) error {
	filter := bson.M{"ClientID": id}
	updateData := bson.M{
		"$set": bson.M{
			"Nome":      update.Nome,
			"CPF":       update.CPF,
			"Email":     update.Email,
			"Celular":   update.Celular,
			"UpdatedAt": time.Now(),
		},
	}

	_, err := client.GetInstance().UpdateOne(ctx, "Client", filter, updateData)
	if err != nil {
		return errors.New("Update Client: problem to update in MongoDB")
	}

	return nil
}

func (c *Client) DeleteClient(ctx context.Context, id string) error {
	filter := map[string]interface{}{"ClientID": id}

	err := client.GetInstance().Remove(ctx, "Client", filter)
	if err != nil {
		return errors.New("Delete Client: problem to delete from MongoDB")
	}

	return nil
}

func (c *Client) GetClientByCpf(ctx context.Context, cpf string) (model.Client, error) {
	var client model.Client
	filter := map[string]interface{}{"Cpf": cpf}

	client, err := repository.GetInstanceClient().FindOne(ctx, "Client", filter)
	if err != nil {
		return client, errors.New("Get Client: problem to find by CPF")
	}
	return client, nil
}
