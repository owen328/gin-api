package route

import (
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	adminGroup := router.Group("/admin")
	adminRouters(adminGroup)
	return router
}