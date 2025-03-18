package handler

import (
	"context"
	"net/http"

	"pix-generation/src/model"
	"pix-generation/src/service"

	"github.com/gin-gonic/gin"
)

func ValidateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Email == "" {
		c.String(http.StatusBadRequest, "Validate User Error: email not find")
		return
	}

	if user.Password == "" {
		c.String(http.StatusBadRequest, "Create User Error: password not find")
		return
	}

	response, err := service.GetInstanceUser().ValidateUser(context.Background(), user.Email, user.Password)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := service.GetInstanceUser().CreateUser(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
