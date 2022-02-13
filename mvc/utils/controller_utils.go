package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, statusCode int, body interface{}) {

	switch c.GetHeader("Accept") {

	case "application/xml":
		c.XML(statusCode, body)
		return
	default:
		c.JSON(statusCode, body)
	}

}
