package domain

import (
	"fmt"
	"github.com/sharkx018/golang-microservices/mvc/utils"
	"net/http"
)

var userMap = map[int64]*User{
	123: &User{
		Id: 123,
		FirstName: "Mukul",
		LastName: "Verma",
		Email: "aryanv018@gmail.com",
	},
}

func GetUser(userId int64) (*User, *utils.ApplicationError)  {

	user := userMap[userId]
	if user == nil{
		return nil, &utils.ApplicationError{
			Message: fmt.Sprintf("User %v not found", userId),
			Status: http.StatusNotFound,
			Code:"not_found",
		}
	}
	return user, nil
}