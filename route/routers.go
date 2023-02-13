package route

import (
	"github.com/gin-gonic/gin"
	"go_learn/middleware"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.Use(middleware.LogRequestMiddleware())
	initGroupRoute(router)
	return router
}

func initGroupRoute(router *gin.Engine) {
	apiGroup := router.Group("/api")
	initAdminRouter(apiGroup)
	initSellerReportRouter(apiGroup)
}
