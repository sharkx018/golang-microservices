package services

import (
	"github.com/sharkx018/golang-microservices/mvc/domain"
	"github.com/sharkx018/golang-microservices/mvc/utils"
)

func GetUser(userId int64 )(*domain.User, *utils.ApplicationError){
	return domain.GetUser(userId)
}