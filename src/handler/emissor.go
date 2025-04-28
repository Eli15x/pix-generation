package handler

import (
	"context"
	"net/http"

	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

// EmissorHandler define o handler com injeção de dependência do service
type EmissorHandler struct {
	service service.ServiceEmissor
}

// NewEmissorHandler retorna um handler com o service injetado
func NewEmissorHandler(s service.ServiceEmissor) *EmissorHandler {
	return &EmissorHandler{service: s}
}

// CreateEmissor godoc
// @Summary      Cria um novo emissor
// @Description  Cria um novo emissor com nome, email, senha, documento e tokenConfraPix
// @Tags         emissor
// @Accept       json
// @Produce      json
// @Param        emissor  body      model.EmissorReceive  true  "Dados do emissor"
// @Success      200      {string}  string "ok"
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /emissor [post]
func (h *EmissorHandler) CreateEmissor(c *gin.Context) {
	var req model.EmissorReceive
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.CreateEmissor(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// GetEmissorByID godoc
// @Summary      Busca emissor por ID
// @Description  Retorna um emissor com base no EmissorID
// @Tags         emissor
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "EmissorID do emissor"
// @Success      200 {object}  model.Emissor
// @Failure      400 {object}  map[string]string
// @Failure      404 {object}  map[string]string
// @Router       /emissor/id/{id} [get]
func (h *EmissorHandler) GetEmissorByID(c *gin.Context) {
	id := c.Param("id")
	emissor, err := h.service.GetEmissorByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, emissor)
}

// GetAllEmissor godoc
// @Summary      Busca todos os emissores
// @Description  Retorna todos os emissores cadastrados
// @Tags         emissor
// @Accept       json
// @Produce      json
// @Success      200 {array}  model.Emissor
// @Failure      500 {object}  map[string]string
// @Router       /emissor [get]
func (h *EmissorHandler) GetAllEmissor(c *gin.Context) {
	emissores, err := h.service.GetAllEmissor(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, emissores)
}

// UpdateEmissor godoc
// @Summary      Atualiza emissor
// @Description  Atualiza dados de um emissor pelo EmissorID
// @Tags         emissor
// @Accept       json
// @Produce      json
// @Param        id        path      string                  true  "EmissorID do emissor"
// @Param        emissor   body      model.EmissorReceive    true  "Novos dados do emissor"
// @Success      200 {object}  map[string]string
// @Failure      400 {object}  map[string]string
// @Failure      500 {object}  map[string]string
// @Router       /emissor/id/{id} [put]
func (h *EmissorHandler) UpdateEmissor(c *gin.Context) {
	id := c.Param("id")
	var req model.EmissorReceive
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.UpdateEmissor(context.Background(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Emissor updated"})
}

// DeleteEmissor godoc
// @Summary      Deleta emissor
// @Description  Remove um emissor com base no EmissorID
// @Tags         emissor
// @Accept       json
// @Produce      json
// @Param        emissor  body      model.EmissorDeleteRequest  true  "EmissorID do emissor"
// @Success      200 {object}  map[string]string
// @Failure      400 {object}  map[string]string
// @Failure      500 {object}  map[string]string
// @Router       /emissor [delete]
func (h *EmissorHandler) DeleteEmissor(c *gin.Context) {
	var req model.EmissorDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.DeleteEmissor(context.Background(), req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Emissor deleted"})
}

// GetEmissorByDocumento godoc
// @Summary      Busca emissor por Documento
// @Description  Retorna um emissor com base no Documento (enviado via body)
// @Tags         emissor
// @Accept       json
// @Produce      json
// @Param        documento  body      model.EmissorDocumentoRequest  true  "Documento do emissor"
// @Success      200 {object}  model.Emissor
// @Failure      400 {object}  map[string]string
// @Failure      404 {object}  map[string]string
// @Router       /emissor/documento [post]
func (h *EmissorHandler) GetEmissorByDocumento(c *gin.Context) {
	var req model.EmissorDocumentoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	emissor, err := h.service.GetEmissorByDocumento(context.Background(), req.Documento)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, emissor)
}
