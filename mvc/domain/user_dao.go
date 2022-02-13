package domain

import (
	"fmt"
	"github.com/sharkx018/golang-microservices/mvc/utils"
	"log"
	"net/http"
)

type userDao struct {
}

type userDaoInterface interface {
	GetUser(userId int64) (*User, *utils.ApplicationError)
}

var (
	userMap = map[int64]*User{
		123: &User{
			Id:        123,
			FirstName: "Mukul",
			LastName:  "Verma",
			Email:     "aryanv018@gmail.com",
		},
	}

	UserDao userDaoInterface
)

func init()  {
	UserDao = &userDao{}
}


func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {

	log.Println("We are using the database")
	user := userMap[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message: fmt.Sprintf("User %v not found", userId),
			Status:  http.StatusNotFound,
			Code:    "not_found",
		}
	}
	return user, nil
}
