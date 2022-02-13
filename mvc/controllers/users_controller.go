package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sharkx018/golang-microservices/mvc/services"
	"github.com/sharkx018/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context)  {

	userIdParam := c.Param("user_id")
	userId, err:= strconv.ParseInt(userIdParam, 10, 64)
	if err!=nil{
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.Respond(c, apiErr.Status, apiErr)
		return
	}
	user, apiErr := services.UserService.GetUser(userId)

	if apiErr!=nil{
		utils.Respond(c, apiErr.Status, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
