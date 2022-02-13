package services

import (
	"github.com/sharkx018/golang-microservices/mvc/domain"
	"github.com/sharkx018/golang-microservices/mvc/utils"
	"net/http"
)

type itemService struct {
}

var (
	ItemService itemService
)

func (u *itemService)GetItem() (*domain.Item, *utils.ApplicationError) {

	return nil, &utils.ApplicationError{
		Message: "Implement me",
		Status: http.StatusInternalServerError,
	}
}
