package service

import (
	"pix-generation/src/client"

	//"go.mongodb.org/mongo-driver/bson"
	"context"
	"errors"
	"sync"

	//"fmt"

	"pix-generation/src/middleware"
	"pix-generation/src/model"
	"pix-generation/src/repository"
	"pix-generation/src/utils"

	"github.com/fatih/structs"
)

var (
	instanceServiceInvoice ServiceInvoice
	onceServiceInvoice     sync.Once
)

//aqui só falta adicionar o util igual do zcom e funciona.
type ServiceInvoice interface {
	ValidateInvoice(ctx context.Context, email string, password string) (model.ResponseInvoice, error)
	GetInvoice(ctx context.Context, id string) (model.Invoice, error)
	GetInvoiceByName(ctx context.Context, name string) (model.Invoice, error)
	GetInvoiceByEmail(ctx context.Context, email string) (model.Invoice, error)
	GetInvoicesByClientId(ctx context.Context, idAcess int) ([]model.Invoice, error)
	GetInvoices(ctx context.Context) ([]model.Invoice, error)

	CreateInvoice(ctx context.Context, Invoice model.Invoice) (model.ResponseInvoice, error)
	EditInvoice(ctx context.Context, Invoice model.Invoice) error
	DeleteInvoice(ctx context.Context, Invoice model.Invoice) error
}

type Invoice struct{}

func GetInstanceInvoice() ServiceInvoice {
	onceServiceInvoice.Do(func() {
		instanceServiceInvoice = &Invoice{}
	})
	return instanceServiceInvoice
}

func (u *Invoice) GetInvoice(ctx context.Context, id string) (model.Invoice, error) {
	var Invoice model.Invoice

	InvoiceId := map[string]interface{}{"InvoiceId": id}

	Invoice, err := repository.GetInstanceInvoice().FindOne(ctx, "Invoice", InvoiceId)
	if err != nil {
		return Invoice, errors.New("Get Invoice: problem to Find Id into MongoDB")
	}

	return Invoice, nil
}

func (u *Invoice) GetInvoicesByCnpj(ctx context.Context, dateStart string, dateEnd string, cnpj int) ([]model.Invoice, error) {
	//criar filtro opcional para date. se for passado pequisar se nao nao pesquisar por ele e sim so pelo cnpj
	//if ..... ver como passar o dateStart como algo especifico para entender que nao esta sendo considerado.

	Cnpj := map[string]interface{}{"cnpjCliente": cnpj}

	Invoices, err := repository.GetInstanceInvoice().Find(ctx, "Invoice", Cnpj)
	if err != nil {
		return nil, errors.New("Get Invoices By Acess: problem to Find cnpj into MongoDB")
	}

	return Invoices, nil
}

func (u *Invoice) CreateInvoice(ctx context.Context, Invoice model.Invoice) (model.ResponseInvoice, error) {

	var responseInvoice model.ResponseInvoice

	var InvoiceId = utils.CreateCodeId()
	Invoice.InvoiceID = InvoiceId

	InvoiceInsert := structs.Map(Invoice)
	_, err := client.GetInstance().Insert(ctx, "Invoice", InvoiceInsert)
	if err != nil {
		return responseInvoice, errors.New("Create Invoice: problem to insert into MongoDB")
	}

	token, err := middleware.GenerateJWT(Invoice.InvoiceID)
	if err != nil {
		return responseInvoice, errors.New("Failed to generate JWT")
	}

	responseInvoice = model.ResponseInvoice{
		InvoiceID: Invoice.InvoiceID,
		JWT:       token,
	}

	return responseInvoice, nil

}

func (u *Invoice) DeleteInvoice(ctx context.Context, dateStart string, dateEnd string, cnpj string) error {

	//varios filtros para busca ver como faz isso.
	InvoiceId := map[string]interface{}{"Invoice_id": Invoice.InvoiceID}

	err := client.GetInstance().Remove(ctx, "Invoice", InvoiceId)
	if err != nil {
		return errors.New("Delete Invoice: problem to delete into MongoDB")
	}

	return nil
}
