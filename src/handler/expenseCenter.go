package handler

import (
	"context"
	"net/http"

	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

// ExpenseCenterHandler trata rotas de centro de custo
type ExpenseCenterHandler struct {
	service     service.ServiceExpenseCenter
	serviceUser service.ServiceUser
}

// NewExpenseCenterHandler injeta a dependência do service
func NewExpenseCenterHandler(s service.ServiceExpenseCenter, su service.ServiceUser) *ExpenseCenterHandler {
	return &ExpenseCenterHandler{
		service:     s,
		serviceUser: su,
	}
}

// CreateExpenseCenter godoc
// @Summary      Cria um centro de custo
// @Description  Cria um novo centro de custo
// @Tags         expense_centers
// @Accept       json
// @Produce      json
// @Param        expense_center  body  model.ExpenseCenterReceive  true  "Dados do centro de custo"
// @Success      200  {string}  string "ok"
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /expense-center [post]
func (h *ExpenseCenterHandler) CreateExpenseCenter(c *gin.Context) {
	var ec model.ExpenseCenterReceive
	if err := c.ShouldBindJSON(&ec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.serviceUser.GetUserByID(context.Background(), ec.UserID)
	if err != nil {
		if err.Error() == "GetUserByID: not exists user with this id" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateExpenseCenter(context.Background(), ec)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// GetExpenseCenterByID godoc
// @Summary      Busca centro de custo
// @Description  Retorna centro de custo por ID
// @Tags         expense_centers
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "CentroExpenseID"
// @Success      200 {object}  model.ExpenseCenter
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

// GetExpenseCenterByID godoc
// @Summary      Busca centro de custo
// @Description  Retorna centro de custo por ID
// @Tags         expense_centers
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "CentroExpenseID"
// @Success      200 {object}  model.ExpenseCenter
// @Failure      404 {object}  map[string]string
// @Router       /expense-center/id/{id} [get]
func (h *ExpenseCenterHandler) GetAllExpenseCenter(c *gin.Context) { //criar para pegar de todos
	center, err := h.service.GetAllExpenseCenter(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, center)
}

// DeleteExpenseCenter godoc
// @Summary      Deleta centro de custo
// @Description  Remove um centro de custo pelo ID
// @Tags         expense_centers
// @Accept       json
// @Produce      json
// @Param        expense_center  body  model.ExpenseCenterDeleteRequest  true  "CentroExpenseID"
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
	err := h.service.DeleteExpenseCenter(context.Background(), req.CentroExpenseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Expense center deleted"})
}

// UpdateExpenseCenter godoc
// @Summary      Atualiza centro de custo
// @Description  Atualiza os dados de um centro de custo pelo ID
// @Tags         expense_centers
// @Accept       json
// @Produce      json
// @Param        id              path  string                     true  "CentroExpenseID"
// @Param        expense_center  body  model.ExpenseCenterReceive true  "Dados atualizados"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /expense-center/id/{id} [put]
func (h *ExpenseCenterHandler) UpdateExpenseCenter(c *gin.Context) {

	var update model.ExpenseCenterReceive
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if update.UserID != "" {
		_, err := h.serviceUser.GetUserByID(context.Background(), update.UserID)
		if err != nil {
			if err.Error() == "GetUserByID: not exists user with this id" {
				c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	err := h.service.UpdateExpenseCenter(context.Background(), update.CentroExpenseID, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Expense center updated"})
}
