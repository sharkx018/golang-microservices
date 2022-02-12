package controllers

import (
	"encoding/json"
	"github.com/sharkx018/golang-microservices/mvc/services"
	"github.com/sharkx018/golang-microservices/mvc/utils"
	"log"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request)  {
	userIdParam := req.URL.Query().Get("user_id")
	log.Printf("About to process the user_id %s", userIdParam)
	userId, err:= strconv.ParseInt(userIdParam, 10, 64)
	if err!=nil{
		//Just return the Bad Request to the Client
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status: http.StatusBadRequest,
			Code: "bad_request",
		}

		jsonUserErr, _ := json.Marshal(apiErr)
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(apiErr.Status)
		resp.Write(jsonUserErr)
		return
	}
	user, apiErr := services.GetUser(userId)

	if apiErr!=nil{
		// Handle the error and return to the client
		jsonUserErr, _ := json.Marshal(apiErr)
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(apiErr.Status)
		resp.Write(jsonUserErr)

		return
	}

	// return user to client
	jsonUser, _ :=  json.Marshal(user)

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonUser)

}