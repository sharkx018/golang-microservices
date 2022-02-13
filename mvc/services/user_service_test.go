package services

import (
	"github.com/sharkx018/golang-microservices/mvc/domain"
	"github.com/sharkx018/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

type usersDaoMock struct {
}

var (
	userDaoMock usersDaoMock
	getUserMockFunction func (userId int64) (*domain.User, *utils.ApplicationError) // this is done to override the functionality according to different usecases
)

func init(){
	domain.UserDao = &userDaoMock
}

func (u *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError){
	return getUserMockFunction(userId)
}


func TestUserService_GetUserNotFoundInDatabase(t *testing.T) {

	getUserMockFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		log.Println("Using the user mock")

		return nil, &utils.ApplicationError{
			Message: "User 0 not found",
			Status: http.StatusNotFound,
		}
	}

	user, err := UserService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "User 0 not found", err.Message)
}

func TestUserService_GetUserSuccess(t *testing.T) {

	getUserMockFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		log.Println("Using the user mock")

		return &domain.User{
			Id: 123,
			FirstName: "Mukul",
			LastName: "Verma",
			Email: "aryanv018@gmail.com",
		}, nil
	}

	user, err := UserService.GetUser(123)

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Verma", user.LastName)
	assert.EqualValues(t, "Mukul", user.FirstName)
	assert.EqualValues(t, "aryanv018@gmail.com", user.Email)
}