package service

import (
	"fmt"
	"os"
	"pix-generation/src/client"
	"time"

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

	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceServiceAdminUser ServiceAdminUser
	onceServiceAdminUser     sync.Once
)

type ServiceAdminUser interface {
	ValidateUser(ctx context.Context, email string, password string) (model.AdminResponseUser, error)
	CreateUser(ctx context.Context, user model.AdminUser) (model.AdminResponseUser, error)
	EditUser(ctx context.Context, user model.AdminUser) error
	DeleteUser(ctx context.Context, document string) error
	GetUserByEmail(ctx context.Context, email string) (model.AdminUser, error)
	GetUserByID(ctx context.Context, id string) (model.AdminUser, error)
	GetUserByDocument(ctx context.Context, document string) (model.AdminUser, error)
}

type adminUser struct{}

func GetInstanceAdminUser() ServiceAdminUser {
	onceServiceAdminUser.Do(func() {
		instanceServiceAdminUser = &adminUser{}
	})
	return instanceServiceAdminUser
}

func (u *adminUser) ValidateUser(ctx context.Context, email string, password string) (model.AdminResponseUser, error) {

	var user model.AdminUser
	var responseUser model.AdminResponseUser
	err := client.GetInstance().Ping(context.Background())

	if err == nil {
		emailValidate := map[string]interface{}{"Email": email}
		user, err = repository.GetInstanceAdminUser().FindOne(ctx, "adminUser", emailValidate)
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

	responseUser = model.AdminResponseUser{
		UserID: user.UserID,
		JWT:    token,
	}

	return responseUser, nil
}

func (u *adminUser) CreateUser(ctx context.Context, user model.AdminUser) (model.AdminResponseUser, error) {

	var responseUser model.AdminResponseUser
	passwordEncrypt := utils.Encrypt(user.Password)
	user.Password = passwordEncrypt
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	var userId = utils.CreateCodeId()
	user.UserID = userId

	userInsert := structs.Map(user)
	_, err := client.GetInstance().Insert(ctx, "adminUser", userInsert)
	if err != nil {
		return responseUser, errors.New("Create Admin User: problem to insert into MongoDB")
	}

	token, err := middleware.GenerateJWT(user.UserID, []byte(os.Getenv("KEY_JWT_ACESS_GROUP")))
	if err != nil {
		return responseUser, errors.New("Failed to generate JWT")
	}

	responseUser = model.AdminResponseUser{
		UserID: user.UserID,
		JWT:    token,
	}

	return responseUser, nil

}

func (u *adminUser) EditUser(ctx context.Context, user model.AdminUser) error {
	var existingUser model.AdminUser
	filter := bson.M{"Email": user.Email}

	emailValidate := map[string]interface{}{"Email": user.Email}
	existingUser, err := repository.GetInstanceAdminUser().FindOne(ctx, "adminUser", emailValidate)
	if err != nil {
		return errors.New("Edit Admin User: could not find existing user")
	}

	if user.Password != "" {
		newEncrypted := utils.Encrypt(user.Password)
		if newEncrypted != existingUser.Password {
			user.Password = newEncrypted
		} else {
			user.Password = existingUser.Password
		}
	}

	user.UpdatedAt = time.Now()
	userUpdate := structs.Map(user)
	filteredUpdate := cleanMap(userUpdate)
	change := bson.M{"$set": filteredUpdate}

	fmt.Println(change)
	_, err = client.GetInstance().UpdateOne(ctx, "adminUser", filter, change)
	if err != nil {
		return errors.New("Edit Admin User: problem to update into MongoDB")
	}

	return nil
}

func (u *adminUser) DeleteUser(ctx context.Context, document string) error {

	documentId := map[string]interface{}{"Document": document}

	err := client.GetInstance().Remove(ctx, "adminUser", documentId)
	if err != nil {
		return errors.New("Delete User: problem to delete into MongoDB")
	}

	return nil
}

func (u *adminUser) GetUserByEmail(ctx context.Context, email string) (model.AdminUser, error) {

	Email := map[string]interface{}{"Email": email}

	user, err := repository.GetInstanceAdminUser().FindOne(ctx, "adminUser", Email)
	if err != nil {
		return model.AdminUser{}, errors.New("Get Users By Acess: problem to Find Id into MongoDB")
	}

	return user, nil
}

func (u *adminUser) GetUserByID(ctx context.Context, id string) (model.AdminUser, error) {
	filter := map[string]interface{}{"UserID": id}
	fmt.Println(id)
	user, err := repository.GetInstanceAdminUser().FindOne(ctx, "adminUser", filter)
	if err != nil {
		return model.AdminUser{}, errors.New("GetUserByID: problem to find user by UserID in MongoDB")
	}

	if user == (model.AdminUser{}) {
		return model.AdminUser{}, errors.New("GetUserByID: not exists user with this id")
	}

	return user, nil
}

func (u *adminUser) GetUserByDocument(ctx context.Context, document string) (model.AdminUser, error) {
	filter := map[string]interface{}{"Document": document}
	user, err := repository.GetInstanceAdminUser().FindOne(ctx, "adminUser", filter)
	if err != nil {
		return model.AdminUser{}, errors.New("GetUserByDocument: problem to find user by document in MongoDB")
	}
	return user, nil
}
