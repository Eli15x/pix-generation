package handler

import (
	"context"
	"net/http"

	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

// SignatureHandler define o handler
type SignatureHandler struct {
	service service.ServiceSignature
}

// NewSignatureHandler cria um novo handler
func NewSignatureHandler(s service.ServiceSignature) *SignatureHandler {
	return &SignatureHandler{service: s}
}

// CreateSignature godoc
// @Summary      Cria uma nova assinatura
// @Description  Cria uma nova assinatura
// @Tags         signature
// @Accept       json
// @Produce      json
// @Param        signature  body      model.SignatureReceive  true  "Dados da assinatura"
// @Success      200        {string}  string "ok"
// @Failure      400        {object}  map[string]string
// @Failure      500        {object}  map[string]string
// @Router       /signature [post]
func (h *SignatureHandler) CreateSignature(c *gin.Context) {
	var req model.SignatureReceive
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.CreateSignature(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// GetSignatureByID godoc
// @Summary      Busca assinatura por ID
// @Description  Retorna uma assinatura pelo SignatureID
// @Tags         signature
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "SignatureID da assinatura"
// @Success      200 {object}  model.Signature
// @Failure      400 {object}  map[string]string
// @Failure      404 {object}  map[string]string
// @Router       /signature/id/{id} [get]
func (h *SignatureHandler) GetSignatureByID(c *gin.Context) {
	id := c.Param("id")
	signature, err := h.service.GetSignatureByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, signature)
}

// GetAllSignature godoc
// @Summary      Busca todas as assinaturas
// @Description  Retorna todas as assinaturas cadastradas
// @Tags         signature
// @Accept       json
// @Produce      json
// @Success      200 {array}  model.Signature
// @Failure      500 {object}  map[string]string
// @Router       /signature [get]
func (h *SignatureHandler) GetAllSignature(c *gin.Context) {
	signatures, err := h.service.GetAllSignature(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, signatures)
}

// UpdateSignature godoc
// @Summary      Atualiza assinatura
// @Description  Atualiza os dados de uma assinatura
// @Tags         signature
// @Accept       json
// @Produce      json
// @Param        id         path      string                   true  "SignatureID da assinatura"
// @Param        signature  body      model.SignatureReceive   true  "Novos dados da assinatura"
// @Success      200 {object}  map[string]string
// @Failure      400 {object}  map[string]string
// @Failure      500 {object}  map[string]string
// @Router       /signature/id/{id} [put]
func (h *SignatureHandler) UpdateSignature(c *gin.Context) {
	id := c.Param("id")
	var req model.SignatureReceive
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.UpdateSignature(context.Background(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Signature updated"})
}

// DeleteSignature godoc
// @Summary      Deleta assinatura
// @Description  Remove uma assinatura pelo SignatureID
// @Tags         signature
// @Accept       json
// @Produce      json
// @Param        signature  body      model.SignatureDeleteRequest  true  "SignatureID da assinatura"
// @Success      200 {object}  map[string]string
// @Failure      400 {object}  map[string]string
// @Failure      500 {object}  map[string]string
// @Router       /signature [delete]
func (h *SignatureHandler) DeleteSignature(c *gin.Context) {
	var req model.SignatureDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.DeleteSignature(context.Background(), req.SignatureID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Signature deleted"})
}

// GetSignatureByClienteID godoc
// @Summary      Busca assinatura pelo ClienteID
// @Description  Retorna assinaturas com base no ClienteID
// @Tags         signature
// @Accept       json
// @Produce      json
// @Param        cliente  body      model.SignatureClienteRequest  true  "ClienteID do cliente"
// @Success      200 {array}  model.Signature
// @Failure      400 {object}  map[string]string
// @Failure      404 {object}  map[string]string
// @Router       /signature/cliente [post]
func (h *SignatureHandler) GetSignatureByClienteID(c *gin.Context) {
	var req model.SignatureClienteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	signatures, err := h.service.GetSignatureByClienteID(context.Background(), req.ClienteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, signatures)
}
