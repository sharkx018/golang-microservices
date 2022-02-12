package app

import (
	"github.com/sharkx018/golang-microservices/mvc/controllers"
	"net/http"
)

func StartApp()  {
	http.HandleFunc("/users", controllers.GetUser)

	if err:= http.ListenAndServe(":8080", nil); err!=nil {
		panic(err)
	}
}