package handler

import (
	"net/http"

	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

// UsuarioHandler define o handler
type UsuarioHandler struct {
	service service.ServiceUsuario
}

// NewUsuarioHandler cria um novo handler
func NewUsuarioHandler(s service.ServiceUsuario) *UsuarioHandler {
	return &UsuarioHandler{service: s}
}

// CreateUsuario godoc
// @Summary      Cria um novo usuário
// @Description  Cria um novo usuário
// @Tags         usuario
// @Accept       json
// @Produce      json
// @Param        usuario  body      model.UsuarioReceive  true  "Dados do usuário"
// @Success      200      {string}  string "ok"
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /usuario [post]
func (h *UsuarioHandler) CreateUsuario(c *gin.Context) {
	var req model.UsuarioReceive
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.CreateUsuario(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// GetUsuarioByID godoc
// @Summary      Busca usuário por ID
// @Description  Retorna um usuário pelo usuarioID
// @Tags         usuario
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "UsuarioID do usuário"
// @Success      200 {object}  model.Usuario
// @Failure      400 {object}  map[string]string
// @Failure      404 {object}  map[string]string
// @Router       /usuario/id/{id} [get]
func (h *UsuarioHandler) GetUsuarioByID(c *gin.Context) {
	id := c.Param("id")
	usuario, err := h.service.GetUsuarioByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

// GetAllUsuario godoc
// @Summary      Busca todos os usuários
// @Description  Retorna todos os usuários cadastrados
// @Tags         usuario
// @Accept       json
// @Produce      json
// @Success      200 {array}  model.Usuario
// @Failure      500 {object}  map[string]string
// @Router       /usuario [get]
func (h *UsuarioHandler) GetAllUsuario(c *gin.Context) {
	usuarios, err := h.service.GetAllUsuario(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

// UpdateUsuario godoc
// @Summary      Atualiza usuário
// @Description  Atualiza dados de um usuário
// @Tags         usuario
// @Accept       json
// @Produce      json
// @Param        id       path      string                true  "UsuarioID do usuário"
// @Param        usuario  body      model.UsuarioReceive  true  "Novos dados do usuário"
// @Success      200 {object}  map[string]string
// @Failure      400 {object}  map[string]string
// @Failure      500 {object}  map[string]string
// @Router       /usuario/id/{id} [put]
func (h *UsuarioHandler) UpdateUsuario(c *gin.Context) {
	id := c.Param("id")
	var req model.UsuarioReceive
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.UpdateUsuario(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Usuario updated"})
}

// DeleteUsuario godoc
// @Summary      Deleta usuário
// @Description  Remove um usuário
// @Tags         usuario
// @Accept       json
// @Produce      json
// @Param        usuario  body      model.UsuarioDeleteRequest  true  "UsuarioID do usuário"
// @Success      200 {object}  map[string]string
// @Failure      400 {object}  map[string]string
// @Failure      500 {object}  map[string]string
// @Router       /usuario [delete]
func (h *UsuarioHandler) DeleteUsuario(c *gin.Context) {
	var req model.UsuarioDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.DeleteUsuario(c.Request.Context(), req.UsuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Usuario deleted"})
}

// GetUsuarioByEmail godoc
// @Summary      Busca usuário por Email
// @Description  Retorna um usuário pelo email
// @Tags         usuario
// @Accept       json
// @Produce      json
// @Param        email  body      model.UsuarioEmailRequest  true  "Email do usuário"
// @Success      200 {object}  model.Usuario
// @Failure      400 {object}  map[string]string
// @Failure      404 {object}  map[string]string
// @Router       /usuario/email [post]
func (h *UsuarioHandler) GetUsuarioByEmail(c *gin.Context) {
	var req model.UsuarioEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usuario, err := h.service.GetUsuarioByEmail(c.Request.Context(), req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuario)
}
