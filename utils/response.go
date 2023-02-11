package utils

import "github.com/gin-gonic/gin"
import "net/http"

func Response(c *gin.Context, code int, message string, data interface{})  {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"message": message,
		"data": data,
	})
}


func Success(c *gin.Context, data interface{}){
	Response(c, 200, "ok", data)
}

func Fail(c *gin.Context, message string) {
	Response(c, 400, message, nil)
}

func UnAuthorization(c *gin.Context)  {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": http.StatusUnauthorized,
		"message": "授权失败",
	})
}