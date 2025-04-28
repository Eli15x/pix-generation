package handler

import (
	"context"
	"net/http"

	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

// ExpenseCenterHandler define o handler com injeção de dependência do service
type ExpenseCenterHandler struct {
	service service.ServiceExpenseCenter
}

// NewExpenseCenterHandler retorna um handler com o service injetado
func NewExpenseCenterHandler(s service.ServiceExpenseCenter) *ExpenseCenterHandler {
	return &ExpenseCenterHandler{service: s}
}

// CreateExpenseCenter godoc
// @Summary      Cria um novo centro de custo
// @Description  Cria um novo centro de custo com nome
// @Tags         expense_center
// @Accept       json
// @Produce      json
// @Param        expense_center  body      model.ExpenseCenterReceive  true  "Dados do centro de custo"
// @Success      200             {string}  string "ok"
// @Failure      400             {object}  map[string]string
// @Failure      500             {object}  map[string]string
// @Router       /expense-center [post]
func (h *ExpenseCenterHandler) CreateExpenseCenter(c *gin.Context) {
	var ec model.ExpenseCenterReceive
	if err := c.ShouldBindJSON(&ec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.CreateExpenseCenter(context.Background(), ec)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// GetExpenseCenterByID godoc
// @Summary      Busca centro de custo por ID
// @Description  Retorna um centro de custo com base no ID
// @Tags         expense_center
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "ID do centro de custo"
// @Success      200 {object}  model.ExpenseCenter
// @Failure      400 {object}  map[string]string
// @Failure      404 {object}  map[string]string
// @Router       /expense-center/id/{id} [get]
func (h *ExpenseCenterHandler) GetExpenseCenterByID(c *gin.Context) {
	id := c.Param("id")
	center, err := h.service.GetExpenseCenterByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, center)
}

// DeleteExpenseCenter godoc
// @Summary      Deleta centro de custo
// @Description  Remove um centro de custo com base no ID
// @Tags         expense_center
// @Accept       json
// @Produce      json
// @Param        expense_center  body      model.ExpenseCenterDeleteRequest  true  "ID do centro de custo"
// @Success      200 {object}  map[string]string
// @Failure      400 {object}  map[string]string
// @Failure      500 {object}  map[string]string
// @Router       /expense-center [delete]
func (h *ExpenseCenterHandler) DeleteExpenseCenter(c *gin.Context) {
	var req model.ExpenseCenterDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.DeleteExpenseCenter(context.Background(), req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Expense center deleted"})
}
