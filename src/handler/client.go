package handler

import (
	"context"
	"net/http"

	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	service     service.ServiceClient
	serviceUser service.ServiceUser
}

func NewClientHandler(s service.ServiceClient, su service.ServiceUser) *ClientHandler {
	return &ClientHandler{
		service:     s,
		serviceUser: su,
	}
}

// CreateClient godoc
// @Summary      Cria um novo cliente
// @Description  Cria um novo cliente vinculado a um usuário já existente
// @Tags         client
// @Accept       json
// @Produce      json
// @Param        client  body      model.ClientReceive  true  "Dados do cliente"
// @Success      200     {string}  string  "ok"
// @Failure      400     {object}  map[string]string  "Erro de validação nos dados enviados"
// @Failure      404     {object}  map[string]string  "Usuário vinculado não encontrado"
// @Failure      500     {object}  map[string]string  "Erro interno ao criar cliente"
// @Router       /client [post]
func (h *ClientHandler) CreateClient(c *gin.Context) {
	var req model.ClientReceive
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.serviceUser.GetUserByID(context.Background(), req.UserID)
	if err != nil {
		if err.Error() == "GetUserByID: not exists user with this id" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateClient(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// GetClientByID godoc
// @Summary Busca cliente por ID
// @Description Retorna um cliente pelo ID
// @Tags client
// @Accept json
// @Produce json
// @Param client body model.ClientRequest true "Dados do cliente"
// @Success 200 {object} model.Client
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /client/id/{id} [get]
func (h *ClientHandler) GetClientByID(c *gin.Context) {
	var req model.ClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := h.service.GetClientByID(context.Background(), req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

// GetClientByUserID godoc
// @Summary Busca cliente por UserID
// @Description Retorna um cliente pelo UserID
// @Tags client
// @Accept json
// @Produce json
// @Param client body model.ClientUserRequest true "Dados do cliente"
// @Success 200 {object} model.Client
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /client/id/{id} [get]
func (h *ClientHandler) GetClientByUserID(c *gin.Context) {
	var req model.ClientUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := h.service.GetClientByUserID(context.Background(), req.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

// GetAllClient godoc
// @Summary Busca todos os clientes
// @Description Retorna todos os clientes
// @Tags client
// @Accept json
// @Produce json
// @Success 200 {array} model.Client
// @Failure 500 {object} map[string]string
// @Router /client [get]
func (h *ClientHandler) GetAllClient(c *gin.Context) {
	clients, err := h.service.GetAllClient(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

// UpdateClient godoc
// @Summary Atualiza cliente
// @Description Atualiza um cliente
// @Tags client
// @Accept json
// @Produce json
// @Param id path string true "ID do cliente"
// @Param client body model.ClientReceive true "Dados do cliente"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /client/id/{id} [put]
func (h *ClientHandler) UpdateClient(c *gin.Context) {
	id := c.Param("id")
	var req model.ClientReceive
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.UpdateClient(context.Background(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Client updated"})
}

// DeleteClient godoc
// @Summary Deleta cliente
// @Description Deleta um cliente
// @Tags client
// @Accept json
// @Produce json
// @Param client body model.ClientRequest true "ID do cliente"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /client [delete]
func (h *ClientHandler) DeleteClient(c *gin.Context) {
	var req model.ClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.DeleteClient(context.Background(), req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Client deleted"})
}

// GetClientByCpf godoc
// @Summary Busca cliente por CPF
// @Description Retorna um cliente pelo CPF
// @Tags client
// @Accept json
// @Produce json
// @Param client body model.ClientCpfRequest true "CPF do cliente"
// @Success 200 {object} model.Client
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /client/cpf [post]
func (h *ClientHandler) GetClientByCpf(c *gin.Context) {
	var req model.ClientCpfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := h.service.GetClientByCpf(context.Background(), req.CPF)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}
