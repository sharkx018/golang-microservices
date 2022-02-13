package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	//router = gin.New()
}

func StartApp() {
	mapUrls()

	//http.HandleFunc("/users", controllers.GetUser)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
