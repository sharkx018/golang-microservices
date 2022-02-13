package app

import (
	"github.com/sharkx018/golang-microservices/mvc/controllers"
)

func mapUrls()  {

	router.GET("/users/:user_id", controllers.GetUser)

}