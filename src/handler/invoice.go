package handler

import (
	"context"
	"net/http"
	"time"

	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(c *gin.Context) {
	var Invoice model.InvoiceReceive
	if err := c.ShouldBindJSON(&Invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.GetInstanceInvoice().CreateInvoice(context.Background(), Invoice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func GetByID(c *gin.Context) {
	var Invoice model.Invoice
	if err := c.ShouldBindJSON(&Invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Invoice, err := service.GetInstanceInvoice().GetInvoice(context.Background(), Invoice.InvoiceID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Invoice)
}

func GetByCnpj(c *gin.Context) {
	// Pega os parâmetros dateStart e dateEnd da query string
	dateStartStr := c.Param("startDate")
	dateEndStr := c.Param("endDate")

	var dateStart, dateEnd time.Time
	var err error

	if dateStartStr != "" && dateEndStr != "" {
		dateStart, err = time.Parse("2006-01-02", dateStartStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido para dateStart. Use YYYY-MM-DD"})
			return
		}

		dateEnd, err = time.Parse("2006-01-02", dateEndStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido para dateEnd. Use YYYY-MM-DD"})
			return
		}
	}

	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O corpo da requisição está vazio"})
		return
	}

	var InvoiceRecieve model.Invoice
	if err := c.ShouldBindJSON(&InvoiceRecieve); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Invoices, err := service.GetInstanceInvoice().GetInvoicesByCnpj(context.Background(), dateStart, dateEnd, InvoiceRecieve.CnpjCliente)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Invoices)
}

func DeleteInvoice(c *gin.Context) {
	// Pega os parâmetros dateStart e dateEnd da query string
	dateStartStr := c.Param("startDate")
	dateEndStr := c.Param("endDate")

	var dateStart, dateEnd time.Time
	var err error

	// Converte as datas, se foram fornecidas
	if dateStartStr != "" && dateEndStr != "" {

		dateStart, err = time.Parse("2006-01-02", dateStartStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido para dateStart. Use YYYY-MM-DD"})
			return
		}

		dateEnd, err = time.Parse("2006-01-02", dateEndStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido para dateEnd. Use YYYY-MM-DD"})
			return
		}
	}

	if dateStart.IsZero() && dateEnd.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Os parâmetros dateStart e dateEnd são obrigatórios."})
		return
	}

	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O corpo da requisição está vazio"})
		return
	}
	// Pega os dados do Invoice no corpo da requisição
	var invoice model.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o serviço para deletar
	err = service.GetInstanceInvoice().DeleteInvoice(context.Background(), dateStart, dateEnd, invoice.CnpjCliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Invoice deleted"})
}
