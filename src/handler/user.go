package handler

import (
	"net/http"
	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

// UserHandler trata rotas relacionadas a usuários
type UserHandler struct {
	service service.ServiceUser
}

// NewUserHandler injeta a dependência do ServiceUser
func NewUserHandler(s service.ServiceUser) *UserHandler {
	return &UserHandler{service: s}
}

// ValidateUser godoc
// @Summary      Valida usuário
// @Description  Verifica se o e-mail e senha são válidos
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        credentials  body  model.UserLoginRequest  true  "E-mail e senha do usuário"
// @Success      200   {object}  model.User
// @Failure      400   {string}  string
// @Router       /login [post]
func (h *UserHandler) ValidateUser(c *gin.Context) {
	var user model.UserLoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email e senha são obrigatórios"})
		return
	}

	ctx := c.Request.Context()
	response, err := h.service.ValidateUser(ctx, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateUser godoc
// @Summary      Cria um novo usuário
// @Description  Cadastra um novo usuário
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "Dados do usuário"
// @Success      200   {object}  model.User
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /register [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	response, err := h.service.CreateUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetUserByID godoc
// @Summary      Busca um usuário por ID
// @Description  Retorna um usuário pelo ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body  model.UserIDRequest  true  "ID do usuário"
// @Success      200   {object}  model.User
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /user [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	var user model.UserIDRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	userValue, err := h.service.GetUser(ctx, user.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userValue)
}

// UpdateUser godoc
// @Summary      Atualiza um usuário
// @Description  Edita os dados de um usuário existente
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "Dados atualizados"
// @Success      200   {string}  string "Usuário atualizado"
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /user [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	err := h.service.EditUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Usuário atualizado")
}

// DeleteUser godoc
// @Summary      Deleta um usuário
// @Description  Remove um usuário do sistema
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserDeleteRequest  true  "Dados do usuário a ser deletado"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /user [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	var user model.UserDeleteRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	err := h.service.DeleteUser(ctx, user.Document)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Usuário deletado"})
}

// GetAllUsers godoc
// @Summary      Lista todos os usuários
// @Description  Retorna uma lista de todos os usuários cadastrados
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.User
// @Failure      500  {object}  map[string]string
// @Router       /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := h.service.GetUsers(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
