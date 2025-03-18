package handler

import (
	"context"
	"net/http"

	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(c *gin.Context) {
	var Invoice model.Invoice
	if err := c.ShouldBindJSON(&Invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := service.GetInstanceInvoice().CreateInvoice(context.Background(), Invoice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetByID(c *gin.Context) {
	var Invoice model.Invoice
	if err := c.ShouldBindJSON(&Invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Invoice, err := service.GetInstanceInvoice().GetInvoice(context.Background(), Invoice.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Invoice)
}

func GetByCnpj(c *gin.Context) {

	var Invoice model.Invoice
	if err := c.ShouldBindJSON(&Invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Invoice, err := service.GetInstanceInvoice().GetInvoicesByCnpj(context.Background(), "", "", Invoice.cnpjCliente) //aqui precisa a variavel ser time.time e ver como passar ela zerada para nao ser usada
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Invoice)
}

func GetByCnpjWithDate(c *gin.Context) {
	//pegar por parametro os dateStart e dateEnd
	//dateStart
	//dateEnd
	var Invoice model.Invoice
	if err := c.ShouldBindJSON(&Invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Invoice, err := service.GetInstanceInvoice().GetInvoicesByCnpj(context.Background(), dateStart, dateEnd, Invoice.cnpjCliente)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Invoice)
}

func DeleteInvoice(c *gin.Context) {
	//aqui tambem pegar paraemtro de dateStart  e dateEnd

	var Invoice model.Invoice
	if err := c.ShouldBindJSON(&Invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.GetInstanceInvoice().DeleteInvoice(context.Background(), dateStart, dateEnd, Invoice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Invoice deleted"})
}
