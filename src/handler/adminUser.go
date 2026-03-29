package handler

import (
	"errors"
	"net/http"
	"pix-generation/src/model"
	apperrors "pix-generation/src/pkg"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

// AdminUserHandler trata rotas relacionadas a usuários administradores
type AdminUserHandler struct {
	service service.ServiceAdminUser
}

// NewAdminUserHandler injeta a dependência do ServiceUser
func NewAdminUserHandler(s service.ServiceAdminUser) *AdminUserHandler {
	return &AdminUserHandler{service: s}
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
func (h *AdminUserHandler) ValidateUser(c *gin.Context) {
	var user model.AdminUserLoginRequest
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
		var ae *apperrors.AppError
		if errors.As(err, &ae) {
			c.JSON(apperrors.HTTPStatus(err), ae.Public())
			return
		}
		c.JSON(http.StatusInternalServerError, apperrors.ErrInternal.Public())
		return
	}
	c.JSON(http.StatusOK, response)
}
