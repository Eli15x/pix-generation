package service

import (
	"os"
	"pix-generation/src/client"

	//"go.mongodb.org/mongo-driver/bson"
	"context"
	"errors"
	"strings"
	"sync"

	//"fmt"

	"pix-generation/src/middleware"
	"pix-generation/src/model"
	"pix-generation/src/repository"
	"pix-generation/src/utils"
)

var (
	instanceServiceAdminUser ServiceAdminUser
	onceServiceAdminUser     sync.Once
)

// aqui só falta adicionar o util igual do zcom e funciona.
type ServiceAdminUser interface {
	ValidateUser(ctx context.Context, email string, password string) (model.ResponseUser, error)
}

type adminUser struct{}

func GetInstanceAdminUser() ServiceAdminUser {
	onceServiceAdminUser.Do(func() {
		instanceServiceAdminUser = &adminUser{}
	})
	return instanceServiceAdminUser
}

func (u *adminUser) ValidateUser(ctx context.Context, email string, password string) (model.ResponseUser, error) {

	var user model.User
	var responseUser model.ResponseUser
	err := client.GetInstance().Ping(context.Background())

	if err == nil {
		emailValidate := map[string]interface{}{"Email": email}
		user, err = repository.GetInstanceUser().FindOne(ctx, "adminUser", emailValidate)
		if err != nil {
			return responseUser, errors.New("Validate user: problem to get information into MongoDB")
		}

	}

	passwordEncrypt := utils.Encrypt(password)

	if strings.Compare(passwordEncrypt, user.Password) != 0 {
		return responseUser, errors.New("Password user: wrong password")
	}

	token, err := middleware.GenerateJWT(user.UserID, []byte(os.Getenv("KEY_JWT_ACESS_GROUP")))
	if err != nil {
		return responseUser, errors.New("Failed to generate JWT")
	}

	responseUser = model.ResponseUser{
		UserID: user.UserID,
		JWT:    token,
	}

	return responseUser, nil
}
