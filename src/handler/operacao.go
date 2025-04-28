package handler

import (
	"context"
	"net/http"

	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

// OperacaoHandler define o handler
type OperacaoHandler struct {
	service service.ServiceOperacao
}

// NewOperacaoHandler cria um novo handler
func NewOperacaoHandler(s service.ServiceOperacao) *OperacaoHandler {
	return &OperacaoHandler{service: s}
}

// CreateOperacao godoc
// @Summary      Cria uma nova operação
// @Description  Cria uma nova operação
// @Tags         operacao
// @Accept       json
// @Produce      json
// @Param        operacao  body      model.OperacaoReceive  true  "Dados da operação"
// @Success      200        {string}  string "ok"
// @Failure      400        {object}  map[string]string
// @Failure      500        {object}  map[string]string
// @Router       /operacao [post]
func (h *OperacaoHandler) CreateOperacao(c *gin.Context) {
	var req model.OperacaoReceive
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.CreateOperacao(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// GetOperacaoByID godoc
// @Summary      Busca operação por OperacaoID
// @Description  Retorna uma operação pelo OperacaoID
// @Tags         operacao
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "OperacaoID da operação"
// @Success      200 {object}  model.Operacao
// @Failure      400 {object}  map[string]string
// @Failure      404 {object}  map[string]string
// @Router       /operacao/id/{id} [get]
func (h *OperacaoHandler) GetOperacaoByID(c *gin.Context) {
	id := c.Param("id")

	operacao, err := h.service.GetOperacaoByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, operacao)
}

// GetAllOperacao godoc
// @Summary      Busca todas as operações
// @Description  Retorna todas as operações cadastradas
// @Tags         operacao
// @Accept       json
// @Produce      json
// @Success      200 {array}  model.Operacao
// @Failure      500 {object}  map[string]string
// @Router       /operacao [get]
func (h *OperacaoHandler) GetAllOperacao(c *gin.Context) {
	operacoes, err := h.service.GetAllOperacao(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, operacoes)
}

// UpdateOperacao godoc
// @Summary      Atualiza operação
// @Description  Atualiza dados de uma operação
// @Tags         operacao
// @Accept       json
// @Produce      json
// @Param        id         path      string                 true  "OperacaoID da operação"
// @Param        operacao   body      model.OperacaoReceive  true  "Novos dados da operação"
// @Success      200 {object}  map[string]string
// @Failure      400 {object}  map[string]string
// @Failure      500 {object}  map[string]string
// @Router       /operacao/id/{id} [put]
func (h *OperacaoHandler) UpdateOperacao(c *gin.Context) {
	id := c.Param("id")

	var req model.OperacaoReceive
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.UpdateOperacao(context.Background(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Operacao updated"})
}

// DeleteOperacao godoc
// @Summary      Deleta operação
// @Description  Remove uma operação
// @Tags         operacao
// @Accept       json
// @Produce      json
// @Param        operacao  body      model.OperacaoDeleteRequest  true  "OperacaoID da operação"
// @Success      200 {object}  map[string]string
// @Failure      400 {object}  map[string]string
// @Failure      500 {object}  map[string]string
// @Router       /operacao [delete]
func (h *OperacaoHandler) DeleteOperacao(c *gin.Context) {
	var req model.OperacaoDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.DeleteOperacao(context.Background(), req.OperacaoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Operacao deleted"})
}
