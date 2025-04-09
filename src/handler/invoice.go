package handler

import (
	"context"
	"net/http"
	"time"

	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

// CreateInvoice godoc
// @Summary      Cria uma nova fatura (Invoice)
// @Description  Cria uma nova fatura a partir dos dados fornecidos
// @Tags         invoice
// @Accept       json
// @Produce      json
// @Param        invoice  body      model.InvoiceReceive  true  "Dados da fatura"
// @Success      200      {string}  string "ok"
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /invoice [post]
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

// GetByID godoc
// @Summary      Busca fatura por ID
// @Description  Busca uma fatura pelo ID
// @Tags         invoice
// @Accept       json
// @Produce      json
// @Param        invoice  body      model.InvoiceIDRequest  true  "ID da fatura"
// @Success      200      {object}  model.Invoice
// @Failure      400      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Router       /invoice/id/{id} [get]
func GetByID(c *gin.Context) {
	var Invoice model.InvoiceIDRequest
	if err := c.ShouldBindJSON(&Invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	InvoiceValue, err := service.GetInstanceInvoice().GetInvoice(context.Background(), Invoice.InvoiceID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, InvoiceValue)
}

// GetByCnpj godoc
// @Summary      Busca faturas por CNPJ e intervalo de datas
// @Description  Busca todas as faturas pelo CNPJ e por data de início e fim
// @Tags         invoice
// @Accept       json
// @Produce      json
// @Param        startDate  path      string  false  "Data inicial (YYYY-MM-DD)"
// @Param        endDate    path      string  false  "Data final (YYYY-MM-DD)"
// @Param        invoice    body      model.InvoiceCNPJRequest  true  "CNPJ do cliente"
// @Success      200        {array}   model.Invoice
// @Failure      400        {object}  map[string]string
// @Failure      404        {object}  map[string]string
// @Router       /invoice/{startDate}/{endDate}/ [post]
func GetByCnpj(c *gin.Context) {
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

	var InvoiceRecieve model.InvoiceCNPJRequest
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

// DeleteInvoiceByData godoc
// @Summary      Deleta faturas por data e CNPJ
// @Description  Remove faturas com base no CNPJ e intervalo de datas
// @Tags         invoice
// @Accept       json
// @Produce      json
// @Param        startDate  path      string  true   "Data inicial (YYYY-MM-DD)"
// @Param        endDate    path      string  true   "Data final (YYYY-MM-DD)"
// @Param        invoice    body      model.InvoiceCNPJRequest  true  "CNPJ do cliente"
// @Success      200        {object}  map[string]string
// @Failure      400        {object}  map[string]string
// @Failure      500        {object}  map[string]string
// @Router       /invoice/{startDate}/{endDate}/ [delete]
func DeleteInvoiceByData(c *gin.Context) {
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

	if dateStart.IsZero() || dateEnd.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Os parâmetros dateStart e dateEnd são obrigatórios."})
		return
	}

	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O corpo da requisição está vazio"})
		return
	}

	var invoice model.InvoiceCNPJRequest
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.GetInstanceInvoice().DeleteInvoiceByData(context.Background(), dateStart, dateEnd, invoice.CnpjCliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Invoice deleted"})
}

// DeleteInvoice godoc
// @Summary      Deleta fatura
// @Description  Remove uma fatura específica
// @Tags         invoice
// @Accept       json
// @Produce      json
// @Param        invoice  body      model.InvoiceDeleteRequest  true  "Dados da fatura"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /invoice [delete]
func DeleteInvoice(c *gin.Context) {
	var invoice model.InvoiceDeleteRequest
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.GetInstanceInvoice().DeleteInvoice(context.Background(), invoice.InvoiceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "User deleted"})
}
