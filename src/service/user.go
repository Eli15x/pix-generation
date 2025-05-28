package service

import (
	"fmt"
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
	instanceServiceUser ServiceUser
	onceServiceUser     sync.Once
)

// aqui só falta adicionar o util igual do zcom e funciona.
type ServiceUser interface {
	ValidateUser(ctx context.Context, email string, password string) (model.ResponseUser, error)
	GetUser(ctx context.Context, id string) (model.User, error)
	GetUserByName(ctx context.Context, name string) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserByDocument(ctx context.Context, document string) (model.User, error)
	GetUserByID(ctx context.Context, id string) (model.User, error)
	GetUsersByClientId(ctx context.Context, idAcess int) ([]model.User, error)
	GetUsers(ctx context.Context) ([]model.User, error)

	CreateUser(ctx context.Context, user model.User) (model.ResponseUser, error)
	EditUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, document string) error
}

type user struct{}

func GetInstanceUser() ServiceUser {
	onceServiceUser.Do(func() {
		instanceServiceUser = &user{}
	})
	return instanceServiceUser
}

func (u *user) ValidateUser(ctx context.Context, email string, password string) (model.ResponseUser, error) {

	var user model.User
	var responseUser model.ResponseUser
	err := client.GetInstance().Ping(context.Background())

	if err == nil {
		emailValidate := map[string]interface{}{"Email": email}
		user, err = repository.GetInstanceUser().FindOne(ctx, "user", emailValidate)
		if err != nil {
			return responseUser, errors.New("Validate user: problem to get information into MongoDB")
		}

	}

	passwordEncrypt := utils.Encrypt(password)

	if strings.Compare(passwordEncrypt, user.Password) != 0 {
		return responseUser, errors.New("Password user: wrong password")
	}

	// Gera o token JWT após validação bem-sucedida
	token, err := middleware.GenerateJWT(user.UserID)
	if err != nil {
		return responseUser, errors.New("Failed to generate JWT")
	}

	responseUser = model.ResponseUser{
		UserID: user.UserID,
		JWT:    token,
	}

	return responseUser, nil
}

func (u *user) GetUser(ctx context.Context, id string) (model.User, error) {
	var user model.User

	userId := map[string]interface{}{"UserID": id}

	user, err := repository.GetInstanceUser().FindOne(ctx, "user", userId)
	if err != nil {
		return user, errors.New("Get user: problem to Find Id into MongoDB")
	}

	return user, nil
}

func (u *user) GetUserByName(ctx context.Context, name string) (model.User, error) {

	Name := map[string]interface{}{"Name": name}
	user, err := repository.GetInstanceUser().FindOne(ctx, "user", Name)
	if err != nil {

		return user, errors.New("Get user by name: problem to Find Id into MongoDB")
	}

	return user, nil
}

func (u *user) GetUsersByClientId(ctx context.Context, clientId int) ([]model.User, error) {

	IdClient := map[string]interface{}{"clientId": clientId}

	users, err := repository.GetInstanceUser().Find(ctx, "user", IdClient)
	if err != nil {
		return nil, errors.New("Get Users By Acess: problem to Find Id into MongoDB")
	}

	return users, nil
}

func (u *user) GetUserByEmail(ctx context.Context, email string) (model.User, error) {

	Email := map[string]interface{}{"Email": email}

	user, err := repository.GetInstanceUser().FindOne(ctx, "user", Email)
	if err != nil {
		return model.User{}, errors.New("Get Users By Acess: problem to Find Id into MongoDB")
	}

	return user, nil
}

func (u *user) GetUserByID(ctx context.Context, id string) (model.User, error) {
	filter := map[string]interface{}{"UserID": id}
	fmt.Println(id)
	user, err := repository.GetInstanceUser().FindOne(ctx, "user", filter)
	if err != nil {
		return model.User{}, errors.New("GetUserByID: problem to find user by UserID in MongoDB")
	}

	if user == (model.User{}) {
		return model.User{}, errors.New("GetUserByID: not exists user with this id")
	}

	return user, nil
}

func (u *user) GetUserByDocument(ctx context.Context, document string) (model.User, error) {
	filter := map[string]interface{}{"Document": document}
	user, err := repository.GetInstanceUser().FindOne(ctx, "user", filter)
	if err != nil {
		return model.User{}, errors.New("GetUserByDocument: problem to find user by document in MongoDB")
	}
	return user, nil
}

func (u *user) GetUsers(ctx context.Context) ([]model.User, error) {

	all := map[string]interface{}{}

	users, err := repository.GetInstanceUser().Find(ctx, "user", all)
	if err != nil {
		return nil, errors.New("Get Users: problem to Find Id into MongoDB")
	}

	return users, nil
}

func (u *user) CreateUser(ctx context.Context, user model.User) (model.ResponseUser, error) {

	var responseUser model.ResponseUser
	passwordEncrypt := utils.Encrypt(user.Password)
	user.Password = passwordEncrypt
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	var userId = utils.CreateCodeId()
	user.UserID = userId

	userInsert := structs.Map(user)
	_, err := client.GetInstance().Insert(ctx, "user", userInsert)
	if err != nil {
		return responseUser, errors.New("Create user: problem to insert into MongoDB")
	}

	token, err := middleware.GenerateJWT(user.UserID)
	if err != nil {
		return responseUser, errors.New("Failed to generate JWT")
	}

	responseUser = model.ResponseUser{
		UserID: user.UserID,
		JWT:    token,
	}

	return responseUser, nil

}

func (u *user) EditUser(ctx context.Context, user model.User) error {
	var existingUser model.User
	filter := bson.M{"Email": user.Email}

	emailValidate := map[string]interface{}{"Email": user.Email}
	existingUser, err := repository.GetInstanceUser().FindOne(ctx, "user", emailValidate)
	if err != nil {
		return errors.New("Edit User: could not find existing user")
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
	_, err = client.GetInstance().UpdateOne(ctx, "user", filter, change)
	if err != nil {
		return errors.New("Edit User: problem to update into MongoDB")
	}

	return nil
}

func (u *user) DeleteUser(ctx context.Context, document string) error {

	documentId := map[string]interface{}{"Document": document}

	err := client.GetInstance().Remove(ctx, "user", documentId)
	if err != nil {
		return errors.New("Delete User: problem to delete into MongoDB")
	}

	return nil
}

func cleanMap(m map[string]interface{}) map[string]interface{} {
	cleaned := make(map[string]interface{})
	for k, v := range m {
		// Ignora campos vazios, zero ou nil
		if v == nil {
			continue
		}
		// Ignora strings vazias
		if str, ok := v.(string); ok && str == "" {
			continue
		}
		// Ignora datas zero
		if t, ok := v.(time.Time); ok && t.IsZero() {
			continue
		}

		cleaned[k] = v
	}
	return cleaned
}
