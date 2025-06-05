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
	instanceServiceSignature ServiceSignature
	onceServiceSignature     sync.Once
)

type ServiceSignature interface {
	CreateSignature(ctx context.Context, signature model.SignatureReceive) error
	GetSignatureByID(ctx context.Context, id string) (model.Signature, error)
	GetAllSignature(ctx context.Context) ([]model.Signature, error)
	UpdateSignature(ctx context.Context, update model.Signature) error
	DeleteSignature(ctx context.Context, id string) error
	GetSignatureByClienteID(ctx context.Context, clienteID string) ([]model.Signature, error)
}

type Signature struct{}

func GetInstanceSignature() ServiceSignature {
	onceServiceSignature.Do(func() {
		instanceServiceSignature = &Signature{}
	})
	return instanceServiceSignature
}

func (s *Signature) CreateSignature(ctx context.Context, signatureReceive model.SignatureReceive) error {

	signature := model.Signature{
		SignatureID:    utils.CreateCodeId(),
		ClienteID:      signatureReceive.ClienteID,
		DiaLancamento:  signatureReceive.DiaLancamento,
		DiaVencimento:  signatureReceive.DiaVencimento,
		QtdParcelas:    signatureReceive.QtdParcelas,
		CentroCustoID:  signatureReceive.CentroCustoID,
		ValorOperacao:  signatureReceive.ValorOperacao,
		EmitidoEsteMes: false,
		VencidoEsteMes: false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	signatureMap := structs.Map(signature)
	_, err := client.GetInstance().Insert(ctx, "Signature", signatureMap)
	if err != nil {
		return errors.New("Create Signature: problem to insert into MongoDB")
	}

	return nil
}

func (s *Signature) GetSignatureByID(ctx context.Context, id string) (model.Signature, error) {
	var signature model.Signature
	filter := map[string]interface{}{"SignatureID": id}

	signature, err := repository.GetInstanceSignature().FindOne(ctx, "Signature", filter)
	if err != nil {
		return signature, errors.New("Get Signature: problem to find by SignatureID")
	}
	return signature, nil
}

func (s *Signature) GetAllSignature(ctx context.Context) ([]model.Signature, error) {
	filter := map[string]interface{}{}

	signatures, err := repository.GetInstanceSignature().Find(ctx, "Signature", filter)
	if err != nil {
		return nil, errors.New("Get All Signatures: problem to find in MongoDB")
	}
	return signatures, nil
}

func (s *Signature) UpdateSignature(ctx context.Context, update model.Signature) error {

	filter := bson.M{"SignatureID": update.SignatureID}
	updateData := bson.M{
		"$set": bson.M{
			"ClienteID":      update.ClienteID,
			"DiaLancamento":  update.DiaLancamento,
			"DiaVencimento":  update.DiaVencimento,
			"QtdParcelas":    update.QtdParcelas,
			"CentroCusto":    update.CentroCustoID,
			"ValorOperacao":  update.ValorOperacao,
			"EmitidoEsteMes": update.EmitidoEsteMes,
			"VencidoEsteMes": update.VencidoEsteMes,
			"UpdatedAt":      time.Now(),
		},
	}

	_, err := client.GetInstance().UpdateOne(ctx, "Signature", filter, updateData)
	if err != nil {
		return errors.New("Update Signature: problem to update in MongoDB")
	}

	return nil
}

func (s *Signature) DeleteSignature(ctx context.Context, id string) error {
	filter := map[string]interface{}{"SignatureID": id}

	err := client.GetInstance().Remove(ctx, "Signature", filter)
	if err != nil {
		return errors.New("Delete Signature: problem to delete from MongoDB")
	}

	return nil
}

func (s *Signature) GetSignatureByClienteID(ctx context.Context, clienteID string) ([]model.Signature, error) {
	filter := map[string]interface{}{"ClienteID": clienteID}

	signatures, err := repository.GetInstanceSignature().Find(ctx, "Signature", filter)
	if err != nil {
		return nil, errors.New("Get Signature: problem to find by ClienteID")
	}
	return signatures, nil
}
