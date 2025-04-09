package service

import (
	"pix-generation/src/client"
	"pix-generation/src/utils"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"context"
	"errors"
	"sync"

	//"fmt"

	"pix-generation/src/model"
	"pix-generation/src/repository"

	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceServiceInvoice ServiceInvoice
	onceServiceInvoice     sync.Once
)

// aqui só falta adicionar o util igual do zcom e funciona.
type ServiceInvoice interface {
	GetInvoice(ctx context.Context, id string) (model.Invoice, error)
	GetInvoicesByCnpj(ctx context.Context, dateStart time.Time, dateEnd time.Time, cnpj string) ([]model.Invoice, error)

	CreateInvoice(ctx context.Context, Invoice model.InvoiceReceive) error
	DeleteInvoiceByData(ctx context.Context, dateStart time.Time, dateEnd time.Time, cnpj string) error
	DeleteInvoice(ctx context.Context, invoiceId string) error
}

type Invoice struct{}

func GetInstanceInvoice() ServiceInvoice {
	onceServiceInvoice.Do(func() {
		instanceServiceInvoice = &Invoice{}
	})
	return instanceServiceInvoice
}

func (i *Invoice) GetInvoice(ctx context.Context, id string) (model.Invoice, error) {
	var Invoice model.Invoice

	InvoiceId := map[string]interface{}{"invoiceID": id}

	Invoice, err := repository.GetInstanceInvoice().FindOne(ctx, "Invoice", InvoiceId)
	if err != nil {
		return Invoice, errors.New("Get Invoice: problem to Find Id into MongoDB")
	}

	return Invoice, nil
}

func (i *Invoice) GetInvoicesByCnpj(ctx context.Context, dateStart time.Time, dateEnd time.Time, cnpj string) ([]model.Invoice, error) {

	filter := map[string]interface{}{
		"CnpjCliente": cnpj,
	}

	if !dateStart.IsZero() && !dateEnd.IsZero() {
		filter["Emitido"] = map[string]interface{}{
			"$gte": dateStart,
			"$lte": dateEnd,
		}
	}

	Invoices, err := repository.GetInstanceInvoice().Find(ctx, "Invoice", filter)
	if err != nil {
		return nil, errors.New("Get Invoices By CNPJ: problema ao buscar no MongoDB")
	}

	return Invoices, nil
}

func (i *Invoice) CreateInvoice(ctx context.Context, invoiceReceive model.InvoiceReceive) error {

	var InvoiceId = utils.CreateCodeId()
	var invoice model.Invoice

	// Copiar os valores normais

	invoice.InvoiceID = InvoiceId
	invoice.CnpjCliente = invoiceReceive.CnpjCliente
	invoice.Amount = invoiceReceive.Amount
	invoice.TxId = invoiceReceive.TxId
	invoice.TaxaPaga = invoiceReceive.TaxaPaga
	invoice.Expira = invoiceReceive.Expira
	invoice.Uuid = invoiceReceive.Uuid
	invoice.Pago = invoiceReceive.Pago

	// Converter "emitido" de string para time.Time
	if invoiceReceive.Emitido != "" {
		parsedEmitido, err := time.Parse("2006-01-02T15:04:05", invoiceReceive.Emitido)

		if err != nil {
			return errors.New("Create Invoice: invalid date format for emitido. Use 'YYYY-MM-DDTHH:MM:SS'")
		}
		invoice.Emitido = parsedEmitido
	}

	InvoiceInsert := structs.Map(invoice)
	_, err := client.GetInstance().Insert(ctx, "Invoice", InvoiceInsert)
	if err != nil {
		return errors.New("Create Invoice: problem to insert into MongoDB")
	}

	return nil

}

func (i *Invoice) DeleteInvoiceByData(ctx context.Context, dateStart time.Time, dateEnd time.Time, cnpj string) error {

	filter := bson.M{
		"CnpjCliente": cnpj,
		"Emitido": bson.M{
			"$gte": dateStart,
			"$lte": dateEnd,
		},
	}

	err := client.GetInstance().RemoveMany(ctx, "Invoice", filter)
	if err != nil {
		return errors.New("Delete Invoice by data: problem to delete into MongoDB")
	}

	return nil
}

func (i *Invoice) DeleteInvoice(ctx context.Context, invoiceId string) error {

	invoiceIdValue := map[string]interface{}{"InvoiceID": invoiceId}

	err := client.GetInstance().Remove(ctx, "invoice", invoiceIdValue)
	if err != nil {
		return errors.New("Delete Invoice: problem to delete into MongoDB")
	}

	return nil
}
